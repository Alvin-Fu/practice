// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto

package clouderrorreporting

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A request for reporting an individual error event.
type ReportErrorEventRequest struct {
	// Required. The resource name of the Google Cloud Platform project. Written
	// as `projects/` plus the
	// [Google Cloud Platform project
	// ID](https://support.google.com/cloud/answer/6158840). Example:
	// `projects/my-project-123`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// Required. The error event to be reported.
	Event                *ReportedErrorEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReportErrorEventRequest) Reset()         { *m = ReportErrorEventRequest{} }
func (m *ReportErrorEventRequest) String() string { return proto.CompactTextString(m) }
func (*ReportErrorEventRequest) ProtoMessage()    {}
func (*ReportErrorEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_575af94d0209aede, []int{0}
}

func (m *ReportErrorEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportErrorEventRequest.Unmarshal(m, b)
}
func (m *ReportErrorEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportErrorEventRequest.Marshal(b, m, deterministic)
}
func (m *ReportErrorEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportErrorEventRequest.Merge(m, src)
}
func (m *ReportErrorEventRequest) XXX_Size() int {
	return xxx_messageInfo_ReportErrorEventRequest.Size(m)
}
func (m *ReportErrorEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportErrorEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportErrorEventRequest proto.InternalMessageInfo

func (m *ReportErrorEventRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *ReportErrorEventRequest) GetEvent() *ReportedErrorEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

// Response for reporting an individual error event.
// Data may be added to this message in the future.
type ReportErrorEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportErrorEventResponse) Reset()         { *m = ReportErrorEventResponse{} }
func (m *ReportErrorEventResponse) String() string { return proto.CompactTextString(m) }
func (*ReportErrorEventResponse) ProtoMessage()    {}
func (*ReportErrorEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_575af94d0209aede, []int{1}
}

func (m *ReportErrorEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportErrorEventResponse.Unmarshal(m, b)
}
func (m *ReportErrorEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportErrorEventResponse.Marshal(b, m, deterministic)
}
func (m *ReportErrorEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportErrorEventResponse.Merge(m, src)
}
func (m *ReportErrorEventResponse) XXX_Size() int {
	return xxx_messageInfo_ReportErrorEventResponse.Size(m)
}
func (m *ReportErrorEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportErrorEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportErrorEventResponse proto.InternalMessageInfo

// An error event which is reported to the Error Reporting system.
type ReportedErrorEvent struct {
	// Optional. Time when the event occurred.
	// If not provided, the time when the event was received by the
	// Error Reporting system will be used.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Required. The service context in which this error has occurred.
	ServiceContext *ServiceContext `protobuf:"bytes,2,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// Required. The error message.
	// If no `context.reportLocation` is provided, the message must contain a
	// header (typically consisting of the exception type name and an error
	// message) and an exception stack trace in one of the supported programming
	// languages and formats.
	// Supported languages are Java, Python, JavaScript, Ruby, C#, PHP, and Go.
	// Supported stack trace formats are:
	//
	// * **Java**: Must be the return value of
	// [`Throwable.printStackTrace()`](https://docs.oracle.com/javase/7/docs/api/java/lang/Throwable.html#printStackTrace%28%29).
	// * **Python**: Must be the return value of
	// [`traceback.format_exc()`](https://docs.python.org/2/library/traceback.html#traceback.format_exc).
	// * **JavaScript**: Must be the value of
	// [`error.stack`](https://github.com/v8/v8/wiki/Stack-Trace-API) as returned
	// by V8.
	// * **Ruby**: Must contain frames returned by
	// [`Exception.backtrace`](https://ruby-doc.org/core-2.2.0/Exception.html#method-i-backtrace).
	// * **C#**: Must be the return value of
	// [`Exception.ToString()`](https://msdn.microsoft.com/en-us/library/system.exception.tostring.aspx).
	// * **PHP**: Must start with `PHP (Notice|Parse error|Fatal error|Warning)`
	// and contain the result of
	// [`(string)$exception`](http://php.net/manual/en/exception.tostring.php).
	// * **Go**: Must be the return value of
	// [`runtime.Stack()`](https://golang.org/pkg/runtime/debug/#Stack).
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Optional. A description of the context in which the error occurred.
	Context              *ErrorContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReportedErrorEvent) Reset()         { *m = ReportedErrorEvent{} }
func (m *ReportedErrorEvent) String() string { return proto.CompactTextString(m) }
func (*ReportedErrorEvent) ProtoMessage()    {}
func (*ReportedErrorEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_575af94d0209aede, []int{2}
}

func (m *ReportedErrorEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportedErrorEvent.Unmarshal(m, b)
}
func (m *ReportedErrorEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportedErrorEvent.Marshal(b, m, deterministic)
}
func (m *ReportedErrorEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportedErrorEvent.Merge(m, src)
}
func (m *ReportedErrorEvent) XXX_Size() int {
	return xxx_messageInfo_ReportedErrorEvent.Size(m)
}
func (m *ReportedErrorEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportedErrorEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReportedErrorEvent proto.InternalMessageInfo

func (m *ReportedErrorEvent) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *ReportedErrorEvent) GetServiceContext() *ServiceContext {
	if m != nil {
		return m.ServiceContext
	}
	return nil
}

func (m *ReportedErrorEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ReportedErrorEvent) GetContext() *ErrorContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*ReportErrorEventRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportErrorEventRequest")
	proto.RegisterType((*ReportErrorEventResponse)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportErrorEventResponse")
	proto.RegisterType((*ReportedErrorEvent)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent")
}

func init() {
	proto.RegisterFile("google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto", fileDescriptor_575af94d0209aede)
}

var fileDescriptor_575af94d0209aede = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0x5d, 0xa0, 0xea, 0x14, 0x01, 0x32, 0x8b, 0x06, 0x0b, 0x44, 0x15, 0x36, 0x15, 0xa5,
	0x1e, 0x92, 0x6e, 0x68, 0x2a, 0x40, 0x4e, 0x89, 0xba, 0x43, 0x91, 0xa9, 0x22, 0x54, 0x45, 0x58,
	0x13, 0xe7, 0xc6, 0x31, 0xb2, 0x67, 0xcc, 0xcc, 0x24, 0x41, 0x42, 0x6c, 0xf8, 0x05, 0xfe, 0x82,
	0x2f, 0xe0, 0x1b, 0x2a, 0xb1, 0x81, 0x1d, 0xab, 0x2e, 0xe0, 0x13, 0x90, 0x10, 0x62, 0x81, 0x3c,
	0x63, 0xb7, 0xce, 0x63, 0x41, 0x60, 0xe9, 0xb9, 0xe7, 0x9e, 0x73, 0xee, 0xcb, 0xe8, 0x30, 0x64,
	0x2c, 0x8c, 0x01, 0xf7, 0x61, 0x2c, 0x19, 0x8b, 0x05, 0x0e, 0x62, 0x36, 0xea, 0x03, 0xe7, 0x8c,
	0x73, 0x48, 0x19, 0x97, 0x11, 0x0d, 0xf1, 0xb8, 0xd6, 0x03, 0x49, 0x6a, 0x58, 0xbf, 0xf8, 0x2a,
	0x2a, 0x7c, 0x01, 0x7c, 0x1c, 0x05, 0xe0, 0xa4, 0x9c, 0x49, 0x66, 0x6d, 0x6b, 0x22, 0xa7, 0x20,
	0x72, 0x16, 0x10, 0x39, 0x39, 0x91, 0x7d, 0x33, 0x57, 0x25, 0x69, 0x84, 0x09, 0xa5, 0x4c, 0x12,
	0x19, 0x31, 0x2a, 0x34, 0x95, 0xbd, 0x51, 0x8a, 0x06, 0x71, 0x04, 0x54, 0xe6, 0x81, 0xdb, 0xa5,
	0xc0, 0x20, 0x82, 0xb8, 0xef, 0xf7, 0x60, 0x48, 0xc6, 0x11, 0xe3, 0x39, 0xe0, 0x46, 0x09, 0xc0,
	0x41, 0xb0, 0x11, 0x2f, 0xfc, 0xd9, 0x0f, 0x96, 0x29, 0x34, 0x60, 0x49, 0xc2, 0xe8, 0x8c, 0xaa,
	0xfa, 0xea, 0x8d, 0x06, 0x58, 0x46, 0x09, 0x08, 0x49, 0x92, 0x54, 0x03, 0xaa, 0x9f, 0x0c, 0xb4,
	0xe1, 0x29, 0x8e, 0x56, 0x46, 0xd7, 0x1a, 0x03, 0x95, 0x1e, 0xbc, 0x1a, 0x81, 0x90, 0x56, 0x07,
	0x5d, 0x4e, 0x39, 0x7b, 0x09, 0x81, 0xf4, 0x29, 0x49, 0xa0, 0x62, 0x6c, 0x1a, 0x5b, 0x6b, 0xcd,
	0xdd, 0x53, 0xd7, 0xfc, 0xe5, 0xee, 0xa0, 0x6d, 0x65, 0xa2, 0x70, 0x9a, 0x10, 0x4a, 0x42, 0xe0,
	0x8e, 0xd6, 0x23, 0x69, 0x24, 0x9c, 0x80, 0x25, 0xb8, 0xad, 0xf3, 0xbd, 0xf5, 0x9c, 0xe8, 0x29,
	0x49, 0xc0, 0x3a, 0x46, 0x17, 0x21, 0xd3, 0xa9, 0x98, 0x9b, 0xc6, 0xd6, 0x7a, 0xfd, 0xb1, 0xb3,
	0x44, 0xfb, 0x1d, 0x6d, 0x16, 0xfa, 0xe7, 0x76, 0x9b, 0x2b, 0xa7, 0xae, 0xe9, 0x69, 0xca, 0xaa,
	0x8d, 0x2a, 0xf3, 0xe5, 0x88, 0x94, 0x51, 0x01, 0xd5, 0x8f, 0x26, 0xb2, 0xe6, 0xd3, 0xad, 0x47,
	0x08, 0xa9, 0x5c, 0x3f, 0xeb, 0x8d, 0x2a, 0x72, 0xbd, 0x6e, 0x17, 0x9e, 0x8a, 0xc6, 0x39, 0x47,
	0x45, 0xe3, 0x32, 0x39, 0xc3, 0x5b, 0x53, 0x29, 0xd9, 0xa3, 0x15, 0xa1, 0xab, 0xf9, 0x3a, 0xf9,
	0x01, 0xa3, 0x12, 0x5e, 0x17, 0x85, 0xed, 0x2f, 0x55, 0xd8, 0x33, 0xcd, 0x71, 0xa0, 0x29, 0x74,
	0x51, 0x57, 0xc4, 0xd4, 0xa3, 0x75, 0x0b, 0xad, 0x26, 0x20, 0x04, 0x09, 0xa1, 0xb2, 0xa2, 0x86,
	0xa1, 0x50, 0xc5, 0x9b, 0xf5, 0x1c, 0xad, 0x16, 0x0e, 0x2e, 0x28, 0x07, 0x7b, 0x4b, 0x39, 0x50,
	0x3d, 0x29, 0xe9, 0x1b, 0x5e, 0x41, 0x57, 0xff, 0x6e, 0xa2, 0xeb, 0xa5, 0xbe, 0x8a, 0xdc, 0xac,
	0xf5, 0xc3, 0x40, 0xd7, 0x66, 0xfb, 0x6d, 0x3d, 0xf9, 0x87, 0x81, 0xce, 0x6d, 0x9f, 0xdd, 0xfa,
	0x4f, 0x96, 0x7c, 0xe8, 0x47, 0x5f, 0x5d, 0xab, 0xbc, 0xc5, 0xf7, 0xd4, 0xe0, 0xde, 0x7d, 0xf9,
	0xf6, 0xde, 0xdc, 0xab, 0xde, 0x3f, 0xbb, 0x9a, 0x37, 0x65, 0xcc, 0xc3, 0xfc, 0x43, 0xe0, 0xbb,
	0x6f, 0xb1, 0xc2, 0x8b, 0x86, 0x96, 0x6c, 0xe8, 0x35, 0xb3, 0x3b, 0x27, 0x6e, 0x75, 0x91, 0x9d,
	0xe9, 0xf5, 0xff, 0xec, 0x3a, 0x43, 0x29, 0x53, 0xd1, 0xc0, 0x78, 0x32, 0x99, 0xcc, 0xde, 0x06,
	0x19, 0xc9, 0xa1, 0xbe, 0xe4, 0x9d, 0x34, 0x26, 0x72, 0xc0, 0x78, 0xd2, 0xfc, 0x6d, 0xa0, 0xec,
	0x80, 0x97, 0x29, 0xbd, 0x59, 0x59, 0x30, 0x98, 0x76, 0xb6, 0xb6, 0x6d, 0xe3, 0xf8, 0x45, 0x4e,
	0x14, 0xb2, 0x98, 0xd0, 0xd0, 0x61, 0x3c, 0xc4, 0x21, 0x50, 0xb5, 0xd4, 0xf8, 0xdc, 0xce, 0x5f,
	0xfd, 0x58, 0xf6, 0x17, 0xc4, 0x7e, 0x1a, 0xc6, 0x07, 0xf3, 0xce, 0xa1, 0xd6, 0x38, 0xc8, 0xe2,
	0x7a, 0x89, 0xbc, 0x33, 0x93, 0x9d, 0x5a, 0x33, 0x4b, 0x3e, 0x29, 0x50, 0x5d, 0x85, 0xea, 0x4e,
	0xa3, 0xba, 0x1d, 0x2d, 0xd1, 0xbb, 0xa4, 0x9c, 0xed, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x93,
	0xd3, 0x64, 0x88, 0xda, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportErrorsServiceClient is the client API for ReportErrorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportErrorsServiceClient interface {
	// Report an individual error event.
	//
	// This endpoint accepts **either** an OAuth token,
	// **or** an [API key](https://support.google.com/cloud/answer/6158862)
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	//
	// `POST
	// https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456`
	ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error)
}

type reportErrorsServiceClient struct {
	cc *grpc.ClientConn
}

func NewReportErrorsServiceClient(cc *grpc.ClientConn) ReportErrorsServiceClient {
	return &reportErrorsServiceClient{cc}
}

func (c *reportErrorsServiceClient) ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error) {
	out := new(ReportErrorEventResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportErrorsServiceServer is the server API for ReportErrorsService service.
type ReportErrorsServiceServer interface {
	// Report an individual error event.
	//
	// This endpoint accepts **either** an OAuth token,
	// **or** an [API key](https://support.google.com/cloud/answer/6158862)
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	//
	// `POST
	// https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456`
	ReportErrorEvent(context.Context, *ReportErrorEventRequest) (*ReportErrorEventResponse, error)
}

// UnimplementedReportErrorsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportErrorsServiceServer struct {
}

func (*UnimplementedReportErrorsServiceServer) ReportErrorEvent(ctx context.Context, req *ReportErrorEventRequest) (*ReportErrorEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportErrorEvent not implemented")
}

func RegisterReportErrorsServiceServer(s *grpc.Server, srv ReportErrorsServiceServer) {
	s.RegisterService(&_ReportErrorsService_serviceDesc, srv)
}

func _ReportErrorsService_ReportErrorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportErrorEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, req.(*ReportErrorEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportErrorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ReportErrorsService",
	HandlerType: (*ReportErrorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportErrorEvent",
			Handler:    _ReportErrorsService_ReportErrorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto",
}
