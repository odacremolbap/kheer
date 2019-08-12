/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	scheme "github.com/kheer/kheer/pkg/crd/api/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/kheer/kheer/pkg/crd/api/kheer/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperationsGetter has a method to return a OperationInterface.
// A group's client should implement this interface.
type OperationsGetter interface {
	Operations(namespace string) OperationInterface
}

// OperationInterface has methods to work with Operation resources.
type OperationInterface interface {
	Create(*v1alpha1.Operation) (*v1alpha1.Operation, error)
	Update(*v1alpha1.Operation) (*v1alpha1.Operation, error)
	UpdateStatus(*v1alpha1.Operation) (*v1alpha1.Operation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Operation, error)
	List(opts v1.ListOptions) (*v1alpha1.OperationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Operation, err error)
	OperationExpansion
}

// operations implements OperationInterface
type operations struct {
	client rest.Interface
	ns     string
}

// newOperations returns a Operations
func newOperations(c *KheerV1alpha1Client, namespace string) *operations {
	return &operations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operation, and returns the corresponding operation object, and an error if there is any.
func (c *operations) Get(name string, options v1.GetOptions) (result *v1alpha1.Operation, err error) {
	result = &v1alpha1.Operation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Operations that match those selectors.
func (c *operations) List(opts v1.ListOptions) (result *v1alpha1.OperationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OperationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operations.
func (c *operations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a operation and creates it.  Returns the server's representation of the operation, and an error, if there is any.
func (c *operations) Create(operation *v1alpha1.Operation) (result *v1alpha1.Operation, err error) {
	result = &v1alpha1.Operation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operations").
		Body(operation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a operation and updates it. Returns the server's representation of the operation, and an error, if there is any.
func (c *operations) Update(operation *v1alpha1.Operation) (result *v1alpha1.Operation, err error) {
	result = &v1alpha1.Operation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operations").
		Name(operation.Name).
		Body(operation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *operations) UpdateStatus(operation *v1alpha1.Operation) (result *v1alpha1.Operation, err error) {
	result = &v1alpha1.Operation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operations").
		Name(operation.Name).
		SubResource("status").
		Body(operation).
		Do().
		Into(result)
	return
}

// Delete takes name of the operation and deletes it. Returns an error if one occurs.
func (c *operations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched operation.
func (c *operations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Operation, err error) {
	result = &v1alpha1.Operation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}