// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/compiler.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Adjacency list for the workflow. This is created as part of the compilation process. Every process after the compilation
// step uses this created ConnectionSet
type ConnectionSet struct {
	// A list of all the node ids that are downstream from a given node id
	Downstream map[string]*ConnectionSet_IdList `protobuf:"bytes,7,rep,name=downstream,proto3" json:"downstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of all the node ids, that are upstream of this node id
	Upstream             map[string]*ConnectionSet_IdList `protobuf:"bytes,8,rep,name=upstream,proto3" json:"upstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ConnectionSet) Reset()         { *m = ConnectionSet{} }
func (m *ConnectionSet) String() string { return proto.CompactTextString(m) }
func (*ConnectionSet) ProtoMessage()    {}
func (*ConnectionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{0}
}

func (m *ConnectionSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSet.Unmarshal(m, b)
}
func (m *ConnectionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSet.Marshal(b, m, deterministic)
}
func (m *ConnectionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSet.Merge(m, src)
}
func (m *ConnectionSet) XXX_Size() int {
	return xxx_messageInfo_ConnectionSet.Size(m)
}
func (m *ConnectionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSet.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSet proto.InternalMessageInfo

func (m *ConnectionSet) GetDownstream() map[string]*ConnectionSet_IdList {
	if m != nil {
		return m.Downstream
	}
	return nil
}

func (m *ConnectionSet) GetUpstream() map[string]*ConnectionSet_IdList {
	if m != nil {
		return m.Upstream
	}
	return nil
}

type ConnectionSet_IdList struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionSet_IdList) Reset()         { *m = ConnectionSet_IdList{} }
func (m *ConnectionSet_IdList) String() string { return proto.CompactTextString(m) }
func (*ConnectionSet_IdList) ProtoMessage()    {}
func (*ConnectionSet_IdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{0, 0}
}

func (m *ConnectionSet_IdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSet_IdList.Unmarshal(m, b)
}
func (m *ConnectionSet_IdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSet_IdList.Marshal(b, m, deterministic)
}
func (m *ConnectionSet_IdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSet_IdList.Merge(m, src)
}
func (m *ConnectionSet_IdList) XXX_Size() int {
	return xxx_messageInfo_ConnectionSet_IdList.Size(m)
}
func (m *ConnectionSet_IdList) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSet_IdList.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSet_IdList proto.InternalMessageInfo

func (m *ConnectionSet_IdList) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Output of the compilation Step. This object represents one workflow. We store more metadata at this layer
type CompiledWorkflow struct {
	// Completely contained Workflow Template
	Template *WorkflowTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	// For internal use only! This field is used by the system and must not be filled in. Any values set will be ignored.
	Connections          *ConnectionSet `protobuf:"bytes,2,opt,name=connections,proto3" json:"connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CompiledWorkflow) Reset()         { *m = CompiledWorkflow{} }
func (m *CompiledWorkflow) String() string { return proto.CompactTextString(m) }
func (*CompiledWorkflow) ProtoMessage()    {}
func (*CompiledWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{1}
}

func (m *CompiledWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledWorkflow.Unmarshal(m, b)
}
func (m *CompiledWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledWorkflow.Marshal(b, m, deterministic)
}
func (m *CompiledWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledWorkflow.Merge(m, src)
}
func (m *CompiledWorkflow) XXX_Size() int {
	return xxx_messageInfo_CompiledWorkflow.Size(m)
}
func (m *CompiledWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledWorkflow proto.InternalMessageInfo

func (m *CompiledWorkflow) GetTemplate() *WorkflowTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CompiledWorkflow) GetConnections() *ConnectionSet {
	if m != nil {
		return m.Connections
	}
	return nil
}

// Output of the Compilation step. This object represent one Task. We store more metadata at this layer
type CompiledTask struct {
	// Completely contained TaskTemplate
	Template             *TaskTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CompiledTask) Reset()         { *m = CompiledTask{} }
func (m *CompiledTask) String() string { return proto.CompactTextString(m) }
func (*CompiledTask) ProtoMessage()    {}
func (*CompiledTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{2}
}

func (m *CompiledTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledTask.Unmarshal(m, b)
}
func (m *CompiledTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledTask.Marshal(b, m, deterministic)
}
func (m *CompiledTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledTask.Merge(m, src)
}
func (m *CompiledTask) XXX_Size() int {
	return xxx_messageInfo_CompiledTask.Size(m)
}
func (m *CompiledTask) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledTask.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledTask proto.InternalMessageInfo

func (m *CompiledTask) GetTemplate() *TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

// A Compiled Workflow Closure contains all the information required to start a new execution, or to visualize a workflow
// and its details. The CompiledWorkflowClosure should always contain a primary workflow, that is the main workflow that
// will being the execution. All subworkflows are denormalized. WorkflowNodes refer to the workflow identifiers of
// compiled subworkflows.
type CompiledWorkflowClosure struct {
	//+required
	Primary *CompiledWorkflow `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	// Guaranteed that there will only exist one and only one workflow with a given id, i.e., every sub workflow has a
	// unique identifier. Also every enclosed subworkflow is used either by a primary workflow or by a subworkflow
	// as an inlined workflow
	//+optional
	SubWorkflows []*CompiledWorkflow `protobuf:"bytes,2,rep,name=sub_workflows,json=subWorkflows,proto3" json:"sub_workflows,omitempty"`
	// Guaranteed that there will only exist one and only one task with a given id, i.e., every task has a unique id
	//+required (at least 1)
	Tasks                []*CompiledTask `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CompiledWorkflowClosure) Reset()         { *m = CompiledWorkflowClosure{} }
func (m *CompiledWorkflowClosure) String() string { return proto.CompactTextString(m) }
func (*CompiledWorkflowClosure) ProtoMessage()    {}
func (*CompiledWorkflowClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{3}
}

func (m *CompiledWorkflowClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledWorkflowClosure.Unmarshal(m, b)
}
func (m *CompiledWorkflowClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledWorkflowClosure.Marshal(b, m, deterministic)
}
func (m *CompiledWorkflowClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledWorkflowClosure.Merge(m, src)
}
func (m *CompiledWorkflowClosure) XXX_Size() int {
	return xxx_messageInfo_CompiledWorkflowClosure.Size(m)
}
func (m *CompiledWorkflowClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledWorkflowClosure.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledWorkflowClosure proto.InternalMessageInfo

func (m *CompiledWorkflowClosure) GetPrimary() *CompiledWorkflow {
	if m != nil {
		return m.Primary
	}
	return nil
}

func (m *CompiledWorkflowClosure) GetSubWorkflows() []*CompiledWorkflow {
	if m != nil {
		return m.SubWorkflows
	}
	return nil
}

func (m *CompiledWorkflowClosure) GetTasks() []*CompiledTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionSet)(nil), "flyteidl.core.ConnectionSet")
	proto.RegisterMapType((map[string]*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.DownstreamEntry")
	proto.RegisterMapType((map[string]*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.UpstreamEntry")
	proto.RegisterType((*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.IdList")
	proto.RegisterType((*CompiledWorkflow)(nil), "flyteidl.core.CompiledWorkflow")
	proto.RegisterType((*CompiledTask)(nil), "flyteidl.core.CompiledTask")
	proto.RegisterType((*CompiledWorkflowClosure)(nil), "flyteidl.core.CompiledWorkflowClosure")
}

func init() { proto.RegisterFile("flyteidl/core/compiler.proto", fileDescriptor_d310dba155e4f065) }

var fileDescriptor_d310dba155e4f065 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0xca, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0x43, 0x2f, 0x9e, 0x34, 0x58, 0x66, 0x63, 0x4c, 0x0b, 0x86, 0xb8, 0x29, 0xa2,
	0x09, 0x56, 0x51, 0xab, 0xe0, 0xc2, 0x56, 0x45, 0xe8, 0x2a, 0x56, 0x04, 0x37, 0x9a, 0xcb, 0x34,
	0x86, 0x5c, 0x26, 0xcc, 0x4c, 0x2c, 0x79, 0x02, 0x97, 0x3e, 0x99, 0xef, 0xf4, 0x91, 0x2b, 0x4d,
	0x4a, 0xcb, 0xb7, 0xf8, 0x76, 0x2d, 0xe7, 0xf7, 0xbf, 0x1c, 0x26, 0x07, 0x16, 0x87, 0x28, 0xe7,
	0x38, 0xf0, 0x22, 0xd3, 0x25, 0x14, 0x9b, 0x2e, 0x89, 0xd3, 0x20, 0xc2, 0xd4, 0x48, 0x29, 0xe1,
	0x04, 0xc9, 0xcd, 0xd4, 0x28, 0xa6, 0x6a, 0x0f, 0x3e, 0x12, 0x1a, 0x1e, 0x22, 0x72, 0xac, 0x60,
	0xf5, 0x61, 0x77, 0xca, 0x6d, 0x16, 0xb2, 0x6a, 0xa4, 0xff, 0x15, 0x41, 0xde, 0x90, 0x24, 0xc1,
	0x2e, 0x0f, 0x48, 0xf2, 0x15, 0x73, 0xb4, 0x03, 0xf0, 0xc8, 0x31, 0x61, 0x9c, 0x62, 0x3b, 0x56,
	0xc6, 0x9a, 0xb8, 0x94, 0x56, 0x4f, 0x8d, 0x4e, 0x9c, 0xd1, 0x51, 0x18, 0xdb, 0x16, 0xff, 0x98,
	0x70, 0x9a, 0x5b, 0x27, 0x7a, 0xf4, 0x09, 0x26, 0x59, 0x5a, 0x7b, 0x4d, 0x4a, 0xaf, 0x27, 0x57,
	0xbd, 0xbe, 0xa5, 0xa7, 0x4e, 0xad, 0x56, 0x55, 0x61, 0xf4, 0xc5, 0xdb, 0x05, 0x8c, 0xa3, 0x19,
	0x88, 0x81, 0xc7, 0x14, 0x41, 0x13, 0x97, 0xf7, 0xac, 0xe2, 0xa7, 0xea, 0xc0, 0xfd, 0x5e, 0x85,
	0x02, 0x0a, 0x71, 0xae, 0x08, 0x9a, 0x50, 0x40, 0x21, 0xce, 0xd1, 0x1a, 0x86, 0x7f, 0xec, 0x28,
	0xc3, 0xca, 0x40, 0x13, 0x96, 0xd2, 0xea, 0xf1, 0xd5, 0x16, 0x55, 0x94, 0x55, 0x29, 0xde, 0x0e,
	0xde, 0x08, 0xea, 0x2f, 0x90, 0x3b, 0xd5, 0xee, 0x3c, 0x41, 0xff, 0x27, 0xc0, 0x6c, 0x53, 0x3d,
	0xb2, 0xf7, 0xbd, 0x7e, 0x3f, 0xf4, 0x0e, 0x26, 0x1c, 0xc7, 0x69, 0x64, 0x73, 0x5c, 0x46, 0x49,
	0xab, 0x47, 0x3d, 0xdb, 0x06, 0xdd, 0xd7, 0x98, 0xd5, 0x0a, 0xd0, 0x7b, 0x90, 0xdc, 0x36, 0x94,
	0xd5, 0xb5, 0x16, 0xd7, 0x6a, 0x59, 0xa7, 0x02, 0xfd, 0x33, 0x4c, 0x9b, 0x42, 0x7b, 0x9b, 0x85,
	0xe8, 0xf5, 0x59, 0x99, 0x79, 0xcf, 0xac, 0xc0, 0xce, 0x8b, 0xe8, 0xff, 0x05, 0x78, 0xd0, 0x5f,
	0x6d, 0x13, 0x11, 0x96, 0x51, 0x8c, 0xd6, 0x30, 0x4e, 0x69, 0x10, 0xdb, 0x34, 0xbf, 0xb0, 0x60,
	0x5f, 0x68, 0x35, 0x3c, 0xda, 0x82, 0xcc, 0x32, 0xe7, 0x67, 0xf3, 0xb1, 0x17, 0x1b, 0x8a, 0xb7,
	0x31, 0x98, 0xb2, 0xcc, 0x69, 0xfe, 0x30, 0xf4, 0x1c, 0x86, 0xe5, 0x41, 0x28, 0x62, 0xa9, 0x9e,
	0x5f, 0x50, 0x17, 0xab, 0x59, 0x15, 0xf9, 0xe1, 0xd5, 0x8f, 0x97, 0x7e, 0xc0, 0x7f, 0x67, 0x8e,
	0xe1, 0x92, 0xd8, 0x2c, 0x79, 0x42, 0x7d, 0xb3, 0xbd, 0x32, 0x1f, 0x27, 0x66, 0xea, 0x3c, 0xf3,
	0x89, 0xd9, 0x39, 0x3c, 0x67, 0x54, 0xde, 0xdc, 0x8b, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x1f, 0x50, 0x2b, 0xdb, 0x03, 0x00, 0x00,
}
