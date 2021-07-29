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

// ClusterVmAffinityRuleLister helps list ClusterVmAffinityRules.
// All objects returned here must be treated as read-only.
type ClusterVmAffinityRuleLister interface {
	// List lists all ClusterVmAffinityRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterVmAffinityRule, err error)
	// ClusterVmAffinityRules returns an object that can list and get ClusterVmAffinityRules.
	ClusterVmAffinityRules(namespace string) ClusterVmAffinityRuleNamespaceLister
	ClusterVmAffinityRuleListerExpansion
}

// clusterVmAffinityRuleLister implements the ClusterVmAffinityRuleLister interface.
type clusterVmAffinityRuleLister struct {
	indexer cache.Indexer
}

// NewClusterVmAffinityRuleLister returns a new ClusterVmAffinityRuleLister.
func NewClusterVmAffinityRuleLister(indexer cache.Indexer) ClusterVmAffinityRuleLister {
	return &clusterVmAffinityRuleLister{indexer: indexer}
}

// List lists all ClusterVmAffinityRules in the indexer.
func (s *clusterVmAffinityRuleLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterVmAffinityRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterVmAffinityRule))
	})
	return ret, err
}

// ClusterVmAffinityRules returns an object that can list and get ClusterVmAffinityRules.
func (s *clusterVmAffinityRuleLister) ClusterVmAffinityRules(namespace string) ClusterVmAffinityRuleNamespaceLister {
	return clusterVmAffinityRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterVmAffinityRuleNamespaceLister helps list and get ClusterVmAffinityRules.
// All objects returned here must be treated as read-only.
type ClusterVmAffinityRuleNamespaceLister interface {
	// List lists all ClusterVmAffinityRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterVmAffinityRule, err error)
	// Get retrieves the ClusterVmAffinityRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterVmAffinityRule, error)
	ClusterVmAffinityRuleNamespaceListerExpansion
}

// clusterVmAffinityRuleNamespaceLister implements the ClusterVmAffinityRuleNamespaceLister
// interface.
type clusterVmAffinityRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterVmAffinityRules in the indexer for a given namespace.
func (s clusterVmAffinityRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterVmAffinityRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterVmAffinityRule))
	})
	return ret, err
}

// Get retrieves the ClusterVmAffinityRule from the indexer for a given namespace and name.
func (s clusterVmAffinityRuleNamespaceLister) Get(name string) (*v1alpha1.ClusterVmAffinityRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustervmaffinityrule"), name)
	}
	return obj.(*v1alpha1.ClusterVmAffinityRule), nil
}
