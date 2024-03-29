// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1/geometry.proto

package vision

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

// A vertex represents a 2D point in the image.
// NOTE: the vertex coordinates are in the same scale as the original image.
type Vertex struct {
	// X coordinate.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d1cf3c6845cfe, []int{0}
}

func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A vertex represents a 2D point in the image.
// NOTE: the normalized vertex coordinates are relative to the original image
// and range from 0 to 1.
type NormalizedVertex struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalizedVertex) Reset()         { *m = NormalizedVertex{} }
func (m *NormalizedVertex) String() string { return proto.CompactTextString(m) }
func (*NormalizedVertex) ProtoMessage()    {}
func (*NormalizedVertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d1cf3c6845cfe, []int{1}
}

func (m *NormalizedVertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedVertex.Unmarshal(m, b)
}
func (m *NormalizedVertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedVertex.Marshal(b, m, deterministic)
}
func (m *NormalizedVertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedVertex.Merge(m, src)
}
func (m *NormalizedVertex) XXX_Size() int {
	return xxx_messageInfo_NormalizedVertex.Size(m)
}
func (m *NormalizedVertex) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedVertex.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedVertex proto.InternalMessageInfo

func (m *NormalizedVertex) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NormalizedVertex) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon for the detected image annotation.
type BoundingPoly struct {
	// The bounding polygon vertices.
	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	// The bounding polygon normalized vertices.
	NormalizedVertices   []*NormalizedVertex `protobuf:"bytes,2,rep,name=normalized_vertices,json=normalizedVertices,proto3" json:"normalized_vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BoundingPoly) Reset()         { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()    {}
func (*BoundingPoly) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d1cf3c6845cfe, []int{2}
}

func (m *BoundingPoly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPoly.Unmarshal(m, b)
}
func (m *BoundingPoly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPoly.Marshal(b, m, deterministic)
}
func (m *BoundingPoly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPoly.Merge(m, src)
}
func (m *BoundingPoly) XXX_Size() int {
	return xxx_messageInfo_BoundingPoly.Size(m)
}
func (m *BoundingPoly) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPoly.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPoly proto.InternalMessageInfo

func (m *BoundingPoly) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *BoundingPoly) GetNormalizedVertices() []*NormalizedVertex {
	if m != nil {
		return m.NormalizedVertices
	}
	return nil
}

// A 3D position in the image, used primarily for Face detection landmarks.
// A valid Position must have both x and y coordinates.
// The position coordinates are in the same scale as the original image.
type Position struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	// Z coordinate (or depth).
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1d1cf3c6845cfe, []int{3}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vertex)(nil), "google.cloud.vision.v1.Vertex")
	proto.RegisterType((*NormalizedVertex)(nil), "google.cloud.vision.v1.NormalizedVertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.vision.v1.BoundingPoly")
	proto.RegisterType((*Position)(nil), "google.cloud.vision.v1.Position")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1/geometry.proto", fileDescriptor_ce1d1cf3c6845cfe)
}

var fileDescriptor_ce1d1cf3c6845cfe = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x49, 0xa7, 0x63, 0xc4, 0x09, 0x52, 0x41, 0xca, 0x10, 0x19, 0x43, 0xa1, 0xa7, 0x84,
	0xa9, 0x27, 0xf5, 0x54, 0x0f, 0xbb, 0x8d, 0xd2, 0xc3, 0x40, 0x2f, 0x12, 0xdb, 0x10, 0x02, 0x6d,
	0x5e, 0x49, 0xd3, 0xb2, 0xf6, 0xcf, 0x11, 0xfc, 0x1f, 0x3d, 0x4a, 0x9b, 0x52, 0x58, 0x71, 0xde,
	0xf2, 0x25, 0xbf, 0xf7, 0xbd, 0x97, 0xef, 0xe1, 0x3b, 0x01, 0x20, 0x52, 0x4e, 0xe3, 0x14, 0xca,
	0x84, 0x56, 0xb2, 0x90, 0xa0, 0x68, 0xb5, 0xa6, 0x82, 0x43, 0xc6, 0x8d, 0xae, 0x49, 0xae, 0xc1,
	0x80, 0x7b, 0x65, 0x31, 0xd2, 0x61, 0xc4, 0x62, 0xa4, 0x5a, 0x2f, 0xae, 0xfb, 0x72, 0x96, 0x4b,
	0xca, 0x94, 0x02, 0xc3, 0x8c, 0x04, 0x55, 0xd8, 0xaa, 0xd5, 0x2d, 0x9e, 0xee, 0xb8, 0x36, 0x7c,
	0xef, 0xce, 0x31, 0xda, 0x7b, 0x68, 0x89, 0xfc, 0xd3, 0x08, 0x75, 0xaa, 0xf6, 0x1c, 0xab, 0xea,
	0x15, 0xc1, 0x17, 0x5b, 0xd0, 0x19, 0x4b, 0x65, 0xc3, 0x93, 0x31, 0xef, 0x1c, 0xf0, 0x4e, 0xcb,
	0x7f, 0x23, 0x3c, 0x0f, 0xa0, 0x54, 0x89, 0x54, 0x22, 0x84, 0xb4, 0x76, 0x9f, 0xf0, 0xac, 0xe2,
	0xda, 0xc8, 0x98, 0x17, 0x1e, 0x5a, 0x4e, 0xfc, 0xb3, 0xfb, 0x1b, 0xf2, 0xf7, 0xbc, 0xc4, 0xda,
	0x47, 0x03, 0xef, 0xbe, 0xe1, 0x4b, 0x35, 0x34, 0xff, 0x18, 0x6c, 0x9c, 0xce, 0xc6, 0x3f, 0x66,
	0x33, 0x9e, 0x37, 0x72, 0xd5, 0xc1, 0x4d, 0xeb, 0xb1, 0x7a, 0xc4, 0xb3, 0x10, 0x0a, 0xd9, 0x06,
	0xf2, 0xdf, 0x7f, 0x5a, 0xd5, 0x78, 0x13, 0xab, 0x9a, 0xa0, 0xc4, 0x8b, 0x18, 0xb2, 0x23, 0x8d,
	0x83, 0xf3, 0x4d, 0xbf, 0x97, 0xb0, 0x0d, 0x38, 0x44, 0xef, 0x2f, 0x3d, 0x28, 0x20, 0x65, 0x4a,
	0x10, 0xd0, 0x82, 0x0a, 0xae, 0xba, 0xf8, 0xa9, 0x7d, 0x62, 0xb9, 0x2c, 0xc6, 0xeb, 0x7d, 0xb6,
	0xa7, 0x1f, 0x84, 0xbe, 0x9c, 0x93, 0xcd, 0xeb, 0x6e, 0xfb, 0x39, 0xed, 0x4a, 0x1e, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0xbb, 0x2b, 0x5f, 0x10, 0x02, 0x00, 0x00,
}
