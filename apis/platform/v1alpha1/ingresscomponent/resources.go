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
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// sampleIngressComponent is a sample containing all fields
const sampleIngressComponent = `apiVersion: platform.addons.nukleros.io/v1alpha1
kind: IngressComponent
metadata:
  name: ingresscomponent-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
  nginx:
    installType: "deployment"
    include: true
    image: "nginx/nginx-ingress"
    version: "2.3.0"
    replicas: 2
  kong:
    include: false
    replicas: 2
    gateway:
      image: "kong/kong-gateway"
      version: "2.8"
    ingressController:
      image: "kong/kubernetes-ingress-controller"
      version: "2.5.0"
  externalDNS:
    provider: "none"
    image: "k8s.gcr.io/external-dns/external-dns"
    version: "v0.12.2"
  namespace: "nukleros-ingress-system"
  domainName: "nukleros.io"
`

// sampleIngressComponentRequired is a sample containing only required fields
const sampleIngressComponentRequired = `apiVersion: platform.addons.nukleros.io/v1alpha1
kind: IngressComponent
metadata:
  name: ingresscomponent-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
  domainName: "nukleros.io"
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleIngressComponentRequired
	}

	return sampleIngressComponent
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj platformv1alpha1.IngressComponent,
	collectionObj setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, &collectionObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte, collectionFile []byte) ([]client.Object, error) {
	var workloadObj platformv1alpha1.IngressComponent
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	var collectionObj setupv1alpha1.SupportServices
	if err := yaml.Unmarshal(collectionFile, &collectionObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into collection, %w", err)
	}

	if err := workload.Validate(&collectionObj); err != nil {
		return nil, fmt.Errorf("error validating collection yaml, %w", err)
	}

	return Generate(workloadObj, collectionObj, nil, nil)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*platformv1alpha1.IngressComponent,
	*setupv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateNamespaceNamespace,
	CreateSecretNamespaceExternalDnsActiveDirectory,
	CreateConfigMapNamespaceExternalDnsActiveDirectoryKerberos,
	CreateSecretNamespaceExternalDnsGoogle,
	CreateSecretNamespaceExternalDnsRoute53,
	CreateDeploymentNamespaceExternalDnsActiveDirectory,
	CreateDeploymentNamespaceExternalDnsGoogle,
	CreateDeploymentNamespaceExternalDnsRoute53,
	CreateClusterRoleBindingExternalDnsViewer,
	CreateServiceAccountNamespaceExternalDns,
	CreateClusterRoleNamespaceExternalDns,
	CreateCertNamespaceNginxDefaultServerSecretLocal,
	CreateCertNamespaceNginxDefaultServerSecretNonProd,
	CreateCertNamespaceNginxDefaultServerSecretProd,
	CreateConfigMapNamespaceNginxConfig,
	CreateCRDDnsendpointsExternaldnsNginxOrg,
	CreateCRDVirtualserverroutesK8sNginxOrg,
	CreateCRDTransportserversK8sNginxOrg,
	CreateCRDPoliciesK8sNginxOrg,
	CreateCRDVirtualserversK8sNginxOrg,
	CreateCRDGlobalconfigurationsK8sNginxOrg,
	CreateDaemonSetNamespaceNginxIngress,
	CreateDeploymentNamespaceNginxIngress,
	CreateIngressClassNginx,
	CreateServiceAccountNamespaceNginxIngress,
	CreateClusterRoleNginxIngress,
	CreateClusterRoleBindingNginxIngress,
	CreateServiceNamespaceNginxIngressAws,
	CreateServiceNamespaceNginxIngressGcpAzure,
	CreateCRDKongclusterpluginsConfigurationKonghqCom,
	CreateCRDKongconsumersConfigurationKonghqCom,
	CreateCRDKongingressesConfigurationKonghqCom,
	CreateCRDKongpluginsConfigurationKonghqCom,
	CreateCRDTcpingressesConfigurationKonghqCom,
	CreateCRDUdpingressesConfigurationKonghqCom,
	CreateDeploymentNamespaceIngressKong,
	CreateIngressClassKong,
	CreateServiceAccountNamespaceKongServiceaccount,
	CreateRoleNamespaceKongLeaderElection,
	CreateClusterRoleKongIngress,
	CreateRoleBindingNamespaceKongLeaderElection,
	CreateClusterRoleBindingKongIngress,
	CreateServiceNamespaceKongProxy,
	CreateServiceNamespaceKongValidationWebhook,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*platformv1alpha1.IngressComponent,
	*setupv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDDnsendpointsExternaldnsNginxOrg,
	CreateCRDVirtualserverroutesK8sNginxOrg,
	CreateCRDTransportserversK8sNginxOrg,
	CreateCRDPoliciesK8sNginxOrg,
	CreateCRDVirtualserversK8sNginxOrg,
	CreateCRDGlobalconfigurationsK8sNginxOrg,
	CreateCRDKongclusterpluginsConfigurationKonghqCom,
	CreateCRDKongconsumersConfigurationKonghqCom,
	CreateCRDKongingressesConfigurationKonghqCom,
	CreateCRDKongpluginsConfigurationKonghqCom,
	CreateCRDTcpingressesConfigurationKonghqCom,
	CreateCRDUdpingressesConfigurationKonghqCom,
}

func ConvertWorkload(component, collection workload.Workload) (
	*platformv1alpha1.IngressComponent,
	*setupv1alpha1.SupportServices,
	error,
) {
	p, ok := component.(*platformv1alpha1.IngressComponent)
	if !ok {
		return nil, nil, platformv1alpha1.ErrUnableToConvertIngressComponent
	}

	c, ok := collection.(*setupv1alpha1.SupportServices)
	if !ok {
		return nil, nil, setupv1alpha1.ErrUnableToConvertSupportServices
	}

	return p, c, nil
}
