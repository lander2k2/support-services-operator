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

package ingresscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/ingresscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// CreateSecretNamespaceExternalDnsActiveDirectory creates the Secret resource with name external-dns-active-directory.
func CreateSecretNamespaceExternalDnsActiveDirectory(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.ExternalDNS.Provider != "active-directory" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalDNS.provider,value="active-directory",include
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "external-dns-active-directory",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "external-dns",
				},
			},
			"stringData": map[string]interface{}{
				"EXTERNAL_DNS_TXT_OWNER_ID":              "external-dns-",
				"EXTERNAL_DNS_TXT_PREFIX":                "external-dns-",
				"EXTERNAL_DNS_PROVIDER":                  "rfc2136",
				"EXTERNAL_DNS_RFC2136_HOST":              "ad.mydomain.com",
				"EXTERNAL_DNS_RFC2136_PORT":              "53",
				"EXTERNAL_DNS_RFC2136_ZONE":              "mydomain.com",
				"EXTERNAL_DNS_DOMAIN_FILTER":             "mydomain.com",
				"EXTERNAL_DNS_RFC2136_KERBEROS_REALM":    "MYDOMAIN.COM",
				"EXTERNAL_DNS_RFC2136_KERBEROS_USERNAME": "administrator",
				"EXTERNAL_DNS_RFC2136_KERBEROS_PASSWORD": "thisisinsecure",
				"EXTERNAL_DNS_RFC2136_GSS_TSIG":          "true",
				"EXTERNAL_DNS_RFC2136_TSIG_AXFR":         "true",
				"EXTERNAL_DNS_POLICY":                    "sync",
			},
		},
	}

	return mutate.MutateSecretNamespaceExternalDnsActiveDirectory(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapNamespaceExternalDnsActiveDirectoryKerberos creates the ConfigMap resource with name external-dns-active-directory-kerberos.
func CreateConfigMapNamespaceExternalDnsActiveDirectoryKerberos(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.ExternalDNS.Provider != "active-directory" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalDNS.provider,value="active-directory",include
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "external-dns-active-directory-kerberos",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "external-dns",
				},
			},
			"data": map[string]interface{}{
				"krb5.conf": `[logging]
default = FILE:/var/log/krb5libs.log
kdc = FILE:/var/log/krb5kdc.log
admin_server = FILE:/var/log/kadmind.log

[libdefaults]
dns_lookup_realm = true
dns_lookup_kdc = true
ticket_lifetime = 24h
renew_lifetime = 7d
forwardable = true
rdns = false
pkinit_anchors = /etc/pki/tls/certs/ca-bundle.crt
default_ccache_name = KEYRING:persistent:%{uid}
default_realm = MYDOMAIN.COM

[realms]
MYDOMAIN.COM = {
  admin_server = mydomain.com
}

[domain_realm]
mydomain.com = MYDOMAIN.COM
.mydomain.com = MYDOMAIN.COM
`,
			},
		},
	}

	return mutate.MutateConfigMapNamespaceExternalDnsActiveDirectoryKerberos(resourceObj, parent, collection, reconciler, req)
}
