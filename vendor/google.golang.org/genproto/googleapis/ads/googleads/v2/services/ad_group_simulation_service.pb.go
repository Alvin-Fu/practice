// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_simulation_service.proto

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

// Request message for [AdGroupSimulationService.GetAdGroupSimulation][google.ads.googleads.v2.services.AdGroupSimulationService.GetAdGroupSimulation].
type GetAdGroupSimulationRequest struct {
	// The resource name of the ad group simulation to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupSimulationRequest) Reset()         { *m = GetAdGroupSimulationRequest{} }
func (m *GetAdGroupSimulationRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupSimulationRequest) ProtoMessage()    {}
func (*GetAdGroupSimulationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e985cbf2545305b, []int{0}
}

func (m *GetAdGroupSimulationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupSimulationRequest.Unmarshal(m, b)
}
func (m *GetAdGroupSimulationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupSimulationRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupSimulationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupSimulationRequest.Merge(m, src)
}
func (m *GetAdGroupSimulationRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupSimulationRequest.Size(m)
}
func (m *GetAdGroupSimulationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupSimulationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupSimulationRequest proto.InternalMessageInfo

func (m *GetAdGroupSimulationRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupSimulationRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupSimulationRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_simulation_service.proto", fileDescriptor_3e985cbf2545305b)
}

var fileDescriptor_3e985cbf2545305b = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4a, 0xe3, 0x40,
	0x18, 0x27, 0x59, 0x58, 0xd8, 0xb0, 0x7b, 0x09, 0x0b, 0x96, 0x54, 0xa1, 0xd4, 0x1e, 0xa4, 0x87,
	0x19, 0x88, 0xc5, 0xc3, 0xd4, 0x1e, 0x92, 0x4b, 0x3c, 0x49, 0x69, 0xa1, 0x07, 0x09, 0x84, 0x31,
	0x19, 0x42, 0x20, 0xc9, 0xc4, 0x7c, 0x49, 0x2f, 0xe2, 0x45, 0x7d, 0x03, 0xdf, 0xc0, 0xa3, 0x6f,
	0x62, 0xaf, 0xbe, 0x82, 0x27, 0x9f, 0x42, 0xd2, 0xc9, 0xa4, 0x6a, 0x8d, 0xbd, 0xfd, 0x98, 0xdf,
	0xbf, 0xf9, 0xbe, 0x19, 0xcd, 0x0e, 0x39, 0x0f, 0x63, 0x86, 0x69, 0x00, 0x58, 0xc0, 0x0a, 0x2d,
	0x4d, 0x0c, 0x2c, 0x5f, 0x46, 0x3e, 0x03, 0x4c, 0x03, 0x2f, 0xcc, 0x79, 0x99, 0x79, 0x10, 0x25,
	0x65, 0x4c, 0x8b, 0x88, 0xa7, 0x5e, 0x4d, 0xa2, 0x2c, 0xe7, 0x05, 0xd7, 0x7b, 0xc2, 0x88, 0x68,
	0x00, 0xa8, 0xc9, 0x40, 0x4b, 0x13, 0xc9, 0x0c, 0x63, 0xdc, 0xd6, 0x92, 0x33, 0xe0, 0x65, 0xde,
	0x52, 0x23, 0xe2, 0x8d, 0x7d, 0x69, 0xce, 0x22, 0x4c, 0xd3, 0x94, 0x17, 0x6b, 0x12, 0x6a, 0x76,
	0xef, 0x03, 0xeb, 0xc7, 0x11, 0x4b, 0x0b, 0x41, 0xf4, 0x6d, 0xad, 0xeb, 0xb0, 0xc2, 0x0a, 0x9c,
	0x2a, 0x75, 0xde, 0x84, 0xce, 0xd8, 0x55, 0xc9, 0xa0, 0xd0, 0x0f, 0xb5, 0x7f, 0xb2, 0xdc, 0x4b,
	0x69, 0xc2, 0x3a, 0x4a, 0x4f, 0x39, 0xfa, 0x33, 0xfb, 0x2b, 0x0f, 0xcf, 0x69, 0xc2, 0xcc, 0x3b,
	0x55, 0xeb, 0x6c, 0x25, 0xcc, 0xc5, 0x54, 0xfa, 0xb3, 0xa2, 0xfd, 0xff, 0xae, 0x41, 0x9f, 0xa0,
	0x5d, 0x0b, 0x41, 0x3f, 0xdc, 0xcc, 0x18, 0xb5, 0xda, 0x9b, 0x6d, 0xa1, 0x2d, 0x73, 0xff, 0xf4,
	0xf6, 0xe5, 0xf5, 0x41, 0x3d, 0xd1, 0x47, 0xd5, 0x5a, 0xaf, 0x3f, 0x8d, 0x36, 0xf1, 0x4b, 0x28,
	0x78, 0xc2, 0x72, 0xc0, 0x43, 0x4c, 0xbf, 0x3a, 0x01, 0x0f, 0x6f, 0x8c, 0xee, 0xca, 0xea, 0x6c,
	0xaa, 0x6a, 0x94, 0x45, 0x80, 0x7c, 0x9e, 0xd8, 0xf7, 0xaa, 0x36, 0xf0, 0x79, 0xb2, 0x73, 0x2a,
	0xfb, 0xa0, 0x6d, 0x57, 0xd3, 0xea, 0x45, 0xa6, 0xca, 0xc5, 0x59, 0x1d, 0x11, 0xf2, 0x98, 0xa6,
	0x21, 0xe2, 0x79, 0x88, 0x43, 0x96, 0xae, 0xdf, 0x0b, 0x6f, 0x4a, 0xdb, 0x3f, 0xe3, 0x58, 0x82,
	0x47, 0xf5, 0x97, 0x63, 0x59, 0x4f, 0x6a, 0xcf, 0x11, 0x81, 0x56, 0x00, 0x48, 0xc0, 0x0a, 0x2d,
	0x4c, 0x54, 0x17, 0xc3, 0x4a, 0x4a, 0x5c, 0x2b, 0x00, 0xb7, 0x91, 0xb8, 0x0b, 0xd3, 0x95, 0x92,
	0x37, 0x75, 0x20, 0xce, 0x09, 0xb1, 0x02, 0x20, 0xa4, 0x11, 0x11, 0xb2, 0x30, 0x09, 0x91, 0xb2,
	0xcb, 0xdf, 0xeb, 0x7b, 0x1e, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x36, 0xdf, 0x10, 0xc0, 0x33,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupSimulationServiceClient is the client API for AdGroupSimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupSimulationServiceClient interface {
	// Returns the requested ad group simulation in full detail.
	GetAdGroupSimulation(ctx context.Context, in *GetAdGroupSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupSimulation, error)
}

type adGroupSimulationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupSimulationServiceClient(cc *grpc.ClientConn) AdGroupSimulationServiceClient {
	return &adGroupSimulationServiceClient{cc}
}

func (c *adGroupSimulationServiceClient) GetAdGroupSimulation(ctx context.Context, in *GetAdGroupSimulationRequest, opts ...grpc.CallOption) (*resources.AdGroupSimulation, error) {
	out := new(resources.AdGroupSimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupSimulationService/GetAdGroupSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupSimulationServiceServer is the server API for AdGroupSimulationService service.
type AdGroupSimulationServiceServer interface {
	// Returns the requested ad group simulation in full detail.
	GetAdGroupSimulation(context.Context, *GetAdGroupSimulationRequest) (*resources.AdGroupSimulation, error)
}

// UnimplementedAdGroupSimulationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupSimulationServiceServer struct {
}

func (*UnimplementedAdGroupSimulationServiceServer) GetAdGroupSimulation(ctx context.Context, req *GetAdGroupSimulationRequest) (*resources.AdGroupSimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupSimulation not implemented")
}

func RegisterAdGroupSimulationServiceServer(s *grpc.Server, srv AdGroupSimulationServiceServer) {
	s.RegisterService(&_AdGroupSimulationService_serviceDesc, srv)
}

func _AdGroupSimulationService_GetAdGroupSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupSimulationServiceServer).GetAdGroupSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupSimulationService/GetAdGroupSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupSimulationServiceServer).GetAdGroupSimulation(ctx, req.(*GetAdGroupSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupSimulationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupSimulationService",
	HandlerType: (*AdGroupSimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupSimulation",
			Handler:    _AdGroupSimulationService_GetAdGroupSimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_simulation_service.proto",
}
