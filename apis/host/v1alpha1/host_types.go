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

type Host struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec,omitempty"`
	Status            HostStatus `json:"status,omitempty"`
}

type HostSpec struct {
	State *HostSpecResource `json:"state,omitempty" tf:"-"`

	Resource HostSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type HostSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the vSphere cluster the host will belong to.
	// +optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster"`
	// Must be set if host is a member of a managed compute_cluster resource.
	// +optional
	ClusterManaged *bool `json:"clusterManaged,omitempty" tf:"cluster_managed"`
	// Set the state of the host. If set to false then the host will be asked to disconnect.
	// +optional
	Connected *bool `json:"connected,omitempty" tf:"connected"`
	// ID of the vSphere datacenter the host will belong to.
	// +optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter"`
	// Force add the host to vsphere, even if it's already managed by a different vSphere instance.
	// +optional
	Force *bool `json:"force,omitempty" tf:"force"`
	// FQDN or IP address of the host.
	Hostname *string `json:"hostname" tf:"hostname"`
	// License key that will be applied to this host.
	// +optional
	License *string `json:"license,omitempty" tf:"license"`
	// Set the host's lockdown status. Default is disabled. Valid options are 'disabled', 'normal', 'strict'
	// +optional
	Lockdown *string `json:"lockdown,omitempty" tf:"lockdown"`
	// Set the host's maintenance mode. Default is false
	// +optional
	Maintenance *bool `json:"maintenance,omitempty" tf:"maintenance"`
	// Password of the administration account of the host.
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// Host's certificate SHA-1 thumbprint. If not set then the CA that signed the host's certificate must be trusted.
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
	// Username of the administration account of the host.
	Username *string `json:"username" tf:"username"`
}

type HostStatus struct {
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

// HostList is a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Host CRD objects
	Items []Host `json:"items,omitempty"`
}