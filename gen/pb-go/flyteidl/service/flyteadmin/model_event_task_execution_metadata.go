/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Holds metadata around how a task was executed. TODO(katrogan): Extend to include freeform fields (https://github.com/flyteorg/flyte/issues/325).
type EventTaskExecutionMetadata struct {
	// Unique, generated name for this task execution used by the backend.
	GeneratedName string `json:"generated_name,omitempty"`
	// Additional data on external resources on other back-ends or platforms (e.g. Hive, Qubole, etc) launched by this task execution.
	ExternalResources []EventExternalResourceInfo `json:"external_resources,omitempty"`
	// Includes additional data on concurrent resource management used during execution.. This is a repeated field because a plugin can request multiple resource allocations during execution.
	ResourcePoolInfo []EventResourcePoolInfo `json:"resource_pool_info,omitempty"`
	// The identifier of the plugin used to execute this task.
	PluginIdentifier string `json:"plugin_identifier,omitempty"`
	InstanceClass *TaskExecutionMetadataInstanceClass `json:"instance_class,omitempty"`
}
