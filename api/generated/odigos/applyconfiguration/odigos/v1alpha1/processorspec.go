/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	common "github.com/odigos-io/odigos/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ProcessorSpecApplyConfiguration represents a declarative configuration of the ProcessorSpec type for use
// with apply.
type ProcessorSpecApplyConfiguration struct {
	Type            *string                        `json:"type,omitempty"`
	ProcessorName   *string                        `json:"processorName,omitempty"`
	Notes           *string                        `json:"notes,omitempty"`
	Disabled        *bool                          `json:"disabled,omitempty"`
	Signals         []common.ObservabilitySignal   `json:"signals,omitempty"`
	CollectorRoles  []v1alpha1.CollectorsGroupRole `json:"collectorRoles,omitempty"`
	OrderHint       *int                           `json:"orderHint,omitempty"`
	ProcessorConfig *runtime.RawExtension          `json:"processorConfig,omitempty"`
}

// ProcessorSpecApplyConfiguration constructs a declarative configuration of the ProcessorSpec type for use with
// apply.
func ProcessorSpec() *ProcessorSpecApplyConfiguration {
	return &ProcessorSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithType(value string) *ProcessorSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithProcessorName sets the ProcessorName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProcessorName field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithProcessorName(value string) *ProcessorSpecApplyConfiguration {
	b.ProcessorName = &value
	return b
}

// WithNotes sets the Notes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Notes field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithNotes(value string) *ProcessorSpecApplyConfiguration {
	b.Notes = &value
	return b
}

// WithDisabled sets the Disabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disabled field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithDisabled(value bool) *ProcessorSpecApplyConfiguration {
	b.Disabled = &value
	return b
}

// WithSignals adds the given value to the Signals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Signals field.
func (b *ProcessorSpecApplyConfiguration) WithSignals(values ...common.ObservabilitySignal) *ProcessorSpecApplyConfiguration {
	for i := range values {
		b.Signals = append(b.Signals, values[i])
	}
	return b
}

// WithCollectorRoles adds the given value to the CollectorRoles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CollectorRoles field.
func (b *ProcessorSpecApplyConfiguration) WithCollectorRoles(values ...v1alpha1.CollectorsGroupRole) *ProcessorSpecApplyConfiguration {
	for i := range values {
		b.CollectorRoles = append(b.CollectorRoles, values[i])
	}
	return b
}

// WithOrderHint sets the OrderHint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OrderHint field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithOrderHint(value int) *ProcessorSpecApplyConfiguration {
	b.OrderHint = &value
	return b
}

// WithProcessorConfig sets the ProcessorConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProcessorConfig field is set to the value of the last call.
func (b *ProcessorSpecApplyConfiguration) WithProcessorConfig(value runtime.RawExtension) *ProcessorSpecApplyConfiguration {
	b.ProcessorConfig = &value
	return b
}
