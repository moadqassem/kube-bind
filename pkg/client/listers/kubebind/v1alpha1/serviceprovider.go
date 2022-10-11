/*
Copyright The Kube Bind Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// ServiceProviderLister helps list ServiceProviders.
// All objects returned here must be treated as read-only.
type ServiceProviderLister interface {
	// List lists all ServiceProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceProvider, err error)
	// ServiceProviders returns an object that can list and get ServiceProviders.
	ServiceProviders(namespace string) ServiceProviderNamespaceLister
	ServiceProviderListerExpansion
}

// serviceProviderLister implements the ServiceProviderLister interface.
type serviceProviderLister struct {
	indexer cache.Indexer
}

// NewServiceProviderLister returns a new ServiceProviderLister.
func NewServiceProviderLister(indexer cache.Indexer) ServiceProviderLister {
	return &serviceProviderLister{indexer: indexer}
}

// List lists all ServiceProviders in the indexer.
func (s *serviceProviderLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceProvider))
	})
	return ret, err
}

// ServiceProviders returns an object that can list and get ServiceProviders.
func (s *serviceProviderLister) ServiceProviders(namespace string) ServiceProviderNamespaceLister {
	return serviceProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceProviderNamespaceLister helps list and get ServiceProviders.
// All objects returned here must be treated as read-only.
type ServiceProviderNamespaceLister interface {
	// List lists all ServiceProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceProvider, err error)
	// Get retrieves the ServiceProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceProvider, error)
	ServiceProviderNamespaceListerExpansion
}

// serviceProviderNamespaceLister implements the ServiceProviderNamespaceLister
// interface.
type serviceProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceProviders in the indexer for a given namespace.
func (s serviceProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceProvider))
	})
	return ret, err
}

// Get retrieves the ServiceProvider from the indexer for a given namespace and name.
func (s serviceProviderNamespaceLister) Get(name string) (*v1alpha1.ServiceProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceprovider"), name)
	}
	return obj.(*v1alpha1.ServiceProvider), nil
}