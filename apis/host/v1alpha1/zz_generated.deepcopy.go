//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Host) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostList) DeepCopyInto(out *HostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Host, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostList.
func (in *HostList) DeepCopy() *HostList {
	if in == nil {
		return nil
	}
	out := new(HostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpec) DeepCopyInto(out *HostSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(HostSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpec.
func (in *HostSpec) DeepCopy() *HostSpec {
	if in == nil {
		return nil
	}
	out := new(HostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpecResource) DeepCopyInto(out *HostSpecResource) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.ClusterManaged != nil {
		in, out := &in.ClusterManaged, &out.ClusterManaged
		*out = new(bool)
		**out = **in
	}
	if in.Connected != nil {
		in, out := &in.Connected, &out.Connected
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(string)
		**out = **in
	}
	if in.Lockdown != nil {
		in, out := &in.Lockdown, &out.Lockdown
		*out = new(string)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpecResource.
func (in *HostSpecResource) DeepCopy() *HostSpecResource {
	if in == nil {
		return nil
	}
	out := new(HostSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroup) DeepCopyInto(out *PortGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroup.
func (in *PortGroup) DeepCopy() *PortGroup {
	if in == nil {
		return nil
	}
	out := new(PortGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroupList) DeepCopyInto(out *PortGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PortGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroupList.
func (in *PortGroupList) DeepCopy() *PortGroupList {
	if in == nil {
		return nil
	}
	out := new(PortGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroupSpec) DeepCopyInto(out *PortGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PortGroupSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroupSpec.
func (in *PortGroupSpec) DeepCopy() *PortGroupSpec {
	if in == nil {
		return nil
	}
	out := new(PortGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroupSpecPorts) DeepCopyInto(out *PortGroupSpecPorts) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.MacAddresses != nil {
		in, out := &in.MacAddresses, &out.MacAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroupSpecPorts.
func (in *PortGroupSpecPorts) DeepCopy() *PortGroupSpecPorts {
	if in == nil {
		return nil
	}
	out := new(PortGroupSpecPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroupSpecResource) DeepCopyInto(out *PortGroupSpecResource) {
	*out = *in
	if in.ActiveNics != nil {
		in, out := &in.ActiveNics, &out.ActiveNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowForgedTransmits != nil {
		in, out := &in.AllowForgedTransmits, &out.AllowForgedTransmits
		*out = new(bool)
		**out = **in
	}
	if in.AllowMACChanges != nil {
		in, out := &in.AllowMACChanges, &out.AllowMACChanges
		*out = new(bool)
		**out = **in
	}
	if in.AllowPromiscuous != nil {
		in, out := &in.AllowPromiscuous, &out.AllowPromiscuous
		*out = new(bool)
		**out = **in
	}
	if in.CheckBeacon != nil {
		in, out := &in.CheckBeacon, &out.CheckBeacon
		*out = new(bool)
		**out = **in
	}
	if in.ComputedPolicy != nil {
		in, out := &in.ComputedPolicy, &out.ComputedPolicy
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Failback != nil {
		in, out := &in.Failback, &out.Failback
		*out = new(bool)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotifySwitches != nil {
		in, out := &in.NotifySwitches, &out.NotifySwitches
		*out = new(bool)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortGroupSpecPorts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShapingAverageBandwidth != nil {
		in, out := &in.ShapingAverageBandwidth, &out.ShapingAverageBandwidth
		*out = new(int64)
		**out = **in
	}
	if in.ShapingBurstSize != nil {
		in, out := &in.ShapingBurstSize, &out.ShapingBurstSize
		*out = new(int64)
		**out = **in
	}
	if in.ShapingEnabled != nil {
		in, out := &in.ShapingEnabled, &out.ShapingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ShapingPeakBandwidth != nil {
		in, out := &in.ShapingPeakBandwidth, &out.ShapingPeakBandwidth
		*out = new(int64)
		**out = **in
	}
	if in.StandbyNics != nil {
		in, out := &in.StandbyNics, &out.StandbyNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TeamingPolicy != nil {
		in, out := &in.TeamingPolicy, &out.TeamingPolicy
		*out = new(string)
		**out = **in
	}
	if in.VirtualSwitchName != nil {
		in, out := &in.VirtualSwitchName, &out.VirtualSwitchName
		*out = new(string)
		**out = **in
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroupSpecResource.
func (in *PortGroupSpecResource) DeepCopy() *PortGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(PortGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortGroupStatus) DeepCopyInto(out *PortGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortGroupStatus.
func (in *PortGroupStatus) DeepCopy() *PortGroupStatus {
	if in == nil {
		return nil
	}
	out := new(PortGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitch) DeepCopyInto(out *VirtualSwitch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitch.
func (in *VirtualSwitch) DeepCopy() *VirtualSwitch {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualSwitch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchList) DeepCopyInto(out *VirtualSwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualSwitch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchList.
func (in *VirtualSwitchList) DeepCopy() *VirtualSwitchList {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualSwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchSpec) DeepCopyInto(out *VirtualSwitchSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VirtualSwitchSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchSpec.
func (in *VirtualSwitchSpec) DeepCopy() *VirtualSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchSpecResource) DeepCopyInto(out *VirtualSwitchSpecResource) {
	*out = *in
	if in.ActiveNics != nil {
		in, out := &in.ActiveNics, &out.ActiveNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowForgedTransmits != nil {
		in, out := &in.AllowForgedTransmits, &out.AllowForgedTransmits
		*out = new(bool)
		**out = **in
	}
	if in.AllowMACChanges != nil {
		in, out := &in.AllowMACChanges, &out.AllowMACChanges
		*out = new(bool)
		**out = **in
	}
	if in.AllowPromiscuous != nil {
		in, out := &in.AllowPromiscuous, &out.AllowPromiscuous
		*out = new(bool)
		**out = **in
	}
	if in.BeaconInterval != nil {
		in, out := &in.BeaconInterval, &out.BeaconInterval
		*out = new(int64)
		**out = **in
	}
	if in.CheckBeacon != nil {
		in, out := &in.CheckBeacon, &out.CheckBeacon
		*out = new(bool)
		**out = **in
	}
	if in.Failback != nil {
		in, out := &in.Failback, &out.Failback
		*out = new(bool)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
	if in.LinkDiscoveryOperation != nil {
		in, out := &in.LinkDiscoveryOperation, &out.LinkDiscoveryOperation
		*out = new(string)
		**out = **in
	}
	if in.LinkDiscoveryProtocol != nil {
		in, out := &in.LinkDiscoveryProtocol, &out.LinkDiscoveryProtocol
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkAdapters != nil {
		in, out := &in.NetworkAdapters, &out.NetworkAdapters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotifySwitches != nil {
		in, out := &in.NotifySwitches, &out.NotifySwitches
		*out = new(bool)
		**out = **in
	}
	if in.NumberOfPorts != nil {
		in, out := &in.NumberOfPorts, &out.NumberOfPorts
		*out = new(int64)
		**out = **in
	}
	if in.ShapingAverageBandwidth != nil {
		in, out := &in.ShapingAverageBandwidth, &out.ShapingAverageBandwidth
		*out = new(int64)
		**out = **in
	}
	if in.ShapingBurstSize != nil {
		in, out := &in.ShapingBurstSize, &out.ShapingBurstSize
		*out = new(int64)
		**out = **in
	}
	if in.ShapingEnabled != nil {
		in, out := &in.ShapingEnabled, &out.ShapingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ShapingPeakBandwidth != nil {
		in, out := &in.ShapingPeakBandwidth, &out.ShapingPeakBandwidth
		*out = new(int64)
		**out = **in
	}
	if in.StandbyNics != nil {
		in, out := &in.StandbyNics, &out.StandbyNics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TeamingPolicy != nil {
		in, out := &in.TeamingPolicy, &out.TeamingPolicy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchSpecResource.
func (in *VirtualSwitchSpecResource) DeepCopy() *VirtualSwitchSpecResource {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchStatus) DeepCopyInto(out *VirtualSwitchStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchStatus.
func (in *VirtualSwitchStatus) DeepCopy() *VirtualSwitchStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchStatus)
	in.DeepCopyInto(out)
	return out
}
