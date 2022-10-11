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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// FakeAPIServiceExportResources implements APIServiceExportResourceInterface
type FakeAPIServiceExportResources struct {
	Fake *FakeKubeBindV1alpha1
	ns   string
}

var apiserviceexportresourcesResource = schema.GroupVersionResource{Group: "kube-bind.io", Version: "v1alpha1", Resource: "apiserviceexportresources"}

var apiserviceexportresourcesKind = schema.GroupVersionKind{Group: "kube-bind.io", Version: "v1alpha1", Kind: "APIServiceExportResource"}

// Get takes name of the aPIServiceExportResource, and returns the corresponding aPIServiceExportResource object, and an error if there is any.
func (c *FakeAPIServiceExportResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIServiceExportResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiserviceexportresourcesResource, c.ns, name), &v1alpha1.APIServiceExportResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportResource), err
}

// List takes label and field selectors, and returns the list of APIServiceExportResources that match those selectors.
func (c *FakeAPIServiceExportResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIServiceExportResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiserviceexportresourcesResource, apiserviceexportresourcesKind, c.ns, opts), &v1alpha1.APIServiceExportResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIServiceExportResourceList{ListMeta: obj.(*v1alpha1.APIServiceExportResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIServiceExportResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIServiceExportResources.
func (c *FakeAPIServiceExportResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiserviceexportresourcesResource, c.ns, opts))

}

// Create takes the representation of a aPIServiceExportResource and creates it.  Returns the server's representation of the aPIServiceExportResource, and an error, if there is any.
func (c *FakeAPIServiceExportResources) Create(ctx context.Context, aPIServiceExportResource *v1alpha1.APIServiceExportResource, opts v1.CreateOptions) (result *v1alpha1.APIServiceExportResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiserviceexportresourcesResource, c.ns, aPIServiceExportResource), &v1alpha1.APIServiceExportResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportResource), err
}

// Update takes the representation of a aPIServiceExportResource and updates it. Returns the server's representation of the aPIServiceExportResource, and an error, if there is any.
func (c *FakeAPIServiceExportResources) Update(ctx context.Context, aPIServiceExportResource *v1alpha1.APIServiceExportResource, opts v1.UpdateOptions) (result *v1alpha1.APIServiceExportResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiserviceexportresourcesResource, c.ns, aPIServiceExportResource), &v1alpha1.APIServiceExportResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIServiceExportResources) UpdateStatus(ctx context.Context, aPIServiceExportResource *v1alpha1.APIServiceExportResource, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiserviceexportresourcesResource, "status", c.ns, aPIServiceExportResource), &v1alpha1.APIServiceExportResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportResource), err
}

// Delete takes name of the aPIServiceExportResource and deletes it. Returns an error if one occurs.
func (c *FakeAPIServiceExportResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apiserviceexportresourcesResource, c.ns, name, opts), &v1alpha1.APIServiceExportResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIServiceExportResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiserviceexportresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIServiceExportResourceList{})
	return err
}

// Patch applies the patch and returns the patched aPIServiceExportResource.
func (c *FakeAPIServiceExportResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceExportResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiserviceexportresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.APIServiceExportResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportResource), err
}