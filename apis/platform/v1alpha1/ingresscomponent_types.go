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

var ErrUnableToConvertIngressComponent = errors.New("unable to convert to IngressComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IngressComponentSpec defines the desired state of IngressComponent.
type IngressComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection IngressComponentCollectionSpec `json:"collection"`

	// +kubebuilder:validation:Optional
	Nginx IngressComponentSpecNginx `json:"nginx,omitempty"`

	// +kubebuilder:validation:Optional
	Kong IngressComponentSpecKong `json:"kong,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalDNS IngressComponentSpecExternalDNS `json:"externalDNS,omitempty"`

	// +kubebuilder:default="nukleros-ingress-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-ingress-system")
	//
	//	Namespace to use for ingress support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	DomainName string `json:"domainName,omitempty"`
}

type IngressComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type IngressComponentSpecNginx struct {
	// +kubebuilder:default="deployment"
	// +kubebuilder:validation:Optional
	// (Default: "deployment")
	//
	//	+kubebuilder:validation:Enum=deployment;daemonset
	//	Method of install nginx ingress controller.  One of: deployment | daemonset.
	InstallType string `json:"installType,omitempty"`

	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	// (Default: true)
	//
	//	Determines whether nginx project is included at installation time.
	Include bool `json:"include,omitempty"`

	// +kubebuilder:default="nginx/nginx-ingress"
	// +kubebuilder:validation:Optional
	// (Default: "nginx/nginx-ingress")
	//
	//	Image repo and name to use for nginx.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="2.3.0"
	// +kubebuilder:validation:Optional
	// (Default: "2.3.0")
	//
	//	Version of nginx to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the nginx ingress controller deployment.
	Replicas int `json:"replicas,omitempty"`
}

type IngressComponentSpecKong struct {
	// +kubebuilder:default=false
	// +kubebuilder:validation:Optional
	// (Default: false)
	//
	//	Determines whether nginx project is included at installation time.
	Include bool `json:"include,omitempty"`

	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	//	Number of replicas to use for the kong ingress deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:validation:Optional
	Gateway IngressComponentSpecKongGateway `json:"gateway,omitempty"`

	// +kubebuilder:validation:Optional
	IngressController IngressComponentSpecKongIngressController `json:"ingressController,omitempty"`
}

type IngressComponentSpecKongGateway struct {
	// +kubebuilder:default="kong/kong-gateway"
	// +kubebuilder:validation:Optional
	// (Default: "kong/kong-gateway")
	//
	//	Image repo and name to use for kong gateway.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="2.8"
	// +kubebuilder:validation:Optional
	// (Default: "2.8")
	//
	//	Version of kong gateway to use.
	Version string `json:"version,omitempty"`
}

type IngressComponentSpecKongIngressController struct {
	// +kubebuilder:default="kong/kubernetes-ingress-controller"
	// +kubebuilder:validation:Optional
	// (Default: "kong/kubernetes-ingress-controller")
	//
	//	Image repo and name to use for kong ingress controller.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="2.5.0"
	// +kubebuilder:validation:Optional
	// (Default: "2.5.0")
	//
	//	Version of kong ingress controller to use.
	Version string `json:"version,omitempty"`
}

type IngressComponentSpecExternalDNS struct {
	// +kubebuilder:default="none"
	// +kubebuilder:validation:Optional
	// (Default: "none")
	//
	//	+kubebuilder:validation:Enum=none;active-directory;google;route53
	//	DNS provider.  One of: none | active-directory | google | route53.
	//	If "none" external DNS will not be installed.
	Provider string `json:"provider,omitempty"`

	// +kubebuilder:default="k8s.gcr.io/external-dns/external-dns"
	// +kubebuilder:validation:Optional
	// (Default: "k8s.gcr.io/external-dns/external-dns")
	//
	//	Image repo and name to use for external-dns.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="v0.12.2"
	// +kubebuilder:validation:Optional
	// (Default: "v0.12.2")
	//
	//	Version of external-dns to use.
	Version string `json:"version,omitempty"`
}

// IngressComponentStatus defines the observed state of IngressComponent.
type IngressComponentStatus struct {
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

// IngressComponent is the Schema for the ingresscomponents API.
type IngressComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressComponentSpec   `json:"spec,omitempty"`
	Status            IngressComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IngressComponentList contains a list of IngressComponent.
type IngressComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *IngressComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *IngressComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *IngressComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *IngressComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *IngressComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *IngressComponent) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *IngressComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *IngressComponent) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*IngressComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{
		&CertificatesComponent{},
	}
}

// GetComponentGVK returns a GVK object for the component.
func (*IngressComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("IngressComponent")
}

func init() {
	SchemeBuilder.Register(&IngressComponent{}, &IngressComponentList{})
}
