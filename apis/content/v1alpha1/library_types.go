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

type Library struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LibrarySpec   `json:"spec,omitempty"`
	Status            LibraryStatus `json:"status,omitempty"`
}

type LibrarySpecPublication struct {
	// +optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method"`
	// +optional
	Password *string `json:"password,omitempty" tf:"password"`
	// +optional
	PublishURL *string `json:"publishURL,omitempty" tf:"publish_url"`
	// +optional
	Published *bool `json:"published,omitempty" tf:"published"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type LibrarySpecSubscription struct {
	// +optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method"`
	// +optional
	AutomaticSync *bool `json:"automaticSync,omitempty" tf:"automatic_sync"`
	// +optional
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand"`
	// +optional
	Password *string `json:"password,omitempty" tf:"password"`
	// +optional
	SubscriptionURL *string `json:"subscriptionURL,omitempty" tf:"subscription_url"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type LibrarySpec struct {
	State *LibrarySpecResource `json:"state,omitempty" tf:"-"`

	Resource LibrarySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LibrarySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional description of the content library.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The name of the content library.
	Name *string `json:"name" tf:"name"`
	// Publication configuration for content library.
	// +optional
	Publication *LibrarySpecPublication `json:"publication,omitempty" tf:"publication"`
	// The name of the content library.
	StorageBacking []string `json:"storageBacking" tf:"storage_backing"`
	// Publication configuration for content library.
	// +optional
	Subscription *LibrarySpecSubscription `json:"subscription,omitempty" tf:"subscription"`
}

type LibraryStatus struct {
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

// LibraryList is a list of Librarys
type LibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Library CRD objects
	Items []Library `json:"items,omitempty"`
}
