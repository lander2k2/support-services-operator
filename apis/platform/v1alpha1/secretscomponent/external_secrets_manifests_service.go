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

package secretscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/secretscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceNamespaceExternalSecretsWebhook creates the Service resource with name external-secrets-webhook.
func CreateServiceNamespaceExternalSecretsWebhook(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.ExternalSecrets.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalSecrets.include,value=true,include
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-webhook",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"external-secrets.io/component": "webhook",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"spec": map[string]interface{}{
				"type": "ClusterIP",
				"ports": []interface{}{
					map[string]interface{}{
						"port":       443,
						"targetPort": 10250,
						"protocol":   "TCP",
						"name":       "webhook",
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/name":     "external-secrets-webhook",
					"app.kubernetes.io/instance": "external-secrets",
				},
			},
		},
	}

	return mutate.MutateServiceNamespaceExternalSecretsWebhook(resourceObj, parent, collection, reconciler, req)
}
