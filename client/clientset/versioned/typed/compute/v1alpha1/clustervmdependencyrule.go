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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/compute/v1alpha1"
	scheme "kubeform.dev/provider-vsphere-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterVmDependencyRulesGetter has a method to return a ClusterVmDependencyRuleInterface.
// A group's client should implement this interface.
type ClusterVmDependencyRulesGetter interface {
	ClusterVmDependencyRules(namespace string) ClusterVmDependencyRuleInterface
}

// ClusterVmDependencyRuleInterface has methods to work with ClusterVmDependencyRule resources.
type ClusterVmDependencyRuleInterface interface {
	Create(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.CreateOptions) (*v1alpha1.ClusterVmDependencyRule, error)
	Update(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.UpdateOptions) (*v1alpha1.ClusterVmDependencyRule, error)
	UpdateStatus(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.UpdateOptions) (*v1alpha1.ClusterVmDependencyRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterVmDependencyRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterVmDependencyRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterVmDependencyRule, err error)
	ClusterVmDependencyRuleExpansion
}

// clusterVmDependencyRules implements ClusterVmDependencyRuleInterface
type clusterVmDependencyRules struct {
	client rest.Interface
	ns     string
}

// newClusterVmDependencyRules returns a ClusterVmDependencyRules
func newClusterVmDependencyRules(c *ComputeV1alpha1Client, namespace string) *clusterVmDependencyRules {
	return &clusterVmDependencyRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterVmDependencyRule, and returns the corresponding clusterVmDependencyRule object, and an error if there is any.
func (c *clusterVmDependencyRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterVmDependencyRule, err error) {
	result = &v1alpha1.ClusterVmDependencyRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterVmDependencyRules that match those selectors.
func (c *clusterVmDependencyRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterVmDependencyRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterVmDependencyRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterVmDependencyRules.
func (c *clusterVmDependencyRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterVmDependencyRule and creates it.  Returns the server's representation of the clusterVmDependencyRule, and an error, if there is any.
func (c *clusterVmDependencyRules) Create(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.CreateOptions) (result *v1alpha1.ClusterVmDependencyRule, err error) {
	result = &v1alpha1.ClusterVmDependencyRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterVmDependencyRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterVmDependencyRule and updates it. Returns the server's representation of the clusterVmDependencyRule, and an error, if there is any.
func (c *clusterVmDependencyRules) Update(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.UpdateOptions) (result *v1alpha1.ClusterVmDependencyRule, err error) {
	result = &v1alpha1.ClusterVmDependencyRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		Name(clusterVmDependencyRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterVmDependencyRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterVmDependencyRules) UpdateStatus(ctx context.Context, clusterVmDependencyRule *v1alpha1.ClusterVmDependencyRule, opts v1.UpdateOptions) (result *v1alpha1.ClusterVmDependencyRule, err error) {
	result = &v1alpha1.ClusterVmDependencyRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		Name(clusterVmDependencyRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterVmDependencyRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterVmDependencyRule and deletes it. Returns an error if one occurs.
func (c *clusterVmDependencyRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterVmDependencyRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterVmDependencyRule.
func (c *clusterVmDependencyRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterVmDependencyRule, err error) {
	result = &v1alpha1.ClusterVmDependencyRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustervmdependencyrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
