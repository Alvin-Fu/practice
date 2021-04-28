// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/cloudtrace/v2/tracing.proto

package cloudtrace

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message for the `BatchWriteSpans` method.
type BatchWriteSpansRequest struct {
	// Required. The name of the project where the spans belong. The format is
	// `projects/[PROJECT_ID]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. A list of new spans. The span names must not match existing
	// spans, or the results are undefined.
	Spans                []*Span  `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchWriteSpansRequest) Reset()         { *m = BatchWriteSpansRequest{} }
func (m *BatchWriteSpansRequest) String() string { return proto.CompactTextString(m) }
func (*BatchWriteSpansRequest) ProtoMessage()    {}
func (*BatchWriteSpansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f9b588db05fdc6, []int{0}
}

func (m *BatchWriteSpansRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchWriteSpansRequest.Unmarshal(m, b)
}
func (m *BatchWriteSpansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchWriteSpansRequest.Marshal(b, m, deterministic)
}
func (m *BatchWriteSpansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchWriteSpansRequest.Merge(m, src)
}
func (m *BatchWriteSpansRequest) XXX_Size() int {
	return xxx_messageInfo_BatchWriteSpansRequest.Size(m)
}
func (m *BatchWriteSpansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchWriteSpansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchWriteSpansRequest proto.InternalMessageInfo

func (m *BatchWriteSpansRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchWriteSpansRequest) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchWriteSpansRequest)(nil), "google.devtools.cloudtrace.v2.BatchWriteSpansRequest")
}

func init() {
	proto.RegisterFile("google/devtools/cloudtrace/v2/tracing.proto", fileDescriptor_d1f9b588db05fdc6)
}

var fileDescriptor_d1f9b588db05fdc6 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0x13, 0x3d,
	0x10, 0xc7, 0xb5, 0xc9, 0xf7, 0x21, 0x61, 0x90, 0x90, 0x56, 0xa2, 0xb4, 0x01, 0x44, 0x09, 0x12,
	0x94, 0x24, 0xb5, 0xc5, 0x56, 0x5c, 0x82, 0x00, 0x6d, 0x2a, 0x94, 0x6b, 0x94, 0xa2, 0x22, 0x41,
	0x24, 0xe4, 0x6c, 0x26, 0x1b, 0xa3, 0x5d, 0xdb, 0xd8, 0xce, 0x46, 0x80, 0x7a, 0xe1, 0xc6, 0x99,
	0x27, 0xe8, 0x15, 0xf1, 0x04, 0x3c, 0x42, 0x8f, 0x70, 0xe3, 0xd4, 0x03, 0x4f, 0xc1, 0x09, 0xad,
	0xbd, 0xdb, 0x84, 0xa8, 0x34, 0xbd, 0xad, 0x67, 0xfe, 0x9e, 0xf9, 0xcd, 0xfc, 0xbd, 0xa8, 0x19,
	0x0b, 0x11, 0x27, 0x40, 0x46, 0x90, 0x19, 0x21, 0x12, 0x4d, 0xa2, 0x44, 0x4c, 0x47, 0x46, 0xd1,
	0x08, 0x48, 0x16, 0x90, 0xfc, 0x83, 0xf1, 0x18, 0x4b, 0x25, 0x8c, 0xf0, 0x6f, 0x3a, 0x31, 0x2e,
	0xc5, 0x78, 0x2e, 0xc6, 0x59, 0x50, 0xbb, 0x51, 0xd4, 0xa2, 0x92, 0x11, 0xca, 0xb9, 0x30, 0xd4,
	0x30, 0xc1, 0xb5, 0xbb, 0x5c, 0xbb, 0xb6, 0x90, 0x8d, 0x12, 0x06, 0xdc, 0x14, 0x89, 0x5b, 0x0b,
	0x89, 0x31, 0x83, 0x64, 0xf4, 0x7a, 0x08, 0x13, 0x9a, 0x31, 0xa1, 0x0a, 0xc1, 0xc6, 0x82, 0x40,
	0x81, 0x16, 0x53, 0x15, 0x41, 0x91, 0xba, 0xbf, 0x1a, 0xbf, 0x94, 0x5e, 0x2f, 0xa4, 0xf6, 0x34,
	0x9c, 0x8e, 0x09, 0xa4, 0xd2, 0xbc, 0x5b, 0x62, 0x38, 0x49, 0x1a, 0x96, 0x82, 0x36, 0x34, 0x95,
	0x4e, 0x50, 0x3f, 0xf4, 0xd0, 0x5a, 0x87, 0x9a, 0x68, 0xf2, 0x42, 0x31, 0x03, 0x7b, 0x92, 0x72,
	0xdd, 0x87, 0xb7, 0x53, 0xd0, 0xc6, 0xef, 0xa2, 0xff, 0x38, 0x4d, 0x61, 0xdd, 0xdb, 0xf4, 0xb6,
	0x2e, 0x76, 0x76, 0x8e, 0xc3, 0xca, 0xef, 0x70, 0x1b, 0x35, 0x2d, 0x49, 0x89, 0x9b, 0x52, 0x4e,
	0x63, 0x50, 0xd8, 0xb5, 0xa1, 0x92, 0x69, 0x1c, 0x89, 0x94, 0xf4, 0x94, 0x78, 0x03, 0x91, 0xe9,
	0xdb, 0x02, 0xfe, 0x13, 0xf4, 0xbf, 0xce, 0x0b, 0xaf, 0x57, 0x36, 0xab, 0x5b, 0x97, 0x82, 0x3b,
	0xf8, 0xcc, 0x75, 0xe3, 0x1c, 0xa2, 0x53, 0x3d, 0x0e, 0x2b, 0x7d, 0x77, 0x2d, 0xf8, 0x56, 0x45,
	0x97, 0x9f, 0xe7, 0xd9, 0x3d, 0x50, 0x19, 0x8b, 0xc0, 0x3f, 0xf4, 0xd0, 0x95, 0x25, 0x68, 0xff,
	0xe1, 0x8a, 0xaa, 0xa7, 0x0f, 0x59, 0x5b, 0x2b, 0xaf, 0x95, 0x1b, 0xc2, 0xcf, 0xf2, 0xf5, 0xd5,
	0x9f, 0xfe, 0x0c, 0x51, 0x0e, 0xdf, 0xb2, 0x08, 0x1f, 0x7f, 0xfc, 0xfa, 0x5c, 0x69, 0xd5, 0xef,
	0xe5, 0xbb, 0xff, 0x90, 0x87, 0x1f, 0x4b, 0x37, 0xa1, 0x26, 0x8d, 0x03, 0xe7, 0x86, 0x6e, 0x0f,
	0x4f, 0x3a, 0xb4, 0xbd, 0x86, 0xff, 0xc9, 0x43, 0x68, 0x57, 0x01, 0x75, 0xfd, 0xfc, 0xf3, 0x0c,
	0x5d, 0x3b, 0x8f, 0xa8, 0xfe, 0xc0, 0xc2, 0x34, 0xeb, 0x77, 0x4f, 0x83, 0x29, 0x58, 0x48, 0x83,
	0x58, 0x6e, 0xd2, 0x38, 0x68, 0x7b, 0x8d, 0xda, 0xfb, 0xa3, 0x70, 0x63, 0xa1, 0xd2, 0xdf, 0x66,
	0x7d, 0x0f, 0x5f, 0x4d, 0x8c, 0x91, 0xba, 0x4d, 0xc8, 0x6c, 0x36, 0x5b, 0x76, 0x92, 0x4e, 0xcd,
	0xc4, 0x3d, 0xbe, 0x6d, 0x99, 0x50, 0x33, 0x16, 0x2a, 0x6d, 0xad, 0x92, 0xbb, 0x2e, 0x54, 0x4a,
	0xe0, 0xa3, 0xce, 0x57, 0x0f, 0xdd, 0x8e, 0x44, 0x7a, 0xf6, 0x64, 0x1d, 0xeb, 0x2f, 0xe3, 0x71,
	0x2f, 0x77, 0xa1, 0xe7, 0xbd, 0xec, 0x16, 0xf2, 0x58, 0x24, 0x94, 0xc7, 0x58, 0xa8, 0x98, 0xc4,
	0xc0, 0xad, 0x47, 0x64, 0xde, 0xf2, 0x1f, 0xbf, 0xc7, 0xa3, 0xf9, 0xe9, 0x4b, 0xe5, 0x6a, 0xd7,
	0x55, 0xda, 0xcd, 0x63, 0xd8, 0x3e, 0x23, 0xbc, 0x1f, 0x1c, 0x95, 0xf1, 0x81, 0x8d, 0x0f, 0x6c,
	0x7c, 0xb0, 0x1f, 0x0c, 0x2f, 0xd8, 0x1e, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xb4,
	0x7d, 0x35, 0x40, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Sends new spans to new or existing traces. You cannot update
	// existing spans.
	BatchWriteSpans(ctx context.Context, in *BatchWriteSpansRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Creates a new span.
	CreateSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*Span, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) BatchWriteSpans(ctx context.Context, in *BatchWriteSpansRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v2.TraceService/BatchWriteSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) CreateSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*Span, error) {
	out := new(Span)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v2.TraceService/CreateSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// Sends new spans to new or existing traces. You cannot update
	// existing spans.
	BatchWriteSpans(context.Context, *BatchWriteSpansRequest) (*empty.Empty, error)
	// Creates a new span.
	CreateSpan(context.Context, *Span) (*Span, error)
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) BatchWriteSpans(ctx context.Context, req *BatchWriteSpansRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchWriteSpans not implemented")
}
func (*UnimplementedTraceServiceServer) CreateSpan(ctx context.Context, req *Span) (*Span, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpan not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_BatchWriteSpans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchWriteSpansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).BatchWriteSpans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v2.TraceService/BatchWriteSpans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).BatchWriteSpans(ctx, req.(*BatchWriteSpansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceService_CreateSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Span)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).CreateSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v2.TraceService/CreateSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).CreateSpan(ctx, req.(*Span))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.cloudtrace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchWriteSpans",
			Handler:    _TraceService_BatchWriteSpans_Handler,
		},
		{
			MethodName: "CreateSpan",
			Handler:    _TraceService_CreateSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/cloudtrace/v2/tracing.proto",
}
