// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package generate

import (
	"net/url"

	v1alpha1 "github.com/talos-systems/talos/pkg/config/types/v1alpha1"
	"github.com/talos-systems/talos/pkg/crypto/x509"
)

func workerUd(in *Input) (*v1alpha1.Config, error) {
	config := &v1alpha1.Config{
		ConfigVersion: "v1alpha1",
		ConfigDebug:   in.Debug,
	}

	machine := &v1alpha1.MachineConfig{
		MachineType:     "worker",
		MachineToken:    in.TrustdInfo.Token,
		MachineCertSANs: in.AdditionalMachineCertSANs,
		MachineKubelet:  &v1alpha1.KubeletConfig{},
		MachineNetwork:  in.NetworkConfig,
		MachineInstall: &v1alpha1.InstallConfig{
			InstallDisk:       in.InstallDisk,
			InstallImage:      in.InstallImage,
			InstallBootloader: true,
		},
		MachineRegistries: v1alpha1.RegistriesConfig{
			RegistryMirrors: in.RegistryMirrors,
		},
	}

	controlPlaneURL, err := url.Parse(in.ControlPlaneEndpoint)
	if err != nil {
		return config, err
	}

	cluster := &v1alpha1.ClusterConfig{
		ClusterCA:      &x509.PEMEncodedCertificateAndKey{Crt: in.Certs.K8s.Crt},
		BootstrapToken: in.Secrets.BootstrapToken,
		ControlPlane: &v1alpha1.ControlPlaneConfig{
			Endpoint: &v1alpha1.Endpoint{URL: controlPlaneURL},
		},
		ClusterNetwork: &v1alpha1.ClusterNetworkConfig{
			DNSDomain:     in.ServiceDomain,
			PodSubnet:     in.PodNet,
			ServiceSubnet: in.ServiceNet,
		},
	}

	config.MachineConfig = machine
	config.ClusterConfig = cluster

	return config, nil
}
