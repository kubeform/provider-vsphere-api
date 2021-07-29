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

	licensev1alpha1 "kubeform.dev/provider-vsphere-api/apis/license/v1alpha1"
	versioned "kubeform.dev/provider-vsphere-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-vsphere-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-vsphere-api/client/listers/license/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LicenseInformer provides access to a shared informer and lister for
// Licenses.
type LicenseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LicenseLister
}

type licenseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLicenseInformer constructs a new informer for License type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLicenseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLicenseInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLicenseInformer constructs a new informer for License type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLicenseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LicenseV1alpha1().Licenses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LicenseV1alpha1().Licenses(namespace).Watch(context.TODO(), options)
			},
		},
		&licensev1alpha1.License{},
		resyncPeriod,
		indexers,
	)
}

func (f *licenseInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLicenseInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *licenseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&licensev1alpha1.License{}, f.defaultInformer)
}

func (f *licenseInformer) Lister() v1alpha1.LicenseLister {
	return v1alpha1.NewLicenseLister(f.Informer().GetIndexer())
}
