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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/ha/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VmOverrideLister helps list VmOverrides.
// All objects returned here must be treated as read-only.
type VmOverrideLister interface {
	// List lists all VmOverrides in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VmOverride, err error)
	// VmOverrides returns an object that can list and get VmOverrides.
	VmOverrides(namespace string) VmOverrideNamespaceLister
	VmOverrideListerExpansion
}

// vmOverrideLister implements the VmOverrideLister interface.
type vmOverrideLister struct {
	indexer cache.Indexer
}

// NewVmOverrideLister returns a new VmOverrideLister.
func NewVmOverrideLister(indexer cache.Indexer) VmOverrideLister {
	return &vmOverrideLister{indexer: indexer}
}

// List lists all VmOverrides in the indexer.
func (s *vmOverrideLister) List(selector labels.Selector) (ret []*v1alpha1.VmOverride, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VmOverride))
	})
	return ret, err
}

// VmOverrides returns an object that can list and get VmOverrides.
func (s *vmOverrideLister) VmOverrides(namespace string) VmOverrideNamespaceLister {
	return vmOverrideNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VmOverrideNamespaceLister helps list and get VmOverrides.
// All objects returned here must be treated as read-only.
type VmOverrideNamespaceLister interface {
	// List lists all VmOverrides in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VmOverride, err error)
	// Get retrieves the VmOverride from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VmOverride, error)
	VmOverrideNamespaceListerExpansion
}

// vmOverrideNamespaceLister implements the VmOverrideNamespaceLister
// interface.
type vmOverrideNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VmOverrides in the indexer for a given namespace.
func (s vmOverrideNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VmOverride, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VmOverride))
	})
	return ret, err
}

// Get retrieves the VmOverride from the indexer for a given namespace and name.
func (s vmOverrideNamespaceLister) Get(name string) (*v1alpha1.VmOverride, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vmoverride"), name)
	}
	return obj.(*v1alpha1.VmOverride), nil
}
