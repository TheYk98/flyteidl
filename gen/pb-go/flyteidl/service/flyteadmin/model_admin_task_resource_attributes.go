/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminTaskResourceAttributes struct {
	Defaults *AdminTaskResourceSpec `json:"defaults,omitempty"`
	Limits *AdminTaskResourceSpec `json:"limits,omitempty"`
}
