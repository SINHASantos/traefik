package v1alpha1

import (
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion

// ServersTransport is the CRD implementation of a ServersTransport.
// If no serversTransport is specified, the default@internal will be used.
// The default@internal serversTransport is created from the static configuration.
// More info: https://doc.traefik.io/traefik/v3.5/routing/services/#serverstransport_1
type ServersTransport struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	Spec ServersTransportSpec `json:"spec"`
}

// +k8s:deepcopy-gen=true

// ServersTransportSpec defines the desired state of a ServersTransport.
type ServersTransportSpec struct {
	// ServerName defines the server name used to contact the server.
	ServerName string `json:"serverName,omitempty"`
	// InsecureSkipVerify disables SSL certificate verification.
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
	// RootCAs defines a list of CA certificate Secrets or ConfigMaps used to validate server certificates.
	RootCAs []RootCA `json:"rootCAs,omitempty"`
	// RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.
	// Deprecated: RootCAsSecrets is deprecated, please use the RootCAs option instead.
	RootCAsSecrets []string `json:"rootCAsSecrets,omitempty"`
	// CertificatesSecrets defines a list of secret storing client certificates for mTLS.
	CertificatesSecrets []string `json:"certificatesSecrets,omitempty"`
	// MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
	// +kubebuilder:validation:Minimum=0
	MaxIdleConnsPerHost int `json:"maxIdleConnsPerHost,omitempty"`
	// ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
	ForwardingTimeouts *ForwardingTimeouts `json:"forwardingTimeouts,omitempty"`
	// DisableHTTP2 disables HTTP/2 for connections with backend servers.
	DisableHTTP2 bool `json:"disableHTTP2,omitempty"`
	// PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
	PeerCertURI string `json:"peerCertURI,omitempty"`
	// Spiffe defines the SPIFFE configuration.
	Spiffe *dynamic.Spiffe `json:"spiffe,omitempty"`
}

// +k8s:deepcopy-gen=true

// ForwardingTimeouts holds the timeout configurations for forwarding requests to the backend servers.
type ForwardingTimeouts struct {
	// DialTimeout is the amount of time to wait until a connection to a backend server can be established.
	// +kubebuilder:validation:Pattern="^([0-9]+(ns|us|µs|ms|s|m|h)?)+$"
	// +kubebuilder:validation:XIntOrString
	DialTimeout *intstr.IntOrString `json:"dialTimeout,omitempty"`
	// ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
	// +kubebuilder:validation:Pattern="^([0-9]+(ns|us|µs|ms|s|m|h)?)+$"
	// +kubebuilder:validation:XIntOrString
	ResponseHeaderTimeout *intstr.IntOrString `json:"responseHeaderTimeout,omitempty"`
	// IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.
	// +kubebuilder:validation:Pattern="^([0-9]+(ns|us|µs|ms|s|m|h)?)+$"
	// +kubebuilder:validation:XIntOrString
	IdleConnTimeout *intstr.IntOrString `json:"idleConnTimeout,omitempty"`
	// ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.
	// +kubebuilder:validation:Pattern="^([0-9]+(ns|us|µs|ms|s|m|h)?)+$"
	// +kubebuilder:validation:XIntOrString
	ReadIdleTimeout *intstr.IntOrString `json:"readIdleTimeout,omitempty"`
	// PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.
	// +kubebuilder:validation:Pattern="^([0-9]+(ns|us|µs|ms|s|m|h)?)+$"
	// +kubebuilder:validation:XIntOrString
	PingTimeout *intstr.IntOrString `json:"pingTimeout,omitempty"`
}

// +k8s:deepcopy-gen=true

// RootCA defines a reference to a Secret or a ConfigMap that holds a CA certificate.
// If both a Secret and a ConfigMap reference are defined, the Secret reference takes precedence.
// +kubebuilder:validation:XValidation:rule="!has(self.secret) || !has(self.configMap)",message="RootCA cannot have both Secret and ConfigMap defined."
type RootCA struct {
	// Secret defines the name of a Secret that holds a CA certificate.
	// The referenced Secret must contain a certificate under either a tls.ca or a ca.crt key.
	Secret *string `json:"secret,omitempty"`
	// ConfigMap defines the name of a ConfigMap that holds a CA certificate.
	// The referenced ConfigMap must contain a certificate under either a tls.ca or a ca.crt key.
	ConfigMap *string `json:"configMap,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServersTransportList is a collection of ServersTransport resources.
type ServersTransportList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of ServersTransport.
	Items []ServersTransport `json:"items"`
}
