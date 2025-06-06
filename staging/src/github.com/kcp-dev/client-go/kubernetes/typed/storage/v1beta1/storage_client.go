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
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/rest"
)

type StorageV1beta1ClusterInterface interface {
	StorageV1beta1ClusterScoper
	StorageClassesClusterGetter
	VolumeAttachmentsClusterGetter
	CSIDriversClusterGetter
	CSINodesClusterGetter
	CSIStorageCapacitiesClusterGetter
	VolumeAttributesClassesClusterGetter
}

type StorageV1beta1ClusterScoper interface {
	Cluster(logicalcluster.Path) storagev1beta1.StorageV1beta1Interface
}

type StorageV1beta1ClusterClient struct {
	clientCache kcpclient.Cache[*storagev1beta1.StorageV1beta1Client]
}

func (c *StorageV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) storagev1beta1.StorageV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *StorageV1beta1ClusterClient) StorageClasses() StorageClassClusterInterface {
	return &storageClassesClusterInterface{clientCache: c.clientCache}
}

func (c *StorageV1beta1ClusterClient) VolumeAttachments() VolumeAttachmentClusterInterface {
	return &volumeAttachmentsClusterInterface{clientCache: c.clientCache}
}

func (c *StorageV1beta1ClusterClient) CSIDrivers() CSIDriverClusterInterface {
	return &cSIDriversClusterInterface{clientCache: c.clientCache}
}

func (c *StorageV1beta1ClusterClient) CSINodes() CSINodeClusterInterface {
	return &cSINodesClusterInterface{clientCache: c.clientCache}
}

func (c *StorageV1beta1ClusterClient) CSIStorageCapacities() CSIStorageCapacityClusterInterface {
	return &cSIStorageCapacitiesClusterInterface{clientCache: c.clientCache}
}

func (c *StorageV1beta1ClusterClient) VolumeAttributesClasses() VolumeAttributesClassClusterInterface {
	return &volumeAttributesClassesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new StorageV1beta1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*StorageV1beta1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new StorageV1beta1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*StorageV1beta1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*storagev1beta1.StorageV1beta1Client]{
		NewForConfigAndClient: storagev1beta1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &StorageV1beta1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new StorageV1beta1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StorageV1beta1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
