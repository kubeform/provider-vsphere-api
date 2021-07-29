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

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/compute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterVmAffinityRules implements ClusterVmAffinityRuleInterface
type FakeClusterVmAffinityRules struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var clustervmaffinityrulesResource = schema.GroupVersionResource{Group: "compute.vsphere.kubeform.com", Version: "v1alpha1", Resource: "clustervmaffinityrules"}

var clustervmaffinityrulesKind = schema.GroupVersionKind{Group: "compute.vsphere.kubeform.com", Version: "v1alpha1", Kind: "ClusterVmAffinityRule"}

// Get takes name of the clusterVmAffinityRule, and returns the corresponding clusterVmAffinityRule object, and an error if there is any.
func (c *FakeClusterVmAffinityRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterVmAffinityRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustervmaffinityrulesResource, c.ns, name), &v1alpha1.ClusterVmAffinityRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), err
}

// List takes label and field selectors, and returns the list of ClusterVmAffinityRules that match those selectors.
func (c *FakeClusterVmAffinityRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterVmAffinityRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustervmaffinityrulesResource, clustervmaffinityrulesKind, c.ns, opts), &v1alpha1.ClusterVmAffinityRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterVmAffinityRuleList{ListMeta: obj.(*v1alpha1.ClusterVmAffinityRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterVmAffinityRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterVmAffinityRules.
func (c *FakeClusterVmAffinityRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustervmaffinityrulesResource, c.ns, opts))

}

// Create takes the representation of a clusterVmAffinityRule and creates it.  Returns the server's representation of the clusterVmAffinityRule, and an error, if there is any.
func (c *FakeClusterVmAffinityRules) Create(ctx context.Context, clusterVmAffinityRule *v1alpha1.ClusterVmAffinityRule, opts v1.CreateOptions) (result *v1alpha1.ClusterVmAffinityRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustervmaffinityrulesResource, c.ns, clusterVmAffinityRule), &v1alpha1.ClusterVmAffinityRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), err
}

// Update takes the representation of a clusterVmAffinityRule and updates it. Returns the server's representation of the clusterVmAffinityRule, and an error, if there is any.
func (c *FakeClusterVmAffinityRules) Update(ctx context.Context, clusterVmAffinityRule *v1alpha1.ClusterVmAffinityRule, opts v1.UpdateOptions) (result *v1alpha1.ClusterVmAffinityRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustervmaffinityrulesResource, c.ns, clusterVmAffinityRule), &v1alpha1.ClusterVmAffinityRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterVmAffinityRules) UpdateStatus(ctx context.Context, clusterVmAffinityRule *v1alpha1.ClusterVmAffinityRule, opts v1.UpdateOptions) (*v1alpha1.ClusterVmAffinityRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustervmaffinityrulesResource, "status", c.ns, clusterVmAffinityRule), &v1alpha1.ClusterVmAffinityRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), err
}

// Delete takes name of the clusterVmAffinityRule and deletes it. Returns an error if one occurs.
func (c *FakeClusterVmAffinityRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustervmaffinityrulesResource, c.ns, name), &v1alpha1.ClusterVmAffinityRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterVmAffinityRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustervmaffinityrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterVmAffinityRuleList{})
	return err
}

// Patch applies the patch and returns the patched clusterVmAffinityRule.
func (c *FakeClusterVmAffinityRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterVmAffinityRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustervmaffinityrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterVmAffinityRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), err
}