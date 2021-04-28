// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1/field.proto

package admin

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents a single field in the database.
//
// Fields are grouped by their "Collection Group", which represent all
// collections in the database with the same id.
type Field struct {
	// A field name of the form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`, which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that needs
	// to be escaped within a quoted field path is the backtick character itself,
	// escaped using a backslash. Special characters in field paths that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//  additional layer of backticks to represent a code block)
	// `\`address.city\`` represents a field named `address.city`, not the map key
	// `city` in the field `address`.
	// `\`*\`` represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all fields.
	// This field's resource name is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which do not
	// have their own `Field` index configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index configuration for this field. If unset, field indexing will
	// revert to the configuration defined by the `ancestor_field`. To
	// explicitly remove all indexes for this field, specify an index config
	// with an empty list of indexes.
	IndexConfig          *Field_IndexConfig `protobuf:"bytes,2,opt,name=index_config,json=indexConfig,proto3" json:"index_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50c2ee90e4658e7, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetIndexConfig() *Field_IndexConfig {
	if m != nil {
		return m.IndexConfig
	}
	return nil
}

// The index configuration for this field.
type Field_IndexConfig struct {
	// The indexes supported for this field.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	// Output only. When true, the `Field`'s index configuration is set from the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `protobuf:"varint,2,opt,name=uses_ancestor_config,json=usesAncestorConfig,proto3" json:"uses_ancestor_config,omitempty"`
	// Output only. Specifies the resource name of the `Field` from which this field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `protobuf:"bytes,3,opt,name=ancestor_field,json=ancestorField,proto3" json:"ancestor_field,omitempty"`
	// Output only
	// When true, the `Field`'s index configuration is in the process of being
	// reverted. Once complete, the index config will transition to the same
	// state as the field specified by `ancestor_field`, at which point
	// `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting            bool     `protobuf:"varint,4,opt,name=reverting,proto3" json:"reverting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field_IndexConfig) Reset()         { *m = Field_IndexConfig{} }
func (m *Field_IndexConfig) String() string { return proto.CompactTextString(m) }
func (*Field_IndexConfig) ProtoMessage()    {}
func (*Field_IndexConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50c2ee90e4658e7, []int{0, 0}
}

func (m *Field_IndexConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field_IndexConfig.Unmarshal(m, b)
}
func (m *Field_IndexConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field_IndexConfig.Marshal(b, m, deterministic)
}
func (m *Field_IndexConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field_IndexConfig.Merge(m, src)
}
func (m *Field_IndexConfig) XXX_Size() int {
	return xxx_messageInfo_Field_IndexConfig.Size(m)
}
func (m *Field_IndexConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Field_IndexConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Field_IndexConfig proto.InternalMessageInfo

func (m *Field_IndexConfig) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *Field_IndexConfig) GetUsesAncestorConfig() bool {
	if m != nil {
		return m.UsesAncestorConfig
	}
	return false
}

func (m *Field_IndexConfig) GetAncestorField() string {
	if m != nil {
		return m.AncestorField
	}
	return ""
}

func (m *Field_IndexConfig) GetReverting() bool {
	if m != nil {
		return m.Reverting
	}
	return false
}

func init() {
	proto.RegisterType((*Field)(nil), "google.firestore.admin.v1.Field")
	proto.RegisterType((*Field_IndexConfig)(nil), "google.firestore.admin.v1.Field.IndexConfig")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1/field.proto", fileDescriptor_a50c2ee90e4658e7)
}

var fileDescriptor_a50c2ee90e4658e7 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0x99, 0x24, 0xfe, 0xd9, 0x8e, 0x7a, 0x18, 0x3c, 0x8c, 0x61, 0xd5, 0x20, 0x2c, 0xe4,
	0x20, 0xdd, 0x66, 0xbd, 0xad, 0x20, 0x64, 0x03, 0x09, 0x9e, 0x5c, 0xa2, 0xec, 0x41, 0x02, 0x4b,
	0x6f, 0x4f, 0xa5, 0x69, 0x99, 0x74, 0x0d, 0xdd, 0x93, 0xa0, 0x84, 0xbc, 0x8c, 0xc7, 0x7d, 0x02,
	0x8f, 0x9e, 0x7d, 0x14, 0x9f, 0x42, 0xa6, 0x7a, 0xfe, 0xe4, 0xb0, 0xbb, 0xa7, 0x54, 0xfa, 0xfb,
	0xd5, 0x57, 0xdf, 0x74, 0x35, 0x3b, 0xd1, 0x88, 0x3a, 0x03, 0xb1, 0x32, 0x0e, 0x7c, 0x81, 0x0e,
	0x84, 0x4c, 0xd7, 0xc6, 0x8a, 0xed, 0x58, 0xac, 0x0c, 0x64, 0x29, 0xcf, 0x1d, 0x16, 0x18, 0xbf,
	0x08, 0x18, 0x6f, 0x30, 0x4e, 0x18, 0xdf, 0x8e, 0x07, 0x95, 0x24, 0x64, 0x6e, 0x84, 0x03, 0x8f,
	0x1b, 0xa7, 0x20, 0x74, 0x0d, 0xee, 0x31, 0x37, 0x36, 0x85, 0x1f, 0x15, 0x76, 0x7c, 0xe0, 0x20,
	0xad, 0xc5, 0x42, 0x16, 0x06, 0xad, 0x0f, 0xea, 0x9b, 0x9b, 0x2e, 0x7b, 0x30, 0x2b, 0xa3, 0xc4,
	0x31, 0xeb, 0x59, 0xb9, 0x86, 0x24, 0x1a, 0x46, 0xa3, 0xa3, 0x05, 0xd5, 0xf1, 0x67, 0xf6, 0x84,
	0xac, 0xae, 0x14, 0xda, 0x95, 0xd1, 0x49, 0x67, 0x18, 0x8d, 0xfa, 0xa7, 0x6f, 0xf9, 0x9d, 0x79,
	0x39, 0x79, 0xf1, 0x4f, 0x65, 0xd3, 0x94, 0x7a, 0x16, 0x7d, 0xd3, 0xfe, 0x19, 0xfc, 0x89, 0x58,
	0xff, 0x40, 0x8c, 0xcf, 0xd8, 0x23, 0x92, 0xc1, 0x27, 0xd1, 0xb0, 0x3b, 0xea, 0x9f, 0x0e, 0xef,
	0xf1, 0xa6, 0xc6, 0x45, 0xdd, 0x10, 0xbf, 0x63, 0xcf, 0x37, 0x1e, 0xfc, 0x95, 0xb4, 0x8a, 0xc8,
	0xc3, 0x90, 0x8f, 0x17, 0x71, 0xa9, 0x4d, 0x2a, 0xa9, 0x9a, 0x76, 0xc2, 0x9e, 0x35, 0x30, 0xdd,
	0x7f, 0xd2, 0xa5, 0x8f, 0x7d, 0x5a, 0x9f, 0x86, 0x9b, 0x38, 0x66, 0x47, 0x0e, 0xb6, 0xe0, 0x0a,
	0x63, 0x75, 0xd2, 0x23, 0xb7, 0xf6, 0xe0, 0xec, 0xe7, 0xbf, 0xc9, 0x96, 0xbd, 0x6a, 0xe3, 0x85,
	0xbc, 0x32, 0x37, 0x9e, 0x2b, 0x5c, 0x8b, 0x60, 0xf1, 0x35, 0x77, 0xf8, 0x1d, 0x54, 0xe1, 0xc5,
	0xae, 0xaa, 0xf6, 0x22, 0x95, 0x85, 0xbc, 0x96, 0x1e, 0xbc, 0xd8, 0xd5, 0xe5, 0x5e, 0x28, 0xcc,
	0x32, 0x50, 0xe5, 0x46, 0xe6, 0x0e, 0x37, 0xb9, 0x17, 0xbb, 0xf6, 0x64, 0x1f, 0x1e, 0x89, 0x17,
	0x3b, 0xfa, 0xdd, 0x9f, 0xff, 0x8e, 0xd8, 0x4b, 0x85, 0xeb, 0xbb, 0xaf, 0xe8, 0x9c, 0xd1, 0xf8,
	0x8b, 0x72, 0xb5, 0x17, 0xd1, 0xb7, 0x8f, 0x15, 0xa8, 0x31, 0x93, 0x56, 0x73, 0x74, 0x5a, 0x68,
	0xb0, 0xb4, 0x78, 0xd1, 0xc6, 0xbe, 0xe5, 0x01, 0x7d, 0xa0, 0xe2, 0x57, 0xa7, 0x37, 0x9f, 0xce,
	0xbe, 0xdc, 0x74, 0x5e, 0xcf, 0x83, 0xcf, 0x34, 0xc3, 0x4d, 0xca, 0x67, 0xcd, 0xd8, 0x09, 0x8d,
	0xbd, 0x1c, 0xff, 0xad, 0x89, 0x25, 0x11, 0xcb, 0x86, 0x58, 0x12, 0xb1, 0xbc, 0x1c, 0x5f, 0x3f,
	0xa4, 0xa9, 0xef, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x95, 0x61, 0x71, 0xb2, 0x12, 0x03, 0x00,
	0x00,
}
