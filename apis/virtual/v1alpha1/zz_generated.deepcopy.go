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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Disk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskList) DeepCopyInto(out *DiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskList.
func (in *DiskList) DeepCopy() *DiskList {
	if in == nil {
		return nil
	}
	out := new(DiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DiskSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpecResource) DeepCopyInto(out *DiskSpecResource) {
	*out = *in
	if in.AdapterType != nil {
		in, out := &in.AdapterType, &out.AdapterType
		*out = new(string)
		**out = **in
	}
	if in.CreateDirectories != nil {
		in, out := &in.CreateDirectories, &out.CreateDirectories
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VmdkPath != nil {
		in, out := &in.VmdkPath, &out.VmdkPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpecResource.
func (in *DiskSpecResource) DeepCopy() *DiskSpecResource {
	if in == nil {
		return nil
	}
	out := new(DiskSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskStatus) DeepCopyInto(out *DiskStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskStatus.
func (in *DiskStatus) DeepCopy() *DiskStatus {
	if in == nil {
		return nil
	}
	out := new(DiskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSnapshot) DeepCopyInto(out *MachineSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSnapshot.
func (in *MachineSnapshot) DeepCopy() *MachineSnapshot {
	if in == nil {
		return nil
	}
	out := new(MachineSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSnapshotList) DeepCopyInto(out *MachineSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSnapshotList.
func (in *MachineSnapshotList) DeepCopy() *MachineSnapshotList {
	if in == nil {
		return nil
	}
	out := new(MachineSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSnapshotSpec) DeepCopyInto(out *MachineSnapshotSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MachineSnapshotSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSnapshotSpec.
func (in *MachineSnapshotSpec) DeepCopy() *MachineSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSnapshotSpecResource) DeepCopyInto(out *MachineSnapshotSpecResource) {
	*out = *in
	if in.Consolidate != nil {
		in, out := &in.Consolidate, &out.Consolidate
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(bool)
		**out = **in
	}
	if in.Quiesce != nil {
		in, out := &in.Quiesce, &out.Quiesce
		*out = new(bool)
		**out = **in
	}
	if in.RemoveChildren != nil {
		in, out := &in.RemoveChildren, &out.RemoveChildren
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineUUID != nil {
		in, out := &in.VirtualMachineUUID, &out.VirtualMachineUUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSnapshotSpecResource.
func (in *MachineSnapshotSpecResource) DeepCopy() *MachineSnapshotSpecResource {
	if in == nil {
		return nil
	}
	out := new(MachineSnapshotSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSnapshotStatus) DeepCopyInto(out *MachineSnapshotStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSnapshotStatus.
func (in *MachineSnapshotStatus) DeepCopy() *MachineSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(MachineSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MachineSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecCdrom) DeepCopyInto(out *MachineSpecCdrom) {
	*out = *in
	if in.ClientDevice != nil {
		in, out := &in.ClientDevice, &out.ClientDevice
		*out = new(bool)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(int64)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecCdrom.
func (in *MachineSpecCdrom) DeepCopy() *MachineSpecCdrom {
	if in == nil {
		return nil
	}
	out := new(MachineSpecCdrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecClone) DeepCopyInto(out *MachineSpecClone) {
	*out = *in
	if in.Customize != nil {
		in, out := &in.Customize, &out.Customize
		*out = new(MachineSpecCloneCustomize)
		(*in).DeepCopyInto(*out)
	}
	if in.LinkedClone != nil {
		in, out := &in.LinkedClone, &out.LinkedClone
		*out = new(bool)
		**out = **in
	}
	if in.OvfNetworkMap != nil {
		in, out := &in.OvfNetworkMap, &out.OvfNetworkMap
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.OvfStorageMap != nil {
		in, out := &in.OvfStorageMap, &out.OvfStorageMap
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TemplateUUID != nil {
		in, out := &in.TemplateUUID, &out.TemplateUUID
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecClone.
func (in *MachineSpecClone) DeepCopy() *MachineSpecClone {
	if in == nil {
		return nil
	}
	out := new(MachineSpecClone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecCloneCustomize) DeepCopyInto(out *MachineSpecCloneCustomize) {
	*out = *in
	if in.DnsServerList != nil {
		in, out := &in.DnsServerList, &out.DnsServerList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DnsSuffixList != nil {
		in, out := &in.DnsSuffixList, &out.DnsSuffixList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ipv4Gateway != nil {
		in, out := &in.Ipv4Gateway, &out.Ipv4Gateway
		*out = new(string)
		**out = **in
	}
	if in.Ipv6Gateway != nil {
		in, out := &in.Ipv6Gateway, &out.Ipv6Gateway
		*out = new(string)
		**out = **in
	}
	if in.LinuxOptions != nil {
		in, out := &in.LinuxOptions, &out.LinuxOptions
		*out = new(MachineSpecCloneCustomizeLinuxOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]MachineSpecCloneCustomizeNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.WindowsOptions != nil {
		in, out := &in.WindowsOptions, &out.WindowsOptions
		*out = new(MachineSpecCloneCustomizeWindowsOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.WindowsSysprepText != nil {
		in, out := &in.WindowsSysprepText, &out.WindowsSysprepText
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecCloneCustomize.
func (in *MachineSpecCloneCustomize) DeepCopy() *MachineSpecCloneCustomize {
	if in == nil {
		return nil
	}
	out := new(MachineSpecCloneCustomize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecCloneCustomizeLinuxOptions) DeepCopyInto(out *MachineSpecCloneCustomizeLinuxOptions) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.HwClockUtc != nil {
		in, out := &in.HwClockUtc, &out.HwClockUtc
		*out = new(bool)
		**out = **in
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecCloneCustomizeLinuxOptions.
func (in *MachineSpecCloneCustomizeLinuxOptions) DeepCopy() *MachineSpecCloneCustomizeLinuxOptions {
	if in == nil {
		return nil
	}
	out := new(MachineSpecCloneCustomizeLinuxOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecCloneCustomizeNetworkInterface) DeepCopyInto(out *MachineSpecCloneCustomizeNetworkInterface) {
	*out = *in
	if in.DnsDomain != nil {
		in, out := &in.DnsDomain, &out.DnsDomain
		*out = new(string)
		**out = **in
	}
	if in.DnsServerList != nil {
		in, out := &in.DnsServerList, &out.DnsServerList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ipv4Address != nil {
		in, out := &in.Ipv4Address, &out.Ipv4Address
		*out = new(string)
		**out = **in
	}
	if in.Ipv4Netmask != nil {
		in, out := &in.Ipv4Netmask, &out.Ipv4Netmask
		*out = new(int64)
		**out = **in
	}
	if in.Ipv6Address != nil {
		in, out := &in.Ipv6Address, &out.Ipv6Address
		*out = new(string)
		**out = **in
	}
	if in.Ipv6Netmask != nil {
		in, out := &in.Ipv6Netmask, &out.Ipv6Netmask
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecCloneCustomizeNetworkInterface.
func (in *MachineSpecCloneCustomizeNetworkInterface) DeepCopy() *MachineSpecCloneCustomizeNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(MachineSpecCloneCustomizeNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecCloneCustomizeWindowsOptions) DeepCopyInto(out *MachineSpecCloneCustomizeWindowsOptions) {
	*out = *in
	if in.AdminPassword != nil {
		in, out := &in.AdminPassword, &out.AdminPassword
		*out = new(string)
		**out = **in
	}
	if in.AutoLogon != nil {
		in, out := &in.AutoLogon, &out.AutoLogon
		*out = new(bool)
		**out = **in
	}
	if in.AutoLogonCount != nil {
		in, out := &in.AutoLogonCount, &out.AutoLogonCount
		*out = new(int64)
		**out = **in
	}
	if in.ComputerName != nil {
		in, out := &in.ComputerName, &out.ComputerName
		*out = new(string)
		**out = **in
	}
	if in.DomainAdminPassword != nil {
		in, out := &in.DomainAdminPassword, &out.DomainAdminPassword
		*out = new(string)
		**out = **in
	}
	if in.DomainAdminUser != nil {
		in, out := &in.DomainAdminUser, &out.DomainAdminUser
		*out = new(string)
		**out = **in
	}
	if in.FullName != nil {
		in, out := &in.FullName, &out.FullName
		*out = new(string)
		**out = **in
	}
	if in.JoinDomain != nil {
		in, out := &in.JoinDomain, &out.JoinDomain
		*out = new(string)
		**out = **in
	}
	if in.OrganizationName != nil {
		in, out := &in.OrganizationName, &out.OrganizationName
		*out = new(string)
		**out = **in
	}
	if in.ProductKey != nil {
		in, out := &in.ProductKey, &out.ProductKey
		*out = new(string)
		**out = **in
	}
	if in.RunOnceCommandList != nil {
		in, out := &in.RunOnceCommandList, &out.RunOnceCommandList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(int64)
		**out = **in
	}
	if in.Workgroup != nil {
		in, out := &in.Workgroup, &out.Workgroup
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecCloneCustomizeWindowsOptions.
func (in *MachineSpecCloneCustomizeWindowsOptions) DeepCopy() *MachineSpecCloneCustomizeWindowsOptions {
	if in == nil {
		return nil
	}
	out := new(MachineSpecCloneCustomizeWindowsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecDisk) DeepCopyInto(out *MachineSpecDisk) {
	*out = *in
	if in.Attach != nil {
		in, out := &in.Attach, &out.Attach
		*out = new(bool)
		**out = **in
	}
	if in.ControllerType != nil {
		in, out := &in.ControllerType, &out.ControllerType
		*out = new(string)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.DiskMode != nil {
		in, out := &in.DiskMode, &out.DiskMode
		*out = new(string)
		**out = **in
	}
	if in.DiskSharing != nil {
		in, out := &in.DiskSharing, &out.DiskSharing
		*out = new(string)
		**out = **in
	}
	if in.EagerlyScrub != nil {
		in, out := &in.EagerlyScrub, &out.EagerlyScrub
		*out = new(bool)
		**out = **in
	}
	if in.IoLimit != nil {
		in, out := &in.IoLimit, &out.IoLimit
		*out = new(int64)
		**out = **in
	}
	if in.IoReservation != nil {
		in, out := &in.IoReservation, &out.IoReservation
		*out = new(int64)
		**out = **in
	}
	if in.IoShareCount != nil {
		in, out := &in.IoShareCount, &out.IoShareCount
		*out = new(int64)
		**out = **in
	}
	if in.IoShareLevel != nil {
		in, out := &in.IoShareLevel, &out.IoShareLevel
		*out = new(string)
		**out = **in
	}
	if in.KeepOnRemove != nil {
		in, out := &in.KeepOnRemove, &out.KeepOnRemove
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(int64)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.StoragePolicyID != nil {
		in, out := &in.StoragePolicyID, &out.StoragePolicyID
		*out = new(string)
		**out = **in
	}
	if in.ThinProvisioned != nil {
		in, out := &in.ThinProvisioned, &out.ThinProvisioned
		*out = new(bool)
		**out = **in
	}
	if in.UnitNumber != nil {
		in, out := &in.UnitNumber, &out.UnitNumber
		*out = new(int64)
		**out = **in
	}
	if in.Uuid != nil {
		in, out := &in.Uuid, &out.Uuid
		*out = new(string)
		**out = **in
	}
	if in.WriteThrough != nil {
		in, out := &in.WriteThrough, &out.WriteThrough
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecDisk.
func (in *MachineSpecDisk) DeepCopy() *MachineSpecDisk {
	if in == nil {
		return nil
	}
	out := new(MachineSpecDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecNetworkInterface) DeepCopyInto(out *MachineSpecNetworkInterface) {
	*out = *in
	if in.AdapterType != nil {
		in, out := &in.AdapterType, &out.AdapterType
		*out = new(string)
		**out = **in
	}
	if in.BandwidthLimit != nil {
		in, out := &in.BandwidthLimit, &out.BandwidthLimit
		*out = new(int64)
		**out = **in
	}
	if in.BandwidthReservation != nil {
		in, out := &in.BandwidthReservation, &out.BandwidthReservation
		*out = new(int64)
		**out = **in
	}
	if in.BandwidthShareCount != nil {
		in, out := &in.BandwidthShareCount, &out.BandwidthShareCount
		*out = new(int64)
		**out = **in
	}
	if in.BandwidthShareLevel != nil {
		in, out := &in.BandwidthShareLevel, &out.BandwidthShareLevel
		*out = new(string)
		**out = **in
	}
	if in.DeviceAddress != nil {
		in, out := &in.DeviceAddress, &out.DeviceAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(int64)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.OvfMapping != nil {
		in, out := &in.OvfMapping, &out.OvfMapping
		*out = new(string)
		**out = **in
	}
	if in.UseStaticMAC != nil {
		in, out := &in.UseStaticMAC, &out.UseStaticMAC
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecNetworkInterface.
func (in *MachineSpecNetworkInterface) DeepCopy() *MachineSpecNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(MachineSpecNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecOvfDeploy) DeepCopyInto(out *MachineSpecOvfDeploy) {
	*out = *in
	if in.AllowUnverifiedSslCert != nil {
		in, out := &in.AllowUnverifiedSslCert, &out.AllowUnverifiedSslCert
		*out = new(bool)
		**out = **in
	}
	if in.DeploymentOption != nil {
		in, out := &in.DeploymentOption, &out.DeploymentOption
		*out = new(string)
		**out = **in
	}
	if in.DiskProvisioning != nil {
		in, out := &in.DiskProvisioning, &out.DiskProvisioning
		*out = new(string)
		**out = **in
	}
	if in.EnableHiddenProperties != nil {
		in, out := &in.EnableHiddenProperties, &out.EnableHiddenProperties
		*out = new(bool)
		**out = **in
	}
	if in.IpAllocationPolicy != nil {
		in, out := &in.IpAllocationPolicy, &out.IpAllocationPolicy
		*out = new(string)
		**out = **in
	}
	if in.IpProtocol != nil {
		in, out := &in.IpProtocol, &out.IpProtocol
		*out = new(string)
		**out = **in
	}
	if in.LocalOvfPath != nil {
		in, out := &in.LocalOvfPath, &out.LocalOvfPath
		*out = new(string)
		**out = **in
	}
	if in.OvfNetworkMap != nil {
		in, out := &in.OvfNetworkMap, &out.OvfNetworkMap
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.RemoteOvfURL != nil {
		in, out := &in.RemoteOvfURL, &out.RemoteOvfURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecOvfDeploy.
func (in *MachineSpecOvfDeploy) DeepCopy() *MachineSpecOvfDeploy {
	if in == nil {
		return nil
	}
	out := new(MachineSpecOvfDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecResource) DeepCopyInto(out *MachineSpecResource) {
	*out = *in
	if in.AlternateGuestName != nil {
		in, out := &in.AlternateGuestName, &out.AlternateGuestName
		*out = new(string)
		**out = **in
	}
	if in.Annotation != nil {
		in, out := &in.Annotation, &out.Annotation
		*out = new(string)
		**out = **in
	}
	if in.BootDelay != nil {
		in, out := &in.BootDelay, &out.BootDelay
		*out = new(int64)
		**out = **in
	}
	if in.BootRetryDelay != nil {
		in, out := &in.BootRetryDelay, &out.BootRetryDelay
		*out = new(int64)
		**out = **in
	}
	if in.BootRetryEnabled != nil {
		in, out := &in.BootRetryEnabled, &out.BootRetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Cdrom != nil {
		in, out := &in.Cdrom, &out.Cdrom
		*out = new(MachineSpecCdrom)
		(*in).DeepCopyInto(*out)
	}
	if in.ChangeVersion != nil {
		in, out := &in.ChangeVersion, &out.ChangeVersion
		*out = new(string)
		**out = **in
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(MachineSpecClone)
		(*in).DeepCopyInto(*out)
	}
	if in.CpuHotAddEnabled != nil {
		in, out := &in.CpuHotAddEnabled, &out.CpuHotAddEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CpuHotRemoveEnabled != nil {
		in, out := &in.CpuHotRemoveEnabled, &out.CpuHotRemoveEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CpuLimit != nil {
		in, out := &in.CpuLimit, &out.CpuLimit
		*out = new(int64)
		**out = **in
	}
	if in.CpuPerformanceCountersEnabled != nil {
		in, out := &in.CpuPerformanceCountersEnabled, &out.CpuPerformanceCountersEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CpuReservation != nil {
		in, out := &in.CpuReservation, &out.CpuReservation
		*out = new(int64)
		**out = **in
	}
	if in.CpuShareCount != nil {
		in, out := &in.CpuShareCount, &out.CpuShareCount
		*out = new(int64)
		**out = **in
	}
	if in.CpuShareLevel != nil {
		in, out := &in.CpuShareLevel, &out.CpuShareLevel
		*out = new(string)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.DatastoreID != nil {
		in, out := &in.DatastoreID, &out.DatastoreID
		*out = new(string)
		**out = **in
	}
	if in.DefaultIPAddress != nil {
		in, out := &in.DefaultIPAddress, &out.DefaultIPAddress
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]MachineSpecDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EfiSecureBootEnabled != nil {
		in, out := &in.EfiSecureBootEnabled, &out.EfiSecureBootEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableDiskUUID != nil {
		in, out := &in.EnableDiskUUID, &out.EnableDiskUUID
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.EptRviMode != nil {
		in, out := &in.EptRviMode, &out.EptRviMode
		*out = new(string)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.ForcePowerOff != nil {
		in, out := &in.ForcePowerOff, &out.ForcePowerOff
		*out = new(bool)
		**out = **in
	}
	if in.GuestID != nil {
		in, out := &in.GuestID, &out.GuestID
		*out = new(string)
		**out = **in
	}
	if in.GuestIPAddresses != nil {
		in, out := &in.GuestIPAddresses, &out.GuestIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HardwareVersion != nil {
		in, out := &in.HardwareVersion, &out.HardwareVersion
		*out = new(int64)
		**out = **in
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
	if in.HvMode != nil {
		in, out := &in.HvMode, &out.HvMode
		*out = new(string)
		**out = **in
	}
	if in.IdeControllerCount != nil {
		in, out := &in.IdeControllerCount, &out.IdeControllerCount
		*out = new(int64)
		**out = **in
	}
	if in.IgnoredGuestIPS != nil {
		in, out := &in.IgnoredGuestIPS, &out.IgnoredGuestIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Imported != nil {
		in, out := &in.Imported, &out.Imported
		*out = new(bool)
		**out = **in
	}
	if in.LatencySensitivity != nil {
		in, out := &in.LatencySensitivity, &out.LatencySensitivity
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(int64)
		**out = **in
	}
	if in.MemoryHotAddEnabled != nil {
		in, out := &in.MemoryHotAddEnabled, &out.MemoryHotAddEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(int64)
		**out = **in
	}
	if in.MemoryReservation != nil {
		in, out := &in.MemoryReservation, &out.MemoryReservation
		*out = new(int64)
		**out = **in
	}
	if in.MemoryShareCount != nil {
		in, out := &in.MemoryShareCount, &out.MemoryShareCount
		*out = new(int64)
		**out = **in
	}
	if in.MemoryShareLevel != nil {
		in, out := &in.MemoryShareLevel, &out.MemoryShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MigrateWaitTimeout != nil {
		in, out := &in.MigrateWaitTimeout, &out.MigrateWaitTimeout
		*out = new(int64)
		**out = **in
	}
	if in.Moid != nil {
		in, out := &in.Moid, &out.Moid
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NestedHvEnabled != nil {
		in, out := &in.NestedHvEnabled, &out.NestedHvEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]MachineSpecNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NumCoresPerSocket != nil {
		in, out := &in.NumCoresPerSocket, &out.NumCoresPerSocket
		*out = new(int64)
		**out = **in
	}
	if in.NumCpus != nil {
		in, out := &in.NumCpus, &out.NumCpus
		*out = new(int64)
		**out = **in
	}
	if in.OvfDeploy != nil {
		in, out := &in.OvfDeploy, &out.OvfDeploy
		*out = new(MachineSpecOvfDeploy)
		(*in).DeepCopyInto(*out)
	}
	if in.PciDeviceID != nil {
		in, out := &in.PciDeviceID, &out.PciDeviceID
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PoweronTimeout != nil {
		in, out := &in.PoweronTimeout, &out.PoweronTimeout
		*out = new(int64)
		**out = **in
	}
	if in.RebootRequired != nil {
		in, out := &in.RebootRequired, &out.RebootRequired
		*out = new(bool)
		**out = **in
	}
	if in.ReplaceTrigger != nil {
		in, out := &in.ReplaceTrigger, &out.ReplaceTrigger
		*out = new(string)
		**out = **in
	}
	if in.ResourcePoolID != nil {
		in, out := &in.ResourcePoolID, &out.ResourcePoolID
		*out = new(string)
		**out = **in
	}
	if in.RunToolsScriptsAfterPowerOn != nil {
		in, out := &in.RunToolsScriptsAfterPowerOn, &out.RunToolsScriptsAfterPowerOn
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsAfterResume != nil {
		in, out := &in.RunToolsScriptsAfterResume, &out.RunToolsScriptsAfterResume
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestReboot != nil {
		in, out := &in.RunToolsScriptsBeforeGuestReboot, &out.RunToolsScriptsBeforeGuestReboot
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestShutdown != nil {
		in, out := &in.RunToolsScriptsBeforeGuestShutdown, &out.RunToolsScriptsBeforeGuestShutdown
		*out = new(bool)
		**out = **in
	}
	if in.RunToolsScriptsBeforeGuestStandby != nil {
		in, out := &in.RunToolsScriptsBeforeGuestStandby, &out.RunToolsScriptsBeforeGuestStandby
		*out = new(bool)
		**out = **in
	}
	if in.SataControllerCount != nil {
		in, out := &in.SataControllerCount, &out.SataControllerCount
		*out = new(int64)
		**out = **in
	}
	if in.ScsiBusSharing != nil {
		in, out := &in.ScsiBusSharing, &out.ScsiBusSharing
		*out = new(string)
		**out = **in
	}
	if in.ScsiControllerCount != nil {
		in, out := &in.ScsiControllerCount, &out.ScsiControllerCount
		*out = new(int64)
		**out = **in
	}
	if in.ScsiType != nil {
		in, out := &in.ScsiType, &out.ScsiType
		*out = new(string)
		**out = **in
	}
	if in.ShutdownWaitTimeout != nil {
		in, out := &in.ShutdownWaitTimeout, &out.ShutdownWaitTimeout
		*out = new(int64)
		**out = **in
	}
	if in.StoragePolicyID != nil {
		in, out := &in.StoragePolicyID, &out.StoragePolicyID
		*out = new(string)
		**out = **in
	}
	if in.SwapPlacementPolicy != nil {
		in, out := &in.SwapPlacementPolicy, &out.SwapPlacementPolicy
		*out = new(string)
		**out = **in
	}
	if in.SyncTimeWithHost != nil {
		in, out := &in.SyncTimeWithHost, &out.SyncTimeWithHost
		*out = new(bool)
		**out = **in
	}
	if in.SyncTimeWithHostPeriodically != nil {
		in, out := &in.SyncTimeWithHostPeriodically, &out.SyncTimeWithHostPeriodically
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Uuid != nil {
		in, out := &in.Uuid, &out.Uuid
		*out = new(string)
		**out = **in
	}
	if in.Vapp != nil {
		in, out := &in.Vapp, &out.Vapp
		*out = new(MachineSpecVapp)
		(*in).DeepCopyInto(*out)
	}
	if in.VappTransport != nil {
		in, out := &in.VappTransport, &out.VappTransport
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VbsEnabled != nil {
		in, out := &in.VbsEnabled, &out.VbsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VmwareToolsStatus != nil {
		in, out := &in.VmwareToolsStatus, &out.VmwareToolsStatus
		*out = new(string)
		**out = **in
	}
	if in.VmxPath != nil {
		in, out := &in.VmxPath, &out.VmxPath
		*out = new(string)
		**out = **in
	}
	if in.VvtdEnabled != nil {
		in, out := &in.VvtdEnabled, &out.VvtdEnabled
		*out = new(bool)
		**out = **in
	}
	if in.WaitForGuestIPTimeout != nil {
		in, out := &in.WaitForGuestIPTimeout, &out.WaitForGuestIPTimeout
		*out = new(int64)
		**out = **in
	}
	if in.WaitForGuestNetRoutable != nil {
		in, out := &in.WaitForGuestNetRoutable, &out.WaitForGuestNetRoutable
		*out = new(bool)
		**out = **in
	}
	if in.WaitForGuestNetTimeout != nil {
		in, out := &in.WaitForGuestNetTimeout, &out.WaitForGuestNetTimeout
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecResource.
func (in *MachineSpecResource) DeepCopy() *MachineSpecResource {
	if in == nil {
		return nil
	}
	out := new(MachineSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpecVapp) DeepCopyInto(out *MachineSpecVapp) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpecVapp.
func (in *MachineSpecVapp) DeepCopy() *MachineSpecVapp {
	if in == nil {
		return nil
	}
	out := new(MachineSpecVapp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}
