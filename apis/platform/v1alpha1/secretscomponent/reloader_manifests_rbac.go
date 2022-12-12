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

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceSecretReloader creates the ServiceAccount resource with name secret-reloader.
func CreateServiceAccountNamespaceSecretReloader(
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
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "reloader",
				},
				"name":      "secret-reloader",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceSecretReloader(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=list;get;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=list;get;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=list;get;update;patch
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=list;get;update;patch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=list;get;update;patch
// +kubebuilder:rbac:groups=extensions,resources=deployments,verbs=list;get;update;patch
// +kubebuilder:rbac:groups=extensions,resources=daemonsets,verbs=list;get;update;patch

// CreateClusterRoleNamespaceSecretReloader creates the ClusterRole resource with name secret-reloader.
func CreateClusterRoleNamespaceSecretReloader(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "reloader",
				},
				"name":      "secret-reloader",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"secrets",
						"configmaps",
					},
					"verbs": []interface{}{
						"list",
						"get",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apps",
					},
					"resources": []interface{}{
						"deployments",
						"daemonsets",
						"statefulsets",
					},
					"verbs": []interface{}{
						"list",
						"get",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"extensions",
					},
					"resources": []interface{}{
						"deployments",
						"daemonsets",
					},
					"verbs": []interface{}{
						"list",
						"get",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleNamespaceSecretReloader(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingNamespaceSecretReloader creates the ClusterRoleBinding resource with name secret-reloader.
func CreateClusterRoleBindingNamespaceSecretReloader(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "reloader",
				},
				"name":      "secret-reloader",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "secret-reloader",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "secret-reloader",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingNamespaceSecretReloader(resourceObj, parent, collection, reconciler, req)
}
