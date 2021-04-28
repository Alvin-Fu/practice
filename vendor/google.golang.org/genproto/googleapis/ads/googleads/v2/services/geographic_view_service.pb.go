// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/geographic_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [GeographicViewService.GetGeographicView][google.ads.googleads.v2.services.GeographicViewService.GetGeographicView].
type GetGeographicViewRequest struct {
	// The resource name of the geographic view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGeographicViewRequest) Reset()         { *m = GetGeographicViewRequest{} }
func (m *GetGeographicViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGeographicViewRequest) ProtoMessage()    {}
func (*GetGeographicViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51498bd55db0eccc, []int{0}
}

func (m *GetGeographicViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGeographicViewRequest.Unmarshal(m, b)
}
func (m *GetGeographicViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGeographicViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGeographicViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGeographicViewRequest.Merge(m, src)
}
func (m *GetGeographicViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGeographicViewRequest.Size(m)
}
func (m *GetGeographicViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGeographicViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGeographicViewRequest proto.InternalMessageInfo

func (m *GetGeographicViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGeographicViewRequest)(nil), "google.ads.googleads.v2.services.GetGeographicViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/geographic_view_service.proto", fileDescriptor_51498bd55db0eccc)
}

var fileDescriptor_51498bd55db0eccc = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4a, 0xf3, 0x40,
	0x1c, 0xc5, 0x49, 0x3e, 0xf8, 0xc0, 0xa0, 0x0b, 0x03, 0x62, 0x89, 0x2e, 0x4a, 0xed, 0x42, 0xba,
	0x98, 0xa1, 0xe9, 0x42, 0x1c, 0x51, 0x49, 0x37, 0x71, 0x25, 0xa5, 0x42, 0x16, 0x12, 0x28, 0x63,
	0x32, 0x8c, 0x03, 0x4d, 0x26, 0x66, 0xd2, 0x74, 0x21, 0x2e, 0xf4, 0x0a, 0xde, 0xc0, 0xa5, 0x77,
	0xf0, 0x02, 0xdd, 0x7a, 0x05, 0x57, 0xae, 0x3d, 0x80, 0xa4, 0x93, 0x49, 0x5b, 0x6c, 0xe8, 0xee,
	0x31, 0xff, 0xf7, 0xfb, 0xbf, 0x99, 0x97, 0x18, 0x17, 0x94, 0x73, 0x3a, 0x26, 0x10, 0x87, 0x02,
	0x4a, 0x59, 0xa8, 0xdc, 0x86, 0x82, 0xa4, 0x39, 0x0b, 0x88, 0x80, 0x94, 0x70, 0x9a, 0xe2, 0xe4,
	0x9e, 0x05, 0xa3, 0x9c, 0x91, 0xe9, 0xa8, 0x1c, 0x80, 0x24, 0xe5, 0x19, 0x37, 0x9b, 0x12, 0x02,
	0x38, 0x14, 0xa0, 0xe2, 0x41, 0x6e, 0x03, 0xc5, 0x5b, 0x27, 0x75, 0x09, 0x29, 0x11, 0x7c, 0x92,
	0xae, 0x89, 0x90, 0xab, 0xad, 0x43, 0x05, 0x26, 0x0c, 0xe2, 0x38, 0xe6, 0x19, 0xce, 0x18, 0x8f,
	0x45, 0x39, 0xdd, 0x5f, 0x9a, 0x06, 0x63, 0x46, 0xe2, 0x4c, 0x0e, 0x5a, 0x97, 0x46, 0xc3, 0x25,
	0x99, 0x5b, 0xad, 0xf4, 0x18, 0x99, 0x0e, 0xc9, 0xc3, 0x84, 0x88, 0xcc, 0x3c, 0x32, 0x76, 0x54,
	0xea, 0x28, 0xc6, 0x11, 0x69, 0x68, 0x4d, 0xed, 0x78, 0x6b, 0xb8, 0xad, 0x0e, 0xaf, 0x71, 0x44,
	0xec, 0x1f, 0xcd, 0xd8, 0x5b, 0xc5, 0x6f, 0xe4, 0x5b, 0xcc, 0x0f, 0xcd, 0xd8, 0xfd, 0xb3, 0xdb,
	0x44, 0x60, 0x53, 0x07, 0xa0, 0xee, 0x42, 0x56, 0xb7, 0x96, 0xad, 0xda, 0x01, 0xab, 0x64, 0xeb,
	0xf4, 0xe5, 0xf3, 0xeb, 0x55, 0xef, 0x99, 0xdd, 0xa2, 0xc3, 0xc7, 0x95, 0xe7, 0x9c, 0x07, 0x13,
	0x91, 0xf1, 0x88, 0xa4, 0x02, 0x76, 0x96, 0x4a, 0x2d, 0x30, 0x01, 0x3b, 0x4f, 0xd6, 0xc1, 0xcc,
	0x69, 0x2c, 0x42, 0x4a, 0x95, 0x30, 0x01, 0x02, 0x1e, 0xf5, 0x9f, 0x75, 0xa3, 0x1d, 0xf0, 0x68,
	0xe3, 0x63, 0xfa, 0xd6, 0xda, 0x72, 0x06, 0x45, 0xf9, 0x03, 0xed, 0xf6, 0xaa, 0xe4, 0x29, 0x1f,
	0xe3, 0x98, 0x02, 0x9e, 0x52, 0x48, 0x49, 0x3c, 0xff, 0x34, 0x70, 0x91, 0x58, 0xff, 0xbf, 0x9d,
	0x29, 0xf1, 0xa6, 0xff, 0x73, 0x1d, 0xe7, 0x5d, 0x6f, 0xba, 0x72, 0xa1, 0x13, 0x0a, 0x20, 0x65,
	0xa1, 0x3c, 0x1b, 0x94, 0xc1, 0x62, 0xa6, 0x2c, 0xbe, 0x13, 0x0a, 0xbf, 0xb2, 0xf8, 0x9e, 0xed,
	0x2b, 0xcb, 0xb7, 0xde, 0x96, 0xe7, 0x08, 0x39, 0xa1, 0x40, 0xa8, 0x32, 0x21, 0xe4, 0xd9, 0x08,
	0x29, 0xdb, 0xdd, 0xff, 0xf9, 0x3d, 0x7b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe8, 0xe9,
	0xd5, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeographicViewServiceClient is the client API for GeographicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeographicViewServiceClient interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error)
}

type geographicViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeographicViewServiceClient(cc *grpc.ClientConn) GeographicViewServiceClient {
	return &geographicViewServiceClient{cc}
}

func (c *geographicViewServiceClient) GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error) {
	out := new(resources.GeographicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GeographicViewService/GetGeographicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographicViewServiceServer is the server API for GeographicViewService service.
type GeographicViewServiceServer interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(context.Context, *GetGeographicViewRequest) (*resources.GeographicView, error)
}

// UnimplementedGeographicViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeographicViewServiceServer struct {
}

func (*UnimplementedGeographicViewServiceServer) GetGeographicView(ctx context.Context, req *GetGeographicViewRequest) (*resources.GeographicView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeographicView not implemented")
}

func RegisterGeographicViewServiceServer(s *grpc.Server, srv GeographicViewServiceServer) {
	s.RegisterService(&_GeographicViewService_serviceDesc, srv)
}

func _GeographicViewService_GetGeographicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeographicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GeographicViewService/GetGeographicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, req.(*GetGeographicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeographicViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GeographicViewService",
	HandlerType: (*GeographicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeographicView",
			Handler:    _GeographicViewService_GetGeographicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/geographic_view_service.proto",
}
