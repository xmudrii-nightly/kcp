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

// Code generated by cluster-client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	kcpv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/apis/wildwest/v1alpha1"
	kcpwildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/typed/wildwest/v1alpha1"
)

// CowboysClusterGetter has a method to return a CowboyClusterInterface.
// A group's cluster client should implement this interface.
type CowboysClusterGetter interface {
	Cowboys() CowboyClusterInterface
}

// CowboyClusterInterface can operate on Cowboys across all clusters,
// or scope down to one cluster and return a CowboysNamespacer.
type CowboyClusterInterface interface {
	Cluster(logicalcluster.Path) CowboysNamespacer
	List(ctx context.Context, opts v1.ListOptions) (*kcpv1alpha1.CowboyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	CowboyClusterExpansion
}

type cowboysClusterInterface struct {
	clientCache kcpclient.Cache[*kcpwildwestv1alpha1.WildwestV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *cowboysClusterInterface) Cluster(clusterPath logicalcluster.Path) CowboysNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cowboysNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all Cowboys across all clusters.
func (c *cowboysClusterInterface) List(ctx context.Context, opts v1.ListOptions) (*kcpv1alpha1.CowboyList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Cowboys(v1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Cowboys across all clusters.
func (c *cowboysClusterInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Cowboys(v1.NamespaceAll).Watch(ctx, opts)
}

// CowboysNamespacer can scope to objects within a namespace, returning a kcpwildwestv1alpha1.CowboyInterface.
type CowboysNamespacer interface {
	Namespace(string) kcpwildwestv1alpha1.CowboyInterface
}

type cowboysNamespacer struct {
	clientCache kcpclient.Cache[*kcpwildwestv1alpha1.WildwestV1alpha1Client]
	clusterPath logicalcluster.Path
}

func (n *cowboysNamespacer) Namespace(namespace string) kcpwildwestv1alpha1.CowboyInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).Cowboys(namespace)
}
