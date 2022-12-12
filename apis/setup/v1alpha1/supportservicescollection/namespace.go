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

package supportservicescollection

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/setup/v1alpha1/supportservicescollection/mutate"
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespaceParentName creates the Namespace resource with name parent.Name.
func CreateNamespaceParentName(
	parent *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// controlled by field: tier
			//  +kubebuilder:validation:Enum=local;development;production
			//  The tier of cluster being used.  One of: local | development | production.
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": parent.Name, //  controlled by field:
			},
		},
	}

	return mutate.MutateNamespaceParentName(resourceObj, parent, reconciler, req)
}
