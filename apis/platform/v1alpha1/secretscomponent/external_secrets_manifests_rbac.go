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

// CreateServiceAccountNamespaceExternalSecretsCertController creates the ServiceAccount resource with name external-secrets-cert-controller.
func CreateServiceAccountNamespaceExternalSecretsCertController(
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
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-cert-controller",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-cert-controller",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceExternalSecretsCertController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceExternalSecrets creates the ServiceAccount resource with name external-secrets.
func CreateServiceAccountNamespaceExternalSecrets(
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
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "external-secrets",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceExternalSecrets(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

// CreateServiceAccountNamespaceExternalSecretsWebhook creates the ServiceAccount resource with name external-secrets-webhook.
func CreateServiceAccountNamespaceExternalSecretsWebhook(
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
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-webhook",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-webhook",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
		},
	}

	return mutate.MutateServiceAccountNamespaceExternalSecretsWebhook(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=list;get;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;update;patch

// CreateClusterRoleExternalSecretsCertController creates the ClusterRole resource with name external-secrets-cert-controller.
func CreateClusterRoleExternalSecretsCertController(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "external-secrets-cert-controller",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-cert-controller",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"apiextensions.k8s.io",
					},
					"resources": []interface{}{
						"customresourcedefinitions",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"admissionregistration.k8s.io",
					},
					"resources": []interface{}{
						"validatingwebhookconfigurations",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"endpoints",
					},
					"verbs": []interface{}{
						"list",
						"get",
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
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleExternalSecretsCertController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=external-secrets.io,resources=secretstores,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=clustersecretstores,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets,verbs=get;list;watch;update;patch;create
// +kubebuilder:rbac:groups=external-secrets.io,resources=clusterexternalsecrets,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets/status,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets/finalizers,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=secretstores/status,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=secretstores/finalizers,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=clustersecretstores/status,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=clustersecretstores/finalizers,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=clusterexternalsecrets/status,verbs=update;patch
// +kubebuilder:rbac:groups=external-secrets.io,resources=clusterexternalsecrets/finalizers,verbs=update;patch
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=core,resources=serviceaccounts/token,verbs=create
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

// CreateClusterRoleExternalSecretsController creates the ClusterRole resource with name external-secrets-controller.
func CreateClusterRoleExternalSecretsController(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "external-secrets-controller",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external-secrets.io",
					},
					"resources": []interface{}{
						"secretstores",
						"clustersecretstores",
						"externalsecrets",
						"clusterexternalsecrets",
					},
					"verbs": []interface{}{
						"get",
						"list",
						"watch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external-secrets.io",
					},
					"resources": []interface{}{
						"externalsecrets",
						"externalsecrets/status",
						"externalsecrets/finalizers",
						"secretstores",
						"secretstores/status",
						"secretstores/finalizers",
						"clustersecretstores",
						"clustersecretstores/status",
						"clustersecretstores/finalizers",
						"clusterexternalsecrets",
						"clusterexternalsecrets/status",
						"clusterexternalsecrets/finalizers",
					},
					"verbs": []interface{}{
						"update",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts",
						"namespaces",
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
						"create",
						"update",
						"delete",
						"patch",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"serviceaccounts/token",
					},
					"verbs": []interface{}{
						"create",
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
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external-secrets.io",
					},
					"resources": []interface{}{
						"externalsecrets",
					},
					"verbs": []interface{}{
						"create",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleExternalSecretsController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets,verbs=get;watch;list
// +kubebuilder:rbac:groups=external-secrets.io,resources=secretstores,verbs=get;watch;list
// +kubebuilder:rbac:groups=external-secrets.io,resources=clustersecretstores,verbs=get;watch;list

// CreateClusterRoleExternalSecretsView creates the ClusterRole resource with name external-secrets-view.
func CreateClusterRoleExternalSecretsView(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "external-secrets-view",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                       "external-secrets",
					"app.kubernetes.io/instance":                   "external-secrets",
					"app.kubernetes.io/version":                    parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"rbac.authorization.k8s.io/aggregate-to-view":  "true",
					"rbac.authorization.k8s.io/aggregate-to-edit":  "true",
					"rbac.authorization.k8s.io/aggregate-to-admin": "true",
					"platform.nukleros.io/category":                "secrets",
					"platform.nukleros.io/project":                 "external-secrets",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external-secrets.io",
					},
					"resources": []interface{}{
						"externalsecrets",
						"secretstores",
						"clustersecretstores",
					},
					"verbs": []interface{}{
						"get",
						"watch",
						"list",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleExternalSecretsView(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=external-secrets.io,resources=secretstores,verbs=create;delete;deletecollection;patch;update
// +kubebuilder:rbac:groups=external-secrets.io,resources=clustersecretstores,verbs=create;delete;deletecollection;patch;update

// CreateClusterRoleExternalSecretsEdit creates the ClusterRole resource with name external-secrets-edit.
func CreateClusterRoleExternalSecretsEdit(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": "external-secrets-edit",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":                       "external-secrets",
					"app.kubernetes.io/instance":                   "external-secrets",
					"app.kubernetes.io/version":                    parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"rbac.authorization.k8s.io/aggregate-to-edit":  "true",
					"rbac.authorization.k8s.io/aggregate-to-admin": "true",
					"platform.nukleros.io/category":                "secrets",
					"platform.nukleros.io/project":                 "external-secrets",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"external-secrets.io",
					},
					"resources": []interface{}{
						"externalsecrets",
						"secretstores",
						"clustersecretstores",
					},
					"verbs": []interface{}{
						"create",
						"delete",
						"deletecollection",
						"patch",
						"update",
					},
				},
			},
		},
	}

	return mutate.MutateClusterRoleExternalSecretsEdit(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingExternalSecretsCertController creates the ClusterRoleBinding resource with name external-secrets-cert-controller.
func CreateClusterRoleBindingExternalSecretsCertController(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "external-secrets-cert-controller",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets-cert-controller",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "external-secrets-cert-controller",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "external-secrets-cert-controller",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingExternalSecretsCertController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateClusterRoleBindingExternalSecretsController creates the ClusterRoleBinding resource with name external-secrets-controller.
func CreateClusterRoleBindingExternalSecretsController(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": "external-secrets-controller",
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     "external-secrets-controller",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"name":      "external-secrets",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
					"kind":      "ServiceAccount",
				},
			},
		},
	}

	return mutate.MutateClusterRoleBindingExternalSecretsController(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;update;patch;create
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update;patch

// CreateRoleNamespaceExternalSecretsLeaderelection creates the Role resource with name external-secrets-leaderelection.
func CreateRoleNamespaceExternalSecretsLeaderelection(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-leaderelection",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{
						"",
					},
					"resources": []interface{}{
						"configmaps",
					},
					"resourceNames": []interface{}{
						"external-secrets-controller",
					},
					"verbs": []interface{}{
						"get",
						"update",
						"patch",
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
						"create",
					},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{
						"coordination.k8s.io",
					},
					"resources": []interface{}{
						"leases",
					},
					"verbs": []interface{}{
						"get",
						"create",
						"update",
						"patch",
					},
				},
			},
		},
	}

	return mutate.MutateRoleNamespaceExternalSecretsLeaderelection(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete

// CreateRoleBindingNamespaceExternalSecretsLeaderelection creates the RoleBinding resource with name external-secrets-leaderelection.
func CreateRoleBindingNamespaceExternalSecretsLeaderelection(
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
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "external-secrets-leaderelection",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "external-secrets",
					"app.kubernetes.io/instance":    "external-secrets",
					"app.kubernetes.io/version":     parent.Spec.ExternalSecrets.Version, //  controlled by field: externalSecrets.version
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "external-secrets",
				},
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "Role",
				"name":     "external-secrets-leaderelection",
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      "external-secrets",
					"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				},
			},
		},
	}

	return mutate.MutateRoleBindingNamespaceExternalSecretsLeaderelection(resourceObj, parent, collection, reconciler, req)
}
