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
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/compute/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterVmGroupLister helps list ClusterVmGroups.
// All objects returned here must be treated as read-only.
type ClusterVmGroupLister interface {
	// List lists all ClusterVmGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterVmGroup, err error)
	// ClusterVmGroups returns an object that can list and get ClusterVmGroups.
	ClusterVmGroups(namespace string) ClusterVmGroupNamespaceLister
	ClusterVmGroupListerExpansion
}

// clusterVmGroupLister implements the ClusterVmGroupLister interface.
type clusterVmGroupLister struct {
	indexer cache.Indexer
}

// NewClusterVmGroupLister returns a new ClusterVmGroupLister.
func NewClusterVmGroupLister(indexer cache.Indexer) ClusterVmGroupLister {
	return &clusterVmGroupLister{indexer: indexer}
}

// List lists all ClusterVmGroups in the indexer.
func (s *clusterVmGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterVmGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterVmGroup))
	})
	return ret, err
}

// ClusterVmGroups returns an object that can list and get ClusterVmGroups.
func (s *clusterVmGroupLister) ClusterVmGroups(namespace string) ClusterVmGroupNamespaceLister {
	return clusterVmGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterVmGroupNamespaceLister helps list and get ClusterVmGroups.
// All objects returned here must be treated as read-only.
type ClusterVmGroupNamespaceLister interface {
	// List lists all ClusterVmGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterVmGroup, err error)
	// Get retrieves the ClusterVmGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterVmGroup, error)
	ClusterVmGroupNamespaceListerExpansion
}

// clusterVmGroupNamespaceLister implements the ClusterVmGroupNamespaceLister
// interface.
type clusterVmGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterVmGroups in the indexer for a given namespace.
func (s clusterVmGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterVmGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterVmGroup))
	})
	return ret, err
}

// Get retrieves the ClusterVmGroup from the indexer for a given namespace and name.
func (s clusterVmGroupNamespaceLister) Get(name string) (*v1alpha1.ClusterVmGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustervmgroup"), name)
	}
	return obj.(*v1alpha1.ClusterVmGroup), nil
}
