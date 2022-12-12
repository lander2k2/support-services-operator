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

package databasecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/application/v1alpha1/databasecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespacePostgresOperator creates the Deployment resource with name postgres-operator.
func CreateDeploymentNamespacePostgresOperator(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "postgres-operator",
				"labels": map[string]interface{}{
					"application":                     "postgres-operator",
					"app.kubernetes.io/name":          "postgres-operator",
					"application.nukleros.io/group":   "database",
					"application.nukleros.io/project": "zalando-postgres",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				// controlled by field: zalandoPostgres.replicas
				//  Number of replicas to use for the postgres-operator deployment.
				"replicas": parent.Spec.ZalandoPostgres.Replicas,
				"strategy": map[string]interface{}{
					"type": "Recreate",
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"name": "postgres-operator",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"name":                            "postgres-operator",
							"app.kubernetes.io/name":          "postgres-operator",
							"application.nukleros.io/group":   "database",
							"application.nukleros.io/project": "zalando-postgres",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "postgres-operator",
						"containers": []interface{}{
							map[string]interface{}{
								"name": "postgres-operator",
								// controlled by field: zalandoPostgres.image
								// controlled by field: zalandoPostgres.version
								//  Image repo and name to use for postgres operator.
								//  Version of postgres operator to use.
								"image":           "" + parent.Spec.ZalandoPostgres.Image + ":" + parent.Spec.ZalandoPostgres.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "250Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "500m",
										"memory": "500Mi",
									},
								},
								"securityContext": map[string]interface{}{
									"runAsUser":                1000,
									"runAsNonRoot":             true,
									"readOnlyRootFilesystem":   true,
									"allowPrivilegeEscalation": false,
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "CONFIG_MAP_NAME",
										"value": "postgres-operator",
									},
								},
							},
						},
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
															"nginx-ingress",
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

	return mutate.MutateDeploymentNamespacePostgresOperator(resourceObj, parent, collection, reconciler, req)
}
