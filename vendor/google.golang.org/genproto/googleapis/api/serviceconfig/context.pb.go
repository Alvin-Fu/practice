// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/context.proto

package serviceconfig

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

// `Context` defines which contexts an API requests.
//
// Example:
//
//     context:
//       rules:
//       - selector: "*"
//         requested:
//         - google.rpc.context.ProjectContext
//         - google.rpc.context.OriginContext
//
// The above specifies that all methods in the API request
// `google.rpc.context.ProjectContext` and
// `google.rpc.context.OriginContext`.
//
// Available context types are defined in package
// `google.rpc.context`.
//
// This also provides mechanism to whitelist any protobuf message extension that
// can be sent in grpc metadata using “x-goog-ext-<extension_id>-bin” and
// “x-goog-ext-<extension_id>-jspb” format. For example, list any service
// specific protobuf types that can appear in grpc metadata as follows in your
// yaml file:
//
// Example:
//
//     context:
//       rules:
//        - selector: "google.example.library.v1.LibraryService.CreateBook"
//          allowed_request_extensions:
//          - google.foo.v1.NewExtension
//          allowed_response_extensions:
//          - google.foo.v1.NewExtension
//
// You can also specify extension ID instead of fully qualified extension name
// here.
type Context struct {
	// A list of RPC context rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules                []*ContextRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_48d8be90143bd46c, []int{0}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetRules() []*ContextRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A context rule provides information about the context for an individual API
// element.
type ContextRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// A list of full type names of requested contexts.
	Requested []string `protobuf:"bytes,2,rep,name=requested,proto3" json:"requested,omitempty"`
	// A list of full type names of provided contexts.
	Provided []string `protobuf:"bytes,3,rep,name=provided,proto3" json:"provided,omitempty"`
	// A list of full type names or extension IDs of extensions allowed in grpc
	// side channel from client to backend.
	AllowedRequestExtensions []string `protobuf:"bytes,4,rep,name=allowed_request_extensions,json=allowedRequestExtensions,proto3" json:"allowed_request_extensions,omitempty"`
	// A list of full type names or extension IDs of extensions allowed in grpc
	// side channel from backend to client.
	AllowedResponseExtensions []string `protobuf:"bytes,5,rep,name=allowed_response_extensions,json=allowedResponseExtensions,proto3" json:"allowed_response_extensions,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *ContextRule) Reset()         { *m = ContextRule{} }
func (m *ContextRule) String() string { return proto.CompactTextString(m) }
func (*ContextRule) ProtoMessage()    {}
func (*ContextRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_48d8be90143bd46c, []int{1}
}

func (m *ContextRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextRule.Unmarshal(m, b)
}
func (m *ContextRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextRule.Marshal(b, m, deterministic)
}
func (m *ContextRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextRule.Merge(m, src)
}
func (m *ContextRule) XXX_Size() int {
	return xxx_messageInfo_ContextRule.Size(m)
}
func (m *ContextRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextRule.DiscardUnknown(m)
}

var xxx_messageInfo_ContextRule proto.InternalMessageInfo

func (m *ContextRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *ContextRule) GetRequested() []string {
	if m != nil {
		return m.Requested
	}
	return nil
}

func (m *ContextRule) GetProvided() []string {
	if m != nil {
		return m.Provided
	}
	return nil
}

func (m *ContextRule) GetAllowedRequestExtensions() []string {
	if m != nil {
		return m.AllowedRequestExtensions
	}
	return nil
}

func (m *ContextRule) GetAllowedResponseExtensions() []string {
	if m != nil {
		return m.AllowedResponseExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Context)(nil), "google.api.Context")
	proto.RegisterType((*ContextRule)(nil), "google.api.ContextRule")
}

func init() { proto.RegisterFile("google/api/context.proto", fileDescriptor_48d8be90143bd46c) }

var fileDescriptor_48d8be90143bd46c = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x95, 0xfe, 0x00, 0x71, 0x11, 0x83, 0x17, 0x4c, 0x61, 0x88, 0x3a, 0x65, 0x21, 0x91,
	0x60, 0x41, 0x02, 0x21, 0x51, 0x54, 0x21, 0xb6, 0x28, 0x23, 0x4b, 0x15, 0x92, 0xc3, 0xb2, 0x64,
	0x7c, 0xc6, 0x76, 0x4a, 0x9f, 0x87, 0x37, 0xe3, 0x4d, 0x50, 0xed, 0x28, 0x09, 0xe3, 0xdd, 0xf7,
	0x7d, 0x83, 0xcf, 0x84, 0x71, 0x44, 0x2e, 0x21, 0xaf, 0xb4, 0xc8, 0x6b, 0x54, 0x0e, 0xf6, 0x2e,
	0xd3, 0x06, 0x1d, 0x52, 0x12, 0x48, 0x56, 0x69, 0xb1, 0xba, 0x23, 0xc7, 0xcf, 0x01, 0xd2, 0x6b,
	0x32, 0x37, 0xad, 0x04, 0xcb, 0xa2, 0x64, 0x9a, 0x2e, 0x6e, 0xce, 0xb3, 0x41, 0xcb, 0x3a, 0xa7,
	0x6c, 0x25, 0x94, 0xc1, 0x5a, 0xfd, 0x46, 0x64, 0x31, 0x5a, 0xd3, 0x25, 0x39, 0xb1, 0x20, 0xa1,
	0x76, 0x68, 0x58, 0x94, 0x44, 0x69, 0x5c, 0xf6, 0x33, 0xbd, 0x22, 0xb1, 0x81, 0xaf, 0x16, 0xac,
	0x83, 0x86, 0x4d, 0x92, 0x69, 0x1a, 0x97, 0xc3, 0xe2, 0x50, 0x6a, 0x83, 0x3b, 0xd1, 0x40, 0xc3,
	0xa6, 0x1e, 0xf6, 0x33, 0x7d, 0x20, 0xcb, 0x4a, 0x4a, 0xfc, 0x86, 0x66, 0xdb, 0x05, 0x5b, 0xd8,
	0x3b, 0x50, 0x56, 0xa0, 0xb2, 0x6c, 0xe6, 0x6d, 0xd6, 0x19, 0x65, 0x10, 0x36, 0x3d, 0xa7, 0x8f,
	0xe4, 0x72, 0xa8, 0xad, 0x46, 0x65, 0x61, 0x9c, 0xcf, 0x7d, 0x7e, 0xd1, 0xe7, 0xc1, 0x18, 0xfa,
	0xb5, 0x22, 0x67, 0x35, 0x7e, 0x8e, 0x0e, 0xb1, 0x3e, 0xed, 0x9e, 0x5c, 0x1c, 0x2e, 0x59, 0x44,
	0x6f, 0x9b, 0x8e, 0x71, 0x94, 0x95, 0xe2, 0x19, 0x1a, 0x9e, 0x73, 0x50, 0xfe, 0xce, 0x79, 0x40,
	0x95, 0x16, 0xd6, 0x7f, 0x82, 0x05, 0xb3, 0x13, 0x35, 0xd4, 0xa8, 0x3e, 0x04, 0xbf, 0xff, 0x37,
	0xfd, 0x4c, 0x66, 0x2f, 0x4f, 0xc5, 0xeb, 0xfb, 0x91, 0x0f, 0x6f, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x00, 0x40, 0x95, 0xa9, 0xbc, 0x01, 0x00, 0x00,
}
