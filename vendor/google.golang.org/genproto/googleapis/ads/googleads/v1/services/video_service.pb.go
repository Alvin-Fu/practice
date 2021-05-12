// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/video_service.proto

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
// [VideoService.GetVideo][google.ads.googleads.v1.services.VideoService.GetVideo].
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
	return fileDescriptor_7856cda63e5cca90, []int{0}
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
	proto.RegisterType((*GetVideoRequest)(nil), "google.ads.googleads.v1.services.GetVideoRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/video_service.proto", fileDescriptor_7856cda63e5cca90)
}

var fileDescriptor_7856cda63e5cca90 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x49, 0xfe, 0xf0, 0x47, 0x43, 0x45, 0xcc, 0x24, 0xc5, 0xa1, 0xd4, 0x0e, 0xa5, 0xe0,
	0x9d, 0x51, 0x71, 0x38, 0x71, 0x48, 0x97, 0x3a, 0x49, 0xa9, 0x90, 0x41, 0x02, 0xe5, 0x6c, 0x5e,
	0x42, 0xa0, 0xc9, 0x5b, 0xf3, 0x5e, 0xb3, 0x88, 0x8b, 0x5f, 0xc1, 0xc9, 0xd5, 0xb1, 0x1f, 0xc5,
	0xd5, 0xaf, 0xe0, 0xe4, 0x27, 0x70, 0x94, 0xe4, 0x7a, 0x41, 0x85, 0xd0, 0xed, 0xc9, 0xe5, 0xf7,
	0x3c, 0xf7, 0xbe, 0xcf, 0x39, 0x67, 0x31, 0x62, 0x3c, 0x07, 0x2e, 0x23, 0xe2, 0x5a, 0x96, 0xaa,
	0xf0, 0x38, 0x41, 0x5e, 0x24, 0x33, 0x20, 0x5e, 0x24, 0x11, 0xe0, 0x74, 0xfd, 0xc9, 0x16, 0x39,
	0x2a, 0x74, 0x3b, 0x1a, 0x65, 0x32, 0x22, 0x56, 0xbb, 0x58, 0xe1, 0x31, 0xe3, 0x6a, 0x1f, 0x35,
	0xe5, 0xe6, 0x40, 0xb8, 0xcc, 0xeb, 0x60, 0x1d, 0xd8, 0x3e, 0x30, 0xf8, 0x22, 0xe1, 0x32, 0xcb,
	0x50, 0x49, 0x95, 0x60, 0x46, 0xfa, 0x6f, 0xf7, 0xdc, 0xd9, 0x1d, 0x81, 0x0a, 0x4a, 0x7e, 0x02,
	0xf7, 0x4b, 0x20, 0xe5, 0x1e, 0x3a, 0x3b, 0x26, 0x69, 0x9a, 0xc9, 0x14, 0xf6, 0xad, 0x8e, 0xd5,
	0xdf, 0x9e, 0xb4, 0xcc, 0xe1, 0xb5, 0x4c, 0xe1, 0x64, 0x65, 0x39, 0xad, 0xca, 0x75, 0xa3, 0xc7,
	0x72, 0x5f, 0x2c, 0x67, 0xcb, 0x24, 0xb9, 0x1e, 0xdb, 0xb4, 0x05, 0xfb, 0x73, 0x6b, 0xbb, 0xdf,
	0x68, 0xa9, 0xd7, 0x62, 0x95, 0xa1, 0x7b, 0xfc, 0xf4, 0xfe, 0xf1, 0x6c, 0x0f, 0xdc, 0x7e, 0xb9,
	0xf3, 0xc3, 0xaf, 0x51, 0x2f, 0x67, 0x4b, 0x52, 0x98, 0x42, 0x4e, 0x7c, 0xa0, 0x4b, 0x20, 0x3e,
	0x78, 0x1c, 0x7e, 0x59, 0x4e, 0x6f, 0x86, 0xe9, 0xc6, 0xa1, 0x86, 0x7b, 0x3f, 0x57, 0x1a, 0x97,
	0x05, 0x8d, 0xad, 0xdb, 0xab, 0xb5, 0x2d, 0xc6, 0xb9, 0xcc, 0x62, 0x86, 0x79, 0xcc, 0x63, 0xc8,
	0xaa, 0xfa, 0x4c, 0xff, 0x8b, 0x84, 0x9a, 0x9f, 0xf9, 0xc2, 0x88, 0x57, 0xfb, 0xdf, 0xc8, 0xf7,
	0x57, 0x76, 0x67, 0xa4, 0x03, 0xfd, 0x88, 0x98, 0x96, 0xa5, 0x0a, 0x3c, 0xb6, 0xbe, 0x98, 0xde,
	0x0c, 0x12, 0xfa, 0x11, 0x85, 0x35, 0x12, 0x06, 0x5e, 0x68, 0x90, 0x4f, 0xbb, 0xa7, 0xcf, 0x85,
	0xf0, 0x23, 0x12, 0xa2, 0x86, 0x84, 0x08, 0x3c, 0x21, 0x0c, 0x76, 0xf7, 0xbf, 0x9a, 0xf3, 0xf4,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5c, 0xd6, 0x52, 0x8d, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.VideoService/GetVideo", in, out, opts...)
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
		FullMethod: "/google.ads.googleads.v1.services.VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/video_service.proto",
}
