// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package firecracker implements Provisioner via Firecracker VMs.
package firecracker

import (
	"context"

	"github.com/talos-systems/talos/internal/pkg/provision"
	"github.com/talos-systems/talos/pkg/config/machine"
	"github.com/talos-systems/talos/pkg/config/types/v1alpha1"
	"github.com/talos-systems/talos/pkg/config/types/v1alpha1/generate"
)

const stateFileName = "state.yaml"

type provisioner struct {
}

// NewProvisioner initializes docker provisioner.
func NewProvisioner(ctx context.Context) (provision.Provisioner, error) {
	p := &provisioner{}

	return p, nil
}

// Close and release resources.
func (p *provisioner) Close() error {
	return nil
}

// GenOptions provides a list of additional config generate options.
func (p *provisioner) GenOptions(networkReq provision.NetworkRequest) []generate.GenOption {
	nameservers := make([]string, len(networkReq.Nameservers))
	for i := range nameservers {
		nameservers[i] = networkReq.Nameservers[i].String()
	}

	return []generate.GenOption{
		generate.WithInstallDisk("/dev/vda"),
		generate.WithNetworkConfig(&v1alpha1.NetworkConfig{
			NameServers: nameservers,
			NetworkInterfaces: []machine.Device{
				{
					Interface: "eth0",
					CIDR:      "169.254.128.128/32", // link-local IP just to trigger the static networkd config
					MTU:       networkReq.MTU,
				},
			},
		}),
	}
}

// GetLoadBalancers returns internal/external loadbalancer endpoints.
func (p *provisioner) GetLoadBalancers(networkReq provision.NetworkRequest) (internalEndpoint, externalEndpoint string) {
	// firecracker runs loadbalancer on the bridge, which is good for both internal & external access
	return networkReq.GatewayAddr.String(), networkReq.GatewayAddr.String()
}
