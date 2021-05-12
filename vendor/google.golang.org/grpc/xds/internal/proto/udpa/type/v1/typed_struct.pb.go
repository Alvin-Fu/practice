// Code generated by protoc-gen-go. DO NOT EDIT.
// source: udpa/type/v1/typed_struct.proto

package udpa_type_v1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// A TypedStruct contains an arbitrary JSON serialized protocol buffer message with a URL that
// describes the type of the serialized message. This is very similar to google.protobuf.Any,
// instead of having protocol buffer binary, this employs google.protobuf.Struct as value.
//
// This message is intended to be embedded inside Any, so it shouldn't be directly referred
// from other UDPA messages.
//
// When packing an opaque extension config, packing the expected type into Any is preferred
// wherever possible for its efficiency. TypedStruct should be used only if a proto descriptor
// is not available, for example if:
// - A control plane sends opaque message that is originally from external source in human readable
//   format such as JSON or YAML.
// - The control plane doesn't have the knowledge of the protocol buffer schema hence it cannot
//   serialize the message in protocol buffer binary format.
// - The DPLB doesn't have have the knowledge of the protocol buffer schema its plugin or extension
//   uses. This has to be indicated in the DPLB capability negotiation.
//
// When a DPLB receives a TypedStruct in Any, it should:
// - Check if the type_url of the TypedStruct matches the type the extension expects.
// - Convert value to the type described in type_url and perform validation.
// TODO(lizan): Figure out how TypeStruct should be used with DPLB extensions that doesn't link
// protobuf descriptor with DPLB itself, (e.g. gRPC LB Plugin, Envoy WASM extensions).
type TypedStruct struct {
	// A URL that uniquely identifies the type of the serialize protocol buffer message.
	// This has same semantics and format described in google.protobuf.Any:
	// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/any.proto
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// A JSON representation of the above specified type.
	Value                *_struct.Struct `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TypedStruct) Reset()         { *m = TypedStruct{} }
func (m *TypedStruct) String() string { return proto.CompactTextString(m) }
func (*TypedStruct) ProtoMessage()    {}
func (*TypedStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_098110268becad8a, []int{0}
}

func (m *TypedStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedStruct.Unmarshal(m, b)
}
func (m *TypedStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedStruct.Marshal(b, m, deterministic)
}
func (m *TypedStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedStruct.Merge(m, src)
}
func (m *TypedStruct) XXX_Size() int {
	return xxx_messageInfo_TypedStruct.Size(m)
}
func (m *TypedStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedStruct.DiscardUnknown(m)
}

var xxx_messageInfo_TypedStruct proto.InternalMessageInfo

func (m *TypedStruct) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *TypedStruct) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedStruct)(nil), "udpa.type.v1.TypedStruct")
}

func init() { proto.RegisterFile("udpa/type/v1/typed_struct.proto", fileDescriptor_098110268becad8a) }

var fileDescriptor_098110268becad8a = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4d, 0x29, 0x48,
	0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0x04, 0xd3, 0x29, 0xf1, 0xc5, 0x25, 0x45, 0xa5,
	0xc9, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x3c, 0x20, 0x05, 0x7a, 0x20, 0x09, 0xbd,
	0x32, 0x43, 0x29, 0xf1, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0x54, 0x7d, 0x18, 0x03, 0xa2,
	0x4c, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f, 0xcc, 0x4b, 0x2a, 0x4d, 0xd3, 0x47,
	0x36, 0x44, 0x29, 0x9c, 0x8b, 0x3b, 0x04, 0x64, 0x74, 0x30, 0x58, 0x50, 0x48, 0x92, 0x8b, 0x03,
	0x64, 0x60, 0x7c, 0x69, 0x51, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x3b, 0x88, 0x1f,
	0x5a, 0x94, 0x23, 0xa4, 0xcb, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xae, 0x07, 0x31, 0x57, 0x0f, 0x66, 0xae, 0x1e, 0xc4, 0x88, 0x20, 0x88, 0x2a,
	0x27, 0x23, 0x2e, 0x99, 0xe4, 0xfc, 0x5c, 0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x3d, 0xb0,
	0x53, 0x91, 0xdd, 0xeb, 0x24, 0x80, 0x64, 0x6d, 0x00, 0xc8, 0x88, 0x00, 0xc6, 0x24, 0x36, 0xb0,
	0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x7c, 0xa9, 0xb5, 0xfb, 0x00, 0x00, 0x00,
}
