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

// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete

// CreateDaemonSetNamespaceNginxIngress creates the DaemonSet resource with name nginx-ingress.
func CreateDaemonSetNamespaceNginxIngress(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	if parent.Spec.Nginx.InstallType != "daemonset" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			// +operator-builder:resource:field=nginx.installType,value="daemonset",include
			"apiVersion": "apps/v1",
			"kind":       "DaemonSet",
			"metadata": map[string]interface{}{
				"name":      "nginx-ingress",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "nginx-ingress",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "nginx-ingress",
							//annotations:
						},
					},
					//prometheus.io/scrape: "true"
					//prometheus.io/port: "9113"
					//prometheus.io/scheme: http
					"spec": map[string]interface{}{
						"serviceAccountName":           "nginx-ingress",
						"automountServiceAccountToken": true,
						"containers": []interface{}{
							map[string]interface{}{
								// controlled by field: nginx.image
								// controlled by field: nginx.version
								//  Image repo and name to use for nginx.
								//  Version of nginx to use.
								"image":           "" + parent.Spec.Nginx.Image + ":" + parent.Spec.Nginx.Version + "",
								"imagePullPolicy": "IfNotPresent",
								"name":            "nginx-ingress",
								"ports": []interface{}{
									map[string]interface{}{
										"name":          "http",
										"containerPort": 80,
										"hostPort":      80,
									},
									map[string]interface{}{
										"name":          "https",
										"containerPort": 443,
										"hostPort":      443,
									},
									map[string]interface{}{
										"name":          "readiness-port",
										"containerPort": 8081,
									},
									map[string]interface{}{
										"name":          "prometheus",
										"containerPort": 9113,
									},
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path": "/nginx-ready",
										"port": "readiness-port",
									},
									"periodSeconds": 1,
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "128Mi",
										//limits:
										//  cpu: "1"
										//  memory: "1Gi"
									},
								},
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": true,
									"runAsUser":                101, // nginx
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
										"add": []interface{}{
											"NET_BIND_SERVICE",
										},
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POD_NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name": "POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
								},
								"args": []interface{}{
									"-nginx-configmaps=$(POD_NAMESPACE)/nginx-config",
									"-default-server-tls-secret=$(POD_NAMESPACE)/default-server-secret",
								},
							},
						},
					},
				},
			},
			//- -v=3 # Enables extensive logging. Useful for troubleshooting.
			//- -report-ingress-status
			//- -external-service=nginx-ingress
			//- -enable-prometheus-metrics
			//- -global-configuration=$(POD_NAMESPACE)/nginx-configuration
		},
	}

	return mutate.MutateDaemonSetNamespaceNginxIngress(resourceObj, parent, collection, reconciler, req)
}
