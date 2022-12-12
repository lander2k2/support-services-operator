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

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceNginxDefaultServerSecretLocal creates the Certificate resource with name nginx-default-server-secret-local.
func CreateCertNamespaceNginxDefaultServerSecretLocal(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	if collection.Spec.Tier != "local" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			// +operator-builder:resource:collectionField=tier,value="local",include
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "nginx-default-server-secret-local",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"secretName": "default-server-secret",
				"dnsNames": []interface{}{
					parent.Spec.DomainName, //  controlled by field: domainName
				},
				"issuerRef": map[string]interface{}{
					"name": "nukleros-local-self-signed",
					"kind": "ClusterIssuer",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceNginxDefaultServerSecretLocal(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceNginxDefaultServerSecretNonProd creates the Certificate resource with name nginx-default-server-secret-non-prod.
func CreateCertNamespaceNginxDefaultServerSecretNonProd(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	if collection.Spec.Tier != "development" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			// +operator-builder:resource:collectionField=tier,value="development",include
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "nginx-default-server-secret-non-prod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"secretName": "default-server-secret",
				"dnsNames": []interface{}{
					parent.Spec.DomainName, //  controlled by field: domainName
				},
				"issuerRef": map[string]interface{}{
					"name": "letsencrypt-staging",
					"kind": "ClusterIssuer",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceNginxDefaultServerSecretNonProd(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceNginxDefaultServerSecretProd creates the Certificate resource with name nginx-default-server-secret-prod.
func CreateCertNamespaceNginxDefaultServerSecretProd(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	if collection.Spec.Tier != "production" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			// +operator-builder:resource:collectionField=tier,value="production",include
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "nginx-default-server-secret-prod",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"secretName": "default-server-secret",
				"dnsNames": []interface{}{
					parent.Spec.DomainName, //  controlled by field: domainName
				},
				"issuerRef": map[string]interface{}{
					"name": "letsencrypt-production",
					"kind": "ClusterIssuer",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceNginxDefaultServerSecretProd(resourceObj, parent, collection, reconciler, req)
}
