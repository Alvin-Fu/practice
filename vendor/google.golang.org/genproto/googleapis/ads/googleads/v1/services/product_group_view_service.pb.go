// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/product_group_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [ProductGroupViewService.GetProductGroupView][google.ads.googleads.v1.services.ProductGroupViewService.GetProductGroupView].
type GetProductGroupViewRequest struct {
	// The resource name of the product group view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductGroupViewRequest) Reset()         { *m = GetProductGroupViewRequest{} }
func (m *GetProductGroupViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductGroupViewRequest) ProtoMessage()    {}
func (*GetProductGroupViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1170300877f7364, []int{0}
}

func (m *GetProductGroupViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductGroupViewRequest.Unmarshal(m, b)
}
func (m *GetProductGroupViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductGroupViewRequest.Marshal(b, m, deterministic)
}
func (m *GetProductGroupViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductGroupViewRequest.Merge(m, src)
}
func (m *GetProductGroupViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductGroupViewRequest.Size(m)
}
func (m *GetProductGroupViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductGroupViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductGroupViewRequest proto.InternalMessageInfo

func (m *GetProductGroupViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductGroupViewRequest)(nil), "google.ads.googleads.v1.services.GetProductGroupViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/product_group_view_service.proto", fileDescriptor_b1170300877f7364)
}

var fileDescriptor_b1170300877f7364 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x4b, 0xfb, 0x40,
	0x1c, 0x25, 0xf9, 0xc3, 0x1f, 0x0c, 0xba, 0xc4, 0x41, 0x09, 0x1d, 0x4a, 0xed, 0x20, 0x1d, 0xee,
	0x88, 0xc5, 0xe5, 0xaa, 0x43, 0xba, 0xc4, 0x49, 0x4a, 0x85, 0x0c, 0x12, 0x08, 0x31, 0x39, 0x8e,
	0x40, 0x93, 0x8b, 0xf7, 0xbb, 0xa4, 0x83, 0xb8, 0xe8, 0x47, 0xf0, 0x1b, 0x38, 0xfa, 0x3d, 0x5c,
	0x5c, 0xfd, 0x04, 0x82, 0x93, 0x9f, 0x42, 0xd2, 0xeb, 0x05, 0x2c, 0x0d, 0xdd, 0x1e, 0x77, 0xef,
	0xbd, 0xdf, 0xbb, 0xf7, 0x3b, 0xcb, 0x63, 0x9c, 0xb3, 0x05, 0xc5, 0x71, 0x0a, 0x58, 0xc1, 0x06,
	0xd5, 0x2e, 0x06, 0x2a, 0xea, 0x2c, 0xa1, 0x80, 0x4b, 0xc1, 0xd3, 0x2a, 0x91, 0x11, 0x13, 0xbc,
	0x2a, 0xa3, 0x3a, 0xa3, 0xcb, 0x68, 0x7d, 0x87, 0x4a, 0xc1, 0x25, 0xb7, 0xfb, 0x4a, 0x87, 0xe2,
	0x14, 0x50, 0x6b, 0x81, 0x6a, 0x17, 0x69, 0x0b, 0x87, 0x74, 0x0d, 0x11, 0x14, 0x78, 0x25, 0xb6,
	0x4f, 0x51, 0xee, 0x4e, 0x4f, 0x6b, 0xcb, 0x0c, 0xc7, 0x45, 0xc1, 0x65, 0x2c, 0x33, 0x5e, 0x80,
	0xba, 0x1d, 0x78, 0x96, 0xe3, 0x53, 0x39, 0x53, 0x62, 0xbf, 0xd1, 0x06, 0x19, 0x5d, 0xce, 0xe9,
	0x7d, 0x45, 0x41, 0xda, 0x27, 0xd6, 0x81, 0x9e, 0x10, 0x15, 0x71, 0x4e, 0x8f, 0x8d, 0xbe, 0x71,
	0xba, 0x37, 0xdf, 0xd7, 0x87, 0xd7, 0x71, 0x4e, 0xcf, 0xbe, 0x0c, 0xeb, 0x68, 0xd3, 0xe0, 0x46,
	0x25, 0xb7, 0xdf, 0x0d, 0xeb, 0x70, 0x8b, 0xbf, 0x7d, 0x81, 0x76, 0xbd, 0x19, 0x75, 0xc7, 0x72,
	0xc6, 0x9d, 0xea, 0xb6, 0x0f, 0xb4, 0xa9, 0x1d, 0x4c, 0x9e, 0x3e, 0xbf, 0x5f, 0xcc, 0x73, 0x7b,
	0xdc, 0xf4, 0xf6, 0xf0, 0xe7, 0x59, 0x97, 0x49, 0x05, 0x92, 0xe7, 0x54, 0x00, 0x1e, 0xe9, 0x22,
	0x5b, 0x21, 0xe0, 0xd1, 0xe3, 0xf4, 0xd9, 0xb4, 0x86, 0x09, 0xcf, 0x77, 0xa6, 0x9e, 0xf6, 0x3a,
	0x9a, 0x98, 0x35, 0x6d, 0xcf, 0x8c, 0xdb, 0xab, 0xb5, 0x03, 0xe3, 0x8b, 0xb8, 0x60, 0x88, 0x0b,
	0x86, 0x19, 0x2d, 0x56, 0xbb, 0xd0, 0x9b, 0x2d, 0x33, 0xe8, 0xfe, 0x4d, 0x13, 0x0d, 0x5e, 0xcd,
	0x7f, 0xbe, 0xe7, 0xbd, 0x99, 0x7d, 0x5f, 0x19, 0x7a, 0x29, 0x20, 0x05, 0x1b, 0x14, 0xb8, 0x68,
	0x3d, 0x18, 0x3e, 0x34, 0x25, 0xf4, 0x52, 0x08, 0x5b, 0x4a, 0x18, 0xb8, 0xa1, 0xa6, 0xfc, 0x98,
	0x43, 0x75, 0x4e, 0x88, 0x97, 0x02, 0x21, 0x2d, 0x89, 0x90, 0xc0, 0x25, 0x44, 0xd3, 0xee, 0xfe,
	0xaf, 0x72, 0x8e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0xc2, 0x72, 0x20, 0xf4, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductGroupViewServiceClient is the client API for ProductGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductGroupViewServiceClient interface {
	// Returns the requested product group view in full detail.
	GetProductGroupView(ctx context.Context, in *GetProductGroupViewRequest, opts ...grpc.CallOption) (*resources.ProductGroupView, error)
}

type productGroupViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductGroupViewServiceClient(cc *grpc.ClientConn) ProductGroupViewServiceClient {
	return &productGroupViewServiceClient{cc}
}

func (c *productGroupViewServiceClient) GetProductGroupView(ctx context.Context, in *GetProductGroupViewRequest, opts ...grpc.CallOption) (*resources.ProductGroupView, error) {
	out := new(resources.ProductGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ProductGroupViewService/GetProductGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductGroupViewServiceServer is the server API for ProductGroupViewService service.
type ProductGroupViewServiceServer interface {
	// Returns the requested product group view in full detail.
	GetProductGroupView(context.Context, *GetProductGroupViewRequest) (*resources.ProductGroupView, error)
}

// UnimplementedProductGroupViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductGroupViewServiceServer struct {
}

func (*UnimplementedProductGroupViewServiceServer) GetProductGroupView(ctx context.Context, req *GetProductGroupViewRequest) (*resources.ProductGroupView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGroupView not implemented")
}

func RegisterProductGroupViewServiceServer(s *grpc.Server, srv ProductGroupViewServiceServer) {
	s.RegisterService(&_ProductGroupViewService_serviceDesc, srv)
}

func _ProductGroupViewService_GetProductGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupViewServiceServer).GetProductGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ProductGroupViewService/GetProductGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupViewServiceServer).GetProductGroupView(ctx, req.(*GetProductGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ProductGroupViewService",
	HandlerType: (*ProductGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductGroupView",
			Handler:    _ProductGroupViewService_GetProductGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/product_group_view_service.proto",
}
