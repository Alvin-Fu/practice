// Code generated by protoc-gen-go. DO NOT EDIT.
// source: codec_perf/perf.proto

package codec_perf

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

// Buffer is a message that contains a body of bytes that is used to exercise
// encoding and decoding overheads.
type Buffer struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_afad72ea7772fe3a, []int{0}
}

func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Buffer)(nil), "codec.perf.Buffer")
}

func init() { proto.RegisterFile("codec_perf/perf.proto", fileDescriptor_afad72ea7772fe3a) }

var fileDescriptor_afad72ea7772fe3a = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0x4f, 0x49,
	0x4d, 0x8e, 0x2f, 0x48, 0x2d, 0x4a, 0xd3, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x5c, 0x60, 0x61, 0x3d, 0x90, 0x88, 0x92, 0x0c, 0x17, 0x9b, 0x53, 0x69, 0x5a, 0x5a, 0x6a, 0x91,
	0x90, 0x10, 0x17, 0x4b, 0x52, 0x7e, 0x4a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98,
	0x9d, 0xc4, 0x06, 0xd6, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x5f, 0x4f, 0x3c, 0x49,
	0x00, 0x00, 0x00,
}
