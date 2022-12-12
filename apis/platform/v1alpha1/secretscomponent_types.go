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

package v1alpha1

import (
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertSecretsComponent = errors.New("unable to convert to SecretsComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecretsComponentSpec defines the desired state of SecretsComponent.
type SecretsComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection SecretsComponentCollectionSpec `json:"collection"`

	// +kubebuilder:validation:Optional
	ExternalSecrets SecretsComponentSpecExternalSecrets `json:"externalSecrets,omitempty"`

	// +kubebuilder:validation:Optional
	Reloader SecretsComponentSpecReloader `json:"reloader,omitempty"`

	// +kubebuilder:default="nukleros-secrets-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-secrets-system")
	//
	//	Namespace to use for secrets support services.
	Namespace string `json:"namespace,omitempty"`
}

type SecretsComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type SecretsComponentSpecExternalSecrets struct {
	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	// (Default: true)
	//
	//	Determines whether external secrets project is included at installation time.
	Include bool `json:"include,omitempty"`

	// +kubebuilder:default="v0.5.9"
	// +kubebuilder:validation:Optional
	// (Default: "v0.5.9")
	//
	//	Version of external-secrets to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:validation:Optional
	CertController SecretsComponentSpecExternalSecretsCertController `json:"certController,omitempty"`

	// +kubebuilder:default="ghcr.io/external-secrets/external-secrets"
	// +kubebuilder:validation:Optional
	// (Default: "ghcr.io/external-secrets/external-secrets")
	//
	//	Image repo and name to use for external-secrets.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Controller SecretsComponentSpecExternalSecretsController `json:"controller,omitempty"`

	// +kubebuilder:validation:Optional
	Webhook SecretsComponentSpecExternalSecretsWebhook `json:"webhook,omitempty"`
}

type SecretsComponentSpecExternalSecretsCertController struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	//	Number of replicas to use for the external-secrets cert-controller deployment.
	Replicas int `json:"replicas,omitempty"`
}

type SecretsComponentSpecExternalSecretsController struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the external-secrets controller deployment.
	Replicas int `json:"replicas,omitempty"`
}

type SecretsComponentSpecExternalSecretsWebhook struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the external-secrets webhook deployment.
	Replicas int `json:"replicas,omitempty"`
}

type SecretsComponentSpecReloader struct {
	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	// (Default: true)
	//
	//	Determines whether reloader project is included at installation time.
	Include bool `json:"include,omitempty"`

	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	//	Number of replicas to use for the reloader deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="stakater/reloader"
	// +kubebuilder:validation:Optional
	// (Default: "stakater/reloader")
	//
	//	Image repo and name to use for reloader.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="v0.0.119"
	// +kubebuilder:validation:Optional
	// (Default: "v0.0.119")
	//
	//	Version of reloader to use.
	Version string `json:"version,omitempty"`
}

// SecretsComponentStatus defines the observed state of SecretsComponent.
type SecretsComponentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// SecretsComponent is the Schema for the secretscomponents API.
type SecretsComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsComponentSpec   `json:"spec,omitempty"`
	Status            SecretsComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsComponentList contains a list of SecretsComponent.
type SecretsComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *SecretsComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *SecretsComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *SecretsComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *SecretsComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *SecretsComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *SecretsComponent) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *SecretsComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *SecretsComponent) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*SecretsComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*SecretsComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("SecretsComponent")
}

func init() {
	SchemeBuilder.Register(&SecretsComponent{}, &SecretsComponentList{})
}
