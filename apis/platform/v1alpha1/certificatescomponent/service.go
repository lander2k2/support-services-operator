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

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/certificatescomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCertManager creates the Service resource with name cert-manager.
func CreateServiceNamespaceCertManager(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cert-manager",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "cert-manager",
					"app.kubernetes.io/name":        "cert-manager",
					"app.kubernetes.io/instance":    "cert-manager",
					"app.kubernetes.io/component":   "controller",
					"app.kubernetes.io/version":     "v1.9.1",
					"platform.nukleros.io/category": "certificates",
					"platform.nukleros.io/project":  "cert-manager",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"protocol":   "TCP",
						"port":       9402,
						"name":       "tcp-prometheus-servicemonitor",
						"targetPort": 9402,
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":      "cert-manager",
					"app.kubernetes.io/instance":  "cert-manager",
					"app.kubernetes.io/component": "controller",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCertManager(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceCertManagerWebhook creates the Service resource with name cert-manager-webhook.
func CreateServiceNamespaceCertManagerWebhook(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "cert-manager-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":                           "webhook",
					"app.kubernetes.io/name":        "webhook",
					"app.kubernetes.io/instance":    "cert-manager",
					"app.kubernetes.io/component":   "webhook",
					"app.kubernetes.io/version":     "v1.9.1",
					"platform.nukleros.io/category": "certificates",
					"platform.nukleros.io/project":  "cert-manager",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "https",
						"port":       443,
						"protocol":   "TCP",
						"targetPort": "https",
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":      "webhook",
					"app.kubernetes.io/instance":  "cert-manager",
					"app.kubernetes.io/component": "webhook",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceCertManagerWebhook(resourceObj, parent, collection, reconciler, req)
}
