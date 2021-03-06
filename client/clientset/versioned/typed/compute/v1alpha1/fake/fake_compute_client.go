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
	v1alpha1 "kubeform.dev/provider-vsphere-api/client/clientset/versioned/typed/compute/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeComputeV1alpha1 struct {
	*testing.Fake
}

func (c *FakeComputeV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterHostGroups(namespace string) v1alpha1.ClusterHostGroupInterface {
	return &FakeClusterHostGroups{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterVmAffinityRules(namespace string) v1alpha1.ClusterVmAffinityRuleInterface {
	return &FakeClusterVmAffinityRules{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterVmAntiAffinityRules(namespace string) v1alpha1.ClusterVmAntiAffinityRuleInterface {
	return &FakeClusterVmAntiAffinityRules{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterVmDependencyRules(namespace string) v1alpha1.ClusterVmDependencyRuleInterface {
	return &FakeClusterVmDependencyRules{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterVmGroups(namespace string) v1alpha1.ClusterVmGroupInterface {
	return &FakeClusterVmGroups{c, namespace}
}

func (c *FakeComputeV1alpha1) ClusterVmHostRules(namespace string) v1alpha1.ClusterVmHostRuleInterface {
	return &FakeClusterVmHostRules{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeComputeV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
