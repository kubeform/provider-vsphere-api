/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ClusterVmAffinityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterVmAffinityRuleSpec   `json:"spec,omitempty"`
	Status            ClusterVmAffinityRuleStatus `json:"status,omitempty"`
}

type ClusterVmAffinityRuleSpec struct {
	State *ClusterVmAffinityRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterVmAffinityRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterVmAffinityRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The managed object ID of the cluster.
	ComputeClusterID *string `json:"computeClusterID" tf:"compute_cluster_id"`
	// Enable this rule in the cluster.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// When true, prevents any virtual machine operations that may violate this rule.
	// +optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory"`
	// The unique name of the virtual machine group in the cluster.
	Name *string `json:"name" tf:"name"`
	// The UUIDs of the virtual machines to run on the same host together.
	VirtualMachineIDS []string `json:"virtualMachineIDS" tf:"virtual_machine_ids"`
}

type ClusterVmAffinityRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClusterVmAffinityRuleList is a list of ClusterVmAffinityRules
type ClusterVmAffinityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterVmAffinityRule CRD objects
	Items []ClusterVmAffinityRule `json:"items,omitempty"`
}
