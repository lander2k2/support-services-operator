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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceIngressKong creates the Deployment resource with name ingress-kong.
func CreateDeploymentNamespaceIngressKong(
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
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":                           "ingress-kong",
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
					"app.kubernetes.io/name":        "kong-ingress",
				},
				"name":      "ingress-kong",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				// controlled by field: kong.replicas
				//  Number of replicas to use for the kong ingress deployment.
				"replicas": parent.Spec.Kong.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "ingress-kong",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"kuma.io/gateway":                              "enabled",
							"kuma.io/service-account-token-volume":         "kong-serviceaccount-token",
							"traffic.sidecar.istio.io/includeInboundPorts": "",
						},
						"labels": map[string]interface{}{
							"app":                           "ingress-kong",
							"platform.nukleros.io/category": "ingress",
							"platform.nukleros.io/project":  "kong-ingress-controller",
						},
					},
					"spec": map[string]interface{}{
						"automountServiceAccountToken": false,
						"containers": []interface{}{
							map[string]interface{}{
								"env": []interface{}{
									map[string]interface{}{
										"name": "KONG_LICENSE_DATA",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"key":      "license",
												"name":     "kong-enterprise-license",
												"optional": true,
											},
										},
									},
									map[string]interface{}{
										"name":  "KONG_PROXY_LISTEN",
										"value": "0.0.0.0:8000, 0.0.0.0:8443 ssl http2",
									},
									map[string]interface{}{
										"name":  "KONG_PORT_MAPS",
										"value": "80:8000, 443:8443",
									},
									map[string]interface{}{
										"name":  "KONG_ADMIN_LISTEN",
										"value": "127.0.0.1:8444 ssl",
									},
									map[string]interface{}{
										"name":  "KONG_STATUS_LISTEN",
										"value": "0.0.0.0:8100",
									},
									map[string]interface{}{
										"name": "KONG_DATABASE",
										"value": `off
`,
									},
									map[string]interface{}{
										"name":  "KONG_NGINX_WORKER_PROCESSES",
										"value": "2",
									},
									map[string]interface{}{
										"name": "KONG_KIC",
										"value": `on
`,
									},
									map[string]interface{}{
										"name":  "KONG_ADMIN_ACCESS_LOG",
										"value": "/dev/stdout",
									},
									map[string]interface{}{
										"name":  "KONG_ADMIN_ERROR_LOG",
										"value": "/dev/stderr",
									},
									map[string]interface{}{
										"name":  "KONG_PROXY_ERROR_LOG",
										"value": "/dev/stderr",
									},
									map[string]interface{}{
										"name":  "KONG_LUA_PACKAGE_PATH",
										"value": "/opt/?.lua;/opt/?/init.lua;;",
									},
									map[string]interface{}{
										"name":  "KONG_PREFIX",
										"value": "/kong_prefix/",
									},
								},
								// controlled by field: kong.gateway.image
								// controlled by field: kong.gateway.version
								//  Image repo and name to use for kong gateway.
								//  Version of kong gateway to use.
								"image": "" + parent.Spec.Kong.Gateway.Image + ":" + parent.Spec.Kong.Gateway.Version + "",
								"lifecycle": map[string]interface{}{
									"preStop": map[string]interface{}{
										"exec": map[string]interface{}{
											"command": []interface{}{
												"/bin/sh",
												"-c",
												"kong quit",
											},
										},
									},
								},
								"livenessProbe": map[string]interface{}{
									"failureThreshold": 3,
									"httpGet": map[string]interface{}{
										"path":   "/status",
										"port":   8100,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
									"successThreshold":    1,
									"timeoutSeconds":      1,
								},
								"name": "proxy",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8000,
										"name":          "proxy",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 8443,
										"name":          "proxy-ssl",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 8100,
										"name":          "metrics",
										"protocol":      "TCP",
									},
								},
								"readinessProbe": map[string]interface{}{
									"failureThreshold": 3,
									"httpGet": map[string]interface{}{
										"path":   "/status",
										"port":   8100,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
									"successThreshold":    1,
									"timeoutSeconds":      1,
								},
								"imagePullPolicy": "IfNotPresent",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "200m",
										"memory": "256Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "400m",
										"memory": "384Mi",
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "kong-prefix-dir",
										"mountPath": "/kong_prefix/",
									},
									map[string]interface{}{
										"name":      "kong-tmp",
										"mountPath": "/tmp",
									},
								},
							},
							map[string]interface{}{
								"env": []interface{}{
									map[string]interface{}{
										"name":  "CONTROLLER_KONG_ADMIN_URL",
										"value": "https://127.0.0.1:8444",
									},
									map[string]interface{}{
										"name":  "CONTROLLER_KONG_ADMIN_TLS_SKIP_VERIFY",
										"value": "true",
									},
									map[string]interface{}{
										"name":  "CONTROLLER_PUBLISH_SERVICE",
										"value": "nukleros-ingress-system/kong-proxy",
									},
									map[string]interface{}{
										"name": "POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "metadata.name",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"apiVersion": "v1",
												"fieldPath":  "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "KONG_LUA_PACKAGE_PATH",
										"value": "/opt/?.lua;/opt/?/init.lua;;",
									},
								},
								// controlled by field: kong.ingressController.image
								// controlled by field: kong.ingressController.version
								//  Image repo and name to use for kong ingress controller.
								//  Version of kong ingress controller to use.
								"image":           "" + parent.Spec.Kong.IngressController.Image + ":" + parent.Spec.Kong.IngressController.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"livenessProbe": map[string]interface{}{
									"failureThreshold": 3,
									"httpGet": map[string]interface{}{
										"path":   "/healthz",
										"port":   10254,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
									"successThreshold":    1,
									"timeoutSeconds":      1,
								},
								"name": "ingress-controller",
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": 8080,
										"name":          "webhook",
										"protocol":      "TCP",
									},
									map[string]interface{}{
										"containerPort": 10255,
										"name":          "cmetrics",
										"protocol":      "TCP",
									},
								},
								"readinessProbe": map[string]interface{}{
									"failureThreshold": 3,
									"httpGet": map[string]interface{}{
										"path":   "/readyz",
										"port":   10254,
										"scheme": "HTTP",
									},
									"initialDelaySeconds": 5,
									"periodSeconds":       10,
									"successThreshold":    1,
									"timeoutSeconds":      1,
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
										"name":      "kong-serviceaccount-token",
										"readOnly":  true,
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "200m",
										"memory": "256Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "400m",
										"memory": "384Mi",
									},
								},
							},
						},
						"serviceAccountName": "kong-serviceaccount",
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "kong-serviceaccount-token",
								"secret": map[string]interface{}{
									"items": []interface{}{
										map[string]interface{}{
											"key":  "token",
											"path": "token",
										},
										map[string]interface{}{
											"key":  "ca.crt",
											"path": "ca.crt",
										},
										map[string]interface{}{
											"key":  "namespace",
											"path": "namespace",
										},
									},
									"secretName": "kong-serviceaccount-token",
								},
							},
							map[string]interface{}{
								"name":     "kong-prefix-dir",
								"emptyDir": map[string]interface{}{},
							},
							map[string]interface{}{
								"name":     "kong-tmp",
								"emptyDir": map[string]interface{}{},
							},
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
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
															"kong-ingress",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceIngressKong(resourceObj, parent, collection, reconciler, req)
}
