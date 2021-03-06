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
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/distributed/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PortGroupLister helps list PortGroups.
// All objects returned here must be treated as read-only.
type PortGroupLister interface {
	// List lists all PortGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PortGroup, err error)
	// PortGroups returns an object that can list and get PortGroups.
	PortGroups(namespace string) PortGroupNamespaceLister
	PortGroupListerExpansion
}

// portGroupLister implements the PortGroupLister interface.
type portGroupLister struct {
	indexer cache.Indexer
}

// NewPortGroupLister returns a new PortGroupLister.
func NewPortGroupLister(indexer cache.Indexer) PortGroupLister {
	return &portGroupLister{indexer: indexer}
}

// List lists all PortGroups in the indexer.
func (s *portGroupLister) List(selector labels.Selector) (ret []*v1alpha1.PortGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PortGroup))
	})
	return ret, err
}

// PortGroups returns an object that can list and get PortGroups.
func (s *portGroupLister) PortGroups(namespace string) PortGroupNamespaceLister {
	return portGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PortGroupNamespaceLister helps list and get PortGroups.
// All objects returned here must be treated as read-only.
type PortGroupNamespaceLister interface {
	// List lists all PortGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PortGroup, err error)
	// Get retrieves the PortGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PortGroup, error)
	PortGroupNamespaceListerExpansion
}

// portGroupNamespaceLister implements the PortGroupNamespaceLister
// interface.
type portGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PortGroups in the indexer for a given namespace.
func (s portGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PortGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PortGroup))
	})
	return ret, err
}

// Get retrieves the PortGroup from the indexer for a given namespace and name.
func (s portGroupNamespaceLister) Get(name string) (*v1alpha1.PortGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("portgroup"), name)
	}
	return obj.(*v1alpha1.PortGroup), nil
}
