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

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDKongclusterpluginsConfigurationKonghqCom creates the CustomResourceDefinition resource with name kongclusterplugins.configuration.konghq.com.
func CreateCRDKongclusterpluginsConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "kongclusterplugins.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "KongClusterPlugin",
					"listKind": "KongClusterPluginList",
					"plural":   "kongclusterplugins",
					"shortNames": []interface{}{
						"kcp",
					},
					"singular": "kongclusterplugin",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Name of the plugin",
								"jsonPath":    ".plugin",
								"name":        "Plugin-Type",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Age",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
							map[string]interface{}{
								"description": "Indicates if the plugin is disabled",
								"jsonPath":    ".disabled",
								"name":        "Disabled",
								"priority":    1,
								"type":        "boolean",
							},
							map[string]interface{}{
								"description": "Configuration of the plugin",
								"jsonPath":    ".config",
								"name":        "Config",
								"priority":    1,
								"type":        "string",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "KongClusterPlugin is the Schema for the kongclusterplugins API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"config": map[string]interface{}{
										"description":                          "Config contains the plugin configuration.",
										"type":                                 "object",
										"x-kubernetes-preserve-unknown-fields": true,
									},
									"configFrom": map[string]interface{}{
										"description": "ConfigFrom references a secret containing the plugin configuration.",
										"properties": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"description": "NamespacedSecretValueFromSource represents the source of a secret value specifying the secret namespace",
												"properties": map[string]interface{}{
													"key": map[string]interface{}{
														"description": "the key containing the value",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "the secret containing the key",
														"type":        "string",
													},
													"namespace": map[string]interface{}{
														"description": "The namespace containing the secret",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"key",
													"name",
													"namespace",
												},
												"type": "object",
											},
										},
										"type": "object",
									},
									"consumerRef": map[string]interface{}{
										"description": "ConsumerRef is a reference to a particular consumer",
										"type":        "string",
									},
									"disabled": map[string]interface{}{
										"description": "Disabled set if the plugin is disabled or not",
										"type":        "boolean",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"plugin": map[string]interface{}{
										"description": "PluginName is the name of the plugin to which to apply the config",
										"type":        "string",
									},
									"protocols": map[string]interface{}{
										"description": "Protocols configures plugin to run on requests received on specific protocols.",
										"items": map[string]interface{}{
											"enum": []interface{}{
												"http",
												"https",
												"grpc",
												"grpcs",
												"tcp",
												"tls",
												"udp",
											},
											"type": "string",
										},
										"type": "array",
									},
									"run_on": map[string]interface{}{
										"description": "RunOn configures the plugin to run on the first or the second or both nodes in case of a service mesh deployment.",
										"enum": []interface{}{
											"first",
											"second",
											"all",
										},
										"type": "string",
									},
								},
								"required": []interface{}{
									"plugin",
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDKongclusterpluginsConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDKongconsumersConfigurationKonghqCom creates the CustomResourceDefinition resource with name kongconsumers.configuration.konghq.com.
func CreateCRDKongconsumersConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "kongconsumers.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "KongConsumer",
					"listKind": "KongConsumerList",
					"plural":   "kongconsumers",
					"shortNames": []interface{}{
						"kc",
					},
					"singular": "kongconsumer",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Username of a Kong Consumer",
								"jsonPath":    ".username",
								"name":        "Username",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Age",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "KongConsumer is the Schema for the kongconsumers API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"credentials": map[string]interface{}{
										"description": "Credentials are references to secrets containing a credential to be provisioned in Kong.",
										"items": map[string]interface{}{
											"type": "string",
										},
										"type": "array",
									},
									"custom_id": map[string]interface{}{
										"description": "CustomID existing unique ID for the consumer - useful for mapping Kong with users in your existing database",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"username": map[string]interface{}{
										"description": "Username unique username of the consumer.",
										"type":        "string",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDKongconsumersConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDKongingressesConfigurationKonghqCom creates the CustomResourceDefinition resource with name kongingresses.configuration.konghq.com.
func CreateCRDKongingressesConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "kongingresses.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "KongIngress",
					"listKind": "KongIngressList",
					"plural":   "kongingresses",
					"shortNames": []interface{}{
						"ki",
					},
					"singular": "kongingress",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "KongIngress is the Schema for the kongingresses API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"proxy": map[string]interface{}{
										"description": "Proxy defines additional connection options for the routes to be configured in the Kong Gateway, e.g. `connection_timeout`, `retries`, e.t.c.",
										"properties": map[string]interface{}{
											"connect_timeout": map[string]interface{}{
												"description": "The timeout in milliseconds for establishing a connection to the upstream server.",
												"minimum":     0,
												"type":        "integer",
											},
											"path": map[string]interface{}{
												"description": "The path to be used in requests to the upstream server.(optional)",
												"pattern":     "^/.*$",
												"type":        "string",
											},
											"protocol": map[string]interface{}{
												"description": "The protocol used to communicate with the upstream.",
												"enum": []interface{}{
													"http",
													"https",
													"grpc",
													"grpcs",
													"tcp",
													"tls",
													"udp",
												},
												"type": "string",
											},
											"read_timeout": map[string]interface{}{
												"description": "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.",
												"minimum":     0,
												"type":        "integer",
											},
											"retries": map[string]interface{}{
												"description": "The number of retries to execute upon failure to proxy.",
												"minimum":     0,
												"type":        "integer",
											},
											"write_timeout": map[string]interface{}{
												"description": "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.",
												"minimum":     0,
												"type":        "integer",
											},
										},
										"type": "object",
									},
									"route": map[string]interface{}{
										"description": "Route define rules to match client requests. Each Route is associated with a Service, and a Service may have multiple Routes associated to it.",
										"properties": map[string]interface{}{
											"headers": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"items": map[string]interface{}{
														"type": "string",
													},
													"type": "array",
												},
												"description": "Headers contains one or more lists of values indexed by header name that will cause this Route to match if present in the request. The Host header cannot be used with this attribute.",
												"type":        "object",
											},
											"https_redirect_status_code": map[string]interface{}{
												"description": "HTTPSRedirectStatusCode is the status code Kong responds with when all properties of a Route match except the protocol.",
												"type":        "integer",
											},
											"methods": map[string]interface{}{
												"description": "Methods is a list of HTTP methods that match this Route.",
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"path_handling": map[string]interface{}{
												"description": "PathHandling controls how the Service path, Route path and requested path are combined when sending a request to the upstream.",
												"enum": []interface{}{
													"v0",
													"v1",
												},
												"type": "string",
											},
											"preserve_host": map[string]interface{}{
												"description": "PreserveHost sets When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers. If set to false, the upstream Host header will be that of the Service’s host.",
												"type":        "boolean",
											},
											"protocols": map[string]interface{}{
												"description": "Protocols is an array of the protocols this Route should allow.",
												"items": map[string]interface{}{
													"enum": []interface{}{
														"http",
														"https",
														"grpc",
														"grpcs",
														"tcp",
														"tls",
														"udp",
													},
													"type": "string",
												},
												"type": "array",
											},
											"regex_priority": map[string]interface{}{
												"description": "RegexPriority is a number used to choose which route resolves a given request when several routes match it using regexes simultaneously.",
												"type":        "integer",
											},
											"request_buffering": map[string]interface{}{
												"description": "RequestBuffering sets whether to enable request body buffering or not.",
												"type":        "boolean",
											},
											"response_buffering": map[string]interface{}{
												"description": "ResponseBuffering sets whether to enable response body buffering or not.",
												"type":        "boolean",
											},
											"snis": map[string]interface{}{
												"description": "SNIs is a list of SNIs that match this Route when using stream routing.",
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"strip_path": map[string]interface{}{
												"description": "StripPath sets When matching a Route via one of the paths strip the matching prefix from the upstream request URL.",
												"type":        "boolean",
											},
										},
										"type": "object",
									},
									"upstream": map[string]interface{}{
										"description": "Upstream represents a virtual hostname and can be used to loadbalance incoming requests over multiple targets (e.g. Kubernetes `Services` can be a target, OR `Endpoints` can be targets).",
										"properties": map[string]interface{}{
											"algorithm": map[string]interface{}{
												"description": "Algorithm is the load balancing algorithm to use.",
												"enum": []interface{}{
													"round-robin",
													"consistent-hashing",
													"least-connections",
												},
												"type": "string",
											},
											"hash_fallback": map[string]interface{}{
												"description": "HashFallback defines What to use as hashing input if the primary hash_on does not return a hash. Accepted values are: \"none\", \"consumer\", \"ip\", \"header\", \"cookie\".",
												"type":        "string",
											},
											"hash_fallback_header": map[string]interface{}{
												"description": "HashFallbackHeader is the header name to take the value from as hash input. Only required when \"hash_fallback\" is set to \"header\".",
												"type":        "string",
											},
											"hash_on": map[string]interface{}{
												"description": "HashOn defines what to use as hashing input. Accepted values are: \"none\", \"consumer\", \"ip\", \"header\", \"cookie\".",
												"type":        "string",
											},
											"hash_on_cookie": map[string]interface{}{
												"description": "The cookie name to take the value from as hash input. Only required when \"hash_on\" or \"hash_fallback\" is set to \"cookie\".",
												"type":        "string",
											},
											"hash_on_cookie_path": map[string]interface{}{
												"description": "The cookie path to set in the response headers. Only required when \"hash_on\" or \"hash_fallback\" is set to \"cookie\".",
												"type":        "string",
											},
											"hash_on_header": map[string]interface{}{
												"description": "HashOnHeader defines the header name to take the value from as hash input. Only required when \"hash_on\" is set to \"header\".",
												"type":        "string",
											},
											"healthchecks": map[string]interface{}{
												"description": "Healthchecks defines the health check configurations in Kong.",
												"properties": map[string]interface{}{
													"active": map[string]interface{}{
														"description": "ActiveHealthcheck configures active health check probing.",
														"properties": map[string]interface{}{
															"concurrency": map[string]interface{}{
																"minimum": 1,
																"type":    "integer",
															},
															"healthy": map[string]interface{}{
																"description": "Healthy configures thresholds and HTTP status codes to mark targets healthy for an upstream.",
																"properties": map[string]interface{}{
																	"http_statuses": map[string]interface{}{
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"interval": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"successes": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																},
																"type": "object",
															},
															"http_path": map[string]interface{}{
																"pattern": "^/.*$",
																"type":    "string",
															},
															"https_sni": map[string]interface{}{
																"type": "string",
															},
															"https_verify_certificate": map[string]interface{}{
																"type": "boolean",
															},
															"timeout": map[string]interface{}{
																"minimum": 0,
																"type":    "integer",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
															"unhealthy": map[string]interface{}{
																"description": "Unhealthy configures thresholds and HTTP status codes to mark targets unhealthy.",
																"properties": map[string]interface{}{
																	"http_failures": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"http_statuses": map[string]interface{}{
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"interval": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"tcp_failures": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"timeouts": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"passive": map[string]interface{}{
														"description": "PassiveHealthcheck configures passive checks around passive health checks.",
														"properties": map[string]interface{}{
															"healthy": map[string]interface{}{
																"description": "Healthy configures thresholds and HTTP status codes to mark targets healthy for an upstream.",
																"properties": map[string]interface{}{
																	"http_statuses": map[string]interface{}{
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"interval": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"successes": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																},
																"type": "object",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
															"unhealthy": map[string]interface{}{
																"description": "Unhealthy configures thresholds and HTTP status codes to mark targets unhealthy.",
																"properties": map[string]interface{}{
																	"http_failures": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"http_statuses": map[string]interface{}{
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"interval": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"tcp_failures": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																	"timeouts": map[string]interface{}{
																		"minimum": 0,
																		"type":    "integer",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"threshold": map[string]interface{}{
														"type": "number",
													},
												},
												"type": "object",
											},
											"host_header": map[string]interface{}{
												"description": "HostHeader is The hostname to be used as Host header when proxying requests through Kong.",
												"type":        "string",
											},
											"slots": map[string]interface{}{
												"description": "Slots is the number of slots in the load balancer algorithm.",
												"minimum":     10,
												"type":        "integer",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDKongingressesConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDKongpluginsConfigurationKonghqCom creates the CustomResourceDefinition resource with name kongplugins.configuration.konghq.com.
func CreateCRDKongpluginsConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "kongplugins.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "KongPlugin",
					"listKind": "KongPluginList",
					"plural":   "kongplugins",
					"shortNames": []interface{}{
						"kp",
					},
					"singular": "kongplugin",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Name of the plugin",
								"jsonPath":    ".plugin",
								"name":        "Plugin-Type",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Age",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
							map[string]interface{}{
								"description": "Indicates if the plugin is disabled",
								"jsonPath":    ".disabled",
								"name":        "Disabled",
								"priority":    1,
								"type":        "boolean",
							},
							map[string]interface{}{
								"description": "Configuration of the plugin",
								"jsonPath":    ".config",
								"name":        "Config",
								"priority":    1,
								"type":        "string",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "KongPlugin is the Schema for the kongplugins API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"config": map[string]interface{}{
										"description":                          "Config contains the plugin configuration.",
										"type":                                 "object",
										"x-kubernetes-preserve-unknown-fields": true,
									},
									"configFrom": map[string]interface{}{
										"description": "ConfigFrom references a secret containing the plugin configuration.",
										"properties": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"description": "SecretValueFromSource represents the source of a secret value",
												"properties": map[string]interface{}{
													"key": map[string]interface{}{
														"description": "the key containing the value",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "the secret containing the key",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"key",
													"name",
												},
												"type": "object",
											},
										},
										"type": "object",
									},
									"consumerRef": map[string]interface{}{
										"description": "ConsumerRef is a reference to a particular consumer",
										"type":        "string",
									},
									"disabled": map[string]interface{}{
										"description": "Disabled set if the plugin is disabled or not",
										"type":        "boolean",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"plugin": map[string]interface{}{
										"description": "PluginName is the name of the plugin to which to apply the config",
										"type":        "string",
									},
									"protocols": map[string]interface{}{
										"description": "Protocols configures plugin to run on requests received on specific protocols.",
										"items": map[string]interface{}{
											"enum": []interface{}{
												"http",
												"https",
												"grpc",
												"grpcs",
												"tcp",
												"tls",
												"udp",
											},
											"type": "string",
										},
										"type": "array",
									},
									"run_on": map[string]interface{}{
										"description": "RunOn configures the plugin to run on the first or the second or both nodes in case of a service mesh deployment.",
										"enum": []interface{}{
											"first",
											"second",
											"all",
										},
										"type": "string",
									},
								},
								"required": []interface{}{
									"plugin",
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDKongpluginsConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDTcpingressesConfigurationKonghqCom creates the CustomResourceDefinition resource with name tcpingresses.configuration.konghq.com.
func CreateCRDTcpingressesConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "tcpingresses.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "TCPIngress",
					"listKind": "TCPIngressList",
					"plural":   "tcpingresses",
					"singular": "tcpingress",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Address of the load balancer",
								"jsonPath":    ".status.loadBalancer.ingress[*].ip",
								"name":        "Address",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Age",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "TCPIngress is the Schema for the tcpingresses API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "TCPIngressSpec defines the desired state of TCPIngress",
										"properties": map[string]interface{}{
											"rules": map[string]interface{}{
												"description": "A list of rules used to configure the Ingress.",
												"items": map[string]interface{}{
													"description": "IngressRule represents a rule to apply against incoming requests. Matching is performed based on an (optional) SNI and port.",
													"properties": map[string]interface{}{
														"backend": map[string]interface{}{
															"description": "Backend defines the referenced service endpoint to which the traffic will be forwarded to.",
															"properties": map[string]interface{}{
																"serviceName": map[string]interface{}{
																	"description": "Specifies the name of the referenced service.",
																	"type":        "string",
																},
																"servicePort": map[string]interface{}{
																	"description": "Specifies the port of the referenced service.",
																	"format":      "int32",
																	"maximum":     65535,
																	"minimum":     1,
																	"type":        "integer",
																},
															},
															"required": []interface{}{
																"serviceName",
																"servicePort",
															},
															"type": "object",
														},
														"host": map[string]interface{}{
															"description": "Host is the fully qualified domain name of a network host, as defined by RFC 3986. If a Host is specified, the protocol must be TLS over TCP. A plain-text TCP request cannot be routed based on Host. It can only be routed based on Port.",
															"type":        "string",
														},
														"port": map[string]interface{}{
															"description": "Port is the port on which to accept TCP or TLS over TCP sessions and route. It is a required field. If a Host is not specified, the requested are routed based only on Port.",
															"format":      "int32",
															"maximum":     65535,
															"minimum":     1,
															"type":        "integer",
														},
													},
													"required": []interface{}{
														"backend",
													},
													"type": "object",
												},
												"type": "array",
											},
											"tls": map[string]interface{}{
												"description": "TLS configuration. This is similar to the `tls` section in the Ingress resource in networking.v1beta1 group. The mapping of SNIs to TLS cert-key pair defined here will be used for HTTP Ingress rules as well. Once can define the mapping in this resource or the original Ingress resource, both have the same effect.",
												"items": map[string]interface{}{
													"description": "IngressTLS describes the transport layer security.",
													"properties": map[string]interface{}{
														"hosts": map[string]interface{}{
															"description": "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
															"items": map[string]interface{}{
																"type": "string",
															},
															"type": "array",
														},
														"secretName": map[string]interface{}{
															"description": "SecretName is the name of the secret used to terminate SSL traffic.",
															"type":        "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "TCPIngressStatus defines the observed state of TCPIngress",
										"properties": map[string]interface{}{
											"loadBalancer": map[string]interface{}{
												"description": "LoadBalancer contains the current status of the load-balancer.",
												"properties": map[string]interface{}{
													"ingress": map[string]interface{}{
														"description": "Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.",
														"items": map[string]interface{}{
															"description": "LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.",
															"properties": map[string]interface{}{
																"hostname": map[string]interface{}{
																	"description": "Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)",
																	"type":        "string",
																},
																"ip": map[string]interface{}{
																	"description": "IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)",
																	"type":        "string",
																},
																"ports": map[string]interface{}{
																	"description": "Ports is a list of records of service ports If used, every port defined in the service should have an entry in it",
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"error": map[string]interface{}{
																				"description": "Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase. --- The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
																				"maxLength":   316,
																				"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`,
																				"type":        "string",
																			},
																			"port": map[string]interface{}{
																				"description": "Port is the port number of the service port of which status is recorded here",
																				"format":      "int32",
																				"type":        "integer",
																			},
																			"protocol": map[string]interface{}{
																				"default":     "TCP",
																				"description": "Protocol is the protocol of the service port of which status is recorded here The supported values are: \"TCP\", \"UDP\", \"SCTP\"",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"port",
																			"protocol",
																		},
																		"type": "object",
																	},
																	"type":                   "array",
																	"x-kubernetes-list-type": "atomic",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
												},
												"type": "object",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDTcpingressesConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDUdpingressesConfigurationKonghqCom creates the CustomResourceDefinition resource with name udpingresses.configuration.konghq.com.
func CreateCRDUdpingressesConfigurationKonghqCom(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Kong.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=kong.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.7.0",
				},
				"name": "udpingresses.configuration.konghq.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "kong-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"kong-ingress-controller",
					},
					"kind":     "UDPIngress",
					"listKind": "UDPIngressList",
					"plural":   "udpingresses",
					"singular": "udpingress",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Address of the load balancer",
								"jsonPath":    ".status.loadBalancer.ingress[*].ip",
								"name":        "Address",
								"type":        "string",
							},
							map[string]interface{}{
								"description": "Age",
								"jsonPath":    ".metadata.creationTimestamp",
								"name":        "Age",
								"type":        "date",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "UDPIngress is the Schema for the udpingresses API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "UDPIngressSpec defines the desired state of UDPIngress",
										"properties": map[string]interface{}{
											"rules": map[string]interface{}{
												"description": "A list of rules used to configure the Ingress.",
												"items": map[string]interface{}{
													"description": "UDPIngressRule represents a rule to apply against incoming requests wherein no Host matching is available for request routing, only the port is used to match requests.",
													"properties": map[string]interface{}{
														"backend": map[string]interface{}{
															"description": "Backend defines the Kubernetes service which accepts traffic from the listening Port defined above.",
															"properties": map[string]interface{}{
																"serviceName": map[string]interface{}{
																	"description": "Specifies the name of the referenced service.",
																	"type":        "string",
																},
																"servicePort": map[string]interface{}{
																	"description": "Specifies the port of the referenced service.",
																	"format":      "int32",
																	"maximum":     65535,
																	"minimum":     1,
																	"type":        "integer",
																},
															},
															"required": []interface{}{
																"serviceName",
																"servicePort",
															},
															"type": "object",
														},
														"port": map[string]interface{}{
															"description": "Port indicates the port for the Kong proxy to accept incoming traffic on, which will then be routed to the service Backend.",
															"type":        "integer",
														},
													},
													"required": []interface{}{
														"backend",
														"port",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "UDPIngressStatus defines the observed state of UDPIngress",
										"properties": map[string]interface{}{
											"loadBalancer": map[string]interface{}{
												"description": "LoadBalancer contains the current status of the load-balancer.",
												"properties": map[string]interface{}{
													"ingress": map[string]interface{}{
														"description": "Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.",
														"items": map[string]interface{}{
															"description": "LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.",
															"properties": map[string]interface{}{
																"hostname": map[string]interface{}{
																	"description": "Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)",
																	"type":        "string",
																},
																"ip": map[string]interface{}{
																	"description": "IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)",
																	"type":        "string",
																},
																"ports": map[string]interface{}{
																	"description": "Ports is a list of records of service ports If used, every port defined in the service should have an entry in it",
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"error": map[string]interface{}{
																				"description": "Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use   CamelCase names - cloud provider specific error values must have names that comply with the   format foo.example.com/CamelCase. --- The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
																				"maxLength":   316,
																				"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`,
																				"type":        "string",
																			},
																			"port": map[string]interface{}{
																				"description": "Port is the port number of the service port of which status is recorded here",
																				"format":      "int32",
																				"type":        "integer",
																			},
																			"protocol": map[string]interface{}{
																				"default":     "TCP",
																				"description": "Protocol is the protocol of the service port of which status is recorded here The supported values are: \"TCP\", \"UDP\", \"SCTP\"",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"port",
																			"protocol",
																		},
																		"type": "object",
																	},
																	"type":                   "array",
																	"x-kubernetes-list-type": "atomic",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
												},
												"type": "object",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
			},
			"status": map[string]interface{}{
				"acceptedNames": map[string]interface{}{
					"kind":   "",
					"plural": "",
				},
				"conditions":     []interface{}{},
				"storedVersions": []interface{}{},
			},
		},
	}

	return mutate.MutateCRDUdpingressesConfigurationKonghqCom(resourceObj, parent, collection, reconciler, req)
}
