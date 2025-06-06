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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	schedulingv1beta1 "k8s.io/api/scheduling/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	schedulingv1beta1listers "k8s.io/client-go/listers/scheduling/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// PriorityClassClusterLister can list PriorityClasses across all workspaces, or scope down to a PriorityClassLister for one workspace.
// All objects returned here must be treated as read-only.
type PriorityClassClusterLister interface {
	// List lists all PriorityClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*schedulingv1beta1.PriorityClass, err error)
	// Cluster returns a lister that can list and get PriorityClasses in one workspace.
	Cluster(clusterName logicalcluster.Name) schedulingv1beta1listers.PriorityClassLister
	PriorityClassClusterListerExpansion
}

type priorityClassClusterLister struct {
	indexer cache.Indexer
}

// NewPriorityClassClusterLister returns a new PriorityClassClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewPriorityClassClusterLister(indexer cache.Indexer) *priorityClassClusterLister {
	return &priorityClassClusterLister{indexer: indexer}
}

// List lists all PriorityClasses in the indexer across all workspaces.
func (s *priorityClassClusterLister) List(selector labels.Selector) (ret []*schedulingv1beta1.PriorityClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*schedulingv1beta1.PriorityClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get PriorityClasses.
func (s *priorityClassClusterLister) Cluster(clusterName logicalcluster.Name) schedulingv1beta1listers.PriorityClassLister {
	return &priorityClassLister{indexer: s.indexer, clusterName: clusterName}
}

// priorityClassLister implements the schedulingv1beta1listers.PriorityClassLister interface.
type priorityClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all PriorityClasses in the indexer for a workspace.
func (s *priorityClassLister) List(selector labels.Selector) (ret []*schedulingv1beta1.PriorityClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*schedulingv1beta1.PriorityClass))
	})
	return ret, err
}

// Get retrieves the PriorityClass from the indexer for a given workspace and name.
func (s *priorityClassLister) Get(name string) (*schedulingv1beta1.PriorityClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schedulingv1beta1.Resource("priorityclasses"), name)
	}
	return obj.(*schedulingv1beta1.PriorityClass), nil
}
