// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/tasks.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskCategory int32

const (
	// Task category that identifies if the system can use default structures in UI, etc to drive the task
	// TODO should we add Container type of tasks as a special Class?
	TaskCategory_SingleStepTask TaskCategory = 0
	TaskCategory_MultiStepTask  TaskCategory = 1
)

var TaskCategory_name = map[int32]string{
	0: "SingleStepTask",
	1: "MultiStepTask",
}
var TaskCategory_value = map[string]int32{
	"SingleStepTask": 0,
	"MultiStepTask":  1,
}

func (x TaskCategory) String() string {
	return proto.EnumName(TaskCategory_name, int32(x))
}
func (TaskCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{0}
}

// Known resource names.
type Resources_ResourceName int32

const (
	Resources_Unknown Resources_ResourceName = 0
	Resources_Cpu     Resources_ResourceName = 1
	Resources_Gpu     Resources_ResourceName = 2
	Resources_Memory  Resources_ResourceName = 3
	Resources_Storage Resources_ResourceName = 4
)

var Resources_ResourceName_name = map[int32]string{
	0: "Unknown",
	1: "Cpu",
	2: "Gpu",
	3: "Memory",
	4: "Storage",
}
var Resources_ResourceName_value = map[string]int32{
	"Unknown": 0,
	"Cpu":     1,
	"Gpu":     2,
	"Memory":  3,
	"Storage": 4,
}

func (x Resources_ResourceName) String() string {
	return proto.EnumName(Resources_ResourceName_name, int32(x))
}
func (Resources_ResourceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{0, 0}
}

type RuntimeMetadata_RuntimeType int32

const (
	RuntimeMetadata_Other    RuntimeMetadata_RuntimeType = 0
	RuntimeMetadata_FlyteSDK RuntimeMetadata_RuntimeType = 1
)

var RuntimeMetadata_RuntimeType_name = map[int32]string{
	0: "Other",
	1: "FlyteSDK",
}
var RuntimeMetadata_RuntimeType_value = map[string]int32{
	"Other":    0,
	"FlyteSDK": 1,
}

func (x RuntimeMetadata_RuntimeType) String() string {
	return proto.EnumName(RuntimeMetadata_RuntimeType_name, int32(x))
}
func (RuntimeMetadata_RuntimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{1, 0}
}

// A customizable interface to convey resources requested for a container. This can be interpretted differently for different
// container engines.
type Resources struct {
	// The desired set of resources requested.
	Requests []*Resources_ResourceEntry `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run.
	Limits               []*Resources_ResourceEntry `protobuf:"bytes,2,rep,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{0}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetRequests() []*Resources_ResourceEntry {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Resources) GetLimits() []*Resources_ResourceEntry {
	if m != nil {
		return m.Limits
	}
	return nil
}

// Encapsulates a resource name and value.
type Resources_ResourceEntry struct {
	// Resource name.
	Name Resources_ResourceName `protobuf:"varint,1,opt,name=name,proto3,enum=flyteidl.core.Resources_ResourceName" json:"name,omitempty"`
	// Value must be a valid k8s quantity.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_ResourceEntry) Reset()         { *m = Resources_ResourceEntry{} }
func (m *Resources_ResourceEntry) String() string { return proto.CompactTextString(m) }
func (*Resources_ResourceEntry) ProtoMessage()    {}
func (*Resources_ResourceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{0, 0}
}
func (m *Resources_ResourceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_ResourceEntry.Unmarshal(m, b)
}
func (m *Resources_ResourceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_ResourceEntry.Marshal(b, m, deterministic)
}
func (dst *Resources_ResourceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_ResourceEntry.Merge(dst, src)
}
func (m *Resources_ResourceEntry) XXX_Size() int {
	return xxx_messageInfo_Resources_ResourceEntry.Size(m)
}
func (m *Resources_ResourceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_ResourceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_ResourceEntry proto.InternalMessageInfo

func (m *Resources_ResourceEntry) GetName() Resources_ResourceName {
	if m != nil {
		return m.Name
	}
	return Resources_Unknown
}

func (m *Resources_ResourceEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Runtime information. This is losely defined to allow for extensibility.
type RuntimeMetadata struct {
	// Type of runtime.
	Type RuntimeMetadata_RuntimeType `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl.core.RuntimeMetadata_RuntimeType" json:"type,omitempty"`
	// Version of the runtime. All versions should be backward compatible. However, certain cases call for version
	// checks to ensure tighter validation or setting expectations.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// +optional It can be used to provide extra information about the runtime (e.g. python, golang... etc.).
	Flavor               string   `protobuf:"bytes,3,opt,name=flavor,proto3" json:"flavor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeMetadata) Reset()         { *m = RuntimeMetadata{} }
func (m *RuntimeMetadata) String() string { return proto.CompactTextString(m) }
func (*RuntimeMetadata) ProtoMessage()    {}
func (*RuntimeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{1}
}
func (m *RuntimeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeMetadata.Unmarshal(m, b)
}
func (m *RuntimeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeMetadata.Marshal(b, m, deterministic)
}
func (dst *RuntimeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeMetadata.Merge(dst, src)
}
func (m *RuntimeMetadata) XXX_Size() int {
	return xxx_messageInfo_RuntimeMetadata.Size(m)
}
func (m *RuntimeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeMetadata proto.InternalMessageInfo

func (m *RuntimeMetadata) GetType() RuntimeMetadata_RuntimeType {
	if m != nil {
		return m.Type
	}
	return RuntimeMetadata_Other
}

func (m *RuntimeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RuntimeMetadata) GetFlavor() string {
	if m != nil {
		return m.Flavor
	}
	return ""
}

// Task Metadata
type TaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `protobuf:"varint,1,opt,name=discoverable,proto3" json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *RuntimeMetadata `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The overall timeout of a task.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries              *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskMetadata) Reset()         { *m = TaskMetadata{} }
func (m *TaskMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskMetadata) ProtoMessage()    {}
func (*TaskMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{2}
}
func (m *TaskMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMetadata.Unmarshal(m, b)
}
func (m *TaskMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMetadata.Marshal(b, m, deterministic)
}
func (dst *TaskMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMetadata.Merge(dst, src)
}
func (m *TaskMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskMetadata.Size(m)
}
func (m *TaskMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMetadata proto.InternalMessageInfo

func (m *TaskMetadata) GetDiscoverable() bool {
	if m != nil {
		return m.Discoverable
	}
	return false
}

func (m *TaskMetadata) GetRuntime() *RuntimeMetadata {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *TaskMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *TaskMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

// A Task structure that uniquely identifies a task in the system
// Tasks are registered as a first step in the system.
type TaskTemplate struct {
	// Auto generated taskId by the system. Task Id uniquely identifies this task globally.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Category of the task. These are predefined and help provide defaults
	Category TaskCategory `protobuf:"varint,2,opt,name=category,proto3,enum=flyteidl.core.TaskCategory" json:"category,omitempty"`
	// A predefined yet extensible Task type identifier. This can be used to customize any of the components. If no
	// extensions are provided in the system, Flyte will resolve the this task to its TaskCategory and default the
	// implementation registered for the TaskCategory.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Extra metadata about the task.
	Metadata *TaskMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A strongly typed interface for the task. This enables others to use this task within a workflow and gauarantees
	// compile-time validation of the workflow to avoid costly runtime failures.
	Interface *TypedInterface `protobuf:"bytes,5,opt,name=interface,proto3" json:"interface,omitempty"`
	// Custom data about the task. This is extensible to allow various plugins in the system.
	Custom []byte `protobuf:"bytes,6,opt,name=custom,proto3" json:"custom,omitempty"`
	// Known target types that the system will guarantee plugins for. Custom SDK plugins are allowed to set these if needed.
	// If no corresponding execution-layer plugins are found, the system will default to handling these using built-in
	// handlers.
	//
	// Types that are valid to be assigned to Target:
	//	*TaskTemplate_Container
	Target               isTaskTemplate_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskTemplate) Reset()         { *m = TaskTemplate{} }
func (m *TaskTemplate) String() string { return proto.CompactTextString(m) }
func (*TaskTemplate) ProtoMessage()    {}
func (*TaskTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{3}
}
func (m *TaskTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTemplate.Unmarshal(m, b)
}
func (m *TaskTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTemplate.Marshal(b, m, deterministic)
}
func (dst *TaskTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTemplate.Merge(dst, src)
}
func (m *TaskTemplate) XXX_Size() int {
	return xxx_messageInfo_TaskTemplate.Size(m)
}
func (m *TaskTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTemplate proto.InternalMessageInfo

func (m *TaskTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskTemplate) GetCategory() TaskCategory {
	if m != nil {
		return m.Category
	}
	return TaskCategory_SingleStepTask
}

func (m *TaskTemplate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskTemplate) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *TaskTemplate) GetCustom() []byte {
	if m != nil {
		return m.Custom
	}
	return nil
}

type isTaskTemplate_Target interface {
	isTaskTemplate_Target()
}

type TaskTemplate_Container struct {
	Container *Container `protobuf:"bytes,7,opt,name=container,proto3,oneof"`
}

func (*TaskTemplate_Container) isTaskTemplate_Target() {}

func (m *TaskTemplate) GetTarget() isTaskTemplate_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TaskTemplate) GetContainer() *Container {
	if x, ok := m.GetTarget().(*TaskTemplate_Container); ok {
		return x.Container
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskTemplate_OneofMarshaler, _TaskTemplate_OneofUnmarshaler, _TaskTemplate_OneofSizer, []interface{}{
		(*TaskTemplate_Container)(nil),
	}
}

func _TaskTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskTemplate)
	// target
	switch x := m.Target.(type) {
	case *TaskTemplate_Container:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskTemplate.Target has unexpected type %T", x)
	}
	return nil
}

func _TaskTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskTemplate)
	switch tag {
	case 7: // target.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Container)
		err := b.DecodeMessage(msg)
		m.Target = &TaskTemplate_Container{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskTemplate)
	// target
	switch x := m.Target.(type) {
	case *TaskTemplate_Container:
		s := proto.Size(x.Container)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Container struct {
	// Container image url. Eg: docker/redis:latest
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Command to be executed, if not provided, the default entrypoint in the container image will be used.
	Command []string `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	// These will default to Flyte given paths. If provided, the system will not append known paths. If the task still
	// needs flyte's inputs and outputs path, add $(FLYTE_INPUT_FILE), $(FLYTE_OUTPUT_FILE) wherever makes sense and the
	// system will populate these before executing the container.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Container resources requirement as specified by the container engine.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// Environment variables will be set as the container is starting up.
	Env []*KeyValuePair `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
	// Allows extra configs to be available for the container.
	// TODO: elaborate on how configs will become available.
	Config               []*KeyValuePair `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_3ed721e905bb82aa, []int{4}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Container) GetEnv() []*KeyValuePair {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetConfig() []*KeyValuePair {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Resources)(nil), "flyteidl.core.Resources")
	proto.RegisterType((*Resources_ResourceEntry)(nil), "flyteidl.core.Resources.ResourceEntry")
	proto.RegisterType((*RuntimeMetadata)(nil), "flyteidl.core.RuntimeMetadata")
	proto.RegisterType((*TaskMetadata)(nil), "flyteidl.core.TaskMetadata")
	proto.RegisterType((*TaskTemplate)(nil), "flyteidl.core.TaskTemplate")
	proto.RegisterType((*Container)(nil), "flyteidl.core.Container")
	proto.RegisterEnum("flyteidl.core.TaskCategory", TaskCategory_name, TaskCategory_value)
	proto.RegisterEnum("flyteidl.core.Resources_ResourceName", Resources_ResourceName_name, Resources_ResourceName_value)
	proto.RegisterEnum("flyteidl.core.RuntimeMetadata_RuntimeType", RuntimeMetadata_RuntimeType_name, RuntimeMetadata_RuntimeType_value)
}

func init() { proto.RegisterFile("flyteidl/core/tasks.proto", fileDescriptor_tasks_3ed721e905bb82aa) }

var fileDescriptor_tasks_3ed721e905bb82aa = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0xae, 0xdb, 0x44,
	0x10, 0xc7, 0xe3, 0x7c, 0x38, 0xf1, 0x24, 0x27, 0x98, 0x15, 0x42, 0x6e, 0x68, 0xab, 0xc8, 0x12,
	0x55, 0x54, 0xa9, 0x36, 0xca, 0x11, 0x6d, 0x11, 0x52, 0x2f, 0x4e, 0x4a, 0x01, 0x55, 0x07, 0xd0,
	0xe6, 0xc0, 0x05, 0x57, 0x6c, 0xec, 0x89, 0xbb, 0x3a, 0xf6, 0xae, 0x59, 0xaf, 0x83, 0xfc, 0x4c,
	0x3c, 0x01, 0x17, 0x3c, 0x08, 0xcf, 0xc0, 0x4b, 0x20, 0x7f, 0xe6, 0x24, 0x7c, 0x1c, 0x71, 0x15,
	0xcf, 0xce, 0xef, 0x3f, 0x99, 0x99, 0x9d, 0x59, 0x78, 0xb0, 0x8f, 0x0b, 0x8d, 0x3c, 0x8c, 0xfd,
	0x40, 0x2a, 0xf4, 0x35, 0xcb, 0x6e, 0x33, 0x2f, 0x55, 0x52, 0x4b, 0x72, 0xd1, 0xba, 0xbc, 0xd2,
	0xb5, 0x78, 0x74, 0x4a, 0x72, 0xa1, 0x51, 0xed, 0x59, 0x80, 0x35, 0xbd, 0x78, 0x78, 0xea, 0x8e,
	0xb9, 0x46, 0xc5, 0xe2, 0x26, 0xd6, 0xe2, 0x71, 0x24, 0x65, 0x14, 0xa3, 0x5f, 0x59, 0xbb, 0x7c,
	0xef, 0x87, 0xb9, 0x62, 0x9a, 0x4b, 0x51, 0xfb, 0xdd, 0xdf, 0xfa, 0x60, 0x51, 0xcc, 0x64, 0xae,
	0x02, 0xcc, 0xc8, 0x15, 0x4c, 0x14, 0xfe, 0x9c, 0x63, 0xa6, 0x33, 0xc7, 0x58, 0x0e, 0x56, 0xd3,
	0xf5, 0x13, 0xef, 0x24, 0x19, 0xaf, 0x63, 0xbb, 0xaf, 0x2f, 0x84, 0x56, 0x05, 0xed, 0x74, 0xe4,
	0x15, 0x98, 0x31, 0x4f, 0xb8, 0xce, 0x9c, 0xfe, 0xff, 0x8a, 0xd0, 0xa8, 0x16, 0x3f, 0xc1, 0xc5,
	0x89, 0x83, 0x7c, 0x06, 0x43, 0xc1, 0x12, 0x74, 0x8c, 0xa5, 0xb1, 0x9a, 0xaf, 0x3f, 0xbe, 0x37,
	0xdc, 0x37, 0x2c, 0x41, 0x5a, 0x49, 0xc8, 0x07, 0x30, 0x3a, 0xb0, 0x38, 0x47, 0xa7, 0xbf, 0x34,
	0x56, 0x16, 0xad, 0x0d, 0xf7, 0x0d, 0xcc, 0xee, 0xb2, 0x64, 0x0a, 0xe3, 0xef, 0xc5, 0xad, 0x90,
	0xbf, 0x08, 0xbb, 0x47, 0xc6, 0x30, 0xd8, 0xa4, 0xb9, 0x6d, 0x94, 0x1f, 0x5f, 0xa6, 0xb9, 0xdd,
	0x27, 0x00, 0xe6, 0x35, 0x26, 0x52, 0x15, 0xf6, 0xa0, 0x44, 0xb7, 0x5a, 0x2a, 0x16, 0xa1, 0x3d,
	0x74, 0x7f, 0x35, 0xe0, 0x3d, 0x9a, 0x0b, 0xcd, 0x13, 0xbc, 0x46, 0xcd, 0x42, 0xa6, 0x19, 0x79,
	0x05, 0x43, 0x5d, 0xa4, 0x6d, 0xb2, 0x4f, 0xcf, 0x93, 0x3d, 0xa5, 0x5b, 0xfb, 0xa6, 0x48, 0x91,
	0x56, 0x3a, 0xe2, 0xc0, 0xf8, 0x80, 0x2a, 0xe3, 0x52, 0x34, 0x39, 0xb7, 0x26, 0xf9, 0x10, 0xcc,
	0x7d, 0xcc, 0x0e, 0x52, 0x39, 0x83, 0xca, 0xd1, 0x58, 0xee, 0x13, 0x98, 0xde, 0x09, 0x43, 0x2c,
	0x18, 0x7d, 0xab, 0xdf, 0xa1, 0xb2, 0x7b, 0x64, 0x06, 0x93, 0x37, 0xe5, 0xdf, 0x6f, 0x5f, 0xbf,
	0xb5, 0x0d, 0xf7, 0x0f, 0x03, 0x66, 0x37, 0x2c, 0xbb, 0xed, 0x52, 0x75, 0x61, 0x16, 0xf2, 0x2c,
	0x90, 0x07, 0x54, 0x6c, 0x17, 0xd7, 0x29, 0x4f, 0xe8, 0xc9, 0x19, 0x79, 0x09, 0x63, 0x55, 0x07,
	0xaf, 0xd2, 0x99, 0xae, 0x1f, 0xff, 0x77, 0x45, 0xb4, 0xc5, 0xc9, 0x25, 0x8c, 0xcb, 0x5f, 0x99,
	0x6b, 0x67, 0x58, 0x29, 0x1f, 0x78, 0xf5, 0x28, 0x7a, 0xed, 0x28, 0x7a, 0xaf, 0x9b, 0x51, 0xa4,
	0x2d, 0x49, 0x9e, 0xc3, 0x58, 0xa1, 0x56, 0x1c, 0x33, 0x67, 0x54, 0x89, 0x1e, 0xfe, 0xed, 0xb6,
	0xb5, 0x2a, 0xb6, 0x5a, 0x31, 0x8d, 0x51, 0x41, 0x5b, 0xd8, 0xfd, 0xbd, 0x5f, 0xd7, 0x76, 0x83,
	0x49, 0x1a, 0x33, 0x8d, 0x64, 0x0e, 0x7d, 0x1e, 0x56, 0x15, 0x59, 0xb4, 0xcf, 0x43, 0xf2, 0x02,
	0x26, 0x41, 0xa9, 0x91, 0xaa, 0xa8, 0x0a, 0x99, 0xaf, 0x3f, 0x3a, 0x8b, 0x5c, 0xca, 0x37, 0x0d,
	0x42, 0x3b, 0x98, 0x90, 0xe6, 0x3e, 0xeb, 0x9e, 0xd7, 0x77, 0xf4, 0x02, 0x26, 0x49, 0x53, 0x6f,
	0x53, 0xdb, 0x3f, 0x05, 0xeb, 0x5a, 0xd2, 0xc1, 0xe4, 0x73, 0xb0, 0xba, 0xed, 0x6d, 0x0a, 0x7c,
	0x74, 0xae, 0x2c, 0x52, 0x0c, 0xbf, 0x6e, 0x21, 0x7a, 0xe4, 0xcb, 0xfb, 0x0f, 0xf2, 0x4c, 0xcb,
	0xc4, 0x31, 0x97, 0xc6, 0x6a, 0x46, 0x1b, 0x8b, 0xbc, 0x04, 0x2b, 0x90, 0x42, 0x33, 0x2e, 0x50,
	0x39, 0xe3, 0x2a, 0xa8, 0x73, 0x16, 0x74, 0xd3, 0xfa, 0xbf, 0xea, 0xd1, 0x23, 0x7c, 0x35, 0x01,
	0x53, 0x33, 0x15, 0xa1, 0x76, 0xff, 0x34, 0xc0, 0xea, 0xa0, 0x72, 0x6b, 0x78, 0xc2, 0x22, 0x6c,
	0xfa, 0x57, 0x1b, 0xe5, 0x64, 0x06, 0x32, 0x49, 0x98, 0x08, 0xab, 0xc5, 0xb6, 0x68, 0x6b, 0x96,
	0x3d, 0x62, 0x2a, 0xca, 0x9c, 0x41, 0x75, 0x5c, 0x7d, 0x93, 0xe7, 0x60, 0xa9, 0x76, 0x33, 0x9b,
	0x26, 0x39, 0xff, 0xb6, 0xb9, 0xf4, 0x88, 0x92, 0x67, 0x30, 0x40, 0x71, 0x70, 0x46, 0xd5, 0xd3,
	0x71, 0xde, 0xd6, 0xb7, 0x58, 0xfc, 0x50, 0x6e, 0xf0, 0x77, 0x8c, 0x2b, 0x5a, 0x72, 0xe4, 0x12,
	0xcc, 0x40, 0x8a, 0x3d, 0x8f, 0x1c, 0xf3, 0x7e, 0x45, 0x83, 0x3e, 0xfd, 0xb4, 0x1e, 0x96, 0xcd,
	0xf1, 0x8e, 0xe7, 0x5b, 0x2e, 0xa2, 0x18, 0xb7, 0x1a, 0xd3, 0xd2, 0x63, 0xf7, 0xc8, 0xfb, 0x70,
	0x71, 0x9d, 0xc7, 0x9a, 0x77, 0x47, 0xc6, 0xd5, 0xfa, 0xc7, 0x4f, 0x22, 0xae, 0xdf, 0xe5, 0x3b,
	0x2f, 0x90, 0x89, 0x1f, 0x17, 0x7b, 0xed, 0x77, 0x4f, 0x6f, 0x84, 0xc2, 0x4f, 0x77, 0xcf, 0x22,
	0xe9, 0x9f, 0xbc, 0xc6, 0x3b, 0xb3, 0x1a, 0xf6, 0xcb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x53, 0x45, 0xae, 0xee, 0x05, 0x00, 0x00,
}
