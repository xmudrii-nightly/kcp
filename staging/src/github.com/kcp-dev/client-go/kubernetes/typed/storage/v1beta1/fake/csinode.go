/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	storagev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var cSINodesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "csinodes"}
var cSINodesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "CSINode"}

type cSINodesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cSINodesClusterClient) Cluster(clusterPath logicalcluster.Path) storagev1beta1client.CSINodeInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cSINodesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CSINodes that match those selectors across all clusters.
func (c *cSINodesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.CSINodeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(cSINodesResource, cSINodesKind, logicalcluster.Wildcard, opts), &storagev1beta1.CSINodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.CSINodeList{ListMeta: obj.(*storagev1beta1.CSINodeList).ListMeta}
	for _, item := range obj.(*storagev1beta1.CSINodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CSINodes across all clusters.
func (c *cSINodesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(cSINodesResource, logicalcluster.Wildcard, opts))
}

type cSINodesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *cSINodesClient) Create(ctx context.Context, cSINode *storagev1beta1.CSINode, opts metav1.CreateOptions) (*storagev1beta1.CSINode, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(cSINodesResource, c.ClusterPath, cSINode), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

func (c *cSINodesClient) Update(ctx context.Context, cSINode *storagev1beta1.CSINode, opts metav1.UpdateOptions) (*storagev1beta1.CSINode, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(cSINodesResource, c.ClusterPath, cSINode), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

func (c *cSINodesClient) UpdateStatus(ctx context.Context, cSINode *storagev1beta1.CSINode, opts metav1.UpdateOptions) (*storagev1beta1.CSINode, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(cSINodesResource, c.ClusterPath, "status", cSINode), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

func (c *cSINodesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(cSINodesResource, c.ClusterPath, name, opts), &storagev1beta1.CSINode{})
	return err
}

func (c *cSINodesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(cSINodesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1beta1.CSINodeList{})
	return err
}

func (c *cSINodesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagev1beta1.CSINode, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(cSINodesResource, c.ClusterPath, name), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

// List takes label and field selectors, and returns the list of CSINodes that match those selectors.
func (c *cSINodesClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.CSINodeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(cSINodesResource, cSINodesKind, c.ClusterPath, opts), &storagev1beta1.CSINodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.CSINodeList{ListMeta: obj.(*storagev1beta1.CSINodeList).ListMeta}
	for _, item := range obj.(*storagev1beta1.CSINodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cSINodesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(cSINodesResource, c.ClusterPath, opts))
}

func (c *cSINodesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1beta1.CSINode, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSINodesResource, c.ClusterPath, name, pt, data, subresources...), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

func (c *cSINodesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.CSINodeApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.CSINode, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSINodesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}

func (c *cSINodesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.CSINodeApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.CSINode, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(cSINodesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &storagev1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.CSINode), err
}
