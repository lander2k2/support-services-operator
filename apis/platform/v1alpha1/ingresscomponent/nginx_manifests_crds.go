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

// CreateCRDDnsendpointsExternaldnsNginxOrg creates the CustomResourceDefinition resource with name dnsendpoints.externaldns.nginx.org.
func CreateCRDDnsendpointsExternaldnsNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "dnsendpoints.externaldns.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "externaldns.nginx.org",
				"names": map[string]interface{}{
					"kind":     "DNSEndpoint",
					"listKind": "DNSEndpointList",
					"plural":   "dnsendpoints",
					"singular": "dnsendpoint",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "DNSEndpoint is the CRD wrapper for Endpoint",
								"type":        "object",
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
										"type": "object",
										"properties": map[string]interface{}{
											"endpoints": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"type": "object",
													"properties": map[string]interface{}{
														"dnsName": map[string]interface{}{
															"description": "The hostname for the DNS record",
															"type":        "string",
														},
														"labels": map[string]interface{}{
															"description": "Labels stores labels defined for the Endpoint",
															"type":        "object",
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
														},
														"providerSpecific": map[string]interface{}{
															"description": "ProviderSpecific stores provider specific config",
															"type":        "array",
															"items": map[string]interface{}{
																"type": "object",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"description": "Name of the property",
																		"type":        "string",
																	},
																	"value": map[string]interface{}{
																		"description": "Value of the property",
																		"type":        "string",
																	},
																},
															},
														},
														"recordTTL": map[string]interface{}{
															"description": "TTL for the record",
															"type":        "integer",
															"format":      "int64",
														},
														"recordType": map[string]interface{}{
															"description": "RecordType type of record, e.g. CNAME, A, SRV, TXT, MX",
															"type":        "string",
														},
														"targets": map[string]interface{}{
															"description": "The targets the DNS service points to",
															"type":        "array",
															"items": map[string]interface{}{
																"type": "string",
															},
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"observedGeneration": map[string]interface{}{
												"description": "The generation observed by by the external-dns controller.",
												"type":        "integer",
												"format":      "int64",
											},
										},
									},
								},
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

	return mutate.MutateCRDDnsendpointsExternaldnsNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDVirtualserverroutesK8sNginxOrg creates the CustomResourceDefinition resource with name virtualserverroutes.k8s.nginx.org.
func CreateCRDVirtualserverroutesK8sNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "virtualserverroutes.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "VirtualServerRoute",
					"listKind": "VirtualServerRouteList",
					"plural":   "virtualserverroutes",
					"shortNames": []interface{}{
						"vsr",
					},
					"singular": "virtualserverroute",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the VirtualServerRoute. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.host",
								"name":     "Host",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ip",
								"name":     "IP",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].hostname",
								"name":     "ExternalHostname",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ports",
								"name":     "Ports",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "VirtualServerRoute defines the VirtualServerRoute resource.",
								"type":        "object",
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
										"description": "VirtualServerRouteSpec is the spec of the VirtualServerRoute resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"host": map[string]interface{}{
												"type": "string",
											},
											"ingressClassName": map[string]interface{}{
												"type": "string",
											},
											"subroutes": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Route defines a route.",
													"type":        "object",
													"properties": map[string]interface{}{
														"action": map[string]interface{}{
															"description": "Action defines an action.",
															"type":        "object",
															"properties": map[string]interface{}{
																"pass": map[string]interface{}{
																	"type": "string",
																},
																"proxy": map[string]interface{}{
																	"description": "ActionProxy defines a proxy in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"requestHeaders": map[string]interface{}{
																			"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"pass": map[string]interface{}{
																					"type": "boolean",
																				},
																				"set": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"description": "Header defines an HTTP Header.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"type": "string",
																							},
																						},
																					},
																				},
																			},
																		},
																		"responseHeaders": map[string]interface{}{
																			"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"add": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"always": map[string]interface{}{
																								"type": "boolean",
																							},
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"type": "string",
																							},
																						},
																					},
																				},
																				"hide": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"ignore": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"pass": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																		"rewritePath": map[string]interface{}{
																			"type": "string",
																		},
																		"upstream": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
																"redirect": map[string]interface{}{
																	"description": "ActionRedirect defines a redirect in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"code": map[string]interface{}{
																			"type": "integer",
																		},
																		"url": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
																"return": map[string]interface{}{
																	"description": "ActionReturn defines a return in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"body": map[string]interface{}{
																			"type": "string",
																		},
																		"code": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
															},
														},
														"dos": map[string]interface{}{
															"type": "string",
														},
														"errorPages": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "ErrorPage defines an ErrorPage in a Route.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"codes": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																	},
																	"redirect": map[string]interface{}{
																		"description": "ErrorPageRedirect defines a redirect for an ErrorPage.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"code": map[string]interface{}{
																				"type": "integer",
																			},
																			"url": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																	"return": map[string]interface{}{
																		"description": "ErrorPageReturn defines a return for an ErrorPage.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"body": map[string]interface{}{
																				"type": "string",
																			},
																			"code": map[string]interface{}{
																				"type": "integer",
																			},
																			"headers": map[string]interface{}{
																				"type": "array",
																				"items": map[string]interface{}{
																					"description": "Header defines an HTTP Header.",
																					"type":        "object",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"type": "string",
																						},
																						"value": map[string]interface{}{
																							"type": "string",
																						},
																					},
																				},
																			},
																			"type": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																},
															},
														},
														"location-snippets": map[string]interface{}{
															"type": "string",
														},
														"matches": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "Match defines a match.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "Action defines an action.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"type": "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "ActionProxy defines a proxy in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"type": "boolean",
																							},
																							"set": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																						},
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"type": "boolean",
																										},
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																							"hide": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"ignore": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"pass": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																					"rewritePath": map[string]interface{}{
																						"type": "string",
																					},
																					"upstream": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"redirect": map[string]interface{}{
																				"description": "ActionRedirect defines a redirect in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"url": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"return": map[string]interface{}{
																				"description": "ActionReturn defines a return in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"type": "string",
																					},
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"type": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																	},
																	"conditions": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"description": "Condition defines a condition in a MatchRule.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"argument": map[string]interface{}{
																					"type": "string",
																				},
																				"cookie": map[string]interface{}{
																					"type": "string",
																				},
																				"header": map[string]interface{}{
																					"type": "string",
																				},
																				"value": map[string]interface{}{
																					"type": "string",
																				},
																				"variable": map[string]interface{}{
																					"type": "string",
																				},
																			},
																		},
																	},
																	"splits": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"description": "Split defines a split.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"action": map[string]interface{}{
																					"description": "Action defines an action.",
																					"type":        "object",
																					"properties": map[string]interface{}{
																						"pass": map[string]interface{}{
																							"type": "string",
																						},
																						"proxy": map[string]interface{}{
																							"description": "ActionProxy defines a proxy in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"requestHeaders": map[string]interface{}{
																									"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"pass": map[string]interface{}{
																											"type": "boolean",
																										},
																										"set": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"description": "Header defines an HTTP Header.",
																												"type":        "object",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"type": "string",
																													},
																												},
																											},
																										},
																									},
																								},
																								"responseHeaders": map[string]interface{}{
																									"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"add": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																												"type":        "object",
																												"properties": map[string]interface{}{
																													"always": map[string]interface{}{
																														"type": "boolean",
																													},
																													"name": map[string]interface{}{
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"type": "string",
																													},
																												},
																											},
																										},
																										"hide": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																										"ignore": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																										"pass": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																									},
																								},
																								"rewritePath": map[string]interface{}{
																									"type": "string",
																								},
																								"upstream": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																						"redirect": map[string]interface{}{
																							"description": "ActionRedirect defines a redirect in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"code": map[string]interface{}{
																									"type": "integer",
																								},
																								"url": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																						"return": map[string]interface{}{
																							"description": "ActionReturn defines a return in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"body": map[string]interface{}{
																									"type": "string",
																								},
																								"code": map[string]interface{}{
																									"type": "integer",
																								},
																								"type": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																				},
																				"weight": map[string]interface{}{
																					"type": "integer",
																				},
																			},
																		},
																	},
																},
															},
														},
														"path": map[string]interface{}{
															"type": "string",
														},
														"policies": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "PolicyReference references a policy by name and an optional namespace.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"type": "string",
																	},
																	"namespace": map[string]interface{}{
																		"type": "string",
																	},
																},
															},
														},
														"route": map[string]interface{}{
															"type": "string",
														},
														"splits": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "Split defines a split.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "Action defines an action.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"type": "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "ActionProxy defines a proxy in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"type": "boolean",
																							},
																							"set": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																						},
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"type": "boolean",
																										},
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																							"hide": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"ignore": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"pass": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																					"rewritePath": map[string]interface{}{
																						"type": "string",
																					},
																					"upstream": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"redirect": map[string]interface{}{
																				"description": "ActionRedirect defines a redirect in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"url": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"return": map[string]interface{}{
																				"description": "ActionReturn defines a return in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"type": "string",
																					},
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"type": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																	},
																	"weight": map[string]interface{}{
																		"type": "integer",
																	},
																},
															},
														},
													},
												},
											},
											"upstreams": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Upstream defines an upstream.",
													"type":        "object",
													"properties": map[string]interface{}{
														"buffer-size": map[string]interface{}{
															"type": "string",
														},
														"buffering": map[string]interface{}{
															"type": "boolean",
														},
														"buffers": map[string]interface{}{
															"description": "UpstreamBuffers defines Buffer Configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"number": map[string]interface{}{
																	"type": "integer",
																},
																"size": map[string]interface{}{
																	"type": "string",
																},
															},
														},
														"client-max-body-size": map[string]interface{}{
															"type": "string",
														},
														"connect-timeout": map[string]interface{}{
															"type": "string",
														},
														"fail-timeout": map[string]interface{}{
															"type": "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "HealthCheck defines the parameters for active Upstream HealthChecks.",
															"type":        "object",
															"properties": map[string]interface{}{
																"connect-timeout": map[string]interface{}{
																	"type": "string",
																},
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"fails": map[string]interface{}{
																	"type": "integer",
																},
																"grpcService": map[string]interface{}{
																	"type": "string",
																},
																"grpcStatus": map[string]interface{}{
																	"type": "integer",
																},
																"headers": map[string]interface{}{
																	"type": "array",
																	"items": map[string]interface{}{
																		"description": "Header defines an HTTP Header.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"value": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																},
																"interval": map[string]interface{}{
																	"type": "string",
																},
																"jitter": map[string]interface{}{
																	"type": "string",
																},
																"mandatory": map[string]interface{}{
																	"type": "boolean",
																},
																"passes": map[string]interface{}{
																	"type": "integer",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"persistent": map[string]interface{}{
																	"type": "boolean",
																},
																"port": map[string]interface{}{
																	"type": "integer",
																},
																"read-timeout": map[string]interface{}{
																	"type": "string",
																},
																"send-timeout": map[string]interface{}{
																	"type": "string",
																},
																"statusMatch": map[string]interface{}{
																	"type": "string",
																},
																"tls": map[string]interface{}{
																	"description": "UpstreamTLS defines a TLS configuration for an Upstream.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"enable": map[string]interface{}{
																			"type": "boolean",
																		},
																	},
																},
															},
														},
														"keepalive": map[string]interface{}{
															"type": "integer",
														},
														"lb-method": map[string]interface{}{
															"type": "string",
														},
														"max-conns": map[string]interface{}{
															"type": "integer",
														},
														"max-fails": map[string]interface{}{
															"type": "integer",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"next-upstream": map[string]interface{}{
															"type": "string",
														},
														"next-upstream-timeout": map[string]interface{}{
															"type": "string",
														},
														"next-upstream-tries": map[string]interface{}{
															"type": "integer",
														},
														"ntlm": map[string]interface{}{
															"type": "boolean",
														},
														"port": map[string]interface{}{
															"type": "integer",
														},
														"queue": map[string]interface{}{
															"description": "UpstreamQueue defines Queue Configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"size": map[string]interface{}{
																	"type": "integer",
																},
																"timeout": map[string]interface{}{
																	"type": "string",
																},
															},
														},
														"read-timeout": map[string]interface{}{
															"type": "string",
														},
														"send-timeout": map[string]interface{}{
															"type": "string",
														},
														"service": map[string]interface{}{
															"type": "string",
														},
														"sessionCookie": map[string]interface{}{
															"description": "SessionCookie defines the parameters for session persistence.",
															"type":        "object",
															"properties": map[string]interface{}{
																"domain": map[string]interface{}{
																	"type": "string",
																},
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"expires": map[string]interface{}{
																	"type": "string",
																},
																"httpOnly": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"secure": map[string]interface{}{
																	"type": "boolean",
																},
															},
														},
														"slow-start": map[string]interface{}{
															"type": "string",
														},
														"subselector": map[string]interface{}{
															"type": "object",
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
														},
														"tls": map[string]interface{}{
															"description": "UpstreamTLS defines a TLS configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
															},
														},
														"type": map[string]interface{}{
															"type": "string",
														},
														"use-cluster-ip": map[string]interface{}{
															"type": "boolean",
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"description": "VirtualServerRouteStatus defines the status for the VirtualServerRoute resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"externalEndpoints": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "ExternalEndpoint defines the IP/ Hostname and ports used to connect to this resource.",
													"type":        "object",
													"properties": map[string]interface{}{
														"hostname": map[string]interface{}{
															"type": "string",
														},
														"ip": map[string]interface{}{
															"type": "string",
														},
														"ports": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
											"message": map[string]interface{}{
												"type": "string",
											},
											"reason": map[string]interface{}{
												"type": "string",
											},
											"referencedBy": map[string]interface{}{
												"type": "string",
											},
											"state": map[string]interface{}{
												"type": "string",
											},
										},
									},
								},
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

	return mutate.MutateCRDVirtualserverroutesK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDTransportserversK8sNginxOrg creates the CustomResourceDefinition resource with name transportservers.k8s.nginx.org.
func CreateCRDTransportserversK8sNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "transportservers.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "TransportServer",
					"listKind": "TransportServerList",
					"plural":   "transportservers",
					"shortNames": []interface{}{
						"ts",
					},
					"singular": "transportserver",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the TransportServer. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.reason",
								"name":     "Reason",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "TransportServer defines the TransportServer resource.",
								"type":        "object",
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
										"description": "TransportServerSpec is the spec of the TransportServer resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"action": map[string]interface{}{
												"description": "Action defines an action.",
												"type":        "object",
												"properties": map[string]interface{}{
													"pass": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"host": map[string]interface{}{
												"type": "string",
											},
											"ingressClassName": map[string]interface{}{
												"type": "string",
											},
											"listener": map[string]interface{}{
												"description": "TransportServerListener defines a listener for a TransportServer.",
												"type":        "object",
												"properties": map[string]interface{}{
													"name": map[string]interface{}{
														"type": "string",
													},
													"protocol": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"serverSnippets": map[string]interface{}{
												"type": "string",
											},
											"sessionParameters": map[string]interface{}{
												"description": "SessionParameters defines session parameters.",
												"type":        "object",
												"properties": map[string]interface{}{
													"timeout": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"streamSnippets": map[string]interface{}{
												"type": "string",
											},
											"upstreamParameters": map[string]interface{}{
												"description": "UpstreamParameters defines parameters for an upstream.",
												"type":        "object",
												"properties": map[string]interface{}{
													"connectTimeout": map[string]interface{}{
														"type": "string",
													},
													"nextUpstream": map[string]interface{}{
														"type": "boolean",
													},
													"nextUpstreamTimeout": map[string]interface{}{
														"type": "string",
													},
													"nextUpstreamTries": map[string]interface{}{
														"type": "integer",
													},
													"udpRequests": map[string]interface{}{
														"type": "integer",
													},
													"udpResponses": map[string]interface{}{
														"type": "integer",
													},
												},
											},
											"upstreams": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Upstream defines an upstream.",
													"type":        "object",
													"properties": map[string]interface{}{
														"failTimeout": map[string]interface{}{
															"type": "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "HealthCheck defines the parameters for active Upstream HealthChecks.",
															"type":        "object",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"fails": map[string]interface{}{
																	"type": "integer",
																},
																"interval": map[string]interface{}{
																	"type": "string",
																},
																"jitter": map[string]interface{}{
																	"type": "string",
																},
																"match": map[string]interface{}{
																	"description": "Match defines the parameters of a custom health check.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"expect": map[string]interface{}{
																			"type": "string",
																		},
																		"send": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
																"passes": map[string]interface{}{
																	"type": "integer",
																},
																"port": map[string]interface{}{
																	"type": "integer",
																},
																"timeout": map[string]interface{}{
																	"type": "string",
																},
															},
														},
														"loadBalancingMethod": map[string]interface{}{
															"type": "string",
														},
														"maxConns": map[string]interface{}{
															"type": "integer",
														},
														"maxFails": map[string]interface{}{
															"type": "integer",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"port": map[string]interface{}{
															"type": "integer",
														},
														"service": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"description": "TransportServerStatus defines the status for the TransportServer resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"message": map[string]interface{}{
												"type": "string",
											},
											"reason": map[string]interface{}{
												"type": "string",
											},
											"state": map[string]interface{}{
												"type": "string",
											},
										},
									},
								},
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

	return mutate.MutateCRDTransportserversK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDPoliciesK8sNginxOrg creates the CustomResourceDefinition resource with name policies.k8s.nginx.org.
func CreateCRDPoliciesK8sNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "policies.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "Policy",
					"listKind": "PolicyList",
					"plural":   "policies",
					"shortNames": []interface{}{
						"pol",
					},
					"singular": "policy",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the Policy. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
								"type":        "object",
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
										"description": "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
										"type":        "object",
										"properties": map[string]interface{}{
											"accessControl": map[string]interface{}{
												"description": "AccessControl defines an access policy based on the source IP of a request.",
												"type":        "object",
												"properties": map[string]interface{}{
													"allow": map[string]interface{}{
														"type": "array",
														"items": map[string]interface{}{
															"type": "string",
														},
													},
													"deny": map[string]interface{}{
														"type": "array",
														"items": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
											"basicAuth": map[string]interface{}{
												"description": "BasicAuth holds HTTP Basic authentication configuration policy status: preview",
												"type":        "object",
												"properties": map[string]interface{}{
													"realm": map[string]interface{}{
														"type": "string",
													},
													"secret": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"egressMTLS": map[string]interface{}{
												"description": "EgressMTLS defines an Egress MTLS policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"ciphers": map[string]interface{}{
														"type": "string",
													},
													"protocols": map[string]interface{}{
														"type": "string",
													},
													"serverName": map[string]interface{}{
														"type": "boolean",
													},
													"sessionReuse": map[string]interface{}{
														"type": "boolean",
													},
													"sslName": map[string]interface{}{
														"type": "string",
													},
													"tlsSecret": map[string]interface{}{
														"type": "string",
													},
													"trustedCertSecret": map[string]interface{}{
														"type": "string",
													},
													"verifyDepth": map[string]interface{}{
														"type": "integer",
													},
													"verifyServer": map[string]interface{}{
														"type": "boolean",
													},
												},
											},
											"ingressClassName": map[string]interface{}{
												"type": "string",
											},
											"ingressMTLS": map[string]interface{}{
												"description": "IngressMTLS defines an Ingress MTLS policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"clientCertSecret": map[string]interface{}{
														"type": "string",
													},
													"verifyClient": map[string]interface{}{
														"type": "string",
													},
													"verifyDepth": map[string]interface{}{
														"type": "integer",
													},
												},
											},
											"jwt": map[string]interface{}{
												"description": "JWTAuth holds JWT authentication configuration.",
												"type":        "object",
												"properties": map[string]interface{}{
													"realm": map[string]interface{}{
														"type": "string",
													},
													"secret": map[string]interface{}{
														"type": "string",
													},
													"token": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"oidc": map[string]interface{}{
												"description": "OIDC defines an Open ID Connect policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"authEndpoint": map[string]interface{}{
														"type": "string",
													},
													"clientID": map[string]interface{}{
														"type": "string",
													},
													"clientSecret": map[string]interface{}{
														"type": "string",
													},
													"jwksURI": map[string]interface{}{
														"type": "string",
													},
													"redirectURI": map[string]interface{}{
														"type": "string",
													},
													"scope": map[string]interface{}{
														"type": "string",
													},
													"tokenEndpoint": map[string]interface{}{
														"type": "string",
													},
													"zoneSyncLeeway": map[string]interface{}{
														"type": "integer",
													},
												},
											},
											"rateLimit": map[string]interface{}{
												"description": "RateLimit defines a rate limit policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"burst": map[string]interface{}{
														"type": "integer",
													},
													"delay": map[string]interface{}{
														"type": "integer",
													},
													"dryRun": map[string]interface{}{
														"type": "boolean",
													},
													"key": map[string]interface{}{
														"type": "string",
													},
													"logLevel": map[string]interface{}{
														"type": "string",
													},
													"noDelay": map[string]interface{}{
														"type": "boolean",
													},
													"rate": map[string]interface{}{
														"type": "string",
													},
													"rejectCode": map[string]interface{}{
														"type": "integer",
													},
													"zoneSize": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"waf": map[string]interface{}{
												"description": "WAF defines an WAF policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"apPolicy": map[string]interface{}{
														"type": "string",
													},
													"enable": map[string]interface{}{
														"type": "boolean",
													},
													"securityLog": map[string]interface{}{
														"description": "SecurityLog defines the security log of a WAF policy.",
														"type":        "object",
														"properties": map[string]interface{}{
															"apLogConf": map[string]interface{}{
																"type": "string",
															},
															"enable": map[string]interface{}{
																"type": "boolean",
															},
															"logDest": map[string]interface{}{
																"type": "string",
															},
														},
													},
													"securityLogs": map[string]interface{}{
														"type": "array",
														"items": map[string]interface{}{
															"description": "SecurityLog defines the security log of a WAF policy.",
															"type":        "object",
															"properties": map[string]interface{}{
																"apLogConf": map[string]interface{}{
																	"type": "string",
																},
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"logDest": map[string]interface{}{
																	"type": "string",
																},
															},
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"description": "PolicyStatus is the status of the policy resource",
										"type":        "object",
										"properties": map[string]interface{}{
											"message": map[string]interface{}{
												"type": "string",
											},
											"reason": map[string]interface{}{
												"type": "string",
											},
											"state": map[string]interface{}{
												"type": "string",
											},
										},
									},
								},
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
								"type":        "object",
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
										"description": "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
										"type":        "object",
										"properties": map[string]interface{}{
											"accessControl": map[string]interface{}{
												"description": "AccessControl defines an access policy based on the source IP of a request.",
												"type":        "object",
												"properties": map[string]interface{}{
													"allow": map[string]interface{}{
														"type": "array",
														"items": map[string]interface{}{
															"type": "string",
														},
													},
													"deny": map[string]interface{}{
														"type": "array",
														"items": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
											"egressMTLS": map[string]interface{}{
												"description": "EgressMTLS defines an Egress MTLS policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"ciphers": map[string]interface{}{
														"type": "string",
													},
													"protocols": map[string]interface{}{
														"type": "string",
													},
													"serverName": map[string]interface{}{
														"type": "boolean",
													},
													"sessionReuse": map[string]interface{}{
														"type": "boolean",
													},
													"sslName": map[string]interface{}{
														"type": "string",
													},
													"tlsSecret": map[string]interface{}{
														"type": "string",
													},
													"trustedCertSecret": map[string]interface{}{
														"type": "string",
													},
													"verifyDepth": map[string]interface{}{
														"type": "integer",
													},
													"verifyServer": map[string]interface{}{
														"type": "boolean",
													},
												},
											},
											"ingressMTLS": map[string]interface{}{
												"description": "IngressMTLS defines an Ingress MTLS policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"clientCertSecret": map[string]interface{}{
														"type": "string",
													},
													"verifyClient": map[string]interface{}{
														"type": "string",
													},
													"verifyDepth": map[string]interface{}{
														"type": "integer",
													},
												},
											},
											"jwt": map[string]interface{}{
												"description": "JWTAuth holds JWT authentication configuration.",
												"type":        "object",
												"properties": map[string]interface{}{
													"realm": map[string]interface{}{
														"type": "string",
													},
													"secret": map[string]interface{}{
														"type": "string",
													},
													"token": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"rateLimit": map[string]interface{}{
												"description": "RateLimit defines a rate limit policy.",
												"type":        "object",
												"properties": map[string]interface{}{
													"burst": map[string]interface{}{
														"type": "integer",
													},
													"delay": map[string]interface{}{
														"type": "integer",
													},
													"dryRun": map[string]interface{}{
														"type": "boolean",
													},
													"key": map[string]interface{}{
														"type": "string",
													},
													"logLevel": map[string]interface{}{
														"type": "string",
													},
													"noDelay": map[string]interface{}{
														"type": "boolean",
													},
													"rate": map[string]interface{}{
														"type": "string",
													},
													"rejectCode": map[string]interface{}{
														"type": "integer",
													},
													"zoneSize": map[string]interface{}{
														"type": "string",
													},
												},
											},
										},
									},
								},
							},
						},
						"served":  true,
						"storage": false,
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

	return mutate.MutateCRDPoliciesK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDVirtualserversK8sNginxOrg creates the CustomResourceDefinition resource with name virtualservers.k8s.nginx.org.
func CreateCRDVirtualserversK8sNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "virtualservers.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "VirtualServer",
					"listKind": "VirtualServerList",
					"plural":   "virtualservers",
					"shortNames": []interface{}{
						"vs",
					},
					"singular": "virtualserver",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the VirtualServer. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.host",
								"name":     "Host",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ip",
								"name":     "IP",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].hostname",
								"name":     "ExternalHostname",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ports",
								"name":     "Ports",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "VirtualServer defines the VirtualServer resource.",
								"type":        "object",
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
										"description": "VirtualServerSpec is the spec of the VirtualServer resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"dos": map[string]interface{}{
												"type": "string",
											},
											"externalDNS": map[string]interface{}{
												"description": "ExternalDNS defines externaldns sub-resource of a virtual server.",
												"type":        "object",
												"properties": map[string]interface{}{
													"enable": map[string]interface{}{
														"type": "boolean",
													},
													"labels": map[string]interface{}{
														"description": "Labels stores labels defined for the Endpoint",
														"type":        "object",
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
													},
													"providerSpecific": map[string]interface{}{
														"description": "ProviderSpecific stores provider specific config",
														"type":        "array",
														"items": map[string]interface{}{
															"description": "ProviderSpecificProperty defines specific property for using with ExternalDNS sub-resource.",
															"type":        "object",
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"description": "Name of the property",
																	"type":        "string",
																},
																"value": map[string]interface{}{
																	"description": "Value of the property",
																	"type":        "string",
																},
															},
														},
													},
													"recordTTL": map[string]interface{}{
														"description": "TTL for the record",
														"type":        "integer",
														"format":      "int64",
													},
													"recordType": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"host": map[string]interface{}{
												"type": "string",
											},
											"http-snippets": map[string]interface{}{
												"type": "string",
											},
											"ingressClassName": map[string]interface{}{
												"type": "string",
											},
											"policies": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "PolicyReference references a policy by name and an optional namespace.",
													"type":        "object",
													"properties": map[string]interface{}{
														"name": map[string]interface{}{
															"type": "string",
														},
														"namespace": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
											"routes": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Route defines a route.",
													"type":        "object",
													"properties": map[string]interface{}{
														"action": map[string]interface{}{
															"description": "Action defines an action.",
															"type":        "object",
															"properties": map[string]interface{}{
																"pass": map[string]interface{}{
																	"type": "string",
																},
																"proxy": map[string]interface{}{
																	"description": "ActionProxy defines a proxy in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"requestHeaders": map[string]interface{}{
																			"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"pass": map[string]interface{}{
																					"type": "boolean",
																				},
																				"set": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"description": "Header defines an HTTP Header.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"type": "string",
																							},
																						},
																					},
																				},
																			},
																		},
																		"responseHeaders": map[string]interface{}{
																			"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"add": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"always": map[string]interface{}{
																								"type": "boolean",
																							},
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"type": "string",
																							},
																						},
																					},
																				},
																				"hide": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"ignore": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"pass": map[string]interface{}{
																					"type": "array",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																		"rewritePath": map[string]interface{}{
																			"type": "string",
																		},
																		"upstream": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
																"redirect": map[string]interface{}{
																	"description": "ActionRedirect defines a redirect in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"code": map[string]interface{}{
																			"type": "integer",
																		},
																		"url": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
																"return": map[string]interface{}{
																	"description": "ActionReturn defines a return in an Action.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"body": map[string]interface{}{
																			"type": "string",
																		},
																		"code": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": map[string]interface{}{
																			"type": "string",
																		},
																	},
																},
															},
														},
														"dos": map[string]interface{}{
															"type": "string",
														},
														"errorPages": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "ErrorPage defines an ErrorPage in a Route.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"codes": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																	},
																	"redirect": map[string]interface{}{
																		"description": "ErrorPageRedirect defines a redirect for an ErrorPage.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"code": map[string]interface{}{
																				"type": "integer",
																			},
																			"url": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																	"return": map[string]interface{}{
																		"description": "ErrorPageReturn defines a return for an ErrorPage.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"body": map[string]interface{}{
																				"type": "string",
																			},
																			"code": map[string]interface{}{
																				"type": "integer",
																			},
																			"headers": map[string]interface{}{
																				"type": "array",
																				"items": map[string]interface{}{
																					"description": "Header defines an HTTP Header.",
																					"type":        "object",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"type": "string",
																						},
																						"value": map[string]interface{}{
																							"type": "string",
																						},
																					},
																				},
																			},
																			"type": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																},
															},
														},
														"location-snippets": map[string]interface{}{
															"type": "string",
														},
														"matches": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "Match defines a match.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "Action defines an action.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"type": "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "ActionProxy defines a proxy in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"type": "boolean",
																							},
																							"set": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																						},
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"type": "boolean",
																										},
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																							"hide": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"ignore": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"pass": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																					"rewritePath": map[string]interface{}{
																						"type": "string",
																					},
																					"upstream": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"redirect": map[string]interface{}{
																				"description": "ActionRedirect defines a redirect in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"url": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"return": map[string]interface{}{
																				"description": "ActionReturn defines a return in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"type": "string",
																					},
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"type": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																	},
																	"conditions": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"description": "Condition defines a condition in a MatchRule.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"argument": map[string]interface{}{
																					"type": "string",
																				},
																				"cookie": map[string]interface{}{
																					"type": "string",
																				},
																				"header": map[string]interface{}{
																					"type": "string",
																				},
																				"value": map[string]interface{}{
																					"type": "string",
																				},
																				"variable": map[string]interface{}{
																					"type": "string",
																				},
																			},
																		},
																	},
																	"splits": map[string]interface{}{
																		"type": "array",
																		"items": map[string]interface{}{
																			"description": "Split defines a split.",
																			"type":        "object",
																			"properties": map[string]interface{}{
																				"action": map[string]interface{}{
																					"description": "Action defines an action.",
																					"type":        "object",
																					"properties": map[string]interface{}{
																						"pass": map[string]interface{}{
																							"type": "string",
																						},
																						"proxy": map[string]interface{}{
																							"description": "ActionProxy defines a proxy in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"requestHeaders": map[string]interface{}{
																									"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"pass": map[string]interface{}{
																											"type": "boolean",
																										},
																										"set": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"description": "Header defines an HTTP Header.",
																												"type":        "object",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"type": "string",
																													},
																												},
																											},
																										},
																									},
																								},
																								"responseHeaders": map[string]interface{}{
																									"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"add": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																												"type":        "object",
																												"properties": map[string]interface{}{
																													"always": map[string]interface{}{
																														"type": "boolean",
																													},
																													"name": map[string]interface{}{
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"type": "string",
																													},
																												},
																											},
																										},
																										"hide": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																										"ignore": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																										"pass": map[string]interface{}{
																											"type": "array",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																										},
																									},
																								},
																								"rewritePath": map[string]interface{}{
																									"type": "string",
																								},
																								"upstream": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																						"redirect": map[string]interface{}{
																							"description": "ActionRedirect defines a redirect in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"code": map[string]interface{}{
																									"type": "integer",
																								},
																								"url": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																						"return": map[string]interface{}{
																							"description": "ActionReturn defines a return in an Action.",
																							"type":        "object",
																							"properties": map[string]interface{}{
																								"body": map[string]interface{}{
																									"type": "string",
																								},
																								"code": map[string]interface{}{
																									"type": "integer",
																								},
																								"type": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																				},
																				"weight": map[string]interface{}{
																					"type": "integer",
																				},
																			},
																		},
																	},
																},
															},
														},
														"path": map[string]interface{}{
															"type": "string",
														},
														"policies": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "PolicyReference references a policy by name and an optional namespace.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"type": "string",
																	},
																	"namespace": map[string]interface{}{
																		"type": "string",
																	},
																},
															},
														},
														"route": map[string]interface{}{
															"type": "string",
														},
														"splits": map[string]interface{}{
															"type": "array",
															"items": map[string]interface{}{
																"description": "Split defines a split.",
																"type":        "object",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "Action defines an action.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"type": "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "ActionProxy defines a proxy in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "ProxyRequestHeaders defines the request headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"type": "boolean",
																							},
																							"set": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																						},
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "ProxyResponseHeaders defines the response headers manipulation in an ActionProxy.",
																						"type":        "object",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"type":        "object",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"type": "boolean",
																										},
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"value": map[string]interface{}{
																											"type": "string",
																										},
																									},
																								},
																							},
																							"hide": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"ignore": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"pass": map[string]interface{}{
																								"type": "array",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																							},
																						},
																					},
																					"rewritePath": map[string]interface{}{
																						"type": "string",
																					},
																					"upstream": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"redirect": map[string]interface{}{
																				"description": "ActionRedirect defines a redirect in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"url": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																			"return": map[string]interface{}{
																				"description": "ActionReturn defines a return in an Action.",
																				"type":        "object",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"type": "string",
																					},
																					"code": map[string]interface{}{
																						"type": "integer",
																					},
																					"type": map[string]interface{}{
																						"type": "string",
																					},
																				},
																			},
																		},
																	},
																	"weight": map[string]interface{}{
																		"type": "integer",
																	},
																},
															},
														},
													},
												},
											},
											"server-snippets": map[string]interface{}{
												"type": "string",
											},
											"tls": map[string]interface{}{
												"description": "TLS defines TLS configuration for a VirtualServer.",
												"type":        "object",
												"properties": map[string]interface{}{
													"cert-manager": map[string]interface{}{
														"description": "CertManager defines a cert manager config for a TLS.",
														"type":        "object",
														"properties": map[string]interface{}{
															"cluster-issuer": map[string]interface{}{
																"type": "string",
															},
															"common-name": map[string]interface{}{
																"type": "string",
															},
															"duration": map[string]interface{}{
																"type": "string",
															},
															"issuer": map[string]interface{}{
																"type": "string",
															},
															"issuer-group": map[string]interface{}{
																"type": "string",
															},
															"issuer-kind": map[string]interface{}{
																"type": "string",
															},
															"renew-before": map[string]interface{}{
																"type": "string",
															},
															"usages": map[string]interface{}{
																"type": "string",
															},
														},
													},
													"redirect": map[string]interface{}{
														"description": "TLSRedirect defines a redirect for a TLS.",
														"type":        "object",
														"properties": map[string]interface{}{
															"basedOn": map[string]interface{}{
																"type": "string",
															},
															"code": map[string]interface{}{
																"type": "integer",
															},
															"enable": map[string]interface{}{
																"type": "boolean",
															},
														},
													},
													"secret": map[string]interface{}{
														"type": "string",
													},
												},
											},
											"upstreams": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Upstream defines an upstream.",
													"type":        "object",
													"properties": map[string]interface{}{
														"buffer-size": map[string]interface{}{
															"type": "string",
														},
														"buffering": map[string]interface{}{
															"type": "boolean",
														},
														"buffers": map[string]interface{}{
															"description": "UpstreamBuffers defines Buffer Configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"number": map[string]interface{}{
																	"type": "integer",
																},
																"size": map[string]interface{}{
																	"type": "string",
																},
															},
														},
														"client-max-body-size": map[string]interface{}{
															"type": "string",
														},
														"connect-timeout": map[string]interface{}{
															"type": "string",
														},
														"fail-timeout": map[string]interface{}{
															"type": "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "HealthCheck defines the parameters for active Upstream HealthChecks.",
															"type":        "object",
															"properties": map[string]interface{}{
																"connect-timeout": map[string]interface{}{
																	"type": "string",
																},
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"fails": map[string]interface{}{
																	"type": "integer",
																},
																"grpcService": map[string]interface{}{
																	"type": "string",
																},
																"grpcStatus": map[string]interface{}{
																	"type": "integer",
																},
																"headers": map[string]interface{}{
																	"type": "array",
																	"items": map[string]interface{}{
																		"description": "Header defines an HTTP Header.",
																		"type":        "object",
																		"properties": map[string]interface{}{
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"value": map[string]interface{}{
																				"type": "string",
																			},
																		},
																	},
																},
																"interval": map[string]interface{}{
																	"type": "string",
																},
																"jitter": map[string]interface{}{
																	"type": "string",
																},
																"mandatory": map[string]interface{}{
																	"type": "boolean",
																},
																"passes": map[string]interface{}{
																	"type": "integer",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"persistent": map[string]interface{}{
																	"type": "boolean",
																},
																"port": map[string]interface{}{
																	"type": "integer",
																},
																"read-timeout": map[string]interface{}{
																	"type": "string",
																},
																"send-timeout": map[string]interface{}{
																	"type": "string",
																},
																"statusMatch": map[string]interface{}{
																	"type": "string",
																},
																"tls": map[string]interface{}{
																	"description": "UpstreamTLS defines a TLS configuration for an Upstream.",
																	"type":        "object",
																	"properties": map[string]interface{}{
																		"enable": map[string]interface{}{
																			"type": "boolean",
																		},
																	},
																},
															},
														},
														"keepalive": map[string]interface{}{
															"type": "integer",
														},
														"lb-method": map[string]interface{}{
															"type": "string",
														},
														"max-conns": map[string]interface{}{
															"type": "integer",
														},
														"max-fails": map[string]interface{}{
															"type": "integer",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"next-upstream": map[string]interface{}{
															"type": "string",
														},
														"next-upstream-timeout": map[string]interface{}{
															"type": "string",
														},
														"next-upstream-tries": map[string]interface{}{
															"type": "integer",
														},
														"ntlm": map[string]interface{}{
															"type": "boolean",
														},
														"port": map[string]interface{}{
															"type": "integer",
														},
														"queue": map[string]interface{}{
															"description": "UpstreamQueue defines Queue Configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"size": map[string]interface{}{
																	"type": "integer",
																},
																"timeout": map[string]interface{}{
																	"type": "string",
																},
															},
														},
														"read-timeout": map[string]interface{}{
															"type": "string",
														},
														"send-timeout": map[string]interface{}{
															"type": "string",
														},
														"service": map[string]interface{}{
															"type": "string",
														},
														"sessionCookie": map[string]interface{}{
															"description": "SessionCookie defines the parameters for session persistence.",
															"type":        "object",
															"properties": map[string]interface{}{
																"domain": map[string]interface{}{
																	"type": "string",
																},
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
																"expires": map[string]interface{}{
																	"type": "string",
																},
																"httpOnly": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"secure": map[string]interface{}{
																	"type": "boolean",
																},
															},
														},
														"slow-start": map[string]interface{}{
															"type": "string",
														},
														"subselector": map[string]interface{}{
															"type": "object",
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
														},
														"tls": map[string]interface{}{
															"description": "UpstreamTLS defines a TLS configuration for an Upstream.",
															"type":        "object",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"type": "boolean",
																},
															},
														},
														"type": map[string]interface{}{
															"type": "string",
														},
														"use-cluster-ip": map[string]interface{}{
															"type": "boolean",
														},
													},
												},
											},
										},
									},
									"status": map[string]interface{}{
										"description": "VirtualServerStatus defines the status for the VirtualServer resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"externalEndpoints": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "ExternalEndpoint defines the IP/ Hostname and ports used to connect to this resource.",
													"type":        "object",
													"properties": map[string]interface{}{
														"hostname": map[string]interface{}{
															"type": "string",
														},
														"ip": map[string]interface{}{
															"type": "string",
														},
														"ports": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
											"message": map[string]interface{}{
												"type": "string",
											},
											"reason": map[string]interface{}{
												"type": "string",
											},
											"state": map[string]interface{}{
												"type": "string",
											},
										},
									},
								},
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

	return mutate.MutateCRDVirtualserversK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDGlobalconfigurationsK8sNginxOrg creates the CustomResourceDefinition resource with name globalconfigurations.k8s.nginx.org.
func CreateCRDGlobalconfigurationsK8sNginxOrg(
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
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.8.0",
				},
				"creationTimestamp": nil,
				"name":              "globalconfigurations.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "GlobalConfiguration",
					"listKind": "GlobalConfigurationList",
					"plural":   "globalconfigurations",
					"shortNames": []interface{}{
						"gc",
					},
					"singular": "globalconfiguration",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "GlobalConfiguration defines the GlobalConfiguration resource.",
								"type":        "object",
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
										"description": "GlobalConfigurationSpec is the spec of the GlobalConfiguration resource.",
										"type":        "object",
										"properties": map[string]interface{}{
											"listeners": map[string]interface{}{
												"type": "array",
												"items": map[string]interface{}{
													"description": "Listener defines a listener.",
													"type":        "object",
													"properties": map[string]interface{}{
														"name": map[string]interface{}{
															"type": "string",
														},
														"port": map[string]interface{}{
															"type": "integer",
														},
														"protocol": map[string]interface{}{
															"type": "string",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"served":  true,
						"storage": true,
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

	return mutate.MutateCRDGlobalconfigurationsK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}
