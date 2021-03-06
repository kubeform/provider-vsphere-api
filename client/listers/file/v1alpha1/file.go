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
	v1alpha1 "kubeform.dev/provider-vsphere-api/apis/file/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FileLister helps list Files.
// All objects returned here must be treated as read-only.
type FileLister interface {
	// List lists all Files in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.File, err error)
	// Files returns an object that can list and get Files.
	Files(namespace string) FileNamespaceLister
	FileListerExpansion
}

// fileLister implements the FileLister interface.
type fileLister struct {
	indexer cache.Indexer
}

// NewFileLister returns a new FileLister.
func NewFileLister(indexer cache.Indexer) FileLister {
	return &fileLister{indexer: indexer}
}

// List lists all Files in the indexer.
func (s *fileLister) List(selector labels.Selector) (ret []*v1alpha1.File, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.File))
	})
	return ret, err
}

// Files returns an object that can list and get Files.
func (s *fileLister) Files(namespace string) FileNamespaceLister {
	return fileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FileNamespaceLister helps list and get Files.
// All objects returned here must be treated as read-only.
type FileNamespaceLister interface {
	// List lists all Files in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.File, err error)
	// Get retrieves the File from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.File, error)
	FileNamespaceListerExpansion
}

// fileNamespaceLister implements the FileNamespaceLister
// interface.
type fileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Files in the indexer for a given namespace.
func (s fileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.File, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.File))
	})
	return ret, err
}

// Get retrieves the File from the indexer for a given namespace and name.
func (s fileNamespaceLister) Get(name string) (*v1alpha1.File, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("file"), name)
	}
	return obj.(*v1alpha1.File), nil
}
