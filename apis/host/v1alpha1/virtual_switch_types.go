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

type VirtualSwitch struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualSwitchSpec   `json:"spec,omitempty"`
	Status            VirtualSwitchStatus `json:"status,omitempty"`
}

type VirtualSwitchSpec struct {
	State *VirtualSwitchSpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualSwitchSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VirtualSwitchSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// List of active network adapters used for load balancing.
	ActiveNics []string `json:"activeNics" tf:"active_nics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	// +optional
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	// +optional
	AllowMACChanges *bool `json:"allowMACChanges,omitempty" tf:"allow_mac_changes"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	// +optional
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous"`
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	// +optional
	BeaconInterval *int64 `json:"beaconInterval,omitempty" tf:"beacon_interval"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used only.
	// +optional
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	// +optional
	Failback *bool `json:"failback,omitempty" tf:"failback"`
	// The managed object ID of the host to set the virtual switch up on.
	HostSystemID *string `json:"hostSystemID" tf:"host_system_id"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	// +optional
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	// +optional
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol"`
	// The maximum transmission unit (MTU) of the virtual switch in bytes.
	// +optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`
	// The name of the virtual switch.
	Name *string `json:"name" tf:"name"`
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []string `json:"networkAdapters" tf:"network_adapters"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	// +optional
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches"`
	// The number of ports that this virtual switch is configured to use.
	// +optional
	NumberOfPorts *int64 `json:"numberOfPorts,omitempty" tf:"number_of_ports"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	// +optional
	ShapingAverageBandwidth *int64 `json:"shapingAverageBandwidth,omitempty" tf:"shaping_average_bandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	// +optional
	ShapingBurstSize *int64 `json:"shapingBurstSize,omitempty" tf:"shaping_burst_size"`
	// Enable traffic shaping on this virtual switch or port group.
	// +optional
	ShapingEnabled *bool `json:"shapingEnabled,omitempty" tf:"shaping_enabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	// +optional
	ShapingPeakBandwidth *int64 `json:"shapingPeakBandwidth,omitempty" tf:"shaping_peak_bandwidth"`
	// List of standby network adapters used for failover.
	StandbyNics []string `json:"standbyNics" tf:"standby_nics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or failover_explicit.
	// +optional
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy"`
}

type VirtualSwitchStatus struct {
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

// VirtualSwitchList is a list of VirtualSwitchs
type VirtualSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualSwitch CRD objects
	Items []VirtualSwitch `json:"items,omitempty"`
}
