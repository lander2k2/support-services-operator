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

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespaceNamespace creates the Namespace resource with name parent.Spec.Namespace.
func CreateNamespaceNamespace(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// controlled by field: nginx.installType
			// controlled by field: nginx.include
			// controlled by field: kong.include
			// controlled by field: externalDNS.provider
			//  +kubebuilder:validation:Enum=deployment;daemonset
			//  Method of install nginx ingress controller.  One of: deployment | daemonset.
			//  Determines whether nginx project is included at installation time.
			//  Determines whether nginx project is included at installation time.
			//  +kubebuilder:validation:Enum=none;active-directory;google;route53
			//  DNS provider.  One of: none | active-directory | google | route53.
			//  If "none" external DNS will not be installed.
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				// controlled by field: namespace
				//  Namespace to use for ingress support services.
				"name": parent.Spec.Namespace,
				"labels": map[string]interface{}{
					"externalDNSProvider": parent.Spec.ExternalDNS.Provider, //  controlled by field: externalDNS.provider
				},
			},
		},
	}

	return mutate.MutateNamespaceNamespace(resourceObj, parent, collection, reconciler, req)
}
