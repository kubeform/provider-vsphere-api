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

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/entity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePermissionses implements PermissionsInterface
type FakePermissionses struct {
	Fake *FakeEntityV1alpha1
	ns   string
}

var permissionsesResource = schema.GroupVersionResource{Group: "entity.vsphere.kubeform.com", Version: "v1alpha1", Resource: "permissionses"}

var permissionsesKind = schema.GroupVersionKind{Group: "entity.vsphere.kubeform.com", Version: "v1alpha1", Kind: "Permissions"}

// Get takes name of the permissions, and returns the corresponding permissions object, and an error if there is any.
func (c *FakePermissionses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Permissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(permissionsesResource, c.ns, name), &v1alpha1.Permissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Permissions), err
}

// List takes label and field selectors, and returns the list of Permissionses that match those selectors.
func (c *FakePermissionses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PermissionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(permissionsesResource, permissionsesKind, c.ns, opts), &v1alpha1.PermissionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PermissionsList{ListMeta: obj.(*v1alpha1.PermissionsList).ListMeta}
	for _, item := range obj.(*v1alpha1.PermissionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested permissionses.
func (c *FakePermissionses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(permissionsesResource, c.ns, opts))

}

// Create takes the representation of a permissions and creates it.  Returns the server's representation of the permissions, and an error, if there is any.
func (c *FakePermissionses) Create(ctx context.Context, permissions *v1alpha1.Permissions, opts v1.CreateOptions) (result *v1alpha1.Permissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(permissionsesResource, c.ns, permissions), &v1alpha1.Permissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Permissions), err
}

// Update takes the representation of a permissions and updates it. Returns the server's representation of the permissions, and an error, if there is any.
func (c *FakePermissionses) Update(ctx context.Context, permissions *v1alpha1.Permissions, opts v1.UpdateOptions) (result *v1alpha1.Permissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(permissionsesResource, c.ns, permissions), &v1alpha1.Permissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Permissions), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePermissionses) UpdateStatus(ctx context.Context, permissions *v1alpha1.Permissions, opts v1.UpdateOptions) (*v1alpha1.Permissions, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(permissionsesResource, "status", c.ns, permissions), &v1alpha1.Permissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Permissions), err
}

// Delete takes name of the permissions and deletes it. Returns an error if one occurs.
func (c *FakePermissionses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(permissionsesResource, c.ns, name), &v1alpha1.Permissions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePermissionses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(permissionsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PermissionsList{})
	return err
}

// Patch applies the patch and returns the patched permissions.
func (c *FakePermissionses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Permissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(permissionsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Permissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Permissions), err
}
