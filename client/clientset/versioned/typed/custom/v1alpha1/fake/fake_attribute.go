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

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/custom/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAttributes implements AttributeInterface
type FakeAttributes struct {
	Fake *FakeCustomV1alpha1
	ns   string
}

var attributesResource = schema.GroupVersionResource{Group: "custom.vsphere.kubeform.com", Version: "v1alpha1", Resource: "attributes"}

var attributesKind = schema.GroupVersionKind{Group: "custom.vsphere.kubeform.com", Version: "v1alpha1", Kind: "Attribute"}

// Get takes name of the attribute, and returns the corresponding attribute object, and an error if there is any.
func (c *FakeAttributes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(attributesResource, c.ns, name), &v1alpha1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attribute), err
}

// List takes label and field selectors, and returns the list of Attributes that match those selectors.
func (c *FakeAttributes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AttributeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(attributesResource, attributesKind, c.ns, opts), &v1alpha1.AttributeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AttributeList{ListMeta: obj.(*v1alpha1.AttributeList).ListMeta}
	for _, item := range obj.(*v1alpha1.AttributeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested attributes.
func (c *FakeAttributes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(attributesResource, c.ns, opts))

}

// Create takes the representation of a attribute and creates it.  Returns the server's representation of the attribute, and an error, if there is any.
func (c *FakeAttributes) Create(ctx context.Context, attribute *v1alpha1.Attribute, opts v1.CreateOptions) (result *v1alpha1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(attributesResource, c.ns, attribute), &v1alpha1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attribute), err
}

// Update takes the representation of a attribute and updates it. Returns the server's representation of the attribute, and an error, if there is any.
func (c *FakeAttributes) Update(ctx context.Context, attribute *v1alpha1.Attribute, opts v1.UpdateOptions) (result *v1alpha1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(attributesResource, c.ns, attribute), &v1alpha1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attribute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAttributes) UpdateStatus(ctx context.Context, attribute *v1alpha1.Attribute, opts v1.UpdateOptions) (*v1alpha1.Attribute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(attributesResource, "status", c.ns, attribute), &v1alpha1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attribute), err
}

// Delete takes name of the attribute and deletes it. Returns an error if one occurs.
func (c *FakeAttributes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(attributesResource, c.ns, name), &v1alpha1.Attribute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAttributes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(attributesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AttributeList{})
	return err
}

// Patch applies the patch and returns the patched attribute.
func (c *FakeAttributes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(attributesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attribute), err
}
