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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// StatefulSetsClusterGetter has a method to return a StatefulSetClusterInterface.
// A group's cluster client should implement this interface.
type StatefulSetsClusterGetter interface {
	StatefulSets() StatefulSetClusterInterface
}

// StatefulSetClusterInterface can operate on StatefulSets across all clusters,
// or scope down to one cluster and return a StatefulSetsNamespacer.
type StatefulSetClusterInterface interface {
	Cluster(logicalcluster.Path) StatefulSetsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*appsv1.StatefulSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type statefulSetsClusterInterface struct {
	clientCache kcpclient.Cache[*appsv1client.AppsV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *statefulSetsClusterInterface) Cluster(clusterPath logicalcluster.Path) StatefulSetsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &statefulSetsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all StatefulSets across all clusters.
func (c *statefulSetsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.StatefulSetList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StatefulSets(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all StatefulSets across all clusters.
func (c *statefulSetsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StatefulSets(metav1.NamespaceAll).Watch(ctx, opts)
}

// StatefulSetsNamespacer can scope to objects within a namespace, returning a appsv1client.StatefulSetInterface.
type StatefulSetsNamespacer interface {
	Namespace(string) appsv1client.StatefulSetInterface
}

type statefulSetsNamespacer struct {
	clientCache kcpclient.Cache[*appsv1client.AppsV1Client]
	clusterPath logicalcluster.Path
}

func (n *statefulSetsNamespacer) Namespace(namespace string) appsv1client.StatefulSetInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).StatefulSets(namespace)
}
