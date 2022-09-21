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

package certificatescomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

const MutatingWebhookConfigurationCertManagerWebhook = "cert-manager-webhook"

// CreateMutatingWebhookConfigurationCertManagerWebhook creates the cert-manager-webhook MutatingWebhookConfiguration resource.
func CreateMutatingWebhookConfigurationCertManagerWebhook(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook",
				"labels": map[string]interface{}{
					"app":                         "webhook",
					"app.kubernetes.io/name":      "webhook",
					"app.kubernetes.io/instance":  "cert-manager",
					"app.kubernetes.io/component": "webhook",
					"app.kubernetes.io/version":   "v1.9.1",
				},
				"annotations": map[string]interface{}{
					"cert-manager.io/inject-ca-from-secret": "certificates-system/cert-manager-webhook-ca",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "webhook.cert-manager.io",
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"cert-manager.io",
								"acme.cert-manager.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"*/*",
							},
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"matchPolicy":    "Equivalent",
					"timeoutSeconds": 10,
					"failurePolicy":  "Fail",
					"sideEffects":    "None",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cert-manager-webhook",
							"namespace": "certificates-system",
							"path":      "/mutate",
						},
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete

const ValidatingWebhookConfigurationCertManagerWebhook = "cert-manager-webhook"

// CreateValidatingWebhookConfigurationCertManagerWebhook creates the cert-manager-webhook ValidatingWebhookConfiguration resource.
func CreateValidatingWebhookConfigurationCertManagerWebhook(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name": "cert-manager-webhook",
				"labels": map[string]interface{}{
					"app":                         "webhook",
					"app.kubernetes.io/name":      "webhook",
					"app.kubernetes.io/instance":  "cert-manager",
					"app.kubernetes.io/component": "webhook",
					"app.kubernetes.io/version":   "v1.9.1",
				},
				"annotations": map[string]interface{}{
					"cert-manager.io/inject-ca-from-secret": "certificates-system/cert-manager-webhook-ca",
				},
			},
			"webhooks": []interface{}{
				map[string]interface{}{
					"name": "webhook.cert-manager.io",
					"namespaceSelector": map[string]interface{}{
						"matchExpressions": []interface{}{
							map[string]interface{}{
								"key":      "cert-manager.io/disable-validation",
								"operator": "NotIn",
								"values": []interface{}{
									"true",
								},
							},
							map[string]interface{}{
								"key":      "name",
								"operator": "NotIn",
								"values": []interface{}{
									"cert-manager",
								},
							},
						},
					},
					"rules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"cert-manager.io",
								"acme.cert-manager.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"*/*",
							},
						},
					},
					"admissionReviewVersions": []interface{}{
						"v1",
					},
					"matchPolicy":    "Equivalent",
					"timeoutSeconds": 10,
					"failurePolicy":  "Fail",
					"sideEffects":    "None",
					"clientConfig": map[string]interface{}{
						"service": map[string]interface{}{
							"name":      "cert-manager-webhook",
							"namespace": "certificates-system",
							"path":      "/validate",
						},
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}