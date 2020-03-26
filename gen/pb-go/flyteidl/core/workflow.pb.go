// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/workflow.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines a condition and the execution unit that should be executed if the condition is satisfied.
type IfBlock struct {
	Condition            *BooleanExpression `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	ThenNode             *Node              `protobuf:"bytes,2,opt,name=then_node,json=thenNode,proto3" json:"then_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IfBlock) Reset()         { *m = IfBlock{} }
func (m *IfBlock) String() string { return proto.CompactTextString(m) }
func (*IfBlock) ProtoMessage()    {}
func (*IfBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{0}
}

func (m *IfBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfBlock.Unmarshal(m, b)
}
func (m *IfBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfBlock.Marshal(b, m, deterministic)
}
func (m *IfBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfBlock.Merge(m, src)
}
func (m *IfBlock) XXX_Size() int {
	return xxx_messageInfo_IfBlock.Size(m)
}
func (m *IfBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfBlock proto.InternalMessageInfo

func (m *IfBlock) GetCondition() *BooleanExpression {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *IfBlock) GetThenNode() *Node {
	if m != nil {
		return m.ThenNode
	}
	return nil
}

// Defines a series of if/else blocks. The first branch whose condition evaluates to true is the one to execute.
// If no conditions were satisfied, the else_node or the error will execute.
type IfElseBlock struct {
	//+required. First condition to evaluate.
	Case *IfBlock `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	//+optional. Additional branches to evaluate.
	Other []*IfBlock `protobuf:"bytes,2,rep,name=other,proto3" json:"other,omitempty"`
	//+required.
	//
	// Types that are valid to be assigned to Default:
	//	*IfElseBlock_ElseNode
	//	*IfElseBlock_Error
	Default              isIfElseBlock_Default `protobuf_oneof:"default"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IfElseBlock) Reset()         { *m = IfElseBlock{} }
func (m *IfElseBlock) String() string { return proto.CompactTextString(m) }
func (*IfElseBlock) ProtoMessage()    {}
func (*IfElseBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{1}
}

func (m *IfElseBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfElseBlock.Unmarshal(m, b)
}
func (m *IfElseBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfElseBlock.Marshal(b, m, deterministic)
}
func (m *IfElseBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfElseBlock.Merge(m, src)
}
func (m *IfElseBlock) XXX_Size() int {
	return xxx_messageInfo_IfElseBlock.Size(m)
}
func (m *IfElseBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfElseBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfElseBlock proto.InternalMessageInfo

func (m *IfElseBlock) GetCase() *IfBlock {
	if m != nil {
		return m.Case
	}
	return nil
}

func (m *IfElseBlock) GetOther() []*IfBlock {
	if m != nil {
		return m.Other
	}
	return nil
}

type isIfElseBlock_Default interface {
	isIfElseBlock_Default()
}

type IfElseBlock_ElseNode struct {
	ElseNode *Node `protobuf:"bytes,3,opt,name=else_node,json=elseNode,proto3,oneof"`
}

type IfElseBlock_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*IfElseBlock_ElseNode) isIfElseBlock_Default() {}

func (*IfElseBlock_Error) isIfElseBlock_Default() {}

func (m *IfElseBlock) GetDefault() isIfElseBlock_Default {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *IfElseBlock) GetElseNode() *Node {
	if x, ok := m.GetDefault().(*IfElseBlock_ElseNode); ok {
		return x.ElseNode
	}
	return nil
}

func (m *IfElseBlock) GetError() *Error {
	if x, ok := m.GetDefault().(*IfElseBlock_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IfElseBlock) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IfElseBlock_ElseNode)(nil),
		(*IfElseBlock_Error)(nil),
	}
}

// BranchNode is a special node that alter the flow of the workflow graph. It allows the control flow to branch at
// runtime based on a series of conditions that get evaluated on various parameters (e.g. inputs, primtives).
type BranchNode struct {
	//+required
	IfElse               *IfElseBlock `protobuf:"bytes,1,opt,name=if_else,json=ifElse,proto3" json:"if_else,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BranchNode) Reset()         { *m = BranchNode{} }
func (m *BranchNode) String() string { return proto.CompactTextString(m) }
func (*BranchNode) ProtoMessage()    {}
func (*BranchNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{2}
}

func (m *BranchNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchNode.Unmarshal(m, b)
}
func (m *BranchNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchNode.Marshal(b, m, deterministic)
}
func (m *BranchNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchNode.Merge(m, src)
}
func (m *BranchNode) XXX_Size() int {
	return xxx_messageInfo_BranchNode.Size(m)
}
func (m *BranchNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchNode.DiscardUnknown(m)
}

var xxx_messageInfo_BranchNode proto.InternalMessageInfo

func (m *BranchNode) GetIfElse() *IfElseBlock {
	if m != nil {
		return m.IfElse
	}
	return nil
}

// Refers to the task that the Node is to execute.
type TaskNode struct {
	// Types that are valid to be assigned to Reference:
	//	*TaskNode_ReferenceId
	Reference            isTaskNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskNode) Reset()         { *m = TaskNode{} }
func (m *TaskNode) String() string { return proto.CompactTextString(m) }
func (*TaskNode) ProtoMessage()    {}
func (*TaskNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{3}
}

func (m *TaskNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNode.Unmarshal(m, b)
}
func (m *TaskNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNode.Marshal(b, m, deterministic)
}
func (m *TaskNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNode.Merge(m, src)
}
func (m *TaskNode) XXX_Size() int {
	return xxx_messageInfo_TaskNode.Size(m)
}
func (m *TaskNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNode.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNode proto.InternalMessageInfo

type isTaskNode_Reference interface {
	isTaskNode_Reference()
}

type TaskNode_ReferenceId struct {
	ReferenceId *Identifier `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3,oneof"`
}

func (*TaskNode_ReferenceId) isTaskNode_Reference() {}

func (m *TaskNode) GetReference() isTaskNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *TaskNode) GetReferenceId() *Identifier {
	if x, ok := m.GetReference().(*TaskNode_ReferenceId); ok {
		return x.ReferenceId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskNode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskNode_ReferenceId)(nil),
	}
}

// Refers to a the workflow the node is to execute.
type WorkflowNode struct {
	// Types that are valid to be assigned to Reference:
	//	*WorkflowNode_LaunchplanRef
	//	*WorkflowNode_SubWorkflowRef
	Reference            isWorkflowNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkflowNode) Reset()         { *m = WorkflowNode{} }
func (m *WorkflowNode) String() string { return proto.CompactTextString(m) }
func (*WorkflowNode) ProtoMessage()    {}
func (*WorkflowNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{4}
}

func (m *WorkflowNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNode.Unmarshal(m, b)
}
func (m *WorkflowNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNode.Marshal(b, m, deterministic)
}
func (m *WorkflowNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNode.Merge(m, src)
}
func (m *WorkflowNode) XXX_Size() int {
	return xxx_messageInfo_WorkflowNode.Size(m)
}
func (m *WorkflowNode) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNode.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNode proto.InternalMessageInfo

type isWorkflowNode_Reference interface {
	isWorkflowNode_Reference()
}

type WorkflowNode_LaunchplanRef struct {
	LaunchplanRef *Identifier `protobuf:"bytes,1,opt,name=launchplan_ref,json=launchplanRef,proto3,oneof"`
}

type WorkflowNode_SubWorkflowRef struct {
	SubWorkflowRef *Identifier `protobuf:"bytes,2,opt,name=sub_workflow_ref,json=subWorkflowRef,proto3,oneof"`
}

func (*WorkflowNode_LaunchplanRef) isWorkflowNode_Reference() {}

func (*WorkflowNode_SubWorkflowRef) isWorkflowNode_Reference() {}

func (m *WorkflowNode) GetReference() isWorkflowNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *WorkflowNode) GetLaunchplanRef() *Identifier {
	if x, ok := m.GetReference().(*WorkflowNode_LaunchplanRef); ok {
		return x.LaunchplanRef
	}
	return nil
}

func (m *WorkflowNode) GetSubWorkflowRef() *Identifier {
	if x, ok := m.GetReference().(*WorkflowNode_SubWorkflowRef); ok {
		return x.SubWorkflowRef
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkflowNode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkflowNode_LaunchplanRef)(nil),
		(*WorkflowNode_SubWorkflowRef)(nil),
	}
}

// Defines extra information about the Node.
type NodeMetadata struct {
	// A friendly name for the Node
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The overall timeout of a task.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	// Identify whether node is interruptible
	//
	// Types that are valid to be assigned to InterruptibleValue:
	//	*NodeMetadata_Interruptible
	InterruptibleValue isNodeMetadata_InterruptibleValue `protobuf_oneof:"interruptible_value"`
	// Total wait time a node can be delayed by queueing.
	// This value cannot be smaller than value set at the workflow level.
	MaxWaitTime          *duration.Duration `protobuf:"bytes,7,opt,name=max_wait_time,json=maxWaitTime,proto3" json:"max_wait_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{5}
}

func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (m *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(m, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *NodeMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

type isNodeMetadata_InterruptibleValue interface {
	isNodeMetadata_InterruptibleValue()
}

type NodeMetadata_Interruptible struct {
	Interruptible bool `protobuf:"varint,6,opt,name=interruptible,proto3,oneof"`
}

func (*NodeMetadata_Interruptible) isNodeMetadata_InterruptibleValue() {}

func (m *NodeMetadata) GetInterruptibleValue() isNodeMetadata_InterruptibleValue {
	if m != nil {
		return m.InterruptibleValue
	}
	return nil
}

func (m *NodeMetadata) GetInterruptible() bool {
	if x, ok := m.GetInterruptibleValue().(*NodeMetadata_Interruptible); ok {
		return x.Interruptible
	}
	return false
}

func (m *NodeMetadata) GetMaxWaitTime() *duration.Duration {
	if m != nil {
		return m.MaxWaitTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeMetadata_Interruptible)(nil),
	}
}

// Links a variable to an alias.
type Alias struct {
	// Must match one of the output variable names on a node.
	Var string `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	// A workflow-level unique alias that downstream nodes can refer to in their input.
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{6}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetVar() string {
	if m != nil {
		return m.Var
	}
	return ""
}

func (m *Alias) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

// A Workflow graph Node. One unit of execution in the graph. Each node can be linked to a Task, a Workflow or a branch
// node.
type Node struct {
	// A workflow-level unique identifier that identifies this node in the workflow. "inputs" and "outputs" are reserved
	// node ids that cannot be used by other nodes.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the node.
	Metadata *NodeMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specifies how to bind the underlying interface's inputs. All required inputs specified in the underlying interface
	// must be fullfilled.
	Inputs []*Binding `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	//+optional Specifies execution depdendency for this node ensuring it will only get scheduled to run after all its
	// upstream nodes have completed. This node will have an implicit depdendency on any node that appears in inputs
	// field.
	UpstreamNodeIds []string `protobuf:"bytes,4,rep,name=upstream_node_ids,json=upstreamNodeIds,proto3" json:"upstream_node_ids,omitempty"`
	//+optional. A node can define aliases for a subset of its outputs. This is particularly useful if different nodes
	// need to conform to the same interface (e.g. all branches in a branch node). Downstream nodes must refer to this
	// nodes outputs using the alias if one's specified.
	OutputAliases []*Alias `protobuf:"bytes,5,rep,name=output_aliases,json=outputAliases,proto3" json:"output_aliases,omitempty"`
	// Information about the target to execute in this node.
	//
	// Types that are valid to be assigned to Target:
	//	*Node_TaskNode
	//	*Node_WorkflowNode
	//	*Node_BranchNode
	Target               isNode_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{7}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetInputs() []*Binding {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Node) GetUpstreamNodeIds() []string {
	if m != nil {
		return m.UpstreamNodeIds
	}
	return nil
}

func (m *Node) GetOutputAliases() []*Alias {
	if m != nil {
		return m.OutputAliases
	}
	return nil
}

type isNode_Target interface {
	isNode_Target()
}

type Node_TaskNode struct {
	TaskNode *TaskNode `protobuf:"bytes,6,opt,name=task_node,json=taskNode,proto3,oneof"`
}

type Node_WorkflowNode struct {
	WorkflowNode *WorkflowNode `protobuf:"bytes,7,opt,name=workflow_node,json=workflowNode,proto3,oneof"`
}

type Node_BranchNode struct {
	BranchNode *BranchNode `protobuf:"bytes,8,opt,name=branch_node,json=branchNode,proto3,oneof"`
}

func (*Node_TaskNode) isNode_Target() {}

func (*Node_WorkflowNode) isNode_Target() {}

func (*Node_BranchNode) isNode_Target() {}

func (m *Node) GetTarget() isNode_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Node) GetTaskNode() *TaskNode {
	if x, ok := m.GetTarget().(*Node_TaskNode); ok {
		return x.TaskNode
	}
	return nil
}

func (m *Node) GetWorkflowNode() *WorkflowNode {
	if x, ok := m.GetTarget().(*Node_WorkflowNode); ok {
		return x.WorkflowNode
	}
	return nil
}

func (m *Node) GetBranchNode() *BranchNode {
	if x, ok := m.GetTarget().(*Node_BranchNode); ok {
		return x.BranchNode
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Node) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Node_TaskNode)(nil),
		(*Node_WorkflowNode)(nil),
		(*Node_BranchNode)(nil),
	}
}

// Metadata for the entire workflow.
// To be used in the future.
type WorkflowMetadata struct {
	// Total wait time a workflow can be delayed by queueing.
	// max_wait_time set at the workflow level will take precedence over max_wait_time
	// set at the node level in the case that the value set at the node level
	// is larger than the workflow level.
	MaxWaitTime          *duration.Duration `protobuf:"bytes,1,opt,name=max_wait_time,json=maxWaitTime,proto3" json:"max_wait_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WorkflowMetadata) Reset()         { *m = WorkflowMetadata{} }
func (m *WorkflowMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowMetadata) ProtoMessage()    {}
func (*WorkflowMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{8}
}

func (m *WorkflowMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMetadata.Unmarshal(m, b)
}
func (m *WorkflowMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMetadata.Merge(m, src)
}
func (m *WorkflowMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowMetadata.Size(m)
}
func (m *WorkflowMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMetadata proto.InternalMessageInfo

func (m *WorkflowMetadata) GetMaxWaitTime() *duration.Duration {
	if m != nil {
		return m.MaxWaitTime
	}
	return nil
}

// Default Workflow Metadata for the entire workflow.
type WorkflowMetadataDefaults struct {
	// Identify whether workflow is interruptible.
	// The value set at the workflow level will be the defualt value used for nodes
	// unless explicitly set at the node level.
	Interruptible        bool     `protobuf:"varint,1,opt,name=interruptible,proto3" json:"interruptible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowMetadataDefaults) Reset()         { *m = WorkflowMetadataDefaults{} }
func (m *WorkflowMetadataDefaults) String() string { return proto.CompactTextString(m) }
func (*WorkflowMetadataDefaults) ProtoMessage()    {}
func (*WorkflowMetadataDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{9}
}

func (m *WorkflowMetadataDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMetadataDefaults.Unmarshal(m, b)
}
func (m *WorkflowMetadataDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMetadataDefaults.Marshal(b, m, deterministic)
}
func (m *WorkflowMetadataDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMetadataDefaults.Merge(m, src)
}
func (m *WorkflowMetadataDefaults) XXX_Size() int {
	return xxx_messageInfo_WorkflowMetadataDefaults.Size(m)
}
func (m *WorkflowMetadataDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMetadataDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMetadataDefaults proto.InternalMessageInfo

func (m *WorkflowMetadataDefaults) GetInterruptible() bool {
	if m != nil {
		return m.Interruptible
	}
	return false
}

// Flyte Workflow Structure that encapsulates task, branch and subworkflow nodes to form a statically analyzable,
// directed acyclic graph.
type WorkflowTemplate struct {
	// A globally unique identifier for the workflow.
	Id *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the workflow.
	Metadata *WorkflowMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Defines a strongly typed interface for the Workflow. This can include some optional parameters.
	Interface *TypedInterface `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	// A list of nodes. In addition, "globals" is a special reserved node id that can be used to consume workflow inputs.
	Nodes []*Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// A list of output bindings that specify how to construct workflow outputs. Bindings can pull node outputs or
	// specify literals. All workflow outputs specified in the interface field must be bound in order for the workflow
	// to be validated. A workflow has an implicit dependency on all of its nodes to execute successfully in order to
	// bind final outputs.
	// Most of these outputs will be Binding's with a BindingData of type OutputReference.  That is, your workflow can
	// just have an output of some constant (`Output(5)`), but usually, the workflow will be pulling
	// outputs from the output of a task.
	Outputs []*Binding `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	//+optional A catch-all node. This node is executed whenever the execution engine determines the workflow has failed.
	// The interface of this node must match the Workflow interface with an additional input named "error" of type
	// pb.lyft.flyte.core.Error.
	FailureNode *Node `protobuf:"bytes,6,opt,name=failure_node,json=failureNode,proto3" json:"failure_node,omitempty"`
	// workflow defaults
	MetadataDefaults     *WorkflowMetadataDefaults `protobuf:"bytes,7,opt,name=metadata_defaults,json=metadataDefaults,proto3" json:"metadata_defaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{10}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WorkflowTemplate) GetMetadata() *WorkflowMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *WorkflowTemplate) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *WorkflowTemplate) GetOutputs() []*Binding {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *WorkflowTemplate) GetFailureNode() *Node {
	if m != nil {
		return m.FailureNode
	}
	return nil
}

func (m *WorkflowTemplate) GetMetadataDefaults() *WorkflowMetadataDefaults {
	if m != nil {
		return m.MetadataDefaults
	}
	return nil
}

func init() {
	proto.RegisterType((*IfBlock)(nil), "flyteidl.core.IfBlock")
	proto.RegisterType((*IfElseBlock)(nil), "flyteidl.core.IfElseBlock")
	proto.RegisterType((*BranchNode)(nil), "flyteidl.core.BranchNode")
	proto.RegisterType((*TaskNode)(nil), "flyteidl.core.TaskNode")
	proto.RegisterType((*WorkflowNode)(nil), "flyteidl.core.WorkflowNode")
	proto.RegisterType((*NodeMetadata)(nil), "flyteidl.core.NodeMetadata")
	proto.RegisterType((*Alias)(nil), "flyteidl.core.Alias")
	proto.RegisterType((*Node)(nil), "flyteidl.core.Node")
	proto.RegisterType((*WorkflowMetadata)(nil), "flyteidl.core.WorkflowMetadata")
	proto.RegisterType((*WorkflowMetadataDefaults)(nil), "flyteidl.core.WorkflowMetadataDefaults")
	proto.RegisterType((*WorkflowTemplate)(nil), "flyteidl.core.WorkflowTemplate")
}

func init() { proto.RegisterFile("flyteidl/core/workflow.proto", fileDescriptor_fccede37486c456e) }

var fileDescriptor_fccede37486c456e = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xce, 0x7f, 0xe2, 0x93, 0xa4, 0x74, 0x67, 0x17, 0x70, 0xcb, 0xee, 0x12, 0x59, 0x08, 0xba,
	0x2b, 0x48, 0xaa, 0xae, 0x54, 0x2e, 0x0a, 0x2b, 0x6a, 0x6d, 0xa5, 0xe4, 0x02, 0x24, 0x86, 0x4a,
	0x95, 0xb8, 0xb1, 0x26, 0xf1, 0x71, 0x32, 0xaa, 0xed, 0xb1, 0xc6, 0xe3, 0x6d, 0x23, 0x9e, 0x81,
	0x97, 0xe0, 0x82, 0x2b, 0x1e, 0x87, 0x07, 0x42, 0x1e, 0xff, 0xa4, 0x71, 0xd3, 0x2d, 0x7b, 0x37,
	0xe3, 0xf3, 0x9d, 0x33, 0xdf, 0x9c, 0xf3, 0x7d, 0xb6, 0xe1, 0xb9, 0xe7, 0xaf, 0x15, 0x72, 0xd7,
	0x9f, 0x2c, 0x84, 0xc4, 0xc9, 0x8d, 0x90, 0xd7, 0x9e, 0x2f, 0x6e, 0xc6, 0x91, 0x14, 0x4a, 0x90,
	0x61, 0x11, 0x1d, 0xa7, 0xd1, 0xc3, 0x17, 0xdb, 0xe0, 0x85, 0x08, 0x5d, 0xae, 0xb8, 0x08, 0x33,
	0xf4, 0xe1, 0xcb, 0xed, 0x30, 0x77, 0x31, 0x54, 0xdc, 0xe3, 0x28, 0xf3, 0x78, 0x25, 0x9d, 0x87,
	0x0a, 0xa5, 0xc7, 0x16, 0x98, 0x87, 0x2b, 0x54, 0x7c, 0xae, 0x50, 0x32, 0x3f, 0xce, 0xa3, 0x07,
	0xdb, 0x51, 0xb5, 0x8e, 0xb0, 0x08, 0xbd, 0x5c, 0x0a, 0xb1, 0xf4, 0x71, 0xa2, 0x77, 0xf3, 0xc4,
	0x9b, 0xb8, 0x89, 0x64, 0x1b, 0x5e, 0xd6, 0x1f, 0xd0, 0x9d, 0x79, 0xb6, 0x2f, 0x16, 0xd7, 0xe4,
	0x2d, 0x18, 0x25, 0x6b, 0xb3, 0x3e, 0xaa, 0x1f, 0xf5, 0x4f, 0x46, 0xe3, 0xad, 0x4b, 0x8e, 0x6d,
	0x21, 0x7c, 0x64, 0xe1, 0xc5, 0x6d, 0x24, 0x31, 0x8e, 0xb9, 0x08, 0xe9, 0x26, 0x85, 0x1c, 0x83,
	0xa1, 0x56, 0x18, 0x3a, 0xa1, 0x70, 0xd1, 0x6c, 0xe8, 0xfc, 0xa7, 0x95, 0xfc, 0x5f, 0x84, 0x8b,
	0xb4, 0x97, 0xa2, 0xd2, 0x95, 0xf5, 0x6f, 0x1d, 0xfa, 0x33, 0xef, 0xc2, 0x8f, 0x31, 0x63, 0xf0,
	0x1a, 0x5a, 0x0b, 0x16, 0x63, 0x7e, 0xf8, 0x67, 0x95, 0xe4, 0x9c, 0x27, 0xd5, 0x18, 0xf2, 0x2d,
	0xb4, 0x85, 0x5a, 0xa1, 0x34, 0x1b, 0xa3, 0xe6, 0x07, 0xc0, 0x19, 0x88, 0x9c, 0x80, 0x81, 0x7e,
	0x8c, 0x19, 0xb7, 0xe6, 0x83, 0xdc, 0xa6, 0x35, 0xda, 0x4b, 0x71, 0xe9, 0x3a, 0x3d, 0x01, 0xa5,
	0x14, 0xd2, 0x6c, 0x69, 0xfc, 0xb3, 0x0a, 0xfe, 0x22, 0x8d, 0x4d, 0x6b, 0x34, 0x03, 0xd9, 0x06,
	0x74, 0x5d, 0xf4, 0x58, 0xe2, 0x2b, 0xeb, 0x1c, 0xc0, 0x96, 0x2c, 0x5c, 0xac, 0x74, 0x99, 0x37,
	0xd0, 0xe5, 0x9e, 0x93, 0x56, 0xcd, 0xef, 0x75, 0x78, 0x8f, 0x6a, 0xd9, 0x01, 0xda, 0xe1, 0x7a,
	0x63, 0x5d, 0x41, 0xef, 0x92, 0xc5, 0xd7, 0xba, 0xc0, 0x5b, 0x18, 0x48, 0xf4, 0x50, 0x62, 0xb8,
	0x40, 0x87, 0xbb, 0x79, 0x95, 0x83, 0x6a, 0x95, 0x52, 0x51, 0xd3, 0x1a, 0xed, 0x97, 0x09, 0x33,
	0xd7, 0xee, 0x83, 0x51, 0x6e, 0xad, 0xbf, 0xeb, 0x30, 0xb8, 0xca, 0x85, 0xac, 0xab, 0xdb, 0xb0,
	0xe7, 0xb3, 0x24, 0x5c, 0xac, 0x22, 0x9f, 0x85, 0x8e, 0x44, 0xef, 0xff, 0xd4, 0x1f, 0x6e, 0x52,
	0x28, 0x7a, 0xe4, 0x02, 0xf6, 0xe3, 0x64, 0xee, 0x14, 0x06, 0xd1, 0x55, 0x1a, 0x8f, 0x57, 0xd9,
	0x8b, 0x93, 0x79, 0xc1, 0x85, 0xa2, 0xb7, 0x4d, 0xf4, 0xcf, 0x06, 0x0c, 0x52, 0x82, 0x3f, 0xa3,
	0x62, 0x2e, 0x53, 0x8c, 0x10, 0x68, 0x85, 0x2c, 0xc8, 0x9a, 0x68, 0x50, 0xbd, 0x4e, 0x7b, 0xab,
	0x78, 0x80, 0x22, 0x51, 0xf9, 0x90, 0x0e, 0xc6, 0x99, 0xde, 0xc7, 0x85, 0xde, 0xc7, 0xef, 0x72,
	0xbd, 0xd3, 0x02, 0x49, 0x4e, 0xa1, 0x2b, 0x51, 0x49, 0x8e, 0xb1, 0xd9, 0xd6, 0x49, 0xcf, 0x2b,
	0x24, 0x29, 0x2a, 0xb9, 0xfe, 0x4d, 0x49, 0xa6, 0x70, 0xb9, 0xa6, 0x05, 0x98, 0x7c, 0x0d, 0x43,
	0x6d, 0x4b, 0x99, 0x44, 0x8a, 0xcf, 0x7d, 0x34, 0x3b, 0xa3, 0xfa, 0x51, 0x2f, 0xed, 0xc6, 0xd6,
	0x63, 0xf2, 0x23, 0x0c, 0x03, 0x76, 0xeb, 0xdc, 0x30, 0xae, 0x9c, 0xf4, 0x4c, 0xb3, 0xfb, 0x18,
	0xb5, 0x7e, 0xc0, 0x6e, 0xaf, 0x18, 0x57, 0x97, 0x3c, 0x40, 0xfb, 0x53, 0x78, 0xba, 0x55, 0xcf,
	0x79, 0xcf, 0xfc, 0x04, 0xad, 0x09, 0xb4, 0xcf, 0x7d, 0xce, 0x62, 0xb2, 0x0f, 0xcd, 0xf7, 0x4c,
	0xe6, 0x6d, 0x48, 0x97, 0xe4, 0x19, 0xb4, 0x59, 0x1a, 0xd2, 0x3d, 0x37, 0x68, 0xb6, 0xb1, 0xfe,
	0x69, 0x42, 0x4b, 0x4f, 0x78, 0x0f, 0x1a, 0xb9, 0x6a, 0x0c, 0xda, 0xe0, 0x2e, 0xf9, 0x1e, 0x7a,
	0x41, 0xde, 0xd4, 0x7c, 0x4a, 0x5f, 0xec, 0xb0, 0x42, 0xd1, 0x77, 0x5a, 0x82, 0xc9, 0x18, 0x3a,
	0x3c, 0x8c, 0x12, 0x15, 0x9b, 0xcd, 0x9d, 0x9e, 0xb3, 0x79, 0xe8, 0xf2, 0x70, 0x49, 0x73, 0x14,
	0x79, 0x0d, 0x4f, 0x92, 0x28, 0x56, 0x12, 0x59, 0xa0, 0x8d, 0xe7, 0x70, 0x37, 0x36, 0x5b, 0xa3,
	0xe6, 0x91, 0x41, 0x3f, 0x29, 0x02, 0xe9, 0x51, 0x33, 0x37, 0x26, 0x67, 0xb0, 0x27, 0x12, 0x15,
	0x25, 0xca, 0xd1, 0xec, 0xf5, 0x6c, 0x9a, 0x3b, 0x5c, 0xa7, 0x7b, 0x40, 0x87, 0x19, 0xf6, 0x3c,
	0x83, 0x92, 0x53, 0x30, 0x14, 0x8b, 0xaf, 0x33, 0x77, 0x77, 0xf4, 0x95, 0x3e, 0xaf, 0xe4, 0x15,
	0x6e, 0x4a, 0x1d, 0xae, 0x0a, 0x67, 0xd9, 0x30, 0x2c, 0x35, 0xab, 0x73, 0xbb, 0x3b, 0xdb, 0x71,
	0xd7, 0x2f, 0xd3, 0x1a, 0x1d, 0xdc, 0xdc, 0xf5, 0xcf, 0x0f, 0xd0, 0x9f, 0x6b, 0xb3, 0x67, 0x15,
	0x7a, 0x3b, 0x65, 0xbf, 0x79, 0x1d, 0x4c, 0x6b, 0x14, 0xe6, 0xe5, 0xce, 0xee, 0x41, 0x47, 0x31,
	0xb9, 0x44, 0x65, 0xfd, 0x0a, 0xfb, 0xc5, 0x39, 0xa5, 0xe4, 0xef, 0x29, 0xa9, 0xfe, 0x31, 0x4a,
	0xb2, 0x7e, 0x02, 0xb3, 0x5a, 0xf2, 0x5d, 0xf6, 0x8a, 0x8a, 0xc9, 0x57, 0x55, 0x31, 0xa7, 0xa5,
	0x7b, 0x15, 0x29, 0x5b, 0x7f, 0x35, 0x37, 0xac, 0x2e, 0x31, 0x88, 0x7c, 0xa6, 0x90, 0xbc, 0x2a,
	0xf5, 0xf4, 0x21, 0x7f, 0x6b, 0xa9, 0x9d, 0xdd, 0x93, 0xda, 0x97, 0x0f, 0xf4, 0x76, 0x87, 0xdc,
	0xce, 0xc0, 0x28, 0x3f, 0x83, 0xf9, 0x3b, 0xfb, 0x45, 0x75, 0xaa, 0xeb, 0x08, 0xdd, 0x59, 0x01,
	0xa2, 0x1b, 0x3c, 0x79, 0x05, 0xed, 0x74, 0x1e, 0x99, 0xde, 0x1e, 0xf8, 0x10, 0x65, 0x08, 0x72,
	0x0c, 0xdd, 0x4c, 0x4e, 0x85, 0xe6, 0x1e, 0xd2, 0x75, 0x01, 0x23, 0xa7, 0x30, 0xf0, 0x18, 0xf7,
	0x13, 0x89, 0x77, 0x25, 0xb7, 0xf3, 0x8c, 0x7e, 0x0e, 0xd4, 0x5a, 0xb9, 0x84, 0x27, 0xc5, 0xed,
	0x9c, 0xfc, 0x63, 0x11, 0xe7, 0x9a, 0xfb, 0xe6, 0x91, 0xbe, 0x14, 0x83, 0xa3, 0xfb, 0x41, 0xe5,
	0x89, 0x7d, 0xf2, 0xfb, 0xf1, 0x92, 0xab, 0x55, 0x32, 0x1f, 0x2f, 0x44, 0x30, 0xf1, 0xd7, 0x9e,
	0x9a, 0x94, 0xff, 0x03, 0x4b, 0x0c, 0x27, 0xd1, 0xfc, 0xbb, 0xa5, 0x98, 0x6c, 0xfd, 0x22, 0xcc,
	0x3b, 0x5a, 0x3a, 0x6f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x89, 0x92, 0x03, 0xa5, 0xe3, 0x08,
	0x00, 0x00,
}
