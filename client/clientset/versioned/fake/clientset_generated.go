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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "kubeform.dev/provider-vsphere-api/client/clientset/versioned"
	computev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/compute/v1alpha1"
	fakecomputev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/compute/v1alpha1/fake"
	contentv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/content/v1alpha1"
	fakecontentv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/content/v1alpha1/fake"
	customv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/custom/v1alpha1"
	fakecustomv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/custom/v1alpha1/fake"
	datacenterv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/datacenter/v1alpha1"
	fakedatacenterv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/datacenter/v1alpha1/fake"
	datastorev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/datastore/v1alpha1"
	fakedatastorev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/datastore/v1alpha1/fake"
	distributedv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/distributed/v1alpha1"
	fakedistributedv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/distributed/v1alpha1/fake"
	dpmv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/dpm/v1alpha1"
	fakedpmv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/dpm/v1alpha1/fake"
	drsv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/drs/v1alpha1"
	fakedrsv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/drs/v1alpha1/fake"
	entityv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/entity/v1alpha1"
	fakeentityv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/entity/v1alpha1/fake"
	filev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/file/v1alpha1"
	fakefilev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/file/v1alpha1/fake"
	folderv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/folder/v1alpha1"
	fakefolderv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/folder/v1alpha1/fake"
	hav1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/ha/v1alpha1"
	fakehav1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/ha/v1alpha1/fake"
	hostv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/host/v1alpha1"
	fakehostv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/host/v1alpha1/fake"
	licensev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/license/v1alpha1"
	fakelicensev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/license/v1alpha1/fake"
	nasv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/nas/v1alpha1"
	fakenasv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/nas/v1alpha1/fake"
	resourcev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/resource/v1alpha1"
	fakeresourcev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/resource/v1alpha1/fake"
	rolev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/role/v1alpha1"
	fakerolev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/role/v1alpha1/fake"
	storagev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/storage/v1alpha1/fake"
	tagv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/tag/v1alpha1"
	faketagv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/tag/v1alpha1/fake"
	vappv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vapp/v1alpha1"
	fakevappv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vapp/v1alpha1/fake"
	virtualv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/virtual/v1alpha1"
	fakevirtualv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/virtual/v1alpha1/fake"
	vmv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vm/v1alpha1"
	fakevmv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vm/v1alpha1/fake"
	vmfsv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vmfs/v1alpha1"
	fakevmfsv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vmfs/v1alpha1/fake"
	vnicv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vnic/v1alpha1"
	fakevnicv1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/vnic/v1alpha1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ComputeV1alpha1 retrieves the ComputeV1alpha1Client
func (c *Clientset) ComputeV1alpha1() computev1alpha1.ComputeV1alpha1Interface {
	return &fakecomputev1alpha1.FakeComputeV1alpha1{Fake: &c.Fake}
}

// ContentV1alpha1 retrieves the ContentV1alpha1Client
func (c *Clientset) ContentV1alpha1() contentv1alpha1.ContentV1alpha1Interface {
	return &fakecontentv1alpha1.FakeContentV1alpha1{Fake: &c.Fake}
}

// CustomV1alpha1 retrieves the CustomV1alpha1Client
func (c *Clientset) CustomV1alpha1() customv1alpha1.CustomV1alpha1Interface {
	return &fakecustomv1alpha1.FakeCustomV1alpha1{Fake: &c.Fake}
}

// DatacenterV1alpha1 retrieves the DatacenterV1alpha1Client
func (c *Clientset) DatacenterV1alpha1() datacenterv1alpha1.DatacenterV1alpha1Interface {
	return &fakedatacenterv1alpha1.FakeDatacenterV1alpha1{Fake: &c.Fake}
}

// DatastoreV1alpha1 retrieves the DatastoreV1alpha1Client
func (c *Clientset) DatastoreV1alpha1() datastorev1alpha1.DatastoreV1alpha1Interface {
	return &fakedatastorev1alpha1.FakeDatastoreV1alpha1{Fake: &c.Fake}
}

// DistributedV1alpha1 retrieves the DistributedV1alpha1Client
func (c *Clientset) DistributedV1alpha1() distributedv1alpha1.DistributedV1alpha1Interface {
	return &fakedistributedv1alpha1.FakeDistributedV1alpha1{Fake: &c.Fake}
}

// DpmV1alpha1 retrieves the DpmV1alpha1Client
func (c *Clientset) DpmV1alpha1() dpmv1alpha1.DpmV1alpha1Interface {
	return &fakedpmv1alpha1.FakeDpmV1alpha1{Fake: &c.Fake}
}

// DrsV1alpha1 retrieves the DrsV1alpha1Client
func (c *Clientset) DrsV1alpha1() drsv1alpha1.DrsV1alpha1Interface {
	return &fakedrsv1alpha1.FakeDrsV1alpha1{Fake: &c.Fake}
}

// EntityV1alpha1 retrieves the EntityV1alpha1Client
func (c *Clientset) EntityV1alpha1() entityv1alpha1.EntityV1alpha1Interface {
	return &fakeentityv1alpha1.FakeEntityV1alpha1{Fake: &c.Fake}
}

// FileV1alpha1 retrieves the FileV1alpha1Client
func (c *Clientset) FileV1alpha1() filev1alpha1.FileV1alpha1Interface {
	return &fakefilev1alpha1.FakeFileV1alpha1{Fake: &c.Fake}
}

// FolderV1alpha1 retrieves the FolderV1alpha1Client
func (c *Clientset) FolderV1alpha1() folderv1alpha1.FolderV1alpha1Interface {
	return &fakefolderv1alpha1.FakeFolderV1alpha1{Fake: &c.Fake}
}

// HaV1alpha1 retrieves the HaV1alpha1Client
func (c *Clientset) HaV1alpha1() hav1alpha1.HaV1alpha1Interface {
	return &fakehav1alpha1.FakeHaV1alpha1{Fake: &c.Fake}
}

// HostV1alpha1 retrieves the HostV1alpha1Client
func (c *Clientset) HostV1alpha1() hostv1alpha1.HostV1alpha1Interface {
	return &fakehostv1alpha1.FakeHostV1alpha1{Fake: &c.Fake}
}

// LicenseV1alpha1 retrieves the LicenseV1alpha1Client
func (c *Clientset) LicenseV1alpha1() licensev1alpha1.LicenseV1alpha1Interface {
	return &fakelicensev1alpha1.FakeLicenseV1alpha1{Fake: &c.Fake}
}

// NasV1alpha1 retrieves the NasV1alpha1Client
func (c *Clientset) NasV1alpha1() nasv1alpha1.NasV1alpha1Interface {
	return &fakenasv1alpha1.FakeNasV1alpha1{Fake: &c.Fake}
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return &fakeresourcev1alpha1.FakeResourceV1alpha1{Fake: &c.Fake}
}

// RoleV1alpha1 retrieves the RoleV1alpha1Client
func (c *Clientset) RoleV1alpha1() rolev1alpha1.RoleV1alpha1Interface {
	return &fakerolev1alpha1.FakeRoleV1alpha1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// TagV1alpha1 retrieves the TagV1alpha1Client
func (c *Clientset) TagV1alpha1() tagv1alpha1.TagV1alpha1Interface {
	return &faketagv1alpha1.FakeTagV1alpha1{Fake: &c.Fake}
}

// VappV1alpha1 retrieves the VappV1alpha1Client
func (c *Clientset) VappV1alpha1() vappv1alpha1.VappV1alpha1Interface {
	return &fakevappv1alpha1.FakeVappV1alpha1{Fake: &c.Fake}
}

// VirtualV1alpha1 retrieves the VirtualV1alpha1Client
func (c *Clientset) VirtualV1alpha1() virtualv1alpha1.VirtualV1alpha1Interface {
	return &fakevirtualv1alpha1.FakeVirtualV1alpha1{Fake: &c.Fake}
}

// VmV1alpha1 retrieves the VmV1alpha1Client
func (c *Clientset) VmV1alpha1() vmv1alpha1.VmV1alpha1Interface {
	return &fakevmv1alpha1.FakeVmV1alpha1{Fake: &c.Fake}
}

// VmfsV1alpha1 retrieves the VmfsV1alpha1Client
func (c *Clientset) VmfsV1alpha1() vmfsv1alpha1.VmfsV1alpha1Interface {
	return &fakevmfsv1alpha1.FakeVmfsV1alpha1{Fake: &c.Fake}
}

// VnicV1alpha1 retrieves the VnicV1alpha1Client
func (c *Clientset) VnicV1alpha1() vnicv1alpha1.VnicV1alpha1Interface {
	return &fakevnicv1alpha1.FakeVnicV1alpha1{Fake: &c.Fake}
}
