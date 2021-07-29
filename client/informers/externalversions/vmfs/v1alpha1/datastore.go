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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	vmfsv1alpha1 "kubeform.dev/provider-vsphere-api/apis/vmfs/v1alpha1"
	versioned "kubeform.dev/provider-vsphere-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-vsphere-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-vsphere-api/client/listers/vmfs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatastoreInformer provides access to a shared informer and lister for
// Datastores.
type DatastoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatastoreLister
}

type datastoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatastoreInformer constructs a new informer for Datastore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatastoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatastoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatastoreInformer constructs a new informer for Datastore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatastoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VmfsV1alpha1().Datastores(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VmfsV1alpha1().Datastores(namespace).Watch(context.TODO(), options)
			},
		},
		&vmfsv1alpha1.Datastore{},
		resyncPeriod,
		indexers,
	)
}

func (f *datastoreInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatastoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *datastoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&vmfsv1alpha1.Datastore{}, f.defaultInformer)
}

func (f *datastoreInformer) Lister() v1alpha1.DatastoreLister {
	return v1alpha1.NewDatastoreLister(f.Informer().GetIndexer())
}
