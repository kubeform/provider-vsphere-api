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

	tagv1alpha1 "kubeform.dev/provider-vsphere-api/apis/tag/v1alpha1"
	versioned "kubeform.dev/provider-vsphere-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-vsphere-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-vsphere-api/client/listers/tag/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CategoryInformer provides access to a shared informer and lister for
// Categories.
type CategoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CategoryLister
}

type categoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCategoryInformer constructs a new informer for Category type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCategoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCategoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCategoryInformer constructs a new informer for Category type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCategoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TagV1alpha1().Categories(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TagV1alpha1().Categories(namespace).Watch(context.TODO(), options)
			},
		},
		&tagv1alpha1.Category{},
		resyncPeriod,
		indexers,
	)
}

func (f *categoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCategoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *categoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tagv1alpha1.Category{}, f.defaultInformer)
}

func (f *categoryInformer) Lister() v1alpha1.CategoryLister {
	return v1alpha1.NewCategoryLister(f.Informer().GetIndexer())
}
