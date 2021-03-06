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
func (in *VmOverride) DeepCopyInto(out *VmOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmOverride.
func (in *VmOverride) DeepCopy() *VmOverride {
	if in == nil {
		return nil
	}
	out := new(VmOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmOverrideList) DeepCopyInto(out *VmOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VmOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmOverrideList.
func (in *VmOverrideList) DeepCopy() *VmOverrideList {
	if in == nil {
		return nil
	}
	out := new(VmOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VmOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmOverrideSpec) DeepCopyInto(out *VmOverrideSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VmOverrideSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmOverrideSpec.
func (in *VmOverrideSpec) DeepCopy() *VmOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(VmOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmOverrideSpecResource) DeepCopyInto(out *VmOverrideSpecResource) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdRecoveryAction != nil {
		in, out := &in.HaDatastoreApdRecoveryAction, &out.HaDatastoreApdRecoveryAction
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdResponse != nil {
		in, out := &in.HaDatastoreApdResponse, &out.HaDatastoreApdResponse
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdResponseDelay != nil {
		in, out := &in.HaDatastoreApdResponseDelay, &out.HaDatastoreApdResponseDelay
		*out = new(int64)
		**out = **in
	}
	if in.HaDatastorePdlResponse != nil {
		in, out := &in.HaDatastorePdlResponse, &out.HaDatastorePdlResponse
		*out = new(string)
		**out = **in
	}
	if in.HaHostIsolationResponse != nil {
		in, out := &in.HaHostIsolationResponse, &out.HaHostIsolationResponse
		*out = new(string)
		**out = **in
	}
	if in.HaVmFailureInterval != nil {
		in, out := &in.HaVmFailureInterval, &out.HaVmFailureInterval
		*out = new(int64)
		**out = **in
	}
	if in.HaVmMaximumFailureWindow != nil {
		in, out := &in.HaVmMaximumFailureWindow, &out.HaVmMaximumFailureWindow
		*out = new(int64)
		**out = **in
	}
	if in.HaVmMaximumResets != nil {
		in, out := &in.HaVmMaximumResets, &out.HaVmMaximumResets
		*out = new(int64)
		**out = **in
	}
	if in.HaVmMinimumUptime != nil {
		in, out := &in.HaVmMinimumUptime, &out.HaVmMinimumUptime
		*out = new(int64)
		**out = **in
	}
	if in.HaVmMonitoring != nil {
		in, out := &in.HaVmMonitoring, &out.HaVmMonitoring
		*out = new(string)
		**out = **in
	}
	if in.HaVmMonitoringUseClusterDefaults != nil {
		in, out := &in.HaVmMonitoringUseClusterDefaults, &out.HaVmMonitoringUseClusterDefaults
		*out = new(bool)
		**out = **in
	}
	if in.HaVmRestartPriority != nil {
		in, out := &in.HaVmRestartPriority, &out.HaVmRestartPriority
		*out = new(string)
		**out = **in
	}
	if in.HaVmRestartTimeout != nil {
		in, out := &in.HaVmRestartTimeout, &out.HaVmRestartTimeout
		*out = new(int64)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmOverrideSpecResource.
func (in *VmOverrideSpecResource) DeepCopy() *VmOverrideSpecResource {
	if in == nil {
		return nil
	}
	out := new(VmOverrideSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmOverrideStatus) DeepCopyInto(out *VmOverrideStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmOverrideStatus.
func (in *VmOverrideStatus) DeepCopy() *VmOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(VmOverrideStatus)
	in.DeepCopyInto(out)
	return out
}
