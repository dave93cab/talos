// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cluster

import (
	"net/url"
	"time"

	"github.com/talos-systems/talos/pkg/config/machine"
	"github.com/talos-systems/talos/pkg/crypto/x509"
)

// Cluster defines the requirements for a config that pertains to cluster
// related options.
type Cluster interface {
	Name() string
	APIServer() APIServer
	ControllerManager() ControllerManager
	Scheduler() Scheduler
	Endpoint() *url.URL
	Token() Token
	CertSANs() []string
	SetCertSANs([]string)
	CA() *x509.PEMEncodedCertificateAndKey
	AESCBCEncryptionSecret() string
	Config(machine.Type) (string, error)
	Etcd() Etcd
	Network() Network
	LocalAPIServerPort() int
	PodCheckpointer() PodCheckpointer
	CoreDNS() CoreDNS
	ExtraManifestURLs() []string
	AdminKubeconfig() AdminKubeconfig
}

// Network defines the requirements for a config that pertains to cluster
// network options.
type Network interface {
	CNI() CNI
	PodCIDR() string
	ServiceCIDR() string
	DNSDomain() string
}

// CNI defines the requirements for a config that pertains to Kubernetes
// cni.
type CNI interface {
	Name() string
	URLs() []string
}

// APIServer defines the requirements for a config that pertains to apiserver related
// options.
type APIServer interface {
	ExtraArgs() map[string]string
}

// ControllerManager defines the requirements for a config that pertains to controller manager related
// options.
type ControllerManager interface {
	ExtraArgs() map[string]string
}

// Scheduler defines the requirements for a config that pertains to scheduler related
// options.
type Scheduler interface {
	ExtraArgs() map[string]string
}

// Etcd defines the requirements for a config that pertains to etcd related
// options.
type Etcd interface {
	Image() string
	CA() *x509.PEMEncodedCertificateAndKey
	ExtraArgs() map[string]string
}

// Token defines the requirements for a config that pertains to Kubernetes
// bootstrap token.
type Token interface {
	ID() string
	Secret() string
}

// PodCheckpointer defines the requirements for a config that pertains to bootkube
// pod-checkpointer options.
type PodCheckpointer interface {
	Image() string
}

// CoreDNS defines the requirements for a config that pertains to bootkube
// coredns options.
type CoreDNS interface {
	Image() string
}

// AdminKubeconfig defines settings for admin kubeconfig.
type AdminKubeconfig interface {
	CertLifetime() time.Duration
}
