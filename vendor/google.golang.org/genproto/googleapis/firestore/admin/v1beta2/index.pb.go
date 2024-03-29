// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/index.proto

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

// Query Scope defines the scope at which a query is run. This is specified on
// a StructuredQuery's `from` field.
type Index_QueryScope int32

const (
	// The query scope is unspecified. Not a valid option.
	Index_QUERY_SCOPE_UNSPECIFIED Index_QueryScope = 0
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified
	// at query time, and that has the collection id specified by the index.
	Index_COLLECTION Index_QueryScope = 1
	// Indexes with a collection group query scope specified allow queries
	// against all collections that has the collection id specified by the
	// index.
	Index_COLLECTION_GROUP Index_QueryScope = 2
)

var Index_QueryScope_name = map[int32]string{
	0: "QUERY_SCOPE_UNSPECIFIED",
	1: "COLLECTION",
	2: "COLLECTION_GROUP",
}

var Index_QueryScope_value = map[string]int32{
	"QUERY_SCOPE_UNSPECIFIED": 0,
	"COLLECTION":              1,
	"COLLECTION_GROUP":        2,
}

func (x Index_QueryScope) String() string {
	return proto.EnumName(Index_QueryScope_name, int32(x))
}

func (Index_QueryScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0, 0}
}

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index creation encounters a problem, the index
// will transition to the `NEEDS_REPAIR` state.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	Index_CREATING Index_State = 1
	// The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	Index_READY Index_State = 2
	// The index was being created, but something went wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	// Use the google.longrunning.Operations API to determine why the operation
	// that last attempted to create this index failed, then re-create the
	// index.
	Index_NEEDS_REPAIR Index_State = 3
)

var Index_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
	3: "NEEDS_REPAIR",
}

var Index_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"CREATING":          1,
	"READY":             2,
	"NEEDS_REPAIR":      3,
}

func (x Index_State) String() string {
	return proto.EnumName(Index_State_name, int32(x))
}

func (Index_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0, 1}
}

// The supported orderings.
type Index_IndexField_Order int32

const (
	// The ordering is unspecified. Not a valid option.
	Index_IndexField_ORDER_UNSPECIFIED Index_IndexField_Order = 0
	// The field is ordered by ascending field value.
	Index_IndexField_ASCENDING Index_IndexField_Order = 1
	// The field is ordered by descending field value.
	Index_IndexField_DESCENDING Index_IndexField_Order = 2
)

var Index_IndexField_Order_name = map[int32]string{
	0: "ORDER_UNSPECIFIED",
	1: "ASCENDING",
	2: "DESCENDING",
}

var Index_IndexField_Order_value = map[string]int32{
	"ORDER_UNSPECIFIED": 0,
	"ASCENDING":         1,
	"DESCENDING":        2,
}

func (x Index_IndexField_Order) String() string {
	return proto.EnumName(Index_IndexField_Order_name, int32(x))
}

func (Index_IndexField_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0, 0, 0}
}

// The supported array value configurations.
type Index_IndexField_ArrayConfig int32

const (
	// The index does not support additional array queries.
	Index_IndexField_ARRAY_CONFIG_UNSPECIFIED Index_IndexField_ArrayConfig = 0
	// The index supports array containment queries.
	Index_IndexField_CONTAINS Index_IndexField_ArrayConfig = 1
)

var Index_IndexField_ArrayConfig_name = map[int32]string{
	0: "ARRAY_CONFIG_UNSPECIFIED",
	1: "CONTAINS",
}

var Index_IndexField_ArrayConfig_value = map[string]int32{
	"ARRAY_CONFIG_UNSPECIFIED": 0,
	"CONTAINS":                 1,
}

func (x Index_IndexField_ArrayConfig) String() string {
	return proto.EnumName(Index_IndexField_ArrayConfig_name, int32(x))
}

func (Index_IndexField_ArrayConfig) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0, 0, 1}
}

// Cloud Firestore indexes enable simple and complex queries against
// documents in a database.
type Index struct {
	// Output only. A server defined name for this index.
	// The form of this name for composite indexes will be:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	// For single field indexes, this field will be empty.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified at
	// query time, and that has the same collection id.
	//
	// Indexes with a collection group query scope specified allow queries against
	// all collections descended from a specific document, specified at query
	// time, and that have the same collection id as this index.
	QueryScope Index_QueryScope `protobuf:"varint,2,opt,name=query_scope,json=queryScope,proto3,enum=google.firestore.admin.v1beta2.Index_QueryScope" json:"query_scope,omitempty"`
	// The fields supported by this index.
	//
	// For composite indexes, this is always 2 or more fields.
	// The last field entry is always for the field path `__name__`. If, on
	// creation, `__name__` was not specified as the last field, it will be added
	// automatically with the same direction as that of the last field defined. If
	// the final field in a composite index is not directional, the `__name__`
	// will be ordered ASCENDING (unless explicitly specified).
	//
	// For single field indexes, this will always be exactly one entry with a
	// field path equal to the field path of the associated field.
	Fields []*Index_IndexField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// Output only. The serving state of the index.
	State                Index_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.firestore.admin.v1beta2.Index_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Index) GetQueryScope() Index_QueryScope {
	if m != nil {
		return m.QueryScope
	}
	return Index_QUERY_SCOPE_UNSPECIFIED
}

func (m *Index) GetFields() []*Index_IndexField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Index) GetState() Index_State {
	if m != nil {
		return m.State
	}
	return Index_STATE_UNSPECIFIED
}

// A field in an index.
// The field_path describes which field is indexed, the value_mode describes
// how the field value is indexed.
type Index_IndexField struct {
	// Can be __name__.
	// For single field indexes, this must match the name of the field or may
	// be omitted.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// How the field value is indexed.
	//
	// Types that are valid to be assigned to ValueMode:
	//	*Index_IndexField_Order_
	//	*Index_IndexField_ArrayConfig_
	ValueMode            isIndex_IndexField_ValueMode `protobuf_oneof:"value_mode"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Index_IndexField) Reset()         { *m = Index_IndexField{} }
func (m *Index_IndexField) String() string { return proto.CompactTextString(m) }
func (*Index_IndexField) ProtoMessage()    {}
func (*Index_IndexField) Descriptor() ([]byte, []int) {
	return fileDescriptor_91374f42b54eaaef, []int{0, 0}
}

func (m *Index_IndexField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_IndexField.Unmarshal(m, b)
}
func (m *Index_IndexField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_IndexField.Marshal(b, m, deterministic)
}
func (m *Index_IndexField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_IndexField.Merge(m, src)
}
func (m *Index_IndexField) XXX_Size() int {
	return xxx_messageInfo_Index_IndexField.Size(m)
}
func (m *Index_IndexField) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_IndexField.DiscardUnknown(m)
}

var xxx_messageInfo_Index_IndexField proto.InternalMessageInfo

func (m *Index_IndexField) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

type isIndex_IndexField_ValueMode interface {
	isIndex_IndexField_ValueMode()
}

type Index_IndexField_Order_ struct {
	Order Index_IndexField_Order `protobuf:"varint,2,opt,name=order,proto3,enum=google.firestore.admin.v1beta2.Index_IndexField_Order,oneof"`
}

type Index_IndexField_ArrayConfig_ struct {
	ArrayConfig Index_IndexField_ArrayConfig `protobuf:"varint,3,opt,name=array_config,json=arrayConfig,proto3,enum=google.firestore.admin.v1beta2.Index_IndexField_ArrayConfig,oneof"`
}

func (*Index_IndexField_Order_) isIndex_IndexField_ValueMode() {}

func (*Index_IndexField_ArrayConfig_) isIndex_IndexField_ValueMode() {}

func (m *Index_IndexField) GetValueMode() isIndex_IndexField_ValueMode {
	if m != nil {
		return m.ValueMode
	}
	return nil
}

func (m *Index_IndexField) GetOrder() Index_IndexField_Order {
	if x, ok := m.GetValueMode().(*Index_IndexField_Order_); ok {
		return x.Order
	}
	return Index_IndexField_ORDER_UNSPECIFIED
}

func (m *Index_IndexField) GetArrayConfig() Index_IndexField_ArrayConfig {
	if x, ok := m.GetValueMode().(*Index_IndexField_ArrayConfig_); ok {
		return x.ArrayConfig
	}
	return Index_IndexField_ARRAY_CONFIG_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Index_IndexField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Index_IndexField_Order_)(nil),
		(*Index_IndexField_ArrayConfig_)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_QueryScope", Index_QueryScope_name, Index_QueryScope_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_State", Index_State_name, Index_State_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_IndexField_Order", Index_IndexField_Order_name, Index_IndexField_Order_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.Index_IndexField_ArrayConfig", Index_IndexField_ArrayConfig_name, Index_IndexField_ArrayConfig_value)
	proto.RegisterType((*Index)(nil), "google.firestore.admin.v1beta2.Index")
	proto.RegisterType((*Index_IndexField)(nil), "google.firestore.admin.v1beta2.Index.IndexField")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/index.proto", fileDescriptor_91374f42b54eaaef)
}

var fileDescriptor_91374f42b54eaaef = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x8b, 0x9b, 0x4c,
	0x14, 0xc6, 0x57, 0xb3, 0x2e, 0x6f, 0x4e, 0xf2, 0x2e, 0x76, 0x68, 0xa9, 0x6c, 0xb7, 0x65, 0x91,
	0x5e, 0x84, 0x16, 0xb4, 0xbb, 0x85, 0x42, 0x69, 0x7b, 0x61, 0x74, 0x92, 0x08, 0x8b, 0x9a, 0x31,
	0x29, 0xa4, 0x37, 0x32, 0x1b, 0x27, 0xae, 0x90, 0x38, 0x59, 0x35, 0x4b, 0xf7, 0xae, 0x9f, 0xa5,
	0xd0, 0x9b, 0x7e, 0xb0, 0x7e, 0x8e, 0xe2, 0x28, 0x49, 0x59, 0xfa, 0x67, 0xf7, 0x46, 0xe6, 0xe8,
	0xf3, 0xfc, 0xe6, 0x99, 0xe3, 0x1c, 0x78, 0x91, 0x70, 0x9e, 0x2c, 0x99, 0xb9, 0x48, 0x73, 0x56,
	0x94, 0x3c, 0x67, 0x26, 0x8d, 0x57, 0x69, 0x66, 0x5e, 0x9f, 0x5e, 0xb0, 0x92, 0x9e, 0x99, 0x69,
	0x16, 0xb3, 0xcf, 0xc6, 0x3a, 0xe7, 0x25, 0x47, 0xcf, 0x6a, 0xad, 0xb1, 0xd5, 0x1a, 0x42, 0x6b,
	0x34, 0xda, 0xa3, 0xe3, 0x86, 0x45, 0xd7, 0xa9, 0x49, 0xb3, 0x8c, 0x97, 0xb4, 0x4c, 0x79, 0x56,
	0xd4, 0x6e, 0xfd, 0xcb, 0x01, 0x28, 0x6e, 0x45, 0x43, 0x08, 0xf6, 0x33, 0xba, 0x62, 0x9a, 0x74,
	0x22, 0xf5, 0xda, 0x44, 0xac, 0xd1, 0x18, 0x3a, 0x57, 0x1b, 0x96, 0xdf, 0x44, 0xc5, 0x9c, 0xaf,
	0x99, 0x26, 0x9f, 0x48, 0xbd, 0xc3, 0xb3, 0x57, 0xc6, 0xdf, 0x77, 0x34, 0x04, 0xcf, 0x18, 0x57,
	0xc6, 0xb0, 0xf2, 0x11, 0xb8, 0xda, 0xae, 0xd1, 0x08, 0x0e, 0x16, 0x29, 0x5b, 0xc6, 0x85, 0xd6,
	0x3a, 0x69, 0xf5, 0x3a, 0x77, 0xa5, 0x89, 0xe7, 0xa0, 0x32, 0x92, 0xc6, 0x8f, 0x2c, 0x50, 0x8a,
	0x92, 0x96, 0x4c, 0xdb, 0x17, 0xb1, 0x5e, 0xde, 0x0d, 0x14, 0x56, 0x16, 0x52, 0x3b, 0x8f, 0x7e,
	0xc8, 0x00, 0x3b, 0x32, 0x7a, 0x0a, 0x20, 0xd8, 0xd1, 0x9a, 0x96, 0x97, 0x4d, 0x23, 0xda, 0xe2,
	0x4d, 0x40, 0xcb, 0x4b, 0xe4, 0x81, 0xc2, 0xf3, 0x98, 0xe5, 0x4d, 0x1f, 0xde, 0xdc, 0x37, 0xb9,
	0xe1, 0x57, 0xee, 0xd1, 0x1e, 0xa9, 0x31, 0x88, 0x42, 0x97, 0xe6, 0x39, 0xbd, 0x89, 0xe6, 0x3c,
	0x5b, 0xa4, 0x89, 0xd6, 0x12, 0xd8, 0xf7, 0xf7, 0xc6, 0x5a, 0x15, 0xc4, 0x16, 0x8c, 0xd1, 0x1e,
	0xe9, 0xd0, 0x5d, 0xa9, 0x7f, 0x00, 0x45, 0x6c, 0x8a, 0x1e, 0xc1, 0x03, 0x9f, 0x38, 0x98, 0x44,
	0x53, 0x2f, 0x0c, 0xb0, 0xed, 0x0e, 0x5c, 0xec, 0xa8, 0x7b, 0xe8, 0x7f, 0x68, 0x5b, 0xa1, 0x8d,
	0x3d, 0xc7, 0xf5, 0x86, 0xaa, 0x84, 0x0e, 0x01, 0x1c, 0xbc, 0xad, 0x65, 0xfd, 0x2d, 0x74, 0x7e,
	0x81, 0xa3, 0x63, 0xd0, 0x2c, 0x42, 0xac, 0x59, 0x64, 0xfb, 0xde, 0xc0, 0x1d, 0xde, 0x62, 0x75,
	0xe1, 0x3f, 0xdb, 0xf7, 0x26, 0x96, 0xeb, 0x85, 0xaa, 0xd4, 0xef, 0x02, 0x5c, 0xd3, 0xe5, 0x86,
	0x45, 0x2b, 0x1e, 0x33, 0xdd, 0x07, 0xd8, 0xdd, 0x07, 0xf4, 0x04, 0x1e, 0x8f, 0xa7, 0x98, 0xcc,
	0xa2, 0xd0, 0xf6, 0x03, 0x7c, 0x0b, 0x73, 0x08, 0x60, 0xfb, 0xe7, 0xe7, 0xd8, 0x9e, 0xb8, 0xbe,
	0xa7, 0x4a, 0xe8, 0x21, 0xa8, 0xbb, 0x3a, 0x1a, 0x12, 0x7f, 0x1a, 0xa8, 0xb2, 0xee, 0x82, 0x22,
	0xfe, 0x64, 0x75, 0xb0, 0x70, 0x62, 0x4d, 0xf0, 0x6f, 0xc2, 0x10, 0x6c, 0x4d, 0xea, 0x73, 0xb5,
	0x41, 0x21, 0xd8, 0x72, 0x66, 0xaa, 0x8c, 0x54, 0xe8, 0x7a, 0x18, 0x3b, 0x61, 0x44, 0x70, 0x60,
	0xb9, 0x44, 0x6d, 0xf5, 0xbf, 0x49, 0xa0, 0xcf, 0xf9, 0xea, 0x1f, 0x6d, 0xef, 0xd7, 0x17, 0x25,
	0xa8, 0xa6, 0x26, 0x90, 0x3e, 0xd9, 0x8d, 0x3a, 0xe1, 0x4b, 0x9a, 0x25, 0x06, 0xcf, 0x13, 0x33,
	0x61, 0x99, 0x98, 0x29, 0xb3, 0xfe, 0x44, 0xd7, 0x69, 0xf1, 0xa7, 0x01, 0x7e, 0x27, 0xaa, 0xaf,
	0xf2, 0xfe, 0xd0, 0x1e, 0x84, 0xdf, 0xe5, 0xe7, 0xc3, 0x1a, 0x66, 0x2f, 0xf9, 0x26, 0x36, 0x06,
	0xdb, 0x00, 0x96, 0x08, 0xf0, 0xf1, 0xb4, 0x5f, 0x79, 0x2e, 0x0e, 0x04, 0xfd, 0xf5, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x38, 0x79, 0x14, 0x1d, 0x04, 0x00, 0x00,
}
