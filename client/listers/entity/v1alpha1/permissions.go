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
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/entity/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PermissionsLister helps list Permissionses.
// All objects returned here must be treated as read-only.
type PermissionsLister interface {
	// List lists all Permissionses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Permissions, err error)
	// Permissionses returns an object that can list and get Permissionses.
	Permissionses(namespace string) PermissionsNamespaceLister
	PermissionsListerExpansion
}

// permissionsLister implements the PermissionsLister interface.
type permissionsLister struct {
	indexer cache.Indexer
}

// NewPermissionsLister returns a new PermissionsLister.
func NewPermissionsLister(indexer cache.Indexer) PermissionsLister {
	return &permissionsLister{indexer: indexer}
}

// List lists all Permissionses in the indexer.
func (s *permissionsLister) List(selector labels.Selector) (ret []*v1alpha1.Permissions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Permissions))
	})
	return ret, err
}

// Permissionses returns an object that can list and get Permissionses.
func (s *permissionsLister) Permissionses(namespace string) PermissionsNamespaceLister {
	return permissionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PermissionsNamespaceLister helps list and get Permissionses.
// All objects returned here must be treated as read-only.
type PermissionsNamespaceLister interface {
	// List lists all Permissionses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Permissions, err error)
	// Get retrieves the Permissions from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Permissions, error)
	PermissionsNamespaceListerExpansion
}

// permissionsNamespaceLister implements the PermissionsNamespaceLister
// interface.
type permissionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Permissionses in the indexer for a given namespace.
func (s permissionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Permissions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Permissions))
	})
	return ret, err
}

// Get retrieves the Permissions from the indexer for a given namespace and name.
func (s permissionsNamespaceLister) Get(name string) (*v1alpha1.Permissions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("permissions"), name)
	}
	return obj.(*v1alpha1.Permissions), nil
}