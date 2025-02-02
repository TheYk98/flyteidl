// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/cluster_assignment.proto

package admin

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

// Defines how a label with a corresponding key and value is selected or excluded.
type Selector_Operator int32

const (
	Selector_EQUALS     Selector_Operator = 0
	Selector_NOT_EQUALS Selector_Operator = 1
	Selector_IN         Selector_Operator = 2
	Selector_NOT_IN     Selector_Operator = 3
	Selector_EXISTS     Selector_Operator = 4
)

var Selector_Operator_name = map[int32]string{
	0: "EQUALS",
	1: "NOT_EQUALS",
	2: "IN",
	3: "NOT_IN",
	4: "EXISTS",
}

var Selector_Operator_value = map[string]int32{
	"EQUALS":     0,
	"NOT_EQUALS": 1,
	"IN":         2,
	"NOT_IN":     3,
	"EXISTS":     4,
}

func (x Selector_Operator) String() string {
	return proto.EnumName(Selector_Operator_name, int32(x))
}

func (Selector_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5699de6e686ea15, []int{3, 0}
}

// Encapsulates specifications for routing an execution onto a specific cluster.
type ClusterAssignment struct {
	Affinity             *Affinity   `protobuf:"bytes,1,opt,name=affinity,proto3" json:"affinity,omitempty"`
	Toleration           *Toleration `protobuf:"bytes,2,opt,name=toleration,proto3" json:"toleration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClusterAssignment) Reset()         { *m = ClusterAssignment{} }
func (m *ClusterAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterAssignment) ProtoMessage()    {}
func (*ClusterAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5699de6e686ea15, []int{0}
}

func (m *ClusterAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterAssignment.Unmarshal(m, b)
}
func (m *ClusterAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterAssignment.Merge(m, src)
}
func (m *ClusterAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterAssignment.Size(m)
}
func (m *ClusterAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterAssignment proto.InternalMessageInfo

func (m *ClusterAssignment) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *ClusterAssignment) GetToleration() *Toleration {
	if m != nil {
		return m.Toleration
	}
	return nil
}

// Defines a set of constraints used to select eligible objects based on labels they possess.
type Affinity struct {
	// Multiples selectors are 'and'-ed together to produce the list of matching, eligible objects.
	Selectors            []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Affinity) Reset()         { *m = Affinity{} }
func (m *Affinity) String() string { return proto.CompactTextString(m) }
func (*Affinity) ProtoMessage()    {}
func (*Affinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5699de6e686ea15, []int{1}
}

func (m *Affinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affinity.Unmarshal(m, b)
}
func (m *Affinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affinity.Marshal(b, m, deterministic)
}
func (m *Affinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affinity.Merge(m, src)
}
func (m *Affinity) XXX_Size() int {
	return xxx_messageInfo_Affinity.Size(m)
}
func (m *Affinity) XXX_DiscardUnknown() {
	xxx_messageInfo_Affinity.DiscardUnknown(m)
}

var xxx_messageInfo_Affinity proto.InternalMessageInfo

func (m *Affinity) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// Defines a set of specific label selectors that the execution can tolerate on a cluster.
type Toleration struct {
	// A toleration selector is similar to that of an affinity but the only valid operators are EQUALS AND EXISTS.
	Selectors            []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Toleration) Reset()         { *m = Toleration{} }
func (m *Toleration) String() string { return proto.CompactTextString(m) }
func (*Toleration) ProtoMessage()    {}
func (*Toleration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5699de6e686ea15, []int{2}
}

func (m *Toleration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Toleration.Unmarshal(m, b)
}
func (m *Toleration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Toleration.Marshal(b, m, deterministic)
}
func (m *Toleration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Toleration.Merge(m, src)
}
func (m *Toleration) XXX_Size() int {
	return xxx_messageInfo_Toleration.Size(m)
}
func (m *Toleration) XXX_DiscardUnknown() {
	xxx_messageInfo_Toleration.DiscardUnknown(m)
}

var xxx_messageInfo_Toleration proto.InternalMessageInfo

func (m *Toleration) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// A Selector is a specification for identifying a set of objects with corresponding labels.
type Selector struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// One or more values used to match labels.
	// For equality (or inequality) requirements, values must contain a single element.
	// For set-based requirements, values may contain one or more elements.
	Value                []string          `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	Operator             Selector_Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=flyteidl.admin.Selector_Operator" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5699de6e686ea15, []int{3}
}

func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Selector) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Selector) GetOperator() Selector_Operator {
	if m != nil {
		return m.Operator
	}
	return Selector_EQUALS
}

func init() {
	proto.RegisterEnum("flyteidl.admin.Selector_Operator", Selector_Operator_name, Selector_Operator_value)
	proto.RegisterType((*ClusterAssignment)(nil), "flyteidl.admin.ClusterAssignment")
	proto.RegisterType((*Affinity)(nil), "flyteidl.admin.Affinity")
	proto.RegisterType((*Toleration)(nil), "flyteidl.admin.Toleration")
	proto.RegisterType((*Selector)(nil), "flyteidl.admin.Selector")
}

func init() {
	proto.RegisterFile("flyteidl/admin/cluster_assignment.proto", fileDescriptor_c5699de6e686ea15)
}

var fileDescriptor_c5699de6e686ea15 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x4a, 0x3b, 0x31,
	0x18, 0xc5, 0xff, 0x99, 0xf9, 0x5b, 0xa6, 0x9f, 0x50, 0xc6, 0xe0, 0x62, 0x70, 0x35, 0xce, 0xc6,
	0x6e, 0x9c, 0x40, 0xbd, 0x81, 0xe0, 0xa2, 0xf5, 0x02, 0x05, 0x69, 0x31, 0x53, 0x41, 0xdc, 0x94,
	0x69, 0x9b, 0x8e, 0xc1, 0x34, 0x29, 0x99, 0x54, 0xe8, 0x03, 0xf8, 0x50, 0xbe, 0x9d, 0xcc, 0xb5,
	0x56, 0x74, 0xe3, 0x2e, 0xdf, 0x97, 0xdf, 0x39, 0x27, 0x84, 0x03, 0x47, 0x73, 0xb1, 0x36, 0x8c,
	0xcf, 0x04, 0x89, 0x67, 0x0b, 0x2e, 0xc9, 0x54, 0xac, 0x52, 0xc3, 0xf4, 0x38, 0x4e, 0x53, 0x9e,
	0xc8, 0x05, 0x93, 0x26, 0x5c, 0x6a, 0x65, 0x14, 0x6e, 0x55, 0x60, 0x98, 0x83, 0xc1, 0x3b, 0x82,
	0xbd, 0xeb, 0x02, 0xee, 0xd6, 0x2c, 0x3e, 0x05, 0x27, 0x9e, 0xcf, 0xb9, 0xe4, 0x66, 0xed, 0x21,
	0x1f, 0xb5, 0x77, 0x3b, 0x5e, 0xb8, 0x2d, 0x0c, 0xbb, 0xe5, 0x3d, 0xad, 0x49, 0x7c, 0x09, 0x60,
	0x94, 0x60, 0x3a, 0x36, 0x5c, 0x49, 0xcf, 0xca, 0x75, 0x07, 0xdf, 0x75, 0xa3, 0x9a, 0xa0, 0x5f,
	0xe8, 0xa0, 0x07, 0x4e, 0xe5, 0x88, 0xcf, 0xa1, 0x99, 0x32, 0xc1, 0xa6, 0x46, 0xe9, 0xd4, 0x43,
	0xbe, 0xfd, 0x53, 0x7c, 0x54, 0x02, 0x74, 0x83, 0x06, 0x37, 0x00, 0x1b, 0xf7, 0x3f, 0xbb, 0x7c,
	0x20, 0x70, 0xaa, 0x3d, 0x76, 0xc1, 0x7e, 0x65, 0xc5, 0x1f, 0x34, 0x69, 0x76, 0xc4, 0xfb, 0xb0,
	0xf3, 0x16, 0x8b, 0x15, 0xf3, 0x2c, 0xdf, 0x6e, 0x37, 0x69, 0x31, 0xe0, 0x2b, 0x70, 0xd4, 0x32,
	0x4b, 0x56, 0xda, 0xb3, 0x7d, 0xd4, 0x6e, 0x75, 0x0e, 0x7f, 0xcb, 0x0a, 0x87, 0x25, 0x48, 0x6b,
	0x49, 0x70, 0x07, 0x4e, 0xb5, 0xc5, 0x00, 0x8d, 0xdb, 0x87, 0xc7, 0xee, 0x7d, 0xe4, 0xfe, 0xc3,
	0x2d, 0x80, 0xc1, 0x70, 0x34, 0x2e, 0x67, 0x84, 0x1b, 0x60, 0xf5, 0x07, 0xae, 0x95, 0x31, 0xd9,
	0xbe, 0x3f, 0x70, 0xed, 0x9c, 0x7f, 0xea, 0x47, 0xa3, 0xc8, 0xfd, 0xdf, 0xbb, 0x78, 0x3e, 0x4b,
	0xb8, 0x79, 0x59, 0x4d, 0xc2, 0xa9, 0x5a, 0x90, 0xfc, 0x01, 0x4a, 0x27, 0xa4, 0x2e, 0x47, 0xc2,
	0x24, 0x59, 0x4e, 0x8e, 0x13, 0x45, 0xb6, 0xfb, 0x32, 0x69, 0xe4, 0xed, 0x38, 0xf9, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x04, 0xfd, 0xc4, 0x48, 0x02, 0x00, 0x00,
}
