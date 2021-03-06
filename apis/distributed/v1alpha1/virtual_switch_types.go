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

type VirtualSwitchSpecHost struct {
	// Name of the physical NIC to be added to the proxy switch.
	// +optional
	Devices []string `json:"devices,omitempty" tf:"devices"`
	// The managed object ID of the host this specification applies to.
	HostSystemID *string `json:"hostSystemID" tf:"host_system_id"`
}

type VirtualSwitchSpecPvlanMapping struct {
	// The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
	PrimaryVLANID *int64 `json:"primaryVLANID" tf:"primary_vlan_id"`
	// The private VLAN type. Valid values are promiscuous, community and isolated.
	PvlanType *string `json:"pvlanType" tf:"pvlan_type"`
	// The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
	SecondaryVLANID *int64 `json:"secondaryVLANID" tf:"secondary_vlan_id"`
}

type VirtualSwitchSpecVlanRange struct {
	// The minimum VLAN to use in the range.
	MaxVLAN *int64 `json:"maxVLAN" tf:"max_vlan"`
	// The minimum VLAN to use in the range.
	MinVLAN *int64 `json:"minVLAN" tf:"min_vlan"`
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

	// List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	// +optional
	ActiveUplinks []string `json:"activeUplinks,omitempty" tf:"active_uplinks"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.
	// +optional
	AllowForgedTransmits *bool `json:"allowForgedTransmits,omitempty" tf:"allow_forged_transmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	// +optional
	AllowMACChanges *bool `json:"allowMACChanges,omitempty" tf:"allow_mac_changes"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	// +optional
	AllowPromiscuous *bool `json:"allowPromiscuous,omitempty" tf:"allow_promiscuous"`
	// The maximum allowed usage for the backupNfc traffic class, in Mbits/sec.
	// +optional
	BackupnfcMaximumMbit *int64 `json:"backupnfcMaximumMbit,omitempty" tf:"backupnfc_maximum_mbit"`
	// The amount of guaranteed bandwidth for the backupNfc traffic class, in Mbits/sec.
	// +optional
	BackupnfcReservationMbit *int64 `json:"backupnfcReservationMbit,omitempty" tf:"backupnfc_reservation_mbit"`
	// The amount of shares to allocate to the backupNfc traffic class for a custom share level.
	// +optional
	BackupnfcShareCount *int64 `json:"backupnfcShareCount,omitempty" tf:"backupnfc_share_count"`
	// The allocation level for the backupNfc traffic class. Can be one of high, low, normal, or custom.
	// +optional
	BackupnfcShareLevel *string `json:"backupnfcShareLevel,omitempty" tf:"backupnfc_share_level"`
	// Indicates whether to block all ports by default.
	// +optional
	BlockAllPorts *bool `json:"blockAllPorts,omitempty" tf:"block_all_ports"`
	// Enable beacon probing on the ports this policy applies to.
	// +optional
	CheckBeacon *bool `json:"checkBeacon,omitempty" tf:"check_beacon"`
	// The version string of the configuration that this spec is trying to change.
	// +optional
	ConfigVersion *string `json:"configVersion,omitempty" tf:"config_version"`
	// The contact detail for this DVS.
	// +optional
	ContactDetail *string `json:"contactDetail,omitempty" tf:"contact_detail"`
	// The contact name for this DVS.
	// +optional
	ContactName *string `json:"contactName,omitempty" tf:"contact_name"`
	// A list of custom attributes to set on this resource.
	// +optional
	CustomAttributes *map[string]string `json:"customAttributes,omitempty" tf:"custom_attributes"`
	// The ID of the datacenter to create this virtual switch in.
	DatacenterID *string `json:"datacenterID" tf:"datacenter_id"`
	// The description of the DVS.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Allow VMDirectPath Gen2 on the ports this policy applies to.
	// +optional
	DirectpathGen2Allowed *bool `json:"directpathGen2Allowed,omitempty" tf:"directpath_gen2_allowed"`
	// The average egress bandwidth in bits per second if egress shaping is enabled on the port.
	// +optional
	EgressShapingAverageBandwidth *int64 `json:"egressShapingAverageBandwidth,omitempty" tf:"egress_shaping_average_bandwidth"`
	// The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
	// +optional
	EgressShapingBurstSize *int64 `json:"egressShapingBurstSize,omitempty" tf:"egress_shaping_burst_size"`
	// True if the traffic shaper is enabled for egress traffic on the port.
	// +optional
	EgressShapingEnabled *bool `json:"egressShapingEnabled,omitempty" tf:"egress_shaping_enabled"`
	// The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
	// +optional
	EgressShapingPeakBandwidth *int64 `json:"egressShapingPeakBandwidth,omitempty" tf:"egress_shaping_peak_bandwidth"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	// +optional
	Failback *bool `json:"failback,omitempty" tf:"failback"`
	// The maximum allowed usage for the faultTolerance traffic class, in Mbits/sec.
	// +optional
	FaulttoleranceMaximumMbit *int64 `json:"faulttoleranceMaximumMbit,omitempty" tf:"faulttolerance_maximum_mbit"`
	// The amount of guaranteed bandwidth for the faultTolerance traffic class, in Mbits/sec.
	// +optional
	FaulttoleranceReservationMbit *int64 `json:"faulttoleranceReservationMbit,omitempty" tf:"faulttolerance_reservation_mbit"`
	// The amount of shares to allocate to the faultTolerance traffic class for a custom share level.
	// +optional
	FaulttoleranceShareCount *int64 `json:"faulttoleranceShareCount,omitempty" tf:"faulttolerance_share_count"`
	// The allocation level for the faultTolerance traffic class. Can be one of high, low, normal, or custom.
	// +optional
	FaulttoleranceShareLevel *string `json:"faulttoleranceShareLevel,omitempty" tf:"faulttolerance_share_level"`
	// The folder to create this virtual switch in, relative to the datacenter.
	// +optional
	Folder *string `json:"folder,omitempty" tf:"folder"`
	// The maximum allowed usage for the hbr traffic class, in Mbits/sec.
	// +optional
	HbrMaximumMbit *int64 `json:"hbrMaximumMbit,omitempty" tf:"hbr_maximum_mbit"`
	// The amount of guaranteed bandwidth for the hbr traffic class, in Mbits/sec.
	// +optional
	HbrReservationMbit *int64 `json:"hbrReservationMbit,omitempty" tf:"hbr_reservation_mbit"`
	// The amount of shares to allocate to the hbr traffic class for a custom share level.
	// +optional
	HbrShareCount *int64 `json:"hbrShareCount,omitempty" tf:"hbr_share_count"`
	// The allocation level for the hbr traffic class. Can be one of high, low, normal, or custom.
	// +optional
	HbrShareLevel *string `json:"hbrShareLevel,omitempty" tf:"hbr_share_level"`
	// A host member specification.
	// +optional
	Host []VirtualSwitchSpecHost `json:"host,omitempty" tf:"host"`
	// Whether to ignore existing PVLAN mappings not managed by this resource. Defaults to false.
	// +optional
	IgnoreOtherPvlanMappings *bool `json:"ignoreOtherPvlanMappings,omitempty" tf:"ignore_other_pvlan_mappings"`
	// The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
	// +optional
	IngressShapingAverageBandwidth *int64 `json:"ingressShapingAverageBandwidth,omitempty" tf:"ingress_shaping_average_bandwidth"`
	// The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
	// +optional
	IngressShapingBurstSize *int64 `json:"ingressShapingBurstSize,omitempty" tf:"ingress_shaping_burst_size"`
	// True if the traffic shaper is enabled for ingress traffic on the port.
	// +optional
	IngressShapingEnabled *bool `json:"ingressShapingEnabled,omitempty" tf:"ingress_shaping_enabled"`
	// The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
	// +optional
	IngressShapingPeakBandwidth *int64 `json:"ingressShapingPeakBandwidth,omitempty" tf:"ingress_shaping_peak_bandwidth"`
	// The IPv4 address of the switch. This can be used to see the DVS as a unique device with NetFlow.
	// +optional
	Ipv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address"`
	// The maximum allowed usage for the iSCSI traffic class, in Mbits/sec.
	// +optional
	IscsiMaximumMbit *int64 `json:"iscsiMaximumMbit,omitempty" tf:"iscsi_maximum_mbit"`
	// The amount of guaranteed bandwidth for the iSCSI traffic class, in Mbits/sec.
	// +optional
	IscsiReservationMbit *int64 `json:"iscsiReservationMbit,omitempty" tf:"iscsi_reservation_mbit"`
	// The amount of shares to allocate to the iSCSI traffic class for a custom share level.
	// +optional
	IscsiShareCount *int64 `json:"iscsiShareCount,omitempty" tf:"iscsi_share_count"`
	// The allocation level for the iSCSI traffic class. Can be one of high, low, normal, or custom.
	// +optional
	IscsiShareLevel *string `json:"iscsiShareLevel,omitempty" tf:"iscsi_share_level"`
	// The Link Aggregation Control Protocol group version in the switch. Can be one of singleLag or multipleLag.
	// +optional
	LacpAPIVersion *string `json:"lacpAPIVersion,omitempty" tf:"lacp_api_version"`
	// Whether or not to enable LACP on all uplink ports.
	// +optional
	LacpEnabled *bool `json:"lacpEnabled,omitempty" tf:"lacp_enabled"`
	// The uplink LACP mode to use. Can be one of active or passive.
	// +optional
	LacpMode *string `json:"lacpMode,omitempty" tf:"lacp_mode"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	// +optional
	LinkDiscoveryOperation *string `json:"linkDiscoveryOperation,omitempty" tf:"link_discovery_operation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	// +optional
	LinkDiscoveryProtocol *string `json:"linkDiscoveryProtocol,omitempty" tf:"link_discovery_protocol"`
	// The maximum allowed usage for the management traffic class, in Mbits/sec.
	// +optional
	ManagementMaximumMbit *int64 `json:"managementMaximumMbit,omitempty" tf:"management_maximum_mbit"`
	// The amount of guaranteed bandwidth for the management traffic class, in Mbits/sec.
	// +optional
	ManagementReservationMbit *int64 `json:"managementReservationMbit,omitempty" tf:"management_reservation_mbit"`
	// The amount of shares to allocate to the management traffic class for a custom share level.
	// +optional
	ManagementShareCount *int64 `json:"managementShareCount,omitempty" tf:"management_share_count"`
	// The allocation level for the management traffic class. Can be one of high, low, normal, or custom.
	// +optional
	ManagementShareLevel *string `json:"managementShareLevel,omitempty" tf:"management_share_level"`
	// The maximum MTU on the switch.
	// +optional
	MaxMTU *int64 `json:"maxMTU,omitempty" tf:"max_mtu"`
	// The multicast filtering mode on the switch. Can be one of legacyFiltering, or snooping.
	// +optional
	MulticastFilteringMode *string `json:"multicastFilteringMode,omitempty" tf:"multicast_filtering_mode"`
	// The name for the DVS. Must be unique in the folder that it is being created in.
	Name *string `json:"name" tf:"name"`
	// The number of seconds after which active flows are forced to be exported to the collector.
	// +optional
	NetflowActiveFlowTimeout *int64 `json:"netflowActiveFlowTimeout,omitempty" tf:"netflow_active_flow_timeout"`
	// IP address for the netflow collector, using IPv4 or IPv6. IPv6 is supported in vSphere Distributed Switch Version 6.0 or later.
	// +optional
	NetflowCollectorIPAddress *string `json:"netflowCollectorIPAddress,omitempty" tf:"netflow_collector_ip_address"`
	// The port for the netflow collector.
	// +optional
	NetflowCollectorPort *int64 `json:"netflowCollectorPort,omitempty" tf:"netflow_collector_port"`
	// Indicates whether to enable netflow on all ports.
	// +optional
	NetflowEnabled *bool `json:"netflowEnabled,omitempty" tf:"netflow_enabled"`
	// The number of seconds after which idle flows are forced to be exported to the collector.
	// +optional
	NetflowIdleFlowTimeout *int64 `json:"netflowIdleFlowTimeout,omitempty" tf:"netflow_idle_flow_timeout"`
	// Whether to limit analysis to traffic that has both source and destination served by the same host.
	// +optional
	NetflowInternalFlowsOnly *bool `json:"netflowInternalFlowsOnly,omitempty" tf:"netflow_internal_flows_only"`
	// The observation Domain ID for the netflow collector.
	// +optional
	NetflowObservationDomainID *int64 `json:"netflowObservationDomainID,omitempty" tf:"netflow_observation_domain_id"`
	// The ratio of total number of packets to the number of packets analyzed. Set to 0 to disable sampling, meaning that all packets are analyzed.
	// +optional
	NetflowSamplingRate *int64 `json:"netflowSamplingRate,omitempty" tf:"netflow_sampling_rate"`
	// Whether or not to enable network resource control, enabling advanced traffic shaping and resource control features.
	// +optional
	NetworkResourceControlEnabled *bool `json:"networkResourceControlEnabled,omitempty" tf:"network_resource_control_enabled"`
	// The network I/O control version to use. Can be one of version2 or version3.
	// +optional
	NetworkResourceControlVersion *string `json:"networkResourceControlVersion,omitempty" tf:"network_resource_control_version"`
	// The maximum allowed usage for the nfs traffic class, in Mbits/sec.
	// +optional
	NfsMaximumMbit *int64 `json:"nfsMaximumMbit,omitempty" tf:"nfs_maximum_mbit"`
	// The amount of guaranteed bandwidth for the nfs traffic class, in Mbits/sec.
	// +optional
	NfsReservationMbit *int64 `json:"nfsReservationMbit,omitempty" tf:"nfs_reservation_mbit"`
	// The amount of shares to allocate to the nfs traffic class for a custom share level.
	// +optional
	NfsShareCount *int64 `json:"nfsShareCount,omitempty" tf:"nfs_share_count"`
	// The allocation level for the nfs traffic class. Can be one of high, low, normal, or custom.
	// +optional
	NfsShareLevel *string `json:"nfsShareLevel,omitempty" tf:"nfs_share_level"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	// +optional
	NotifySwitches *bool `json:"notifySwitches,omitempty" tf:"notify_switches"`
	// The secondary VLAN ID for this port.
	// +optional
	PortPrivateSecondaryVLANID *int64 `json:"portPrivateSecondaryVLANID,omitempty" tf:"port_private_secondary_vlan_id"`
	// A private VLAN (PVLAN) mapping.
	// +optional
	PvlanMapping []VirtualSwitchSpecPvlanMapping `json:"pvlanMapping,omitempty" tf:"pvlan_mapping"`
	// List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
	// +optional
	StandbyUplinks []string `json:"standbyUplinks,omitempty" tf:"standby_uplinks"`
	// A list of tag IDs to apply to this object.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, failover_explicit, or loadbalance_loadbased.
	// +optional
	TeamingPolicy *string `json:"teamingPolicy,omitempty" tf:"teaming_policy"`
	// If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet forwarded done by the switch.
	// +optional
	TxUplink *bool `json:"txUplink,omitempty" tf:"tx_uplink"`
	// A list of uplink ports. The contents of this list control both the uplink count and names of the uplinks on the DVS across hosts.
	// +optional
	Uplinks []string `json:"uplinks,omitempty" tf:"uplinks"`
	// The maximum allowed usage for the vdp traffic class, in Mbits/sec.
	// +optional
	VdpMaximumMbit *int64 `json:"vdpMaximumMbit,omitempty" tf:"vdp_maximum_mbit"`
	// The amount of guaranteed bandwidth for the vdp traffic class, in Mbits/sec.
	// +optional
	VdpReservationMbit *int64 `json:"vdpReservationMbit,omitempty" tf:"vdp_reservation_mbit"`
	// The amount of shares to allocate to the vdp traffic class for a custom share level.
	// +optional
	VdpShareCount *int64 `json:"vdpShareCount,omitempty" tf:"vdp_share_count"`
	// The allocation level for the vdp traffic class. Can be one of high, low, normal, or custom.
	// +optional
	VdpShareLevel *string `json:"vdpShareLevel,omitempty" tf:"vdp_share_level"`
	// The version of this virtual switch. Allowed versions are 7.0.3, 7.0.0, 6.6.0, 6.5.0, 6.0.0, 5.5.0, 5.1.0, and 5.0.0.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// The maximum allowed usage for the virtualMachine traffic class, in Mbits/sec.
	// +optional
	VirtualmachineMaximumMbit *int64 `json:"virtualmachineMaximumMbit,omitempty" tf:"virtualmachine_maximum_mbit"`
	// The amount of guaranteed bandwidth for the virtualMachine traffic class, in Mbits/sec.
	// +optional
	VirtualmachineReservationMbit *int64 `json:"virtualmachineReservationMbit,omitempty" tf:"virtualmachine_reservation_mbit"`
	// The amount of shares to allocate to the virtualMachine traffic class for a custom share level.
	// +optional
	VirtualmachineShareCount *int64 `json:"virtualmachineShareCount,omitempty" tf:"virtualmachine_share_count"`
	// The allocation level for the virtualMachine traffic class. Can be one of high, low, normal, or custom.
	// +optional
	VirtualmachineShareLevel *string `json:"virtualmachineShareLevel,omitempty" tf:"virtualmachine_share_level"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	// +optional
	VlanID *int64 `json:"vlanID,omitempty" tf:"vlan_id"`
	// The VLAN ID for single VLAN mode. 0 denotes no VLAN.
	// +optional
	VlanRange []VirtualSwitchSpecVlanRange `json:"vlanRange,omitempty" tf:"vlan_range"`
	// The maximum allowed usage for the vmotion traffic class, in Mbits/sec.
	// +optional
	VmotionMaximumMbit *int64 `json:"vmotionMaximumMbit,omitempty" tf:"vmotion_maximum_mbit"`
	// The amount of guaranteed bandwidth for the vmotion traffic class, in Mbits/sec.
	// +optional
	VmotionReservationMbit *int64 `json:"vmotionReservationMbit,omitempty" tf:"vmotion_reservation_mbit"`
	// The amount of shares to allocate to the vmotion traffic class for a custom share level.
	// +optional
	VmotionShareCount *int64 `json:"vmotionShareCount,omitempty" tf:"vmotion_share_count"`
	// The allocation level for the vmotion traffic class. Can be one of high, low, normal, or custom.
	// +optional
	VmotionShareLevel *string `json:"vmotionShareLevel,omitempty" tf:"vmotion_share_level"`
	// The maximum allowed usage for the vsan traffic class, in Mbits/sec.
	// +optional
	VsanMaximumMbit *int64 `json:"vsanMaximumMbit,omitempty" tf:"vsan_maximum_mbit"`
	// The amount of guaranteed bandwidth for the vsan traffic class, in Mbits/sec.
	// +optional
	VsanReservationMbit *int64 `json:"vsanReservationMbit,omitempty" tf:"vsan_reservation_mbit"`
	// The amount of shares to allocate to the vsan traffic class for a custom share level.
	// +optional
	VsanShareCount *int64 `json:"vsanShareCount,omitempty" tf:"vsan_share_count"`
	// The allocation level for the vsan traffic class. Can be one of high, low, normal, or custom.
	// +optional
	VsanShareLevel *string `json:"vsanShareLevel,omitempty" tf:"vsan_share_level"`
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
