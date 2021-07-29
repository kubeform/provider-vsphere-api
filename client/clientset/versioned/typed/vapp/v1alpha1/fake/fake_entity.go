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

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/vapp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEntities implements EntityInterface
type FakeEntities struct {
	Fake *FakeVappV1alpha1
	ns   string
}

var entitiesResource = schema.GroupVersionResource{Group: "vapp.vsphere.kubeform.com", Version: "v1alpha1", Resource: "entities"}

var entitiesKind = schema.GroupVersionKind{Group: "vapp.vsphere.kubeform.com", Version: "v1alpha1", Kind: "Entity"}

// Get takes name of the entity, and returns the corresponding entity object, and an error if there is any.
func (c *FakeEntities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(entitiesResource, c.ns, name), &v1alpha1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Entity), err
}

// List takes label and field selectors, and returns the list of Entities that match those selectors.
func (c *FakeEntities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EntityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(entitiesResource, entitiesKind, c.ns, opts), &v1alpha1.EntityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EntityList{ListMeta: obj.(*v1alpha1.EntityList).ListMeta}
	for _, item := range obj.(*v1alpha1.EntityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested entities.
func (c *FakeEntities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(entitiesResource, c.ns, opts))

}

// Create takes the representation of a entity and creates it.  Returns the server's representation of the entity, and an error, if there is any.
func (c *FakeEntities) Create(ctx context.Context, entity *v1alpha1.Entity, opts v1.CreateOptions) (result *v1alpha1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(entitiesResource, c.ns, entity), &v1alpha1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Entity), err
}

// Update takes the representation of a entity and updates it. Returns the server's representation of the entity, and an error, if there is any.
func (c *FakeEntities) Update(ctx context.Context, entity *v1alpha1.Entity, opts v1.UpdateOptions) (result *v1alpha1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(entitiesResource, c.ns, entity), &v1alpha1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Entity), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEntities) UpdateStatus(ctx context.Context, entity *v1alpha1.Entity, opts v1.UpdateOptions) (*v1alpha1.Entity, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(entitiesResource, "status", c.ns, entity), &v1alpha1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Entity), err
}

// Delete takes name of the entity and deletes it. Returns an error if one occurs.
func (c *FakeEntities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(entitiesResource, c.ns, name), &v1alpha1.Entity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEntities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(entitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EntityList{})
	return err
}

// Patch applies the patch and returns the patched entity.
func (c *FakeEntities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entitiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Entity), err
}
