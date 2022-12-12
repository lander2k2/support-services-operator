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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceSecretReloader creates the Deployment resource with name secret-reloader.
func CreateDeploymentNamespaceSecretReloader(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Reloader.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=reloader.include,value=true,include
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "secret-reloader",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "reloader",
				},
				"name":      "secret-reloader",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				// controlled by field: reloader.replicas
				//  Number of replicas to use for the reloader deployment.
				"replicas":             parent.Spec.Reloader.Replicas,
				"revisionHistoryLimit": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":        "secret-reloader",
						"platform.nukleros.io/category": "secrets",
						"platform.nukleros.io/project":  "reloader",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "secret-reloader",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "reloader",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								// controlled by field: reloader.image
								// controlled by field: reloader.version
								//  Image repo and name to use for reloader.
								//  Version of reloader to use.
								"image":           "" + parent.Spec.Reloader.Image + ":" + parent.Spec.Reloader.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"name":            "secret-reloader",
								"ports": []interface{}{
									map[string]interface{}{
										"name":          "http",
										"containerPort": 9090,
									},
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/metrics",
										"port": "http",
									},
									"timeoutSeconds":   5,
									"failureThreshold": 5,
									"periodSeconds":    10,
									"successThreshold": 1,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/metrics",
										"port": "http",
									},
									"timeoutSeconds":   5,
									"failureThreshold": 5,
									"periodSeconds":    10,
									"successThreshold": 1,
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"runAsNonRoot":             true,
									"readOnlyRootFilesystem":   true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
									"runAsUser": 65534,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "25m",
										"memory": "32Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "50m",
										"memory": "64Mi",
									},
								},
							},
						},
						"securityContext": map[string]interface{}{
							"runAsNonRoot": true,
							"runAsUser":    65534,
						},
						"serviceAccountName": "secret-reloader",
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"reloader",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceSecretReloader(resourceObj, parent, collection, reconciler, req)
}
