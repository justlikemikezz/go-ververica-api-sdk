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

type Failure struct {
	Message  string `json:"message,omitempty"`
	Reason   string `json:"reason,omitempty"`
	FailedAt string `json:"failedAt,omitempty"`
}
