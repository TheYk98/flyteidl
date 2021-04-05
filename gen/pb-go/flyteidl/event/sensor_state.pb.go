// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/sensor_state.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Defines custom state an event sensor of a specific type can durably persist.
type EventSensorState struct {
	// Types that are valid to be assigned to State:
	//	*EventSensorState_Uri
	//	*EventSensorState_CustomInfo
	State                isEventSensorState_State `protobuf_oneof:"state"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EventSensorState) Reset()         { *m = EventSensorState{} }
func (m *EventSensorState) String() string { return proto.CompactTextString(m) }
func (*EventSensorState) ProtoMessage()    {}
func (*EventSensorState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c90cbc8ab4108e95, []int{0}
}

func (m *EventSensorState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSensorState.Unmarshal(m, b)
}
func (m *EventSensorState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSensorState.Marshal(b, m, deterministic)
}
func (m *EventSensorState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSensorState.Merge(m, src)
}
func (m *EventSensorState) XXX_Size() int {
	return xxx_messageInfo_EventSensorState.Size(m)
}
func (m *EventSensorState) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSensorState.DiscardUnknown(m)
}

var xxx_messageInfo_EventSensorState proto.InternalMessageInfo

type isEventSensorState_State interface {
	isEventSensorState_State()
}

type EventSensorState_Uri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3,oneof"`
}

type EventSensorState_CustomInfo struct {
	CustomInfo *_struct.Struct `protobuf:"bytes,2,opt,name=custom_info,json=customInfo,proto3,oneof"`
}

func (*EventSensorState_Uri) isEventSensorState_State() {}

func (*EventSensorState_CustomInfo) isEventSensorState_State() {}

func (m *EventSensorState) GetState() isEventSensorState_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *EventSensorState) GetUri() string {
	if x, ok := m.GetState().(*EventSensorState_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *EventSensorState) GetCustomInfo() *_struct.Struct {
	if x, ok := m.GetState().(*EventSensorState_CustomInfo); ok {
		return x.CustomInfo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventSensorState) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventSensorState_Uri)(nil),
		(*EventSensorState_CustomInfo)(nil),
	}
}

// A request to fetch saved state for a specific event type.
type EventSensorStateGetRequest struct {
	// Indicates the event sensor type by unique name for which to fetch event sensor state.
	EventType            string   `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSensorStateGetRequest) Reset()         { *m = EventSensorStateGetRequest{} }
func (m *EventSensorStateGetRequest) String() string { return proto.CompactTextString(m) }
func (*EventSensorStateGetRequest) ProtoMessage()    {}
func (*EventSensorStateGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c90cbc8ab4108e95, []int{1}
}

func (m *EventSensorStateGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSensorStateGetRequest.Unmarshal(m, b)
}
func (m *EventSensorStateGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSensorStateGetRequest.Marshal(b, m, deterministic)
}
func (m *EventSensorStateGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSensorStateGetRequest.Merge(m, src)
}
func (m *EventSensorStateGetRequest) XXX_Size() int {
	return xxx_messageInfo_EventSensorStateGetRequest.Size(m)
}
func (m *EventSensorStateGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSensorStateGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventSensorStateGetRequest proto.InternalMessageInfo

func (m *EventSensorStateGetRequest) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

// Saved state for a specific event sensor type.
type EventSensorStateGetResponse struct {
	State *EventSensorState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// Event sensor provided timestamp for when this state was last specified.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventSensorStateGetResponse) Reset()         { *m = EventSensorStateGetResponse{} }
func (m *EventSensorStateGetResponse) String() string { return proto.CompactTextString(m) }
func (*EventSensorStateGetResponse) ProtoMessage()    {}
func (*EventSensorStateGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c90cbc8ab4108e95, []int{2}
}

func (m *EventSensorStateGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSensorStateGetResponse.Unmarshal(m, b)
}
func (m *EventSensorStateGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSensorStateGetResponse.Marshal(b, m, deterministic)
}
func (m *EventSensorStateGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSensorStateGetResponse.Merge(m, src)
}
func (m *EventSensorStateGetResponse) XXX_Size() int {
	return xxx_messageInfo_EventSensorStateGetResponse.Size(m)
}
func (m *EventSensorStateGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSensorStateGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventSensorStateGetResponse proto.InternalMessageInfo

func (m *EventSensorStateGetResponse) GetState() *EventSensorState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *EventSensorStateGetResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Updates saved state for a specific event sensor type.
type EventSensorStateUpdateRequest struct {
	// Indicates the event sensor type by unique name for which to fetch event sensor state.
	EventType string            `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	State     *EventSensorState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// Event sensor provided timestamp for when this state was last specified.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventSensorStateUpdateRequest) Reset()         { *m = EventSensorStateUpdateRequest{} }
func (m *EventSensorStateUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*EventSensorStateUpdateRequest) ProtoMessage()    {}
func (*EventSensorStateUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c90cbc8ab4108e95, []int{3}
}

func (m *EventSensorStateUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSensorStateUpdateRequest.Unmarshal(m, b)
}
func (m *EventSensorStateUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSensorStateUpdateRequest.Marshal(b, m, deterministic)
}
func (m *EventSensorStateUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSensorStateUpdateRequest.Merge(m, src)
}
func (m *EventSensorStateUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_EventSensorStateUpdateRequest.Size(m)
}
func (m *EventSensorStateUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSensorStateUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventSensorStateUpdateRequest proto.InternalMessageInfo

func (m *EventSensorStateUpdateRequest) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventSensorStateUpdateRequest) GetState() *EventSensorState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *EventSensorStateUpdateRequest) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type EventSensorStateUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSensorStateUpdateResponse) Reset()         { *m = EventSensorStateUpdateResponse{} }
func (m *EventSensorStateUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*EventSensorStateUpdateResponse) ProtoMessage()    {}
func (*EventSensorStateUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c90cbc8ab4108e95, []int{4}
}

func (m *EventSensorStateUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSensorStateUpdateResponse.Unmarshal(m, b)
}
func (m *EventSensorStateUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSensorStateUpdateResponse.Marshal(b, m, deterministic)
}
func (m *EventSensorStateUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSensorStateUpdateResponse.Merge(m, src)
}
func (m *EventSensorStateUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_EventSensorStateUpdateResponse.Size(m)
}
func (m *EventSensorStateUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSensorStateUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventSensorStateUpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventSensorState)(nil), "flyteidl.event.EventSensorState")
	proto.RegisterType((*EventSensorStateGetRequest)(nil), "flyteidl.event.EventSensorStateGetRequest")
	proto.RegisterType((*EventSensorStateGetResponse)(nil), "flyteidl.event.EventSensorStateGetResponse")
	proto.RegisterType((*EventSensorStateUpdateRequest)(nil), "flyteidl.event.EventSensorStateUpdateRequest")
	proto.RegisterType((*EventSensorStateUpdateResponse)(nil), "flyteidl.event.EventSensorStateUpdateResponse")
}

func init() { proto.RegisterFile("flyteidl/event/sensor_state.proto", fileDescriptor_c90cbc8ab4108e95) }

var fileDescriptor_c90cbc8ab4108e95 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x9b, 0x96, 0xf7, 0x1e, 0xbd, 0x85, 0x87, 0x64, 0x63, 0x89, 0x56, 0x63, 0x56, 0xdd,
	0x38, 0x03, 0x15, 0x15, 0x75, 0x65, 0x41, 0xac, 0xdb, 0xb4, 0x6e, 0xdc, 0x84, 0xa4, 0xbd, 0x89,
	0xc1, 0x26, 0x33, 0x66, 0xee, 0x08, 0xfd, 0x17, 0xfe, 0x15, 0xff, 0xa1, 0x74, 0xa6, 0x55, 0x1a,
	0x15, 0x14, 0x97, 0xc9, 0x3d, 0xf7, 0x7c, 0xe7, 0x5c, 0x06, 0x0e, 0xd2, 0xf9, 0x82, 0x30, 0x9f,
	0xcd, 0x39, 0x3e, 0x61, 0x49, 0x5c, 0x61, 0xa9, 0x44, 0x15, 0x29, 0x8a, 0x09, 0x99, 0xac, 0x04,
	0x09, 0xf7, 0xff, 0x5a, 0xc2, 0x8c, 0xc4, 0xdb, 0xcd, 0x84, 0xc8, 0xe6, 0xc8, 0xcd, 0x34, 0xd1,
	0x29, 0x57, 0x54, 0xe9, 0x29, 0x59, 0xb5, 0xb7, 0x5f, 0x9f, 0x52, 0x5e, 0xa0, 0xa2, 0xb8, 0x90,
	0x56, 0x10, 0x3c, 0xc0, 0xd6, 0xd5, 0xd2, 0x67, 0x6c, 0x48, 0xe3, 0x25, 0xc8, 0x75, 0xa1, 0xa5,
	0xab, 0xbc, 0xeb, 0xf8, 0x4e, 0xbf, 0x3d, 0x6a, 0x84, 0xcb, 0x0f, 0xf7, 0x1c, 0x3a, 0x53, 0xad,
	0x48, 0x14, 0x51, 0x5e, 0xa6, 0xa2, 0xdb, 0xf4, 0x9d, 0x7e, 0x67, 0xb0, 0xcd, 0xac, 0x3d, 0x5b,
	0xdb, 0xb3, 0xb1, 0x81, 0x8f, 0x1a, 0x21, 0x58, 0xf5, 0x4d, 0x99, 0x8a, 0xe1, 0x3f, 0xf8, 0x63,
	0x1a, 0x04, 0x17, 0xe0, 0xd5, 0x61, 0xd7, 0x48, 0x21, 0x3e, 0x6a, 0x54, 0xe4, 0xf6, 0x00, 0x4c,
	0xa5, 0x88, 0x16, 0x12, 0x2d, 0x3d, 0x6c, 0x9b, 0x3f, 0x93, 0x85, 0xc4, 0xe0, 0xd9, 0x81, 0x9d,
	0x4f, 0xb7, 0x95, 0x14, 0xa5, 0x42, 0xf7, 0x64, 0x45, 0x31, 0x9b, 0x9d, 0x81, 0xcf, 0x36, 0x0f,
	0xc5, 0xea, 0xbb, 0xa1, 0x95, 0xbb, 0x67, 0x00, 0x5a, 0xce, 0x62, 0xc2, 0x59, 0x14, 0xd3, 0xaa,
	0x98, 0xf7, 0xa1, 0xd8, 0x64, 0x7d, 0xb7, 0xb0, 0xbd, 0x52, 0x5f, 0x52, 0xf0, 0xe2, 0x40, 0xaf,
	0x6e, 0x7b, 0x6b, 0xa6, 0xdf, 0xeb, 0xf4, 0x9e, 0xb9, 0xf9, 0x9b, 0xcc, 0xad, 0x9f, 0x64, 0xf6,
	0x61, 0xef, 0xab, 0xc8, 0xf6, 0x90, 0xc3, 0xd3, 0xbb, 0xe3, 0x2c, 0xa7, 0x7b, 0x9d, 0xb0, 0xa9,
	0x28, 0xb8, 0x49, 0x24, 0xaa, 0x8c, 0xbf, 0x3d, 0xcd, 0x0c, 0x4b, 0x2e, 0x93, 0xc3, 0x4c, 0xf0,
	0xcd, 0xd7, 0x9a, 0xfc, 0x35, 0xe4, 0xa3, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x50, 0x85,
	0x4b, 0xc6, 0x02, 0x00, 0x00,
}
