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

type JobDeploymentTarget struct {
	Kubernetes *KubernetesTarget `json:"kubernetes,omitempty"`
}
