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

type Permissions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionsSpec   `json:"spec,omitempty"`
	Status            PermissionsStatus `json:"status,omitempty"`
}

type PermissionsSpecPermissions struct {
	// Whether user_or_group field refers to a user or a group. True for a group and false for a user.
	IsGroup *bool `json:"isGroup" tf:"is_group"`
	// Whether or not this permission propagates down the hierarchy to sub-entities.
	Propagate *bool `json:"propagate" tf:"propagate"`
	// Reference to the role providing the access.
	RoleID *string `json:"roleID" tf:"role_id"`
	// User or group receiving access.
	UserOrGroup *string `json:"userOrGroup" tf:"user_or_group"`
}

type PermissionsSpec struct {
	State *PermissionsSpecResource `json:"state,omitempty" tf:"-"`

	Resource PermissionsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PermissionsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The managed object id or uuid of the entity.
	EntityID *string `json:"entityID" tf:"entity_id"`
	// The entity managed object type.
	EntityType *string `json:"entityType" tf:"entity_type"`
	// Permissions to be given to the entity.
	// +kubebuilder:validation:MinItems=1
	Permissions []PermissionsSpecPermissions `json:"permissions" tf:"permissions"`
}

type PermissionsStatus struct {
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

// PermissionsList is a list of Permissionss
type PermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Permissions CRD objects
	Items []Permissions `json:"items,omitempty"`
}
