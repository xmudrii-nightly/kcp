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

package v1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamcoordinationv1informers "k8s.io/client-go/informers/coordination/v1"
	upstreamcoordinationv1listers "k8s.io/client-go/listers/coordination/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	clientset "github.com/kcp-dev/client-go/kubernetes"
	coordinationv1listers "github.com/kcp-dev/client-go/listers/coordination/v1"
)

// LeaseClusterInformer provides access to a shared informer and lister for
// Leases.
type LeaseClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamcoordinationv1informers.LeaseInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() coordinationv1listers.LeaseClusterLister
}

type leaseClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLeaseClusterInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLeaseClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredLeaseClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLeaseClusterInformer constructs a new informer for Lease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLeaseClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1().Leases().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoordinationV1().Leases().Watch(context.TODO(), options)
			},
		},
		&coordinationv1.Lease{},
		resyncPeriod,
		indexers,
	)
}

func (f *leaseClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredLeaseClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *leaseClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&coordinationv1.Lease{}, f.defaultInformer)
}

func (f *leaseClusterInformer) Lister() coordinationv1listers.LeaseClusterLister {
	return coordinationv1listers.NewLeaseClusterLister(f.Informer().GetIndexer())
}

func (f *leaseClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamcoordinationv1informers.LeaseInformer {
	return &leaseInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type leaseInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamcoordinationv1listers.LeaseLister
}

func (f *leaseInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *leaseInformer) Lister() upstreamcoordinationv1listers.LeaseLister {
	return f.lister
}
