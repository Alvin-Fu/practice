// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/carrier_constant_service.proto

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
// [CarrierConstantService.GetCarrierConstant][google.ads.googleads.v1.services.CarrierConstantService.GetCarrierConstant].
type GetCarrierConstantRequest struct {
	// Resource name of the carrier constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCarrierConstantRequest) Reset()         { *m = GetCarrierConstantRequest{} }
func (m *GetCarrierConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetCarrierConstantRequest) ProtoMessage()    {}
func (*GetCarrierConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad8e2a96eee0407f, []int{0}
}

func (m *GetCarrierConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCarrierConstantRequest.Unmarshal(m, b)
}
func (m *GetCarrierConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCarrierConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetCarrierConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCarrierConstantRequest.Merge(m, src)
}
func (m *GetCarrierConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetCarrierConstantRequest.Size(m)
}
func (m *GetCarrierConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCarrierConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCarrierConstantRequest proto.InternalMessageInfo

func (m *GetCarrierConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCarrierConstantRequest)(nil), "google.ads.googleads.v1.services.GetCarrierConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/carrier_constant_service.proto", fileDescriptor_ad8e2a96eee0407f)
}

var fileDescriptor_ad8e2a96eee0407f = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xf3, 0x40,
	0x18, 0x24, 0xf9, 0xe1, 0x07, 0x83, 0x5e, 0x72, 0x10, 0xad, 0x1e, 0x4a, 0x2d, 0x52, 0x3c, 0xec,
	0x92, 0x7a, 0x91, 0x2d, 0xa2, 0x69, 0x0f, 0xf5, 0x24, 0xa5, 0x42, 0x0f, 0x12, 0x28, 0x6b, 0xb2,
	0x2c, 0x81, 0x76, 0xb7, 0xee, 0xb7, 0xcd, 0x45, 0xbc, 0xf4, 0x15, 0x7c, 0x03, 0x8f, 0xde, 0x7d,
	0x09, 0x4f, 0x82, 0xaf, 0xe0, 0xc9, 0xa7, 0x90, 0x74, 0xb3, 0x81, 0xd6, 0x86, 0xde, 0x86, 0xfd,
	0x66, 0xe6, 0x9b, 0x6f, 0x12, 0xef, 0x8a, 0x4b, 0xc9, 0x27, 0x0c, 0xd3, 0x04, 0xb0, 0x81, 0x39,
	0xca, 0x02, 0x0c, 0x4c, 0x65, 0x69, 0xcc, 0x00, 0xc7, 0x54, 0xa9, 0x94, 0xa9, 0x71, 0x2c, 0x05,
	0x68, 0x2a, 0xf4, 0xb8, 0x98, 0xa0, 0x99, 0x92, 0x5a, 0xfa, 0x75, 0xa3, 0x42, 0x34, 0x01, 0x54,
	0x1a, 0xa0, 0x2c, 0x40, 0xd6, 0xa0, 0x76, 0x51, 0xb5, 0x42, 0x31, 0x90, 0x73, 0xb5, 0x69, 0x87,
	0xf1, 0xae, 0x1d, 0x5b, 0xe5, 0x2c, 0xc5, 0x54, 0x08, 0xa9, 0xa9, 0x4e, 0xa5, 0x00, 0x33, 0x6d,
	0x5c, 0x7b, 0x87, 0x7d, 0xa6, 0x7b, 0x46, 0xda, 0x2b, 0x94, 0x43, 0xf6, 0x38, 0x67, 0xa0, 0xfd,
	0x13, 0x6f, 0xcf, 0xda, 0x8f, 0x05, 0x9d, 0xb2, 0x03, 0xa7, 0xee, 0xb4, 0x76, 0x86, 0xbb, 0xf6,
	0xf1, 0x96, 0x4e, 0x59, 0xfb, 0xd3, 0xf1, 0xf6, 0xd7, 0xf4, 0x77, 0x26, 0xb5, 0xff, 0xee, 0x78,
	0xfe, 0x5f, 0x77, 0xbf, 0x83, 0xb6, 0x9d, 0x8b, 0x2a, 0x33, 0xd5, 0xda, 0x95, 0xe2, 0xb2, 0x09,
	0xb4, 0x26, 0x6d, 0xa0, 0xc5, 0xd7, 0xf7, 0x8b, 0xdb, 0xf2, 0x4f, 0xf3, 0xc2, 0x9e, 0x56, 0x4e,
	0xba, 0x8c, 0x57, 0xb9, 0x80, 0xcf, 0x9e, 0xbb, 0x0b, 0xd7, 0x6b, 0xc6, 0x72, 0xba, 0x35, 0x66,
	0xf7, 0x68, 0xf3, 0xe1, 0x83, 0xbc, 0xda, 0x81, 0x73, 0x7f, 0x53, 0x18, 0x70, 0x39, 0xa1, 0x82,
	0x23, 0xa9, 0x38, 0xe6, 0x4c, 0x2c, 0x8b, 0xb7, 0x1f, 0x71, 0x96, 0x42, 0xf5, 0x6f, 0xd3, 0xb1,
	0xe0, 0xd5, 0xfd, 0xd7, 0x0f, 0xc3, 0x37, 0xb7, 0xde, 0x37, 0x86, 0x61, 0x02, 0xc8, 0xc0, 0x1c,
	0x8d, 0x02, 0x54, 0x2c, 0x86, 0x0f, 0x4b, 0x89, 0xc2, 0x04, 0xa2, 0x92, 0x12, 0x8d, 0x82, 0xc8,
	0x52, 0x7e, 0xdc, 0xa6, 0x79, 0x27, 0x24, 0x4c, 0x80, 0x90, 0x92, 0x44, 0xc8, 0x28, 0x20, 0xc4,
	0xd2, 0x1e, 0xfe, 0x2f, 0x73, 0x9e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xf6, 0x00, 0xfb,
	0xdd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CarrierConstantServiceClient is the client API for CarrierConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarrierConstantServiceClient interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error)
}

type carrierConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewCarrierConstantServiceClient(cc *grpc.ClientConn) CarrierConstantServiceClient {
	return &carrierConstantServiceClient{cc}
}

func (c *carrierConstantServiceClient) GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error) {
	out := new(resources.CarrierConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CarrierConstantService/GetCarrierConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierConstantServiceServer is the server API for CarrierConstantService service.
type CarrierConstantServiceServer interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error)
}

// UnimplementedCarrierConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarrierConstantServiceServer struct {
}

func (*UnimplementedCarrierConstantServiceServer) GetCarrierConstant(ctx context.Context, req *GetCarrierConstantRequest) (*resources.CarrierConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarrierConstant not implemented")
}

func RegisterCarrierConstantServiceServer(s *grpc.Server, srv CarrierConstantServiceServer) {
	s.RegisterService(&_CarrierConstantService_serviceDesc, srv)
}

func _CarrierConstantService_GetCarrierConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarrierConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CarrierConstantService/GetCarrierConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, req.(*GetCarrierConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarrierConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CarrierConstantService",
	HandlerType: (*CarrierConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarrierConstant",
			Handler:    _CarrierConstantService_GetCarrierConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/carrier_constant_service.proto",
}
