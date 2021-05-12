// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_criterion_simulation_service.proto

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
// [CampaignCriterionSimulationService.GetCampaignCriterionSimulation][google.ads.googleads.v1.services.CampaignCriterionSimulationService.GetCampaignCriterionSimulation].
type GetCampaignCriterionSimulationRequest struct {
	// The resource name of the campaign criterion simulation to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignCriterionSimulationRequest) Reset()         { *m = GetCampaignCriterionSimulationRequest{} }
func (m *GetCampaignCriterionSimulationRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignCriterionSimulationRequest) ProtoMessage()    {}
func (*GetCampaignCriterionSimulationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892d70bcad64eee6, []int{0}
}

func (m *GetCampaignCriterionSimulationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignCriterionSimulationRequest.Unmarshal(m, b)
}
func (m *GetCampaignCriterionSimulationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignCriterionSimulationRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignCriterionSimulationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignCriterionSimulationRequest.Merge(m, src)
}
func (m *GetCampaignCriterionSimulationRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignCriterionSimulationRequest.Size(m)
}
func (m *GetCampaignCriterionSimulationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignCriterionSimulationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignCriterionSimulationRequest proto.InternalMessageInfo

func (m *GetCampaignCriterionSimulationRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignCriterionSimulationRequest)(nil), "google.ads.googleads.v1.services.GetCampaignCriterionSimulationRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_criterion_simulation_service.proto", fileDescriptor_892d70bcad64eee6)
}

var fileDescriptor_892d70bcad64eee6 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x2e, 0x5c, 0xb8, 0xe1, 0xde, 0x4d, 0x56, 0x97, 0x22, 0x52, 0x6a, 0x45, 0xe9,
	0x62, 0x86, 0xe8, 0x6e, 0xc4, 0x62, 0x5a, 0x34, 0x2e, 0x8a, 0x94, 0x16, 0xba, 0x90, 0x40, 0x19,
	0x93, 0x21, 0x04, 0x9a, 0x99, 0x38, 0xff, 0xa4, 0x1b, 0x71, 0xe3, 0x03, 0xb8, 0x71, 0xe3, 0xda,
	0xa5, 0x8f, 0xe2, 0xd6, 0x57, 0x70, 0xe5, 0xce, 0x37, 0x90, 0x64, 0x3a, 0x01, 0xc1, 0x36, 0xee,
	0x0e, 0x33, 0x87, 0xef, 0xcc, 0x9c, 0xff, 0x77, 0x46, 0x89, 0x10, 0xc9, 0x82, 0x61, 0x1a, 0x03,
	0xd6, 0xb2, 0x54, 0x4b, 0x0f, 0x03, 0x93, 0xcb, 0x34, 0x62, 0x80, 0x23, 0x9a, 0xe5, 0x34, 0x4d,
	0xf8, 0x3c, 0x92, 0xa9, 0x62, 0x32, 0x15, 0x7c, 0x0e, 0x69, 0x56, 0x2c, 0xa8, 0xaa, 0xa4, 0xb6,
	0xa1, 0x5c, 0x0a, 0x25, 0xdc, 0xb6, 0x46, 0x20, 0x1a, 0x03, 0xaa, 0x69, 0x68, 0xe9, 0x21, 0x43,
	0x6b, 0x9d, 0xae, 0xcb, 0x93, 0x0c, 0x44, 0x21, 0x1b, 0x03, 0x75, 0x50, 0x6b, 0xcb, 0x60, 0xf2,
	0x14, 0x53, 0xce, 0x85, 0xaa, 0x2e, 0x41, 0xdf, 0x76, 0x46, 0xce, 0x6e, 0xc0, 0xd4, 0x70, 0xc5,
	0x19, 0x1a, 0xcc, 0xb4, 0xa6, 0x4c, 0xd8, 0x75, 0xc1, 0x40, 0xb9, 0x3b, 0xce, 0x3f, 0x93, 0x3b,
	0xe7, 0x34, 0x63, 0xff, 0xad, 0xb6, 0xb5, 0xff, 0x67, 0xf2, 0xd7, 0x1c, 0x5e, 0xd0, 0x8c, 0x1d,
	0x3c, 0xda, 0x4e, 0x67, 0x03, 0x6b, 0xaa, 0xbf, 0xe6, 0x7e, 0x58, 0xce, 0xf6, 0xe6, 0x54, 0x37,
	0x40, 0x4d, 0xfd, 0xa0, 0x1f, 0xbd, 0xbb, 0xd5, 0x5f, 0x0b, 0xaa, 0x6b, 0x44, 0x1b, 0x30, 0x9d,
	0xb3, 0xbb, 0xd7, 0xb7, 0x07, 0xfb, 0xc4, 0xed, 0x97, 0xcd, 0xdf, 0x7c, 0xa9, 0xe0, 0x38, 0x2a,
	0x40, 0x89, 0x8c, 0x49, 0xc0, 0xbd, 0x7a, 0x14, 0xdf, 0x30, 0x00, 0xf7, 0x6e, 0x07, 0xf7, 0xb6,
	0xd3, 0x8d, 0x44, 0xd6, 0xf8, 0xad, 0xc1, 0x5e, 0x73, 0x81, 0xe3, 0x72, 0x74, 0x63, 0xeb, 0xf2,
	0x7c, 0x05, 0x4b, 0xc4, 0x82, 0xf2, 0x04, 0x09, 0x99, 0xe0, 0x84, 0xf1, 0x6a, 0xb0, 0x66, 0x63,
	0xf2, 0x14, 0xd6, 0x2f, 0xec, 0x91, 0x11, 0x4f, 0xf6, 0xaf, 0xc0, 0xf7, 0x9f, 0xed, 0x76, 0xa0,
	0x81, 0x7e, 0x0c, 0x48, 0xcb, 0x52, 0xcd, 0x3c, 0xb4, 0x0a, 0x86, 0x17, 0x63, 0x09, 0xfd, 0x18,
	0xc2, 0xda, 0x12, 0xce, 0xbc, 0xd0, 0x58, 0xde, 0xed, 0xae, 0x3e, 0x27, 0xc4, 0x8f, 0x81, 0x90,
	0xda, 0x44, 0xc8, 0xcc, 0x23, 0xc4, 0xd8, 0xae, 0x7e, 0x57, 0xef, 0x3c, 0xfc, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x54, 0x12, 0xce, 0x57, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignCriterionSimulationServiceClient is the client API for CampaignCriterionSimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignCriterionSimulationServiceClient interface {
	// Returns the requested campaign criterion simulation in full detail.
	GetCampaignCriterionSimulation(ctx context.Context, in *GetCampaignCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.CampaignCriterionSimulation, error)
}

type campaignCriterionSimulationServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignCriterionSimulationServiceClient(cc *grpc.ClientConn) CampaignCriterionSimulationServiceClient {
	return &campaignCriterionSimulationServiceClient{cc}
}

func (c *campaignCriterionSimulationServiceClient) GetCampaignCriterionSimulation(ctx context.Context, in *GetCampaignCriterionSimulationRequest, opts ...grpc.CallOption) (*resources.CampaignCriterionSimulation, error) {
	out := new(resources.CampaignCriterionSimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignCriterionSimulationService/GetCampaignCriterionSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionSimulationServiceServer is the server API for CampaignCriterionSimulationService service.
type CampaignCriterionSimulationServiceServer interface {
	// Returns the requested campaign criterion simulation in full detail.
	GetCampaignCriterionSimulation(context.Context, *GetCampaignCriterionSimulationRequest) (*resources.CampaignCriterionSimulation, error)
}

// UnimplementedCampaignCriterionSimulationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignCriterionSimulationServiceServer struct {
}

func (*UnimplementedCampaignCriterionSimulationServiceServer) GetCampaignCriterionSimulation(ctx context.Context, req *GetCampaignCriterionSimulationRequest) (*resources.CampaignCriterionSimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignCriterionSimulation not implemented")
}

func RegisterCampaignCriterionSimulationServiceServer(s *grpc.Server, srv CampaignCriterionSimulationServiceServer) {
	s.RegisterService(&_CampaignCriterionSimulationService_serviceDesc, srv)
}

func _CampaignCriterionSimulationService_GetCampaignCriterionSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignCriterionSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionSimulationServiceServer).GetCampaignCriterionSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignCriterionSimulationService/GetCampaignCriterionSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionSimulationServiceServer).GetCampaignCriterionSimulation(ctx, req.(*GetCampaignCriterionSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignCriterionSimulationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignCriterionSimulationService",
	HandlerType: (*CampaignCriterionSimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignCriterionSimulation",
			Handler:    _CampaignCriterionSimulationService_GetCampaignCriterionSimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_criterion_simulation_service.proto",
}
