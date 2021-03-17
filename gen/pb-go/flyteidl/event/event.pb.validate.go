// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = core.WorkflowExecution_Phase(0)

	_ = core.NodeExecution_Phase(0)

	_ = core.CatalogCacheStatus(0)

	_ = core.TaskExecution_Phase(0)
)

// define the regex for a UUID once up-front
var _event_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on WorkflowExecutionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowExecutionEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionEventValidationError{
				field:  "ExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProducerId

	// no validation rules for Phase

	if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionEventValidationError{
				field:  "OccurredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.OutputResult.(type) {

	case *WorkflowExecutionEvent_OutputUri:
		// no validation rules for OutputUri

	case *WorkflowExecutionEvent_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowExecutionEventValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WorkflowExecutionEventValidationError is the validation error returned by
// WorkflowExecutionEvent.Validate if the designated constraints aren't met.
type WorkflowExecutionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowExecutionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowExecutionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowExecutionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowExecutionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowExecutionEventValidationError) ErrorName() string {
	return "WorkflowExecutionEventValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowExecutionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowExecutionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowExecutionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowExecutionEventValidationError{}

// Validate checks the field values on NodeExecutionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NodeExecutionEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionEventValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProducerId

	// no validation rules for Phase

	if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionEventValidationError{
				field:  "OccurredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InputUri

	if v, ok := interface{}(m.GetParentTaskMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionEventValidationError{
				field:  "ParentTaskMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetParentNodeMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeExecutionEventValidationError{
				field:  "ParentNodeMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RetryGroup

	// no validation rules for SpecNodeId

	// no validation rules for NodeName

	switch m.OutputResult.(type) {

	case *NodeExecutionEvent_OutputUri:
		// no validation rules for OutputUri

	case *NodeExecutionEvent_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionEventValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.TargetMetadata.(type) {

	case *NodeExecutionEvent_WorkflowNodeMetadata:

		if v, ok := interface{}(m.GetWorkflowNodeMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionEventValidationError{
					field:  "WorkflowNodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *NodeExecutionEvent_TaskNodeMetadata:

		if v, ok := interface{}(m.GetTaskNodeMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeExecutionEventValidationError{
					field:  "TaskNodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeExecutionEventValidationError is the validation error returned by
// NodeExecutionEvent.Validate if the designated constraints aren't met.
type NodeExecutionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeExecutionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeExecutionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeExecutionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeExecutionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeExecutionEventValidationError) ErrorName() string {
	return "NodeExecutionEventValidationError"
}

// Error satisfies the builtin error interface
func (e NodeExecutionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeExecutionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeExecutionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeExecutionEventValidationError{}

// Validate checks the field values on WorkflowNodeMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowNodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowNodeMetadataValidationError{
				field:  "ExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowNodeMetadataValidationError is the validation error returned by
// WorkflowNodeMetadata.Validate if the designated constraints aren't met.
type WorkflowNodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowNodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowNodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowNodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowNodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowNodeMetadataValidationError) ErrorName() string {
	return "WorkflowNodeMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowNodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowNodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowNodeMetadataValidationError{}

// Validate checks the field values on TaskNodeMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TaskNodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CacheStatus

	if v, ok := interface{}(m.GetCatalogKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskNodeMetadataValidationError{
				field:  "CatalogKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TaskNodeMetadataValidationError is the validation error returned by
// TaskNodeMetadata.Validate if the designated constraints aren't met.
type TaskNodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskNodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskNodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskNodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskNodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskNodeMetadataValidationError) ErrorName() string { return "TaskNodeMetadataValidationError" }

// Error satisfies the builtin error interface
func (e TaskNodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskNodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskNodeMetadataValidationError{}

// Validate checks the field values on ParentTaskExecutionMetadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ParentTaskExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ParentTaskExecutionMetadataValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ParentTaskExecutionMetadataValidationError is the validation error returned
// by ParentTaskExecutionMetadata.Validate if the designated constraints
// aren't met.
type ParentTaskExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParentTaskExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParentTaskExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParentTaskExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParentTaskExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParentTaskExecutionMetadataValidationError) ErrorName() string {
	return "ParentTaskExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e ParentTaskExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParentTaskExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParentTaskExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParentTaskExecutionMetadataValidationError{}

// Validate checks the field values on ParentNodeExecutionMetadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ParentNodeExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NodeId

	return nil
}

// ParentNodeExecutionMetadataValidationError is the validation error returned
// by ParentNodeExecutionMetadata.Validate if the designated constraints
// aren't met.
type ParentNodeExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParentNodeExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParentNodeExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParentNodeExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParentNodeExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParentNodeExecutionMetadataValidationError) ErrorName() string {
	return "ParentNodeExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e ParentNodeExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParentNodeExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParentNodeExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParentNodeExecutionMetadataValidationError{}

// Validate checks the field values on TaskExecutionEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTaskId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionEventValidationError{
				field:  "TaskId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetParentNodeExecutionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionEventValidationError{
				field:  "ParentNodeExecutionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RetryAttempt

	// no validation rules for Phase

	// no validation rules for ProducerId

	for idx, item := range m.GetLogs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionEventValidationError{
					field:  fmt.Sprintf("Logs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionEventValidationError{
				field:  "OccurredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InputUri

	if v, ok := interface{}(m.GetCustomInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionEventValidationError{
				field:  "CustomInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PhaseVersion

	// no validation rules for Reason

	// no validation rules for TaskType

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskExecutionEventValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.OutputResult.(type) {

	case *TaskExecutionEvent_OutputUri:
		// no validation rules for OutputUri

	case *TaskExecutionEvent_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionEventValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TaskExecutionEventValidationError is the validation error returned by
// TaskExecutionEvent.Validate if the designated constraints aren't met.
type TaskExecutionEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionEventValidationError) ErrorName() string {
	return "TaskExecutionEventValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionEventValidationError{}

// Validate checks the field values on ExternalResourceInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExternalResourceInfo) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ExternalResourceInfoValidationError is the validation error returned by
// ExternalResourceInfo.Validate if the designated constraints aren't met.
type ExternalResourceInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExternalResourceInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExternalResourceInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExternalResourceInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExternalResourceInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExternalResourceInfoValidationError) ErrorName() string {
	return "ExternalResourceInfoValidationError"
}

// Error satisfies the builtin error interface
func (e ExternalResourceInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExternalResourceInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExternalResourceInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExternalResourceInfoValidationError{}

// Validate checks the field values on ResourcePoolInfo with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ResourcePoolInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AllocationToken

	// no validation rules for Namespace

	return nil
}

// ResourcePoolInfoValidationError is the validation error returned by
// ResourcePoolInfo.Validate if the designated constraints aren't met.
type ResourcePoolInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourcePoolInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourcePoolInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourcePoolInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourcePoolInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourcePoolInfoValidationError) ErrorName() string { return "ResourcePoolInfoValidationError" }

// Error satisfies the builtin error interface
func (e ResourcePoolInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourcePoolInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourcePoolInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourcePoolInfoValidationError{}

// Validate checks the field values on TaskExecutionMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TaskExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GeneratedName

	for idx, item := range m.GetExternalResources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionMetadataValidationError{
					field:  fmt.Sprintf("ExternalResources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResourcePoolInfo() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskExecutionMetadataValidationError{
					field:  fmt.Sprintf("ResourcePoolInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PluginIdentifier

	// no validation rules for InstanceClass

	return nil
}

// TaskExecutionMetadataValidationError is the validation error returned by
// TaskExecutionMetadata.Validate if the designated constraints aren't met.
type TaskExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskExecutionMetadataValidationError) ErrorName() string {
	return "TaskExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e TaskExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskExecutionMetadataValidationError{}
