package instrumentation_ebpf

import (
	"context"

	"github.com/keyval-dev/odigos/common"
	"github.com/keyval-dev/odigos/odiglet/pkg/ebpf"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeploymentsReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	Directors map[common.ProgrammingLanguage]ebpf.Director
}

func (d *DeploymentsReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	err := ApplyEbpfToPodWorkload(ctx, d.Client, d.Directors, &common.PodWorkload{
		Name:      request.Name,
		Namespace: request.Namespace,
		Kind:      "Deployment",
	})

	return ctrl.Result{}, err
}