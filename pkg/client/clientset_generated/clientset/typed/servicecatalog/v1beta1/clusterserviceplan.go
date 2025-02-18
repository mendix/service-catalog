/*
Copyright 2023 The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scheme "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterServicePlansGetter has a method to return a ClusterServicePlanInterface.
// A group's client should implement this interface.
type ClusterServicePlansGetter interface {
	ClusterServicePlans() ClusterServicePlanInterface
}

// ClusterServicePlanInterface has methods to work with ClusterServicePlan resources.
type ClusterServicePlanInterface interface {
	Create(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.CreateOptions) (*v1beta1.ClusterServicePlan, error)
	Update(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.UpdateOptions) (*v1beta1.ClusterServicePlan, error)
	UpdateStatus(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.UpdateOptions) (*v1beta1.ClusterServicePlan, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterServicePlan, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterServicePlanList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterServicePlan, err error)
	ClusterServicePlanExpansion
}

// clusterServicePlans implements ClusterServicePlanInterface
type clusterServicePlans struct {
	client rest.Interface
}

// newClusterServicePlans returns a ClusterServicePlans
func newClusterServicePlans(c *ServicecatalogV1beta1Client) *clusterServicePlans {
	return &clusterServicePlans{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterServicePlan, and returns the corresponding clusterServicePlan object, and an error if there is any.
func (c *clusterServicePlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterServicePlan, err error) {
	result = &v1beta1.ClusterServicePlan{}
	err = c.client.Get().
		Resource("clusterserviceplans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterServicePlans that match those selectors.
func (c *clusterServicePlans) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterServicePlanList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ClusterServicePlanList{}
	err = c.client.Get().
		Resource("clusterserviceplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterServicePlans.
func (c *clusterServicePlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterserviceplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterServicePlan and creates it.  Returns the server's representation of the clusterServicePlan, and an error, if there is any.
func (c *clusterServicePlans) Create(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.CreateOptions) (result *v1beta1.ClusterServicePlan, err error) {
	result = &v1beta1.ClusterServicePlan{}
	err = c.client.Post().
		Resource("clusterserviceplans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServicePlan).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterServicePlan and updates it. Returns the server's representation of the clusterServicePlan, and an error, if there is any.
func (c *clusterServicePlans) Update(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.UpdateOptions) (result *v1beta1.ClusterServicePlan, err error) {
	result = &v1beta1.ClusterServicePlan{}
	err = c.client.Put().
		Resource("clusterserviceplans").
		Name(clusterServicePlan.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServicePlan).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterServicePlans) UpdateStatus(ctx context.Context, clusterServicePlan *v1beta1.ClusterServicePlan, opts v1.UpdateOptions) (result *v1beta1.ClusterServicePlan, err error) {
	result = &v1beta1.ClusterServicePlan{}
	err = c.client.Put().
		Resource("clusterserviceplans").
		Name(clusterServicePlan.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServicePlan).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterServicePlan and deletes it. Returns an error if one occurs.
func (c *clusterServicePlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterserviceplans").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterServicePlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterserviceplans").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterServicePlan.
func (c *clusterServicePlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterServicePlan, err error) {
	result = &v1beta1.ClusterServicePlan{}
	err = c.client.Patch(pt).
		Resource("clusterserviceplans").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
