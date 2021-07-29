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
func (in *Datastore) DeepCopyInto(out *Datastore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datastore.
func (in *Datastore) DeepCopy() *Datastore {
	if in == nil {
		return nil
	}
	out := new(Datastore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datastore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreList) DeepCopyInto(out *DatastoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datastore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreList.
func (in *DatastoreList) DeepCopy() *DatastoreList {
	if in == nil {
		return nil
	}
	out := new(DatastoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreSpec) DeepCopyInto(out *DatastoreSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DatastoreSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreSpec.
func (in *DatastoreSpec) DeepCopy() *DatastoreSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreSpecResource) DeepCopyInto(out *DatastoreSpecResource) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.Accessible != nil {
		in, out := &in.Accessible, &out.Accessible
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
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
	if in.DatastoreClusterID != nil {
		in, out := &in.DatastoreClusterID, &out.DatastoreClusterID
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.FreeSpace != nil {
		in, out := &in.FreeSpace, &out.FreeSpace
		*out = new(int64)
		**out = **in
	}
	if in.HostSystemIDS != nil {
		in, out := &in.HostSystemIDS, &out.HostSystemIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaintenanceMode != nil {
		in, out := &in.MaintenanceMode, &out.MaintenanceMode
		*out = new(string)
		**out = **in
	}
	if in.MultipleHostAccess != nil {
		in, out := &in.MultipleHostAccess, &out.MultipleHostAccess
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProtocolEndpoint != nil {
		in, out := &in.ProtocolEndpoint, &out.ProtocolEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.RemoteHosts != nil {
		in, out := &in.RemoteHosts, &out.RemoteHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemotePath != nil {
		in, out := &in.RemotePath, &out.RemotePath
		*out = new(string)
		**out = **in
	}
	if in.SecurityType != nil {
		in, out := &in.SecurityType, &out.SecurityType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UncommittedSpace != nil {
		in, out := &in.UncommittedSpace, &out.UncommittedSpace
		*out = new(int64)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreSpecResource.
func (in *DatastoreSpecResource) DeepCopy() *DatastoreSpecResource {
	if in == nil {
		return nil
	}
	out := new(DatastoreSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreStatus) DeepCopyInto(out *DatastoreStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreStatus.
func (in *DatastoreStatus) DeepCopy() *DatastoreStatus {
	if in == nil {
		return nil
	}
	out := new(DatastoreStatus)
	in.DeepCopyInto(out)
	return out
}