// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/video_service.proto

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

// Request message for [VideoService.GetVideo][google.ads.googleads.v2.services.VideoService.GetVideo].
type GetVideoRequest struct {
	// The resource name of the video to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVideoRequest) Reset()         { *m = GetVideoRequest{} }
func (m *GetVideoRequest) String() string { return proto.CompactTextString(m) }
func (*GetVideoRequest) ProtoMessage()    {}
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94f6646097cdc15, []int{0}
}

func (m *GetVideoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoRequest.Unmarshal(m, b)
}
func (m *GetVideoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoRequest.Marshal(b, m, deterministic)
}
func (m *GetVideoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoRequest.Merge(m, src)
}
func (m *GetVideoRequest) XXX_Size() int {
	return xxx_messageInfo_GetVideoRequest.Size(m)
}
func (m *GetVideoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoRequest proto.InternalMessageInfo

func (m *GetVideoRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetVideoRequest)(nil), "google.ads.googleads.v2.services.GetVideoRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/video_service.proto", fileDescriptor_c94f6646097cdc15)
}

var fileDescriptor_c94f6646097cdc15 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x25, 0x11, 0x44, 0x43, 0x45, 0xcc, 0x62, 0x89, 0x0e, 0xa5, 0x76, 0x28, 0x05, 0xef, 0x34,
	0x8a, 0xc3, 0x89, 0x43, 0xba, 0xd4, 0x49, 0x4a, 0x85, 0x0c, 0x12, 0x28, 0x67, 0x72, 0x84, 0x40,
	0x73, 0x57, 0xf3, 0x5d, 0xb3, 0x88, 0x8b, 0x7f, 0xc1, 0xc9, 0xd5, 0xd1, 0x3f, 0x22, 0x74, 0xf5,
	0x2f, 0x38, 0xf9, 0x0b, 0x1c, 0x25, 0xb9, 0x5c, 0x5a, 0x85, 0xd2, 0xed, 0xe5, 0xbe, 0xf7, 0xde,
	0xf7, 0xbe, 0x47, 0xac, 0xf3, 0x58, 0x88, 0x78, 0xc2, 0x30, 0x8d, 0x00, 0x2b, 0x58, 0xa0, 0xdc,
	0xc5, 0xc0, 0xb2, 0x3c, 0x09, 0x19, 0xe0, 0x3c, 0x89, 0x98, 0x18, 0x57, 0x9f, 0x68, 0x9a, 0x09,
	0x29, 0xec, 0x96, 0xa2, 0x22, 0x1a, 0x01, 0xaa, 0x55, 0x28, 0x77, 0x91, 0x56, 0x39, 0xc7, 0xab,
	0x7c, 0x33, 0x06, 0x62, 0x96, 0xd5, 0xc6, 0xca, 0xd0, 0x39, 0xd4, 0xf4, 0x69, 0x82, 0x29, 0xe7,
	0x42, 0x52, 0x99, 0x08, 0x0e, 0xd5, 0x74, 0x7f, 0x69, 0x1a, 0x4e, 0x12, 0xc6, 0xa5, 0x1a, 0xb4,
	0x2f, 0xac, 0xdd, 0x01, 0x93, 0x7e, 0x61, 0x34, 0x62, 0x0f, 0x33, 0x06, 0xd2, 0x3e, 0xb2, 0x76,
	0xf4, 0x8a, 0x31, 0xa7, 0x29, 0x6b, 0x1a, 0x2d, 0xa3, 0xbb, 0x3d, 0x6a, 0xe8, 0xc7, 0x1b, 0x9a,
	0x32, 0xf7, 0xc3, 0xb0, 0x1a, 0xa5, 0xea, 0x56, 0xe5, 0xb5, 0x5f, 0x0d, 0x6b, 0x4b, 0x3b, 0xd9,
	0xa7, 0x68, 0xdd, 0x79, 0xe8, 0xdf, 0x56, 0xa7, 0xbb, 0x52, 0x52, 0xdf, 0x8b, 0x4a, 0x41, 0xfb,
	0xe4, 0xf9, 0xf3, 0xeb, 0xc5, 0xec, 0xd9, 0xdd, 0xa2, 0x8c, 0xc7, 0x3f, 0x51, 0xaf, 0xc2, 0x19,
	0x48, 0x91, 0xb2, 0x0c, 0x70, 0x4f, 0xb5, 0x03, 0xb8, 0xf7, 0xe4, 0x1c, 0xcc, 0xbd, 0xe6, 0xc2,
	0xb2, 0x42, 0xd3, 0x04, 0x50, 0x28, 0xd2, 0xfe, 0x8f, 0x61, 0x75, 0x42, 0x91, 0xae, 0x4d, 0xdc,
	0xdf, 0x5b, 0xbe, 0x77, 0x58, 0xb4, 0x37, 0x34, 0xee, 0xae, 0x2b, 0x59, 0x2c, 0x26, 0x94, 0xc7,
	0x48, 0x64, 0x31, 0x8e, 0x19, 0x2f, 0xbb, 0xc5, 0x8b, 0x45, 0xab, 0x7f, 0x8e, 0x4b, 0x0d, 0xde,
	0xcc, 0x8d, 0x81, 0xe7, 0xbd, 0x9b, 0xad, 0x81, 0x32, 0xf4, 0x22, 0x40, 0x0a, 0x16, 0xc8, 0x77,
	0x51, 0xb5, 0x18, 0xe6, 0x9a, 0x12, 0x78, 0x11, 0x04, 0x35, 0x25, 0xf0, 0xdd, 0x40, 0x53, 0xbe,
	0xcd, 0x8e, 0x7a, 0x27, 0xc4, 0x8b, 0x80, 0x90, 0x9a, 0x44, 0x88, 0xef, 0x12, 0xa2, 0x69, 0xf7,
	0x9b, 0x65, 0xce, 0xb3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7f, 0x34, 0x0c, 0xc3, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoServiceClient interface {
	// Returns the requested video in full detail.
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error)
}

type videoServiceClient struct {
	cc *grpc.ClientConn
}

func NewVideoServiceClient(cc *grpc.ClientConn) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error) {
	out := new(resources.Video)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.VideoService/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
type VideoServiceServer interface {
	// Returns the requested video in full detail.
	GetVideo(context.Context, *GetVideoRequest) (*resources.Video, error)
}

// UnimplementedVideoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (*UnimplementedVideoServiceServer) GetVideo(ctx context.Context, req *GetVideoRequest) (*resources.Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}

func RegisterVideoServiceServer(s *grpc.Server, srv VideoServiceServer) {
	s.RegisterService(&_VideoService_serviceDesc, srv)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/video_service.proto",
}
