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
func (in *Vnic) DeepCopyInto(out *Vnic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vnic.
func (in *Vnic) DeepCopy() *Vnic {
	if in == nil {
		return nil
	}
	out := new(Vnic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vnic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicList) DeepCopyInto(out *VnicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vnic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicList.
func (in *VnicList) DeepCopy() *VnicList {
	if in == nil {
		return nil
	}
	out := new(VnicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VnicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicSpec) DeepCopyInto(out *VnicSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VnicSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicSpec.
func (in *VnicSpec) DeepCopy() *VnicSpec {
	if in == nil {
		return nil
	}
	out := new(VnicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicSpecIpv4) DeepCopyInto(out *VnicSpecIpv4) {
	*out = *in
	if in.Dhcp != nil {
		in, out := &in.Dhcp, &out.Dhcp
		*out = new(bool)
		**out = **in
	}
	if in.Gw != nil {
		in, out := &in.Gw, &out.Gw
		*out = new(string)
		**out = **in
	}
	if in.Ip != nil {
		in, out := &in.Ip, &out.Ip
		*out = new(string)
		**out = **in
	}
	if in.Netmask != nil {
		in, out := &in.Netmask, &out.Netmask
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicSpecIpv4.
func (in *VnicSpecIpv4) DeepCopy() *VnicSpecIpv4 {
	if in == nil {
		return nil
	}
	out := new(VnicSpecIpv4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicSpecIpv6) DeepCopyInto(out *VnicSpecIpv6) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Autoconfig != nil {
		in, out := &in.Autoconfig, &out.Autoconfig
		*out = new(bool)
		**out = **in
	}
	if in.Dhcp != nil {
		in, out := &in.Dhcp, &out.Dhcp
		*out = new(bool)
		**out = **in
	}
	if in.Gw != nil {
		in, out := &in.Gw, &out.Gw
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicSpecIpv6.
func (in *VnicSpecIpv6) DeepCopy() *VnicSpecIpv6 {
	if in == nil {
		return nil
	}
	out := new(VnicSpecIpv6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicSpecResource) DeepCopyInto(out *VnicSpecResource) {
	*out = *in
	if in.DistributedPortGroup != nil {
		in, out := &in.DistributedPortGroup, &out.DistributedPortGroup
		*out = new(string)
		**out = **in
	}
	if in.DistributedSwitchPort != nil {
		in, out := &in.DistributedSwitchPort, &out.DistributedSwitchPort
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Ipv4 != nil {
		in, out := &in.Ipv4, &out.Ipv4
		*out = new(VnicSpecIpv4)
		(*in).DeepCopyInto(*out)
	}
	if in.Ipv6 != nil {
		in, out := &in.Ipv6, &out.Ipv6
		*out = new(VnicSpecIpv6)
		(*in).DeepCopyInto(*out)
	}
	if in.Mac != nil {
		in, out := &in.Mac, &out.Mac
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(int64)
		**out = **in
	}
	if in.Netstack != nil {
		in, out := &in.Netstack, &out.Netstack
		*out = new(string)
		**out = **in
	}
	if in.Portgroup != nil {
		in, out := &in.Portgroup, &out.Portgroup
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicSpecResource.
func (in *VnicSpecResource) DeepCopy() *VnicSpecResource {
	if in == nil {
		return nil
	}
	out := new(VnicSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnicStatus) DeepCopyInto(out *VnicStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnicStatus.
func (in *VnicStatus) DeepCopy() *VnicStatus {
	if in == nil {
		return nil
	}
	out := new(VnicStatus)
	in.DeepCopyInto(out)
	return out
}