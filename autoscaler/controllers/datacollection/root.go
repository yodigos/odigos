package datacollection

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"sync"
	"time"

	odigosv1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var dm = &DelayManager{}

const (
	DAEMONSET_PATCH_DELAY = 20 * time.Second
	PATCH_DAEMONSET_RETRY = 3
)

func Sync(ctx context.Context, c client.Client, scheme *runtime.Scheme, imagePullSecrets []string, odigosVersion string) error {
	logger := log.FromContext(ctx)
	var collectorGroups odigosv1.CollectorsGroupList
	if err := c.List(ctx, &collectorGroups); err != nil {
		logger.Error(err, "Failed to list collectors groups")
		return err
	}

	var dataCollectionCollectorGroup *odigosv1.CollectorsGroup
	for _, collectorGroup := range collectorGroups.Items {
		if collectorGroup.Spec.Role == odigosv1.CollectorsGroupRoleNodeCollector {
			dataCollectionCollectorGroup = &collectorGroup
			break
		}
	}

	if dataCollectionCollectorGroup == nil {
		logger.V(3).Info("Data collection collector group doesn't exist, nothing to sync")
		return nil
	}

	var instApps odigosv1.InstrumentedApplicationList
	if err := c.List(ctx, &instApps); err != nil {
		logger.Error(err, "Failed to list instrumented apps")
		return err
	}

	var dests odigosv1.DestinationList
	if err := c.List(ctx, &dests); err != nil {
		logger.Error(err, "Failed to list destinations")
		return err
	}

	var processors odigosv1.ProcessorList
	if err := c.List(ctx, &processors); err != nil {
		logger.Error(err, "Failed to list processors")
		return err
	}

	return syncDataCollection(&instApps, &dests, &processors, dataCollectionCollectorGroup, ctx, c, scheme, imagePullSecrets, odigosVersion)
}

func syncDataCollection(instApps *odigosv1.InstrumentedApplicationList, dests *odigosv1.DestinationList, processors *odigosv1.ProcessorList,
	dataCollection *odigosv1.CollectorsGroup, ctx context.Context, c client.Client,
	scheme *runtime.Scheme, imagePullSecrets []string, odigosVersion string) error {
	logger := log.FromContext(ctx)
	logger.V(0).Info("Syncing data collection")

	_, err := syncConfigMap(instApps, dests, processors, dataCollection, ctx, c, scheme)
	if err != nil {
		logger.Error(err, "Failed to sync config map")
		return err
	}

	syncDaemonSetFunction := func(args ...interface{}) (*appsv1.DaemonSet, error) {
		return syncDaemonSet(dests, dataCollection, ctx, c, scheme, imagePullSecrets, odigosVersion)
	}
	dm.runFunctionWithDelayAndSkipNewCalls(DAEMONSET_PATCH_DELAY, syncDaemonSetFunction,
		dests, dataCollection, ctx, c, scheme, imagePullSecrets, odigosVersion)

	return nil
}

type DelayManager struct {
	mu         sync.Mutex
	inProgress bool
}

// runFunctionWithDelayAndSkipNewCalls runs the function with the specified delay and skips new calls until the function execution is finished
func (dm *DelayManager) runFunctionWithDelayAndSkipNewCalls(delay time.Duration, fn func(args ...interface{}) (*appsv1.DaemonSet, error), fnArgs ...interface{}) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	logger := log.FromContext(fnArgs[2].(context.Context))
	if dm.inProgress {
		logger.Info("Function execution in progress. Skipping...")
		return
	}

	dm.inProgress = true

	time.AfterFunc(delay, func() {
		dm.mu.Lock()
		defer dm.mu.Unlock()

		logger.Info("Sync DaemonSet function execution started...")
		for i := 0; i < PATCH_DAEMONSET_RETRY; i++ {
			_, err := fn(fnArgs...)
			if err == nil {
				dm.inProgress = false
				return
			}
		}

		dm.inProgress = false
	})
}
