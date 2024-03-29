// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/google_ads_field_service.proto

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
// [GoogleAdsFieldService.GetGoogleAdsField][google.ads.googleads.v1.services.GoogleAdsFieldService.GetGoogleAdsField].
type GetGoogleAdsFieldRequest struct {
	// The resource name of the field to get.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGoogleAdsFieldRequest) Reset()         { *m = GetGoogleAdsFieldRequest{} }
func (m *GetGoogleAdsFieldRequest) String() string { return proto.CompactTextString(m) }
func (*GetGoogleAdsFieldRequest) ProtoMessage()    {}
func (*GetGoogleAdsFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc4a77920e176fd, []int{0}
}

func (m *GetGoogleAdsFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGoogleAdsFieldRequest.Unmarshal(m, b)
}
func (m *GetGoogleAdsFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGoogleAdsFieldRequest.Marshal(b, m, deterministic)
}
func (m *GetGoogleAdsFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGoogleAdsFieldRequest.Merge(m, src)
}
func (m *GetGoogleAdsFieldRequest) XXX_Size() int {
	return xxx_messageInfo_GetGoogleAdsFieldRequest.Size(m)
}
func (m *GetGoogleAdsFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGoogleAdsFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGoogleAdsFieldRequest proto.InternalMessageInfo

func (m *GetGoogleAdsFieldRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [GoogleAdsFieldService.SearchGoogleAdsFields][google.ads.googleads.v1.services.GoogleAdsFieldService.SearchGoogleAdsFields].
type SearchGoogleAdsFieldsRequest struct {
	// The query string.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Token of the page to retrieve. If not specified, the first page of
	// results will be returned. Use the value obtained from `next_page_token`
	// in the previous response in order to request the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Number of elements to retrieve in a single page.
	// When too large a page is requested, the server may decide to further
	// limit the number of returned resources.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchGoogleAdsFieldsRequest) Reset()         { *m = SearchGoogleAdsFieldsRequest{} }
func (m *SearchGoogleAdsFieldsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchGoogleAdsFieldsRequest) ProtoMessage()    {}
func (*SearchGoogleAdsFieldsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc4a77920e176fd, []int{1}
}

func (m *SearchGoogleAdsFieldsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGoogleAdsFieldsRequest.Unmarshal(m, b)
}
func (m *SearchGoogleAdsFieldsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGoogleAdsFieldsRequest.Marshal(b, m, deterministic)
}
func (m *SearchGoogleAdsFieldsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGoogleAdsFieldsRequest.Merge(m, src)
}
func (m *SearchGoogleAdsFieldsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchGoogleAdsFieldsRequest.Size(m)
}
func (m *SearchGoogleAdsFieldsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGoogleAdsFieldsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGoogleAdsFieldsRequest proto.InternalMessageInfo

func (m *SearchGoogleAdsFieldsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchGoogleAdsFieldsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchGoogleAdsFieldsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for
// [GoogleAdsFieldService.SearchGoogleAdsFields][google.ads.googleads.v1.services.GoogleAdsFieldService.SearchGoogleAdsFields].
type SearchGoogleAdsFieldsResponse struct {
	// The list of fields that matched the query.
	Results []*resources.GoogleAdsField `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Pagination token used to retrieve the next page of results. Pass the
	// content of this string as the `page_token` attribute of the next request.
	// `next_page_token` is not returned for the last page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of results that match the query ignoring the LIMIT clause.
	TotalResultsCount    int64    `protobuf:"varint,3,opt,name=total_results_count,json=totalResultsCount,proto3" json:"total_results_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchGoogleAdsFieldsResponse) Reset()         { *m = SearchGoogleAdsFieldsResponse{} }
func (m *SearchGoogleAdsFieldsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchGoogleAdsFieldsResponse) ProtoMessage()    {}
func (*SearchGoogleAdsFieldsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc4a77920e176fd, []int{2}
}

func (m *SearchGoogleAdsFieldsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchGoogleAdsFieldsResponse.Unmarshal(m, b)
}
func (m *SearchGoogleAdsFieldsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchGoogleAdsFieldsResponse.Marshal(b, m, deterministic)
}
func (m *SearchGoogleAdsFieldsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchGoogleAdsFieldsResponse.Merge(m, src)
}
func (m *SearchGoogleAdsFieldsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchGoogleAdsFieldsResponse.Size(m)
}
func (m *SearchGoogleAdsFieldsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchGoogleAdsFieldsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchGoogleAdsFieldsResponse proto.InternalMessageInfo

func (m *SearchGoogleAdsFieldsResponse) GetResults() []*resources.GoogleAdsField {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SearchGoogleAdsFieldsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *SearchGoogleAdsFieldsResponse) GetTotalResultsCount() int64 {
	if m != nil {
		return m.TotalResultsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GetGoogleAdsFieldRequest)(nil), "google.ads.googleads.v1.services.GetGoogleAdsFieldRequest")
	proto.RegisterType((*SearchGoogleAdsFieldsRequest)(nil), "google.ads.googleads.v1.services.SearchGoogleAdsFieldsRequest")
	proto.RegisterType((*SearchGoogleAdsFieldsResponse)(nil), "google.ads.googleads.v1.services.SearchGoogleAdsFieldsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/google_ads_field_service.proto", fileDescriptor_bdc4a77920e176fd)
}

var fileDescriptor_bdc4a77920e176fd = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0x29, 0xab, 0xee, 0xe8, 0x22, 0x3b, 0xba, 0x50, 0xe2, 0x2e, 0x96, 0xb8, 0xab, 0x65,
	0xc1, 0x09, 0x59, 0x2f, 0x32, 0xa2, 0x25, 0x2b, 0x58, 0x41, 0x90, 0xd2, 0x4a, 0x0f, 0x52, 0x08,
	0x63, 0xf3, 0x8c, 0xc1, 0x76, 0x26, 0x9b, 0x99, 0x16, 0x5d, 0xf1, 0xa0, 0x77, 0x4f, 0xfe, 0x03,
	0x8f, 0xde, 0xfc, 0x05, 0x5e, 0x3c, 0x79, 0xf5, 0x2f, 0x78, 0xf2, 0x57, 0xc8, 0x64, 0x3a, 0xc5,
	0xee, 0x36, 0x16, 0xbd, 0x4d, 0xde, 0xf7, 0xbe, 0xef, 0x7b, 0xf3, 0xde, 0x9b, 0xa0, 0x56, 0x2a,
	0x44, 0x3a, 0x82, 0x80, 0x25, 0x32, 0x30, 0x47, 0x7d, 0x9a, 0x86, 0x81, 0x84, 0x62, 0x9a, 0x0d,
	0xc1, 0x46, 0x63, 0x96, 0xc8, 0xf8, 0x79, 0x06, 0xa3, 0x24, 0x9e, 0x21, 0x24, 0x2f, 0x84, 0x12,
	0xb8, 0x61, 0x70, 0xc2, 0x12, 0x49, 0xe6, 0x02, 0x64, 0x1a, 0x12, 0x2b, 0xe0, 0xdd, 0xae, 0xb2,
	0x28, 0x40, 0x8a, 0x49, 0xb1, 0xcc, 0xc3, 0x68, 0x7b, 0xdb, 0x96, 0x99, 0x67, 0x01, 0xe3, 0x5c,
	0x28, 0xa6, 0x32, 0xc1, 0xa5, 0x41, 0xfd, 0x16, 0xaa, 0xb7, 0x41, 0xb5, 0xcb, 0x94, 0x28, 0x91,
	0x0f, 0x34, 0xb1, 0x0b, 0x47, 0x13, 0x90, 0x0a, 0x5f, 0x43, 0x1b, 0x56, 0x3d, 0xe6, 0x6c, 0x0c,
	0x75, 0xa7, 0xe1, 0x34, 0xd7, 0xbb, 0x17, 0x6c, 0xf0, 0x31, 0x1b, 0x83, 0x9f, 0xa3, 0xed, 0x1e,
	0xb0, 0x62, 0xf8, 0x62, 0x51, 0x43, 0x5a, 0x91, 0xcb, 0x68, 0xed, 0x68, 0x02, 0xc5, 0xeb, 0x19,
	0xd9, 0x7c, 0xe0, 0x1d, 0x84, 0x72, 0x96, 0x42, 0xac, 0xc4, 0x4b, 0xe0, 0x75, 0xb7, 0x84, 0xd6,
	0x75, 0xe4, 0x89, 0x0e, 0xe0, 0x2b, 0xa8, 0xfc, 0x88, 0x65, 0x76, 0x0c, 0xf5, 0x5a, 0xc3, 0x69,
	0xae, 0x75, 0xcf, 0xe9, 0x40, 0x2f, 0x3b, 0x06, 0xff, 0x9b, 0x83, 0x76, 0x2a, 0x2c, 0x65, 0x2e,
	0xb8, 0x04, 0xfc, 0x08, 0x9d, 0x2d, 0x40, 0x4e, 0x46, 0x4a, 0xd6, 0x9d, 0x46, 0xad, 0x79, 0xfe,
	0x20, 0x24, 0x55, 0x0d, 0x9e, 0xb7, 0x8f, 0x9c, 0xe8, 0x81, 0x55, 0xc0, 0xd7, 0xd1, 0x45, 0x0e,
	0xaf, 0x54, 0x7c, 0xaa, 0xde, 0x0d, 0x1d, 0xee, 0xcc, 0x6b, 0x26, 0xe8, 0x92, 0x12, 0x8a, 0x8d,
	0xe2, 0x19, 0x31, 0x1e, 0x8a, 0x09, 0x57, 0x65, 0xf5, 0xb5, 0xee, 0x66, 0x09, 0x75, 0x0d, 0x72,
	0x5f, 0x03, 0x07, 0x1f, 0x6a, 0x68, 0x6b, 0xd1, 0xb3, 0x67, 0x86, 0x8d, 0xbf, 0x38, 0x68, 0xf3,
	0xd4, 0x50, 0x30, 0x25, 0xab, 0x96, 0x84, 0x54, 0x4d, 0xd2, 0xfb, 0xf7, 0xfb, 0xfb, 0x37, 0xdf,
	0xff, 0xf8, 0xf9, 0xd1, 0xbd, 0x81, 0xf7, 0xf4, 0x92, 0xbd, 0x59, 0xd8, 0x83, 0xbb, 0xe9, 0x62,
	0xdf, 0x83, 0xfd, 0xb7, 0xf8, 0xab, 0x83, 0xb6, 0x96, 0x0e, 0x05, 0xdf, 0x5b, 0x5d, 0xf7, 0xdf,
	0x16, 0xc8, 0x6b, 0xfd, 0x37, 0xdf, 0x6c, 0x83, 0xbf, 0x57, 0xde, 0xe4, 0xaa, 0xef, 0xe9, 0x9b,
	0x9c, 0x28, 0x9d, 0xca, 0x92, 0x4a, 0x9d, 0xfd, 0xc3, 0x77, 0x2e, 0xda, 0x1d, 0x8a, 0xf1, 0x4a,
	0xb7, 0x43, 0x6f, 0xe9, 0xd4, 0x3a, 0xfa, 0x39, 0x75, 0x9c, 0xa7, 0x0f, 0x67, 0xfc, 0x54, 0x8c,
	0x18, 0x4f, 0x89, 0x28, 0xd2, 0x20, 0x05, 0x5e, 0x3e, 0x36, 0xfb, 0x70, 0xf3, 0x4c, 0x56, 0xff,
	0x2a, 0xee, 0xd8, 0xc3, 0x27, 0xb7, 0xd6, 0x8e, 0xa2, 0xcf, 0x6e, 0xc3, 0xd8, 0x91, 0x28, 0xf9,
	0x63, 0x46, 0xa4, 0x1f, 0x92, 0x99, 0xb1, 0xfc, 0x6e, 0x53, 0x06, 0x51, 0x22, 0x07, 0xf3, 0x94,
	0x41, 0x3f, 0x1c, 0xd8, 0x94, 0x5f, 0xee, 0xae, 0x89, 0x53, 0x1a, 0x25, 0x92, 0xd2, 0x79, 0x12,
	0xa5, 0xfd, 0x90, 0x52, 0x9b, 0xf6, 0xec, 0x4c, 0x59, 0xe7, 0xad, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xea, 0xe2, 0xe3, 0x8f, 0xd1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoogleAdsFieldServiceClient is the client API for GoogleAdsFieldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoogleAdsFieldServiceClient interface {
	// Returns just the requested field.
	GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error)
}

type googleAdsFieldServiceClient struct {
	cc *grpc.ClientConn
}

func NewGoogleAdsFieldServiceClient(cc *grpc.ClientConn) GoogleAdsFieldServiceClient {
	return &googleAdsFieldServiceClient{cc}
}

func (c *googleAdsFieldServiceClient) GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error) {
	out := new(resources.GoogleAdsField)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.GoogleAdsFieldService/GetGoogleAdsField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleAdsFieldServiceClient) SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error) {
	out := new(SearchGoogleAdsFieldsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.GoogleAdsFieldService/SearchGoogleAdsFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleAdsFieldServiceServer is the server API for GoogleAdsFieldService service.
type GoogleAdsFieldServiceServer interface {
	// Returns just the requested field.
	GetGoogleAdsField(context.Context, *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	SearchGoogleAdsFields(context.Context, *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error)
}

// UnimplementedGoogleAdsFieldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGoogleAdsFieldServiceServer struct {
}

func (*UnimplementedGoogleAdsFieldServiceServer) GetGoogleAdsField(ctx context.Context, req *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleAdsField not implemented")
}
func (*UnimplementedGoogleAdsFieldServiceServer) SearchGoogleAdsFields(ctx context.Context, req *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGoogleAdsFields not implemented")
}

func RegisterGoogleAdsFieldServiceServer(s *grpc.Server, srv GoogleAdsFieldServiceServer) {
	s.RegisterService(&_GoogleAdsFieldService_serviceDesc, srv)
}

func _GoogleAdsFieldService_GetGoogleAdsField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleAdsFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.GoogleAdsFieldService/GetGoogleAdsField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, req.(*GetGoogleAdsFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleAdsFieldService_SearchGoogleAdsFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGoogleAdsFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.GoogleAdsFieldService/SearchGoogleAdsFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, req.(*SearchGoogleAdsFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoogleAdsFieldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.GoogleAdsFieldService",
	HandlerType: (*GoogleAdsFieldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoogleAdsField",
			Handler:    _GoogleAdsFieldService_GetGoogleAdsField_Handler,
		},
		{
			MethodName: "SearchGoogleAdsFields",
			Handler:    _GoogleAdsFieldService_SearchGoogleAdsFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/google_ads_field_service.proto",
}
