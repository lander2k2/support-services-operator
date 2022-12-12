/*
Copyright 2022.

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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	NamespaceNamespace                                   = "parent.Spec.Namespace"
	SecretNamespaceExternalDnsActiveDirectory            = "external-dns-active-directory"
	ConfigMapNamespaceExternalDnsActiveDirectoryKerberos = "external-dns-active-directory-kerberos"
	SecretNamespaceExternalDnsGoogle                     = "external-dns-google"
	SecretNamespaceExternalDnsRoute53                    = "external-dns-route53"
	DeploymentNamespaceExternalDnsActiveDirectory        = "external-dns-active-directory"
	DeploymentNamespaceExternalDnsGoogle                 = "external-dns-google"
	DeploymentNamespaceExternalDnsRoute53                = "external-dns-route53"
	ClusterRoleBindingExternalDnsViewer                  = "external-dns-viewer"
	ServiceAccountNamespaceExternalDns                   = "external-dns"
	ClusterRoleNamespaceExternalDns                      = "external-dns"
	CertNamespaceNginxDefaultServerSecretLocal           = "nginx-default-server-secret-local"
	CertNamespaceNginxDefaultServerSecretNonProd         = "nginx-default-server-secret-non-prod"
	CertNamespaceNginxDefaultServerSecretProd            = "nginx-default-server-secret-prod"
	ConfigMapNamespaceNginxConfig                        = "nginx-config"
	CRDDnsendpointsExternaldnsNginxOrg                   = "dnsendpoints.externaldns.nginx.org"
	CRDVirtualserverroutesK8sNginxOrg                    = "virtualserverroutes.k8s.nginx.org"
	CRDTransportserversK8sNginxOrg                       = "transportservers.k8s.nginx.org"
	CRDPoliciesK8sNginxOrg                               = "policies.k8s.nginx.org"
	CRDVirtualserversK8sNginxOrg                         = "virtualservers.k8s.nginx.org"
	CRDGlobalconfigurationsK8sNginxOrg                   = "globalconfigurations.k8s.nginx.org"
	DaemonSetNamespaceNginxIngress                       = "nginx-ingress"
	DeploymentNamespaceNginxIngress                      = "nginx-ingress"
	IngressClassNginx                                    = "nginx"
	ServiceAccountNamespaceNginxIngress                  = "nginx-ingress"
	ClusterRoleNginxIngress                              = "nginx-ingress"
	ClusterRoleBindingNginxIngress                       = "nginx-ingress"
	ServiceNamespaceNginxIngressAws                      = "nginx-ingress-aws"
	ServiceNamespaceNginxIngressGcpAzure                 = "nginx-ingress-gcp-azure"
	CRDKongclusterpluginsConfigurationKonghqCom          = "kongclusterplugins.configuration.konghq.com"
	CRDKongconsumersConfigurationKonghqCom               = "kongconsumers.configuration.konghq.com"
	CRDKongingressesConfigurationKonghqCom               = "kongingresses.configuration.konghq.com"
	CRDKongpluginsConfigurationKonghqCom                 = "kongplugins.configuration.konghq.com"
	CRDTcpingressesConfigurationKonghqCom                = "tcpingresses.configuration.konghq.com"
	CRDUdpingressesConfigurationKonghqCom                = "udpingresses.configuration.konghq.com"
	DeploymentNamespaceIngressKong                       = "ingress-kong"
	IngressClassKong                                     = "kong"
	ServiceAccountNamespaceKongServiceaccount            = "kong-serviceaccount"
	RoleNamespaceKongLeaderElection                      = "kong-leader-election"
	ClusterRoleKongIngress                               = "kong-ingress"
	RoleBindingNamespaceKongLeaderElection               = "kong-leader-election"
	ClusterRoleBindingKongIngress                        = "kong-ingress"
	ServiceNamespaceKongProxy                            = "kong-proxy"
	ServiceNamespaceKongValidationWebhook                = "kong-validation-webhook"
)
