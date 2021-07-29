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

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/datacenter/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatacenters implements DatacenterInterface
type FakeDatacenters struct {
	Fake *FakeDatacenterV1alpha1
	ns   string
}

var datacentersResource = schema.GroupVersionResource{Group: "datacenter.vsphere.kubeform.com", Version: "v1alpha1", Resource: "datacenters"}

var datacentersKind = schema.GroupVersionKind{Group: "datacenter.vsphere.kubeform.com", Version: "v1alpha1", Kind: "Datacenter"}

// Get takes name of the datacenter, and returns the corresponding datacenter object, and an error if there is any.
func (c *FakeDatacenters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Datacenter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datacentersResource, c.ns, name), &v1alpha1.Datacenter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Datacenter), err
}

// List takes label and field selectors, and returns the list of Datacenters that match those selectors.
func (c *FakeDatacenters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatacenterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datacentersResource, datacentersKind, c.ns, opts), &v1alpha1.DatacenterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatacenterList{ListMeta: obj.(*v1alpha1.DatacenterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatacenterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datacenters.
func (c *FakeDatacenters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datacentersResource, c.ns, opts))

}

// Create takes the representation of a datacenter and creates it.  Returns the server's representation of the datacenter, and an error, if there is any.
func (c *FakeDatacenters) Create(ctx context.Context, datacenter *v1alpha1.Datacenter, opts v1.CreateOptions) (result *v1alpha1.Datacenter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datacentersResource, c.ns, datacenter), &v1alpha1.Datacenter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Datacenter), err
}

// Update takes the representation of a datacenter and updates it. Returns the server's representation of the datacenter, and an error, if there is any.
func (c *FakeDatacenters) Update(ctx context.Context, datacenter *v1alpha1.Datacenter, opts v1.UpdateOptions) (result *v1alpha1.Datacenter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datacentersResource, c.ns, datacenter), &v1alpha1.Datacenter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Datacenter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatacenters) UpdateStatus(ctx context.Context, datacenter *v1alpha1.Datacenter, opts v1.UpdateOptions) (*v1alpha1.Datacenter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datacentersResource, "status", c.ns, datacenter), &v1alpha1.Datacenter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Datacenter), err
}

// Delete takes name of the datacenter and deletes it. Returns an error if one occurs.
func (c *FakeDatacenters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datacentersResource, c.ns, name), &v1alpha1.Datacenter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatacenters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datacentersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatacenterList{})
	return err
}

// Patch applies the patch and returns the patched datacenter.
func (c *FakeDatacenters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Datacenter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datacentersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Datacenter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Datacenter), err
}
