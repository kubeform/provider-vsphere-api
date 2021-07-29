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
	"context"

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/vm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStoragePolicies implements StoragePolicyInterface
type FakeStoragePolicies struct {
	Fake *FakeVmV1alpha1
	ns   string
}

var storagepoliciesResource = schema.GroupVersionResource{Group: "vm.vsphere.kubeform.com", Version: "v1alpha1", Resource: "storagepolicies"}

var storagepoliciesKind = schema.GroupVersionKind{Group: "vm.vsphere.kubeform.com", Version: "v1alpha1", Kind: "StoragePolicy"}

// Get takes name of the storagePolicy, and returns the corresponding storagePolicy object, and an error if there is any.
func (c *FakeStoragePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StoragePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagepoliciesResource, c.ns, name), &v1alpha1.StoragePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragePolicy), err
}

// List takes label and field selectors, and returns the list of StoragePolicies that match those selectors.
func (c *FakeStoragePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StoragePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagepoliciesResource, storagepoliciesKind, c.ns, opts), &v1alpha1.StoragePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoragePolicyList{ListMeta: obj.(*v1alpha1.StoragePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoragePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storagePolicies.
func (c *FakeStoragePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagepoliciesResource, c.ns, opts))

}

// Create takes the representation of a storagePolicy and creates it.  Returns the server's representation of the storagePolicy, and an error, if there is any.
func (c *FakeStoragePolicies) Create(ctx context.Context, storagePolicy *v1alpha1.StoragePolicy, opts v1.CreateOptions) (result *v1alpha1.StoragePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagepoliciesResource, c.ns, storagePolicy), &v1alpha1.StoragePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragePolicy), err
}

// Update takes the representation of a storagePolicy and updates it. Returns the server's representation of the storagePolicy, and an error, if there is any.
func (c *FakeStoragePolicies) Update(ctx context.Context, storagePolicy *v1alpha1.StoragePolicy, opts v1.UpdateOptions) (result *v1alpha1.StoragePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagepoliciesResource, c.ns, storagePolicy), &v1alpha1.StoragePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStoragePolicies) UpdateStatus(ctx context.Context, storagePolicy *v1alpha1.StoragePolicy, opts v1.UpdateOptions) (*v1alpha1.StoragePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagepoliciesResource, "status", c.ns, storagePolicy), &v1alpha1.StoragePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragePolicy), err
}

// Delete takes name of the storagePolicy and deletes it. Returns an error if one occurs.
func (c *FakeStoragePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagepoliciesResource, c.ns, name), &v1alpha1.StoragePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStoragePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoragePolicyList{})
	return err
}

// Patch applies the patch and returns the patched storagePolicy.
func (c *FakeStoragePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StoragePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StoragePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragePolicy), err
}
