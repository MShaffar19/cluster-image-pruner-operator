// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	prunev1alpha1 "github.com/openshift/cluster-image-pruner-operator/pkg/generated/clientset/versioned/typed/prune/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	PruneV1alpha1() prunev1alpha1.PruneV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Prune() prunev1alpha1.PruneV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	pruneV1alpha1 *prunev1alpha1.PruneV1alpha1Client
}

// PruneV1alpha1 retrieves the PruneV1alpha1Client
func (c *Clientset) PruneV1alpha1() prunev1alpha1.PruneV1alpha1Interface {
	return c.pruneV1alpha1
}

// Deprecated: Prune retrieves the default version of PruneClient.
// Please explicitly pick a version.
func (c *Clientset) Prune() prunev1alpha1.PruneV1alpha1Interface {
	return c.pruneV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.pruneV1alpha1, err = prunev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.pruneV1alpha1 = prunev1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.pruneV1alpha1 = prunev1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
