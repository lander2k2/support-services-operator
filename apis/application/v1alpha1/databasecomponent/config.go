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

package databasecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	applicationv1alpha1 "github.com/nukleros/support-services-operator/apis/application/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/application/v1alpha1/databasecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapNamespacePostgresOperator creates the ConfigMap resource with name postgres-operator.
func CreateConfigMapNamespacePostgresOperator(
	parent *applicationv1alpha1.DatabaseComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "postgres-operator",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"data": map[string]interface{}{
				"api_port":                            "8080",
				"aws_region":                          "eu-central-1",
				"cluster_domain":                      "cluster.local",
				"cluster_history_entries":             "1000",
				"cluster_labels":                      "application:spilo",
				"cluster_name_label":                  "cluster-name",
				"connection_pooler_image":             "registry.opensource.zalan.do/acid/pgbouncer:master-22",
				"crd_categories":                      "all",
				"db_hosted_zone":                      "db.example.com",
				"debug_logging":                       "true",
				"docker_image":                        "registry.opensource.zalan.do/acid/spilo-14:2.1-p6",
				"enable_ebs_gp3_migration":            "false",
				"enable_master_load_balancer":         "false",
				"enable_master_pooler_load_balancer":  "false",
				"enable_password_rotation":            "false",
				"enable_pgversion_env_var":            "true",
				"enable_replica_load_balancer":        "false",
				"enable_replica_pooler_load_balancer": "false",
				"enable_spilo_wal_path_compat":        "true",
				"enable_team_member_deprecation":      "false",
				"enable_teams_api":                    "false",
				"external_traffic_policy":             "Cluster",
				"logical_backup_docker_image":         "registry.opensource.zalan.do/acid/logical-backup:v1.8.2",
				"logical_backup_job_prefix":           "logical-backup-",
				"logical_backup_provider":             "s3",
				"logical_backup_s3_bucket":            "my-bucket-url",
				"logical_backup_s3_sse":               "AES256",
				"logical_backup_schedule":             "30 00 * * *",
				"major_version_upgrade_mode":          "manual",
				"master_dns_name_format":              "{cluster}.{team}.{hostedzone}",
				"patroni_api_check_interval":          "1s",
				"patroni_api_check_timeout":           "5s",
				"pdb_name_format":                     "postgres-{cluster}-pdb",
				"pod_deletion_wait_timeout":           "10m",
				"pod_label_wait_timeout":              "10m",
				"pod_management_policy":               "ordered_ready",
				"pod_role_label":                      "spilo-role",
				"pod_service_account_name":            "postgres-pod",
				"pod_terminate_grace_period":          "5m",
				"ready_wait_interval":                 "3s",
				"ready_wait_timeout":                  "30s",
				"repair_period":                       "5m",
				"replica_dns_name_format":             "{cluster}-repl.{team}.{hostedzone}",
				"replication_username":                "standby",
				"resource_check_interval":             "3s",
				"resource_check_timeout":              "10m",
				"resync_period":                       "30m",
				"ring_log_lines":                      "100",
				"role_deletion_suffix":                "_deleted",
				"secret_name_template":                "{username}.{cluster}.credentials.{tprkind}.{tprgroup}",
				"spilo_allow_privilege_escalation":    "true",
				"spilo_privileged":                    "false",
				"storage_resize_mode":                 "pvc",
				"super_username":                      "postgres",
				"watched_namespace":                   "*",
				"workers":                             "8",
			},
		},
	}

	return mutate.MutateConfigMapNamespacePostgresOperator(resourceObj, parent, collection, reconciler, req)
}
