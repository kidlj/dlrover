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

package fake

import (
	"context"

	v1alpha1 "github.com/intelligent-machine-learning/easydl/dlrover/go/operator/api/elastic.iml.github.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticJobs implements ElasticJobInterface
type FakeElasticJobs struct {
	Fake *FakeElasticV1alpha1
	ns   string
}

var elasticjobsResource = v1alpha1.SchemeGroupVersion.WithResource("elasticjobs")

var elasticjobsKind = v1alpha1.SchemeGroupVersion.WithKind("ElasticJob")

// Get takes name of the elasticJob, and returns the corresponding elasticJob object, and an error if there is any.
func (c *FakeElasticJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ElasticJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticjobsResource, c.ns, name), &v1alpha1.ElasticJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticJob), err
}

// List takes label and field selectors, and returns the list of ElasticJobs that match those selectors.
func (c *FakeElasticJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ElasticJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticjobsResource, elasticjobsKind, c.ns, opts), &v1alpha1.ElasticJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticJobList{ListMeta: obj.(*v1alpha1.ElasticJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElasticJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticJobs.
func (c *FakeElasticJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticjobsResource, c.ns, opts))

}

// Create takes the representation of a elasticJob and creates it.  Returns the server's representation of the elasticJob, and an error, if there is any.
func (c *FakeElasticJobs) Create(ctx context.Context, elasticJob *v1alpha1.ElasticJob, opts v1.CreateOptions) (result *v1alpha1.ElasticJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticjobsResource, c.ns, elasticJob), &v1alpha1.ElasticJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticJob), err
}

// Update takes the representation of a elasticJob and updates it. Returns the server's representation of the elasticJob, and an error, if there is any.
func (c *FakeElasticJobs) Update(ctx context.Context, elasticJob *v1alpha1.ElasticJob, opts v1.UpdateOptions) (result *v1alpha1.ElasticJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticjobsResource, c.ns, elasticJob), &v1alpha1.ElasticJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticJobs) UpdateStatus(ctx context.Context, elasticJob *v1alpha1.ElasticJob, opts v1.UpdateOptions) (*v1alpha1.ElasticJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticjobsResource, "status", c.ns, elasticJob), &v1alpha1.ElasticJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticJob), err
}

// Delete takes name of the elasticJob and deletes it. Returns an error if one occurs.
func (c *FakeElasticJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(elasticjobsResource, c.ns, name, opts), &v1alpha1.ElasticJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticJobList{})
	return err
}

// Patch applies the patch and returns the patched elasticJob.
func (c *FakeElasticJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ElasticJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticJob), err
}