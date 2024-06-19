/*
Copyright 2022.

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
	"context"
	"time"

	v1alpha1 "github.com/intelligent-machine-learning/easydl/dlrover/go/operator/api/elastic.iml.github.io/v1alpha1"
	scheme "github.com/intelligent-machine-learning/easydl/dlrover/go/operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScalePlansGetter has a method to return a ScalePlanInterface.
// A group's client should implement this interface.
type ScalePlansGetter interface {
	ScalePlans(namespace string) ScalePlanInterface
}

// ScalePlanInterface has methods to work with ScalePlan resources.
type ScalePlanInterface interface {
	Create(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.CreateOptions) (*v1alpha1.ScalePlan, error)
	Update(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.UpdateOptions) (*v1alpha1.ScalePlan, error)
	UpdateStatus(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.UpdateOptions) (*v1alpha1.ScalePlan, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ScalePlan, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ScalePlanList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScalePlan, err error)
	ScalePlanExpansion
}

// scalePlans implements ScalePlanInterface
type scalePlans struct {
	client rest.Interface
	ns     string
}

// newScalePlans returns a ScalePlans
func newScalePlans(c *ElasticV1alpha1Client, namespace string) *scalePlans {
	return &scalePlans{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scalePlan, and returns the corresponding scalePlan object, and an error if there is any.
func (c *scalePlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScalePlan, err error) {
	result = &v1alpha1.ScalePlan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scaleplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScalePlans that match those selectors.
func (c *scalePlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScalePlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ScalePlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scaleplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scalePlans.
func (c *scalePlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scaleplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a scalePlan and creates it.  Returns the server's representation of the scalePlan, and an error, if there is any.
func (c *scalePlans) Create(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.CreateOptions) (result *v1alpha1.ScalePlan, err error) {
	result = &v1alpha1.ScalePlan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scaleplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scalePlan).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a scalePlan and updates it. Returns the server's representation of the scalePlan, and an error, if there is any.
func (c *scalePlans) Update(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.UpdateOptions) (result *v1alpha1.ScalePlan, err error) {
	result = &v1alpha1.ScalePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scaleplans").
		Name(scalePlan.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scalePlan).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *scalePlans) UpdateStatus(ctx context.Context, scalePlan *v1alpha1.ScalePlan, opts v1.UpdateOptions) (result *v1alpha1.ScalePlan, err error) {
	result = &v1alpha1.ScalePlan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scaleplans").
		Name(scalePlan.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scalePlan).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the scalePlan and deletes it. Returns an error if one occurs.
func (c *scalePlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scaleplans").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scalePlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scaleplans").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched scalePlan.
func (c *scalePlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScalePlan, err error) {
	result = &v1alpha1.ScalePlan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scaleplans").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
