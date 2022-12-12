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

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceNginxIngress creates the ServiceAccount resource with name nginx-ingress.
func CreateServiceAccountNamespaceNginxIngress(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "nginx-ingress",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceNginxIngress(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;update;create
// +kubebuilder:rbac:groups=core,resources=pods,verbs=list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch;list
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=list;watch;get
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=update
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=virtualservers,verbs=list;watch;get
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=virtualserverroutes,verbs=list;watch;get
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=globalconfigurations,verbs=list;watch;get
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=transportservers,verbs=list;watch;get
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=policies,verbs=list;watch;get
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=virtualservers/status,verbs=update
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=virtualserverroutes/status,verbs=update
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=policies/status,verbs=update
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=transportservers/status,verbs=update
// +kubebuilder:rbac:groups=k8s.nginx.org,resources=dnsendpoints/status,verbs=update
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingressclasses,verbs=get
// +kubebuilder:rbac:groups=cis.f5.com,resources=ingresslinks,verbs=list;watch;get
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=list;watch;get;update;create;delete
// +kubebuilder:rbac:groups=externaldns.nginx.org,resources=dnsendpoints,verbs=list;watch;get;update;create;delete
// +kubebuilder:rbac:groups=externaldns.nginx.org,resources=dnsendpoints/status,verbs=update

// CreateClusterRoleNginxIngress creates the ClusterRole resource with name nginx-ingress.
func CreateClusterRoleNginxIngress(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"kind":       "ClusterRole",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "nginx-ingress",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"services",
						"endpoints",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"pods",
					},
					"verbs": []interface{}{
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"events",
					},
					"verbs": []interface{}{
						"create",
						"patch",
						"list",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses",
					},
					"verbs": []interface{}{
						"list",
						"watch",
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingresses/status",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"k8s.nginx.org",
					},
					"resources": []interface{}{
						"virtualservers",
						"virtualserverroutes",
						"globalconfigurations",
						"transportservers",
						"policies",
					},
					"verbs": []interface{}{
						"list",
						"watch",
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"k8s.nginx.org",
					},
					"resources": []interface{}{
						"virtualservers/status",
						"virtualserverroutes/status",
						"policies/status",
						"transportservers/status",
						"dnsendpoints/status",
					},
					"verbs": []interface{}{
						"update",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"networking.k8s.io",
					},
					"resources": []interface{}{
						"ingressclasses",
					},
					"verbs": []interface{}{
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cis.f5.com",
					},
					"resources": []interface{}{
						"ingresslinks",
					},
					"verbs": []interface{}{
						"list",
						"watch",
						"get",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"cert-manager.io",
					},
					"resources": []interface{}{
						"certificates",
					},
					"verbs": []interface{}{
						"list",
						"watch",
						"get",
						"update",
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"externaldns.nginx.org",
					},
					"resources": []interface{}{
						"dnsendpoints",
					},
					"verbs": []interface{}{
						"list",
						"watch",
						"get",
						"update",
						"create",
						"delete",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"externaldns.nginx.org",
					},
					"resources": []interface{}{
						"dnsendpoints/status",
					},
					"verbs": []interface{}{
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleNginxIngress(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingNginxIngress creates the ClusterRoleBinding resource with name nginx-ingress.
func CreateClusterRoleBindingNginxIngress(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"kind":       "ClusterRoleBinding",
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"metadata": map[string]interface{}{
				"name": "nginx-ingress",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "nginx-ingress",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
			"roleRef": map[string]interface{}{
				"kind":     "ClusterRole",
				"name":     "nginx-ingress",
				"apiGroup": "rbac.authorization.k8s.io",
			},
		},
	}

	return mutate.MutateClusterRoleBindingNginxIngress(resourceObj, parent, collection, reconciler, req)
}
