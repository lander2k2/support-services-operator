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

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceKongProxy creates the Service resource with name kong-proxy.
func CreateServiceNamespaceKongProxy(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"service.beta.kubernetes.io/aws-load-balancer-backend-protocol": "tcp",
					"service.beta.kubernetes.io/aws-load-balancer-type":             "nlb",
				},
				"name":      "kong-proxy",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "proxy",
						"port":       80,
						"protocol":   "TCP",
						"targetPort": 8000,
					},
					map[string]interface{}{
						"name":       "proxy-ssl",
						"port":       443,
						"protocol":   "TCP",
						"targetPort": 8443,
					},
				},
				"selector": map[string]interface{}{
					"app": "ingress-kong",
				},
				"type": "LoadBalancer",
			},
		},
	}

	return mutate.MutateServiceNamespaceKongProxy(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceKongValidationWebhook creates the Service resource with name kong-validation-webhook.
func CreateServiceNamespaceKongValidationWebhook(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "kong-validation-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"name":       "webhook",
						"port":       443,
						"protocol":   "TCP",
						"targetPort": 8080,
					},
				},
				"selector": map[string]interface{}{
					"app": "ingress-kong",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceKongValidationWebhook(resourceObj, parent, collection, reconciler, req)
}
