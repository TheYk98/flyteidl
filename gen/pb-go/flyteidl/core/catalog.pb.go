// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/catalog.proto

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

// Indicates the status of CatalogCaching. The reason why this is not embedded in TaskNodeMetadata is, that we may use for other types of nodes as well in the future
type CatalogCacheStatus int32

const (
	// Used to indicate that caching was disabled
	CatalogCacheStatus_CACHE_DISABLED CatalogCacheStatus = 0
	// Used to indicate that the cache lookup resulted in no matches
	CatalogCacheStatus_CACHE_MISS CatalogCacheStatus = 1
	// used to indicate that the associated artifact was a result of a previous execution
	CatalogCacheStatus_CACHE_HIT CatalogCacheStatus = 2
	// used to indicate that the resultant artifact was added to the cache
	CatalogCacheStatus_CACHE_POPULATED CatalogCacheStatus = 3
	// Used to indicate that cache lookup failed because of an error
	CatalogCacheStatus_CACHE_LOOKUP_FAILURE CatalogCacheStatus = 4
	// Used to indicate that cache lookup failed because of an error
	CatalogCacheStatus_CACHE_PUT_FAILURE CatalogCacheStatus = 5
)

var CatalogCacheStatus_name = map[int32]string{
	0: "CACHE_DISABLED",
	1: "CACHE_MISS",
	2: "CACHE_HIT",
	3: "CACHE_POPULATED",
	4: "CACHE_LOOKUP_FAILURE",
	5: "CACHE_PUT_FAILURE",
}

var CatalogCacheStatus_value = map[string]int32{
	"CACHE_DISABLED":       0,
	"CACHE_MISS":           1,
	"CACHE_HIT":            2,
	"CACHE_POPULATED":      3,
	"CACHE_LOOKUP_FAILURE": 4,
	"CACHE_PUT_FAILURE":    5,
}

func (x CatalogCacheStatus) String() string {
	return proto.EnumName(CatalogCacheStatus_name, int32(x))
}

func (CatalogCacheStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_534f5d1443565574, []int{0}
}

type CatalogArtifactTag struct {
	// Artifact ID is generated name
	ArtifactId string `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	// Flyte computes the tag automatically, as the hash of the values
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatalogArtifactTag) Reset()         { *m = CatalogArtifactTag{} }
func (m *CatalogArtifactTag) String() string { return proto.CompactTextString(m) }
func (*CatalogArtifactTag) ProtoMessage()    {}
func (*CatalogArtifactTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_534f5d1443565574, []int{0}
}

func (m *CatalogArtifactTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogArtifactTag.Unmarshal(m, b)
}
func (m *CatalogArtifactTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogArtifactTag.Marshal(b, m, deterministic)
}
func (m *CatalogArtifactTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogArtifactTag.Merge(m, src)
}
func (m *CatalogArtifactTag) XXX_Size() int {
	return xxx_messageInfo_CatalogArtifactTag.Size(m)
}
func (m *CatalogArtifactTag) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogArtifactTag.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogArtifactTag proto.InternalMessageInfo

func (m *CatalogArtifactTag) GetArtifactId() string {
	if m != nil {
		return m.ArtifactId
	}
	return ""
}

func (m *CatalogArtifactTag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Catalog artifact information with specific metadata
type CatalogMetadata struct {
	// Dataset ID in the catalog
	DatasetId *Identifier `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Artifact tag in the catalog
	ArtifactTag *CatalogArtifactTag `protobuf:"bytes,2,opt,name=artifact_tag,json=artifactTag,proto3" json:"artifact_tag,omitempty"`
	// Optional: Source Execution identifier, if this dataset was generated by another execution in Flyte. This is a one-of field and will depend on the caching context
	//
	// Types that are valid to be assigned to SourceExecution:
	//	*CatalogMetadata_SourceTaskExecution
	SourceExecution      isCatalogMetadata_SourceExecution `protobuf_oneof:"source_execution"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CatalogMetadata) Reset()         { *m = CatalogMetadata{} }
func (m *CatalogMetadata) String() string { return proto.CompactTextString(m) }
func (*CatalogMetadata) ProtoMessage()    {}
func (*CatalogMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_534f5d1443565574, []int{1}
}

func (m *CatalogMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogMetadata.Unmarshal(m, b)
}
func (m *CatalogMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogMetadata.Marshal(b, m, deterministic)
}
func (m *CatalogMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogMetadata.Merge(m, src)
}
func (m *CatalogMetadata) XXX_Size() int {
	return xxx_messageInfo_CatalogMetadata.Size(m)
}
func (m *CatalogMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogMetadata proto.InternalMessageInfo

func (m *CatalogMetadata) GetDatasetId() *Identifier {
	if m != nil {
		return m.DatasetId
	}
	return nil
}

func (m *CatalogMetadata) GetArtifactTag() *CatalogArtifactTag {
	if m != nil {
		return m.ArtifactTag
	}
	return nil
}

type isCatalogMetadata_SourceExecution interface {
	isCatalogMetadata_SourceExecution()
}

type CatalogMetadata_SourceTaskExecution struct {
	SourceTaskExecution *TaskExecutionIdentifier `protobuf:"bytes,3,opt,name=source_task_execution,json=sourceTaskExecution,proto3,oneof"`
}

func (*CatalogMetadata_SourceTaskExecution) isCatalogMetadata_SourceExecution() {}

func (m *CatalogMetadata) GetSourceExecution() isCatalogMetadata_SourceExecution {
	if m != nil {
		return m.SourceExecution
	}
	return nil
}

func (m *CatalogMetadata) GetSourceTaskExecution() *TaskExecutionIdentifier {
	if x, ok := m.GetSourceExecution().(*CatalogMetadata_SourceTaskExecution); ok {
		return x.SourceTaskExecution
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CatalogMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CatalogMetadata_SourceTaskExecution)(nil),
	}
}

func init() {
	proto.RegisterEnum("flyteidl.core.CatalogCacheStatus", CatalogCacheStatus_name, CatalogCacheStatus_value)
	proto.RegisterType((*CatalogArtifactTag)(nil), "flyteidl.core.CatalogArtifactTag")
	proto.RegisterType((*CatalogMetadata)(nil), "flyteidl.core.CatalogMetadata")
}

func init() { proto.RegisterFile("flyteidl/core/catalog.proto", fileDescriptor_534f5d1443565574) }

var fileDescriptor_534f5d1443565574 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0xeb, 0xda, 0x40,
	0x10, 0xc5, 0x8d, 0xda, 0x82, 0x63, 0xd5, 0x74, 0xad, 0x90, 0xb6, 0xd0, 0x1f, 0x1e, 0x4a, 0x29,
	0x34, 0x01, 0x5b, 0x4a, 0xaf, 0x31, 0x49, 0x71, 0x69, 0x44, 0xc9, 0x8f, 0x4b, 0x29, 0x84, 0x35,
	0x59, 0xe3, 0xa2, 0x66, 0x25, 0xd9, 0x40, 0x7b, 0xee, 0xb9, 0xff, 0x73, 0x31, 0x89, 0x5b, 0xe2,
	0xf7, 0x94, 0xcc, 0x9b, 0x37, 0x9f, 0xbc, 0x09, 0x03, 0x2f, 0xf7, 0xa7, 0xdf, 0x82, 0xb2, 0xe4,
	0x64, 0xc4, 0x3c, 0xa7, 0x46, 0x4c, 0x04, 0x39, 0xf1, 0x54, 0xbf, 0xe4, 0x5c, 0x70, 0x34, 0xba,
	0x35, 0xf5, 0x6b, 0xf3, 0xc5, 0xab, 0xb6, 0x97, 0x25, 0x34, 0x13, 0x6c, 0xcf, 0x68, 0x5e, 0xdb,
	0xe7, 0x18, 0x90, 0x55, 0xcf, 0x9b, 0xb9, 0x60, 0x7b, 0x12, 0x8b, 0x80, 0xa4, 0xe8, 0x35, 0x0c,
	0x49, 0x53, 0x46, 0x2c, 0xd1, 0x94, 0x37, 0xca, 0xfb, 0x81, 0x07, 0x37, 0x09, 0x27, 0x08, 0x41,
	0x3f, 0x23, 0x67, 0xaa, 0x75, 0xab, 0x4e, 0xf5, 0x3e, 0xff, 0xd3, 0x85, 0x49, 0xc3, 0x5a, 0x53,
	0x41, 0x12, 0x22, 0x08, 0xfa, 0x0a, 0x70, 0x7d, 0x16, 0x54, 0x72, 0x86, 0x8b, 0xe7, 0x7a, 0x2b,
	0xa2, 0x8e, 0x65, 0x26, 0x6f, 0xd0, 0x98, 0x71, 0x82, 0x6c, 0x78, 0x22, 0x23, 0x08, 0x92, 0x56,
	0x5f, 0x1a, 0x2e, 0xde, 0xde, 0xcd, 0x3e, 0xcc, 0xee, 0xc9, 0xe4, 0xd7, 0x45, 0x7e, 0xc2, 0xac,
	0xe0, 0x65, 0x1e, 0xd3, 0x48, 0x90, 0xe2, 0x18, 0xd1, 0x5f, 0x34, 0x2e, 0x05, 0xe3, 0x99, 0xd6,
	0xab, 0x70, 0xef, 0xee, 0x70, 0x01, 0x29, 0x8e, 0xce, 0xcd, 0xf3, 0x3f, 0xd7, 0xaa, 0xe3, 0x4d,
	0x6b, 0x4c, 0xcb, 0xb0, 0x44, 0xa0, 0x36, 0x74, 0x09, 0xfe, 0xf0, 0x57, 0x91, 0x7f, 0xd4, 0x22,
	0xf1, 0x81, 0xfa, 0x82, 0x88, 0xb2, 0x40, 0x08, 0xc6, 0x96, 0x69, 0xad, 0x9c, 0xc8, 0xc6, 0xbe,
	0xb9, 0x74, 0x1d, 0x5b, 0xed, 0xa0, 0x31, 0x40, 0xad, 0xad, 0xb1, 0xef, 0xab, 0x0a, 0x1a, 0xc1,
	0xa0, 0xae, 0x57, 0x38, 0x50, 0xbb, 0x68, 0x0a, 0x93, 0xba, 0xdc, 0x6e, 0xb6, 0xa1, 0x6b, 0x06,
	0x8e, 0xad, 0xf6, 0x90, 0x06, 0xcf, 0x6a, 0xd1, 0xdd, 0x6c, 0xbe, 0x87, 0xdb, 0xe8, 0x9b, 0x89,
	0xdd, 0xd0, 0x73, 0xd4, 0x3e, 0x9a, 0xc1, 0xd3, 0xc6, 0x1e, 0x06, 0x52, 0x7e, 0xb4, 0xfc, 0xf2,
	0xe3, 0x73, 0xca, 0xc4, 0xa1, 0xdc, 0xe9, 0x31, 0x3f, 0x1b, 0xd5, 0xba, 0x3c, 0x4f, 0x0d, 0x79,
	0x16, 0x29, 0xcd, 0x8c, 0xcb, 0xee, 0x63, 0xca, 0x8d, 0xd6, 0xa5, 0xec, 0x1e, 0x57, 0xf7, 0xf1,
	0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0xfe, 0xcc, 0x9d, 0x6d, 0x02, 0x00, 0x00,
}
