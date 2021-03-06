// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package firecracker

import (
	"context"
	"fmt"
	"os"

	"github.com/talos-systems/talos/internal/pkg/provision"
)

// Destroy Talos cluster as set of Firecracker VMs.
func (p *provisioner) Destroy(ctx context.Context, cluster provision.Cluster, opts ...provision.Option) error {
	options := provision.DefaultOptions()

	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return err
		}
	}

	fmt.Fprintln(options.LogWriter, "stopping VMs")

	if err := p.destroyNodes(cluster.Info(), &options); err != nil {
		return err
	}

	state, ok := cluster.(*state)
	if !ok {
		return fmt.Errorf("error inspecting firecracker state, %#+v", cluster)
	}

	fmt.Fprintln(options.LogWriter, "removing load balancer")

	if err := p.destroyLoadBalancer(state); err != nil {
		return fmt.Errorf("error stopping loadbalancer: %w", err)
	}

	fmt.Fprintln(options.LogWriter, "removing network")

	if err := p.destroyNetwork(state); err != nil {
		return err
	}

	fmt.Fprintln(options.LogWriter, "removing state directory")

	stateDirectoryPath, err := cluster.StatePath()
	if err != nil {
		return err
	}

	return os.RemoveAll(stateDirectoryPath)
}
