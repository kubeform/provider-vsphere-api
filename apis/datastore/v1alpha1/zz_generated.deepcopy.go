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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
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
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SdrsAdvancedOptions != nil {
		in, out := &in.SdrsAdvancedOptions, &out.SdrsAdvancedOptions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.SdrsAutomationLevel != nil {
		in, out := &in.SdrsAutomationLevel, &out.SdrsAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsDefaultIntraVmAffinity != nil {
		in, out := &in.SdrsDefaultIntraVmAffinity, &out.SdrsDefaultIntraVmAffinity
		*out = new(bool)
		**out = **in
	}
	if in.SdrsEnabled != nil {
		in, out := &in.SdrsEnabled, &out.SdrsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SdrsFreeSpaceThreshold != nil {
		in, out := &in.SdrsFreeSpaceThreshold, &out.SdrsFreeSpaceThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsFreeSpaceThresholdMode != nil {
		in, out := &in.SdrsFreeSpaceThresholdMode, &out.SdrsFreeSpaceThresholdMode
		*out = new(string)
		**out = **in
	}
	if in.SdrsFreeSpaceUtilizationDifference != nil {
		in, out := &in.SdrsFreeSpaceUtilizationDifference, &out.SdrsFreeSpaceUtilizationDifference
		*out = new(int64)
		**out = **in
	}
	if in.SdrsIoBalanceAutomationLevel != nil {
		in, out := &in.SdrsIoBalanceAutomationLevel, &out.SdrsIoBalanceAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsIoLatencyThreshold != nil {
		in, out := &in.SdrsIoLatencyThreshold, &out.SdrsIoLatencyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsIoLoadBalanceEnabled != nil {
		in, out := &in.SdrsIoLoadBalanceEnabled, &out.SdrsIoLoadBalanceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SdrsIoLoadImbalanceThreshold != nil {
		in, out := &in.SdrsIoLoadImbalanceThreshold, &out.SdrsIoLoadImbalanceThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsIoReservableIopsThreshold != nil {
		in, out := &in.SdrsIoReservableIopsThreshold, &out.SdrsIoReservableIopsThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsIoReservablePercentThreshold != nil {
		in, out := &in.SdrsIoReservablePercentThreshold, &out.SdrsIoReservablePercentThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsIoReservableThresholdMode != nil {
		in, out := &in.SdrsIoReservableThresholdMode, &out.SdrsIoReservableThresholdMode
		*out = new(string)
		**out = **in
	}
	if in.SdrsLoadBalanceInterval != nil {
		in, out := &in.SdrsLoadBalanceInterval, &out.SdrsLoadBalanceInterval
		*out = new(int64)
		**out = **in
	}
	if in.SdrsPolicyEnforcementAutomationLevel != nil {
		in, out := &in.SdrsPolicyEnforcementAutomationLevel, &out.SdrsPolicyEnforcementAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsRuleEnforcementAutomationLevel != nil {
		in, out := &in.SdrsRuleEnforcementAutomationLevel, &out.SdrsRuleEnforcementAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsSpaceBalanceAutomationLevel != nil {
		in, out := &in.SdrsSpaceBalanceAutomationLevel, &out.SdrsSpaceBalanceAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsSpaceUtilizationThreshold != nil {
		in, out := &in.SdrsSpaceUtilizationThreshold, &out.SdrsSpaceUtilizationThreshold
		*out = new(int64)
		**out = **in
	}
	if in.SdrsVmEvacuationAutomationLevel != nil {
		in, out := &in.SdrsVmEvacuationAutomationLevel, &out.SdrsVmEvacuationAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVmAntiAffinityRule) DeepCopyInto(out *ClusterVmAntiAffinityRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVmAntiAffinityRule.
func (in *ClusterVmAntiAffinityRule) DeepCopy() *ClusterVmAntiAffinityRule {
	if in == nil {
		return nil
	}
	out := new(ClusterVmAntiAffinityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVmAntiAffinityRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVmAntiAffinityRuleList) DeepCopyInto(out *ClusterVmAntiAffinityRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVmAntiAffinityRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVmAntiAffinityRuleList.
func (in *ClusterVmAntiAffinityRuleList) DeepCopy() *ClusterVmAntiAffinityRuleList {
	if in == nil {
		return nil
	}
	out := new(ClusterVmAntiAffinityRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVmAntiAffinityRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVmAntiAffinityRuleSpec) DeepCopyInto(out *ClusterVmAntiAffinityRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterVmAntiAffinityRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVmAntiAffinityRuleSpec.
func (in *ClusterVmAntiAffinityRuleSpec) DeepCopy() *ClusterVmAntiAffinityRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterVmAntiAffinityRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVmAntiAffinityRuleSpecResource) DeepCopyInto(out *ClusterVmAntiAffinityRuleSpecResource) {
	*out = *in
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Mandatory != nil {
		in, out := &in.Mandatory, &out.Mandatory
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineIDS != nil {
		in, out := &in.VirtualMachineIDS, &out.VirtualMachineIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVmAntiAffinityRuleSpecResource.
func (in *ClusterVmAntiAffinityRuleSpecResource) DeepCopy() *ClusterVmAntiAffinityRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterVmAntiAffinityRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVmAntiAffinityRuleStatus) DeepCopyInto(out *ClusterVmAntiAffinityRuleStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVmAntiAffinityRuleStatus.
func (in *ClusterVmAntiAffinityRuleStatus) DeepCopy() *ClusterVmAntiAffinityRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterVmAntiAffinityRuleStatus)
	in.DeepCopyInto(out)
	return out
}