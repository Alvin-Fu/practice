// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_audience_view_service.proto

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

// Request message for [AdGroupAudienceViewService.GetAdGoupAudienceView][].
type GetAdGroupAudienceViewRequest struct {
	// The resource name of the ad group audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAudienceViewRequest) Reset()         { *m = GetAdGroupAudienceViewRequest{} }
func (m *GetAdGroupAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAudienceViewRequest) ProtoMessage()    {}
func (*GetAdGroupAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e713f94d9c114b9, []int{0}
}

func (m *GetAdGroupAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.Merge(m, src)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Size(m)
}
func (m *GetAdGroupAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAudienceViewRequest proto.InternalMessageInfo

func (m *GetAdGroupAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAudienceViewRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_audience_view_service.proto", fileDescriptor_1e713f94d9c114b9)
}

var fileDescriptor_1e713f94d9c114b9 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0xcb, 0xd3, 0x40,
	0x10, 0x25, 0x11, 0x04, 0x83, 0x5e, 0x72, 0xd0, 0x12, 0x15, 0x4b, 0xed, 0x41, 0x7a, 0xd8, 0x85,
	0x08, 0x45, 0x56, 0xac, 0x6c, 0x51, 0xe2, 0x49, 0x4a, 0x85, 0x1c, 0x24, 0x10, 0xd6, 0xec, 0x10,
	0x16, 0x9a, 0xdd, 0x98, 0x4d, 0xd2, 0x83, 0x78, 0xf1, 0xe0, 0xcd, 0x93, 0xff, 0xc0, 0xa3, 0x3f,
	0xa5, 0x37, 0xf1, 0x2f, 0x78, 0xf2, 0x57, 0x48, 0xba, 0xd9, 0x54, 0xb1, 0xf9, 0x7a, 0x7b, 0xec,
	0xbc, 0x79, 0x6f, 0xe6, 0xcd, 0x7a, 0x2f, 0x73, 0xa5, 0xf2, 0x1d, 0x60, 0xc6, 0x35, 0x36, 0xb0,
	0x43, 0x6d, 0x88, 0x35, 0x54, 0xad, 0xc8, 0x40, 0x63, 0xc6, 0xd3, 0xbc, 0x52, 0x4d, 0x99, 0xb2,
	0x86, 0x0b, 0x90, 0x19, 0xa4, 0xad, 0x80, 0x7d, 0xda, 0xd7, 0x51, 0x59, 0xa9, 0x5a, 0xf9, 0x53,
	0xd3, 0x8b, 0x18, 0xd7, 0x68, 0x90, 0x41, 0x6d, 0x88, 0xac, 0x4c, 0xb0, 0x1a, 0x33, 0xaa, 0x40,
	0xab, 0xa6, 0x1a, 0x77, 0x32, 0x0e, 0xc1, 0x3d, 0xdb, 0x5f, 0x0a, 0xcc, 0xa4, 0x54, 0x35, 0xab,
	0x85, 0x92, 0xba, 0xaf, 0xde, 0xf9, 0xab, 0x9a, 0xed, 0x04, 0xc8, 0xda, 0x14, 0x66, 0x2f, 0xbc,
	0xfb, 0x11, 0xd4, 0x94, 0x47, 0x9d, 0x30, 0xed, 0x75, 0x63, 0x01, 0xfb, 0x2d, 0xbc, 0x6f, 0x40,
	0xd7, 0xfe, 0x43, 0xef, 0x96, 0x9d, 0x20, 0x95, 0xac, 0x80, 0x89, 0x33, 0x75, 0x1e, 0xdd, 0xd8,
	0xde, 0xb4, 0x8f, 0xaf, 0x59, 0x01, 0xe1, 0x17, 0xd7, 0x0b, 0xce, 0x68, 0xbc, 0x31, 0xcb, 0xf9,
	0x3f, 0x1c, 0xef, 0xf6, 0x79, 0x17, 0xff, 0x39, 0xba, 0x94, 0x0c, 0xba, 0x72, 0xbe, 0x60, 0x39,
	0x2a, 0x30, 0x04, 0x87, 0xce, 0xb4, 0xcf, 0x56, 0x9f, 0x7e, 0xfe, 0xfa, 0xea, 0x3e, 0xf1, 0x97,
	0x5d, 0xc6, 0x1f, 0xfe, 0x59, 0xf1, 0x59, 0xd6, 0xe8, 0x5a, 0x15, 0x50, 0x69, 0xbc, 0xc0, 0xec,
	0xff, 0x5e, 0x8d, 0x17, 0x1f, 0x83, 0xbb, 0x07, 0x3a, 0x39, 0xd9, 0xf5, 0xa8, 0x14, 0x1a, 0x65,
	0xaa, 0x58, 0x7f, 0x76, 0xbd, 0x79, 0xa6, 0x8a, 0x8b, 0xbb, 0xad, 0x1f, 0x8c, 0xa7, 0xb6, 0xe9,
	0xee, 0xb3, 0x71, 0xde, 0xbe, 0xea, 0x45, 0x72, 0xb5, 0x63, 0x32, 0x47, 0xaa, 0xca, 0x71, 0x0e,
	0xf2, 0x78, 0x3d, 0x7c, 0xb2, 0x1d, 0xff, 0xa0, 0x4f, 0x2d, 0xf8, 0xe6, 0x5e, 0x8b, 0x28, 0xfd,
	0xee, 0x4e, 0x23, 0x23, 0x48, 0xb9, 0x46, 0x06, 0x76, 0x28, 0x0e, 0x51, 0x6f, 0xac, 0x0f, 0x96,
	0x92, 0x50, 0xae, 0x93, 0x81, 0x92, 0xc4, 0x61, 0x62, 0x29, 0xbf, 0xdd, 0xb9, 0x79, 0x27, 0x84,
	0x72, 0x4d, 0xc8, 0x40, 0x22, 0x24, 0x0e, 0x09, 0xb1, 0xb4, 0x77, 0xd7, 0x8f, 0x73, 0x3e, 0xfe,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x91, 0x6b, 0x3c, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupAudienceViewServiceClient is the client API for AdGroupAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAudienceViewServiceClient interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error)
}

type adGroupAudienceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupAudienceViewServiceClient(cc *grpc.ClientConn) AdGroupAudienceViewServiceClient {
	return &adGroupAudienceViewServiceClient{cc}
}

func (c *adGroupAudienceViewServiceClient) GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error) {
	out := new(resources.AdGroupAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAudienceViewService/GetAdGroupAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAudienceViewServiceServer is the server API for AdGroupAudienceViewService service.
type AdGroupAudienceViewServiceServer interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(context.Context, *GetAdGroupAudienceViewRequest) (*resources.AdGroupAudienceView, error)
}

// UnimplementedAdGroupAudienceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAudienceViewServiceServer struct {
}

func (*UnimplementedAdGroupAudienceViewServiceServer) GetAdGroupAudienceView(ctx context.Context, req *GetAdGroupAudienceViewRequest) (*resources.AdGroupAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupAudienceView not implemented")
}

func RegisterAdGroupAudienceViewServiceServer(s *grpc.Server, srv AdGroupAudienceViewServiceServer) {
	s.RegisterService(&_AdGroupAudienceViewService_serviceDesc, srv)
}

func _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAudienceViewService/GetAdGroupAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, req.(*GetAdGroupAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupAudienceViewService",
	HandlerType: (*AdGroupAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAudienceView",
			Handler:    _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_audience_view_service.proto",
}
