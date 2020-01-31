/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.0.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vapi

type DeploymentTemplateSpec struct {
	Artifact *Artifact `json:"artifact,omitempty"`
	Parallelism int32 `json:"parallelism,omitempty"`
	NumberOfTaskManagers int32 `json:"numberOfTaskManagers,omitempty"`
	Resources map[string]ResourceSpec `json:"resources,omitempty"`
	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	Kubernetes *KubernetesOptions `json:"kubernetes,omitempty"`
}
