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

// CreateSecretNamespaceExternalDnsGoogle creates the Secret resource with name external-dns-google.
func CreateSecretNamespaceExternalDnsGoogle(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.ExternalDNS.Provider != "google" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalDNS.provider,value="google",include
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "external-dns-google",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "external-dns",
				},
			},
			"stringData": map[string]interface{}{
				"EXTERNAL_DNS_TXT_OWNER_ID":           "external-dns-",
				"EXTERNAL_DNS_TXT_PREFIX":             "external-dns-",
				"EXTERNAL_DNS_PROVIDER":               "google",
				"EXTERNAL_DNS_GOOGLE_ZONE_VISIBILITY": "private",
				"EXTERNAL_DNS_GOOGLE_PROJECT":         "my-project",
				"EXTERNAL_DNS_DOMAIN_FILTER":          "mydomain.com",
				"EXTERNAL_DNS_POLICY":                 "sync",
			},
		},
	}

	return mutate.MutateSecretNamespaceExternalDnsGoogle(resourceObj, parent, collection, reconciler, req)
}
