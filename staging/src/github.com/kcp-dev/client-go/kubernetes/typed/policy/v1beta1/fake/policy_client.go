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
	"github.com/kcp-dev/logicalcluster/v3"

	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	"k8s.io/client-go/rest"

	kcppolicyv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/policy/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcppolicyv1beta1.PolicyV1beta1ClusterInterface = (*PolicyV1beta1ClusterClient)(nil)

type PolicyV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *PolicyV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) policyv1beta1.PolicyV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &PolicyV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *PolicyV1beta1ClusterClient) PodDisruptionBudgets() kcppolicyv1beta1.PodDisruptionBudgetClusterInterface {
	return &podDisruptionBudgetsClusterClient{Fake: c.Fake}
}

func (c *PolicyV1beta1ClusterClient) Evictions() kcppolicyv1beta1.EvictionClusterInterface {
	return &evictionsClusterClient{Fake: c.Fake}
}

var _ policyv1beta1.PolicyV1beta1Interface = (*PolicyV1beta1Client)(nil)

type PolicyV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *PolicyV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *PolicyV1beta1Client) PodDisruptionBudgets(namespace string) policyv1beta1.PodDisruptionBudgetInterface {
	return &podDisruptionBudgetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *PolicyV1beta1Client) Evictions(namespace string) policyv1beta1.EvictionInterface {
	return &evictionsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
