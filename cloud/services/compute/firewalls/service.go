/*
Copyright 2021 The Kubernetes Authors.

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

package firewalls

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-cloud-provider/pkg/cloud/meta"
	"github.com/newrelic-forks/cluster-api-provider-gcp/cloud"
	"google.golang.org/api/compute/v1"
)

type firewallsInterface interface {
	Get(ctx context.Context, key *meta.Key) (*compute.Firewall, error)
	Insert(ctx context.Context, key *meta.Key, obj *compute.Firewall) error
	Update(ctx context.Context, key *meta.Key, obj *compute.Firewall) error
	Delete(ctx context.Context, key *meta.Key) error
}

// Scope is an interfaces that hold used methods.
type Scope interface {
	cloud.ClusterGetter
	FirewallRulesSpec() []*compute.Firewall
}

// Service implements firewalls reconciler.
type Service struct {
	scope     Scope
	firewalls firewallsInterface
}

var _ cloud.Reconciler = &Service{}

// New returns Service from given scope.
func New(scope Scope) *Service {
	return &Service{
		scope:     scope,
		firewalls: scope.Cloud().Firewalls(),
	}
}
