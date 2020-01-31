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

type Savepoint struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Kind string `json:"kind,omitempty"`
	Metadata *SavepointMetadata `json:"metadata,omitempty"`
	Spec *SavepointSpec `json:"spec,omitempty"`
	Status *SavepointStatus `json:"status,omitempty"`
}
