// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1beta1/source.proto

package expr

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Source information collected at parse time.
type SourceInfo struct {
	// The location name. All position information attached to an expression is
	// relative to this location.
	//
	// The location could be a file, UI element, or similar. For example,
	// `acme/app/AnvilPolicy.cel`.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// Monotonically increasing list of character offsets where newlines appear.
	//
	// The line number of a given position is the index `i` where for a given
	// `id` the `line_offsets[i] < id_positions[id] < line_offsets[i+1]`. The
	// column may be derivd from `id_positions[id] - line_offsets[i]`.
	LineOffsets []int32 `protobuf:"varint,3,rep,packed,name=line_offsets,json=lineOffsets,proto3" json:"line_offsets,omitempty"`
	// A map from the parse node id (e.g. `Expr.id`) to the character offset
	// within source.
	Positions            map[int32]int32 `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SourceInfo) Reset()         { *m = SourceInfo{} }
func (m *SourceInfo) String() string { return proto.CompactTextString(m) }
func (*SourceInfo) ProtoMessage()    {}
func (*SourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_26d280d77790e409, []int{0}
}

func (m *SourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceInfo.Unmarshal(m, b)
}
func (m *SourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceInfo.Marshal(b, m, deterministic)
}
func (m *SourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceInfo.Merge(m, src)
}
func (m *SourceInfo) XXX_Size() int {
	return xxx_messageInfo_SourceInfo.Size(m)
}
func (m *SourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SourceInfo proto.InternalMessageInfo

func (m *SourceInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SourceInfo) GetLineOffsets() []int32 {
	if m != nil {
		return m.LineOffsets
	}
	return nil
}

func (m *SourceInfo) GetPositions() map[int32]int32 {
	if m != nil {
		return m.Positions
	}
	return nil
}

// A specific position in source.
type SourcePosition struct {
	// The soucre location name (e.g. file name).
	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The character offset.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// The 1-based index of the starting line in the source text
	// where the issue occurs, or 0 if unknown.
	Line int32 `protobuf:"varint,3,opt,name=line,proto3" json:"line,omitempty"`
	// The 0-based index of the starting position within the line of source text
	// where the issue occurs.  Only meaningful if line is nonzer..
	Column               int32    `protobuf:"varint,4,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourcePosition) Reset()         { *m = SourcePosition{} }
func (m *SourcePosition) String() string { return proto.CompactTextString(m) }
func (*SourcePosition) ProtoMessage()    {}
func (*SourcePosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_26d280d77790e409, []int{1}
}

func (m *SourcePosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourcePosition.Unmarshal(m, b)
}
func (m *SourcePosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourcePosition.Marshal(b, m, deterministic)
}
func (m *SourcePosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourcePosition.Merge(m, src)
}
func (m *SourcePosition) XXX_Size() int {
	return xxx_messageInfo_SourcePosition.Size(m)
}
func (m *SourcePosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SourcePosition.DiscardUnknown(m)
}

var xxx_messageInfo_SourcePosition proto.InternalMessageInfo

func (m *SourcePosition) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SourcePosition) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SourcePosition) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *SourcePosition) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

func init() {
	proto.RegisterType((*SourceInfo)(nil), "google.api.expr.v1beta1.SourceInfo")
	proto.RegisterMapType((map[int32]int32)(nil), "google.api.expr.v1beta1.SourceInfo.PositionsEntry")
	proto.RegisterType((*SourcePosition)(nil), "google.api.expr.v1beta1.SourcePosition")
}

func init() {
	proto.RegisterFile("google/api/expr/v1beta1/source.proto", fileDescriptor_26d280d77790e409)
}

var fileDescriptor_26d280d77790e409 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x49, 0xb7, 0x5b, 0xfe, 0x9d, 0xfe, 0x29, 0x12, 0x44, 0x97, 0x7a, 0x59, 0x8b, 0x87,
	0x3d, 0x65, 0x69, 0xbd, 0x88, 0xf5, 0x54, 0xf0, 0xe0, 0xc9, 0xb2, 0xde, 0xbc, 0x48, 0xba, 0xa4,
	0x4b, 0x30, 0xcd, 0x84, 0xcd, 0xb6, 0xd8, 0xcf, 0xea, 0x17, 0xf1, 0x28, 0x49, 0xb6, 0x96, 0x2a,
	0xbd, 0xe5, 0xcd, 0xfc, 0x32, 0x33, 0x8f, 0x07, 0x37, 0x15, 0x62, 0xa5, 0x44, 0xce, 0x8d, 0xcc,
	0xc5, 0x87, 0xa9, 0xf3, 0xed, 0x64, 0x29, 0x1a, 0x3e, 0xc9, 0x2d, 0x6e, 0xea, 0x52, 0x30, 0x53,
	0x63, 0x83, 0xf4, 0x32, 0x50, 0x8c, 0x1b, 0xc9, 0x1c, 0xc5, 0x5a, 0x6a, 0xfc, 0x49, 0x00, 0x5e,
	0x3c, 0xf9, 0xa4, 0x57, 0x48, 0x47, 0xf0, 0x4f, 0x61, 0xc9, 0x1b, 0x89, 0x3a, 0xe9, 0xa4, 0x24,
	0xeb, 0x17, 0x3f, 0x9a, 0x5e, 0xc3, 0x7f, 0x25, 0xb5, 0x78, 0xc3, 0xd5, 0xca, 0x8a, 0xc6, 0x26,
	0x51, 0x1a, 0x65, 0x71, 0x31, 0x70, 0xb5, 0xe7, 0x50, 0xa2, 0x0b, 0xe8, 0x1b, 0xb4, 0xd2, 0xe1,
	0x36, 0xe9, 0xa6, 0x51, 0x36, 0x98, 0x4e, 0xd9, 0x89, 0xd5, 0xec, 0xb0, 0x96, 0x2d, 0xf6, 0x9f,
	0x1e, 0x75, 0x53, 0xef, 0x8a, 0xc3, 0x90, 0xd1, 0x03, 0x0c, 0x8f, 0x9b, 0xf4, 0x0c, 0xa2, 0x77,
	0xb1, 0x4b, 0x48, 0x4a, 0xb2, 0xb8, 0x70, 0x4f, 0x7a, 0x0e, 0xf1, 0x96, 0xab, 0x8d, 0xf0, 0x17,
	0xc7, 0x45, 0x10, 0xf7, 0x9d, 0x3b, 0x32, 0x36, 0x30, 0x0c, 0x5b, 0xf6, 0x33, 0x8e, 0x0c, 0x92,
	0x5f, 0x06, 0x2f, 0xa0, 0x17, 0xbc, 0xb5, 0x83, 0x5a, 0x45, 0x29, 0x74, 0x9d, 0xc9, 0x24, 0xf2,
	0x55, 0xff, 0x76, 0x6c, 0x89, 0x6a, 0xb3, 0xd6, 0x49, 0x37, 0xb0, 0x41, 0xcd, 0x15, 0x5c, 0x95,
	0xb8, 0x3e, 0xe5, 0x79, 0x3e, 0x68, 0xcf, 0x71, 0xa1, 0x2c, 0xc8, 0xeb, 0xac, 0xe5, 0x2a, 0x54,
	0x5c, 0x57, 0x0c, 0xeb, 0x2a, 0xaf, 0x84, 0xf6, 0x91, 0xe5, 0xa1, 0xc5, 0x8d, 0xb4, 0x7f, 0xb2,
	0x9d, 0x39, 0xf1, 0x45, 0xc8, 0xb2, 0xe7, 0xd1, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x8a, 0x1a, 0x6c, 0x05, 0x02, 0x00, 0x00,
}
