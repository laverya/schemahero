/*
Copyright 2019 Replicated, Inc.

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
// Code generated by main. DO NOT EDIT.

package schemaheroclientset

import (
	databasesv1alpha2 "github.com/schemahero/schemahero/pkg/client/schemaheroclientset/typed/databases/v1alpha2"
	schemasv1alpha2 "github.com/schemahero/schemahero/pkg/client/schemaheroclientset/typed/schemas/v1alpha2"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	DatabasesV1alpha2() databasesv1alpha2.DatabasesV1alpha2Interface
	// Deprecated: please explicitly pick a version if possible.
	Databases() databasesv1alpha2.DatabasesV1alpha2Interface
	SchemasV1alpha2() schemasv1alpha2.SchemasV1alpha2Interface
	// Deprecated: please explicitly pick a version if possible.
	Schemas() schemasv1alpha2.SchemasV1alpha2Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	databasesV1alpha2 *databasesv1alpha2.DatabasesV1alpha2Client
	schemasV1alpha2   *schemasv1alpha2.SchemasV1alpha2Client
}

// DatabasesV1alpha2 retrieves the DatabasesV1alpha2Client
func (c *Clientset) DatabasesV1alpha2() databasesv1alpha2.DatabasesV1alpha2Interface {
	return c.databasesV1alpha2
}

// Deprecated: Databases retrieves the default version of DatabasesClient.
// Please explicitly pick a version.
func (c *Clientset) Databases() databasesv1alpha2.DatabasesV1alpha2Interface {
	return c.databasesV1alpha2
}

// SchemasV1alpha2 retrieves the SchemasV1alpha2Client
func (c *Clientset) SchemasV1alpha2() schemasv1alpha2.SchemasV1alpha2Interface {
	return c.schemasV1alpha2
}

// Deprecated: Schemas retrieves the default version of SchemasClient.
// Please explicitly pick a version.
func (c *Clientset) Schemas() schemasv1alpha2.SchemasV1alpha2Interface {
	return c.schemasV1alpha2
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
	cs.databasesV1alpha2, err = databasesv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schemasV1alpha2, err = schemasv1alpha2.NewForConfig(&configShallowCopy)
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
	cs.databasesV1alpha2 = databasesv1alpha2.NewForConfigOrDie(c)
	cs.schemasV1alpha2 = schemasv1alpha2.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.databasesV1alpha2 = databasesv1alpha2.New(c)
	cs.schemasV1alpha2 = schemasv1alpha2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
