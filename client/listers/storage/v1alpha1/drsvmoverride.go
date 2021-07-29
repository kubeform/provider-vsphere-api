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
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/storage/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DrsVmOverrideLister helps list DrsVmOverrides.
// All objects returned here must be treated as read-only.
type DrsVmOverrideLister interface {
	// List lists all DrsVmOverrides in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DrsVmOverride, err error)
	// DrsVmOverrides returns an object that can list and get DrsVmOverrides.
	DrsVmOverrides(namespace string) DrsVmOverrideNamespaceLister
	DrsVmOverrideListerExpansion
}

// drsVmOverrideLister implements the DrsVmOverrideLister interface.
type drsVmOverrideLister struct {
	indexer cache.Indexer
}

// NewDrsVmOverrideLister returns a new DrsVmOverrideLister.
func NewDrsVmOverrideLister(indexer cache.Indexer) DrsVmOverrideLister {
	return &drsVmOverrideLister{indexer: indexer}
}

// List lists all DrsVmOverrides in the indexer.
func (s *drsVmOverrideLister) List(selector labels.Selector) (ret []*v1alpha1.DrsVmOverride, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DrsVmOverride))
	})
	return ret, err
}

// DrsVmOverrides returns an object that can list and get DrsVmOverrides.
func (s *drsVmOverrideLister) DrsVmOverrides(namespace string) DrsVmOverrideNamespaceLister {
	return drsVmOverrideNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DrsVmOverrideNamespaceLister helps list and get DrsVmOverrides.
// All objects returned here must be treated as read-only.
type DrsVmOverrideNamespaceLister interface {
	// List lists all DrsVmOverrides in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DrsVmOverride, err error)
	// Get retrieves the DrsVmOverride from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DrsVmOverride, error)
	DrsVmOverrideNamespaceListerExpansion
}

// drsVmOverrideNamespaceLister implements the DrsVmOverrideNamespaceLister
// interface.
type drsVmOverrideNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DrsVmOverrides in the indexer for a given namespace.
func (s drsVmOverrideNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DrsVmOverride, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DrsVmOverride))
	})
	return ret, err
}

// Get retrieves the DrsVmOverride from the indexer for a given namespace and name.
func (s drsVmOverrideNamespaceLister) Get(name string) (*v1alpha1.DrsVmOverride, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("drsvmoverride"), name)
	}
	return obj.(*v1alpha1.DrsVmOverride), nil
}