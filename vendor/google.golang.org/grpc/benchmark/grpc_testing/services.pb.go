// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package grpc_testing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x0f, 0x42, 0x16, 0x53, 0x64, 0x4f, 0xba, 0xfa, 0x00, 0x9e, 0x82, 0x54, 0x5f,
	0xc0, 0x06, 0x3d, 0x0a, 0x36, 0x54, 0x0f, 0x9e, 0xd6, 0x74, 0x88, 0x4b, 0x93, 0x99, 0x38, 0x33,
	0x11, 0x7c, 0x02, 0x1f, 0xc1, 0xd7, 0x15, 0xb3, 0x56, 0x6a, 0xc8, 0xcd, 0x1e, 0xe7, 0xff, 0x86,
	0x8f, 0xfd, 0x77, 0xd7, 0xcc, 0x04, 0xf8, 0x2d, 0x94, 0x20, 0x59, 0xcb, 0xa4, 0x64, 0x0f, 0x2b,
	0x6e, 0xcb, 0x4c, 0x41, 0x34, 0x60, 0xe5, 0x66, 0x0d, 0x88, 0xf8, 0x6a, 0x4b, 0x5d, 0x5a, 0x12,
	0x2a, 0x53, 0x1d, 0xc7, 0xf9, 0xc7, 0xd4, 0x1c, 0x2d, 0x00, 0xcb, 0x97, 0xc6, 0xf3, 0xa6, 0x88,
	0x22, 0x7b, 0x6b, 0x92, 0x15, 0x7a, 0x7e, 0xcf, 0x7d, 0x5d, 0xdb, 0xd3, 0x6c, 0xd7, 0x97, 0x15,
	0xa1, 0x69, 0x6b, 0x58, 0xc2, 0x6b, 0x07, 0xa2, 0xee, 0x6c, 0x1c, 0x4a, 0x4b, 0x28, 0x60, 0xef,
	0x4c, 0x5a, 0x28, 0x83, 0x6f, 0x02, 0x56, 0xff, 0x74, 0x9d, 0x4f, 0x2e, 0x26, 0xf6, 0xc9, 0xb8,
	0x15, 0x96, 0x84, 0xa2, 0xec, 0x03, 0xc2, 0x7a, 0x9f, 0xf2, 0xf9, 0xe7, 0xd4, 0xa4, 0x8f, 0xc4,
	0x1b, 0xe0, 0xed, 0x35, 0xdc, 0x98, 0x64, 0xd9, 0xe1, 0xf7, 0x04, 0x6c, 0x8f, 0x07, 0x82, 0x3e,
	0xbd, 0xe6, 0x4a, 0x9c, 0x1b, 0x23, 0x85, 0x7a, 0xed, 0xa4, 0x3f, 0x75, 0xd4, 0xe4, 0x75, 0x00,
	0xd4, 0xa1, 0x26, 0xa6, 0x63, 0x9a, 0x48, 0x76, 0x34, 0x0b, 0x93, 0xe4, 0xc4, 0x90, 0x53, 0x87,
	0x6a, 0x4f, 0x06, 0xcb, 0xc4, 0xbf, 0x4d, 0xdd, 0x18, 0xfa, 0x79, 0x90, 0x2b, 0x63, 0xee, 0xbb,
	0xa0, 0xb1, 0xa6, 0xb5, 0x7f, 0x37, 0x1f, 0x28, 0xac, 0xdd, 0x48, 0xf6, 0x7c, 0xd0, 0x7f, 0x95,
	0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xb4, 0x19, 0x36, 0x69, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BenchmarkServiceClient is the client API for BenchmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BenchmarkServiceClient interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_StreamingCallClient, error)
	// Unconstrainted streaming.
	// Both server and client keep sending & receiving simultaneously.
	UnconstrainedStreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_UnconstrainedStreamingCallClient, error)
}

type benchmarkServiceClient struct {
	cc *grpc.ClientConn
}

func NewBenchmarkServiceClient(cc *grpc.ClientConn) BenchmarkServiceClient {
	return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.BenchmarkService/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchmarkServiceClient) StreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_StreamingCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BenchmarkService_serviceDesc.Streams[0], "/grpc.testing.BenchmarkService/StreamingCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkServiceStreamingCallClient{stream}
	return x, nil
}

type BenchmarkService_StreamingCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type benchmarkServiceStreamingCallClient struct {
	grpc.ClientStream
}

func (x *benchmarkServiceStreamingCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *benchmarkServiceStreamingCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *benchmarkServiceClient) UnconstrainedStreamingCall(ctx context.Context, opts ...grpc.CallOption) (BenchmarkService_UnconstrainedStreamingCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BenchmarkService_serviceDesc.Streams[1], "/grpc.testing.BenchmarkService/UnconstrainedStreamingCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkServiceUnconstrainedStreamingCallClient{stream}
	return x, nil
}

type BenchmarkService_UnconstrainedStreamingCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type benchmarkServiceUnconstrainedStreamingCallClient struct {
	grpc.ClientStream
}

func (x *benchmarkServiceUnconstrainedStreamingCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *benchmarkServiceUnconstrainedStreamingCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BenchmarkServiceServer is the server API for BenchmarkService service.
type BenchmarkServiceServer interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	StreamingCall(BenchmarkService_StreamingCallServer) error
	// Unconstrainted streaming.
	// Both server and client keep sending & receiving simultaneously.
	UnconstrainedStreamingCall(BenchmarkService_UnconstrainedStreamingCallServer) error
}

// UnimplementedBenchmarkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBenchmarkServiceServer struct {
}

func (*UnimplementedBenchmarkServiceServer) UnaryCall(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (*UnimplementedBenchmarkServiceServer) StreamingCall(srv BenchmarkService_StreamingCallServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingCall not implemented")
}
func (*UnimplementedBenchmarkServiceServer) UnconstrainedStreamingCall(srv BenchmarkService_UnconstrainedStreamingCallServer) error {
	return status.Errorf(codes.Unimplemented, "method UnconstrainedStreamingCall not implemented")
}

func RegisterBenchmarkServiceServer(s *grpc.Server, srv BenchmarkServiceServer) {
	s.RegisterService(&_BenchmarkService_serviceDesc, srv)
}

func _BenchmarkService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.BenchmarkService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchmarkService_StreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).StreamingCall(&benchmarkServiceStreamingCallServer{stream})
}

type BenchmarkService_StreamingCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type benchmarkServiceStreamingCallServer struct {
	grpc.ServerStream
}

func (x *benchmarkServiceStreamingCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *benchmarkServiceStreamingCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BenchmarkService_UnconstrainedStreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkServiceServer).UnconstrainedStreamingCall(&benchmarkServiceUnconstrainedStreamingCallServer{stream})
}

type BenchmarkService_UnconstrainedStreamingCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type benchmarkServiceUnconstrainedStreamingCallServer struct {
	grpc.ServerStream
}

func (x *benchmarkServiceUnconstrainedStreamingCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *benchmarkServiceUnconstrainedStreamingCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BenchmarkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _BenchmarkService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingCall",
			Handler:       _BenchmarkService_StreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UnconstrainedStreamingCall",
			Handler:       _BenchmarkService_UnconstrainedStreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerServiceClient interface {
	// Start server with specified workload.
	// First request sent specifies the ServerConfig followed by ServerStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test server
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunServer(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunServerClient, error)
	// Start client with specified workload.
	// First request sent specifies the ClientConfig followed by ClientStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test client
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunClient(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunClientClient, error)
	// Just return the core count - unary call
	CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error)
	// Quit this worker
	QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
}

type workerServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkerServiceClient(cc *grpc.ClientConn) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) RunServer(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkerService_serviceDesc.Streams[0], "/grpc.testing.WorkerService/RunServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerServiceRunServerClient{stream}
	return x, nil
}

type WorkerService_RunServerClient interface {
	Send(*ServerArgs) error
	Recv() (*ServerStatus, error)
	grpc.ClientStream
}

type workerServiceRunServerClient struct {
	grpc.ClientStream
}

func (x *workerServiceRunServerClient) Send(m *ServerArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerServiceRunServerClient) Recv() (*ServerStatus, error) {
	m := new(ServerStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerServiceClient) RunClient(ctx context.Context, opts ...grpc.CallOption) (WorkerService_RunClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkerService_serviceDesc.Streams[1], "/grpc.testing.WorkerService/RunClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerServiceRunClientClient{stream}
	return x, nil
}

type WorkerService_RunClientClient interface {
	Send(*ClientArgs) error
	Recv() (*ClientStatus, error)
	grpc.ClientStream
}

type workerServiceRunClientClient struct {
	grpc.ClientStream
}

func (x *workerServiceRunClientClient) Send(m *ClientArgs) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerServiceRunClientClient) Recv() (*ClientStatus, error) {
	m := new(ClientStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerServiceClient) CoreCount(ctx context.Context, in *CoreRequest, opts ...grpc.CallOption) (*CoreResponse, error) {
	out := new(CoreResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.WorkerService/CoreCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) QuitWorker(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.testing.WorkerService/QuitWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
type WorkerServiceServer interface {
	// Start server with specified workload.
	// First request sent specifies the ServerConfig followed by ServerStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test server
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunServer(WorkerService_RunServerServer) error
	// Start client with specified workload.
	// First request sent specifies the ClientConfig followed by ClientStatus
	// response. After that, a "Mark" can be sent anytime to request the latest
	// stats. Closing the stream will initiate shutdown of the test client
	// and once the shutdown has finished, the OK status is sent to terminate
	// this RPC.
	RunClient(WorkerService_RunClientServer) error
	// Just return the core count - unary call
	CoreCount(context.Context, *CoreRequest) (*CoreResponse, error)
	// Quit this worker
	QuitWorker(context.Context, *Void) (*Void, error)
}

// UnimplementedWorkerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (*UnimplementedWorkerServiceServer) RunServer(srv WorkerService_RunServerServer) error {
	return status.Errorf(codes.Unimplemented, "method RunServer not implemented")
}
func (*UnimplementedWorkerServiceServer) RunClient(srv WorkerService_RunClientServer) error {
	return status.Errorf(codes.Unimplemented, "method RunClient not implemented")
}
func (*UnimplementedWorkerServiceServer) CoreCount(ctx context.Context, req *CoreRequest) (*CoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoreCount not implemented")
}
func (*UnimplementedWorkerServiceServer) QuitWorker(ctx context.Context, req *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitWorker not implemented")
}

func RegisterWorkerServiceServer(s *grpc.Server, srv WorkerServiceServer) {
	s.RegisterService(&_WorkerService_serviceDesc, srv)
}

func _WorkerService_RunServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).RunServer(&workerServiceRunServerServer{stream})
}

type WorkerService_RunServerServer interface {
	Send(*ServerStatus) error
	Recv() (*ServerArgs, error)
	grpc.ServerStream
}

type workerServiceRunServerServer struct {
	grpc.ServerStream
}

func (x *workerServiceRunServerServer) Send(m *ServerStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerServiceRunServerServer) Recv() (*ServerArgs, error) {
	m := new(ServerArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WorkerService_RunClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServiceServer).RunClient(&workerServiceRunClientServer{stream})
}

type WorkerService_RunClientServer interface {
	Send(*ClientStatus) error
	Recv() (*ClientArgs, error)
	grpc.ServerStream
}

type workerServiceRunClientServer struct {
	grpc.ServerStream
}

func (x *workerServiceRunClientServer) Send(m *ClientStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerServiceRunClientServer) Recv() (*ClientArgs, error) {
	m := new(ClientArgs)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WorkerService_CoreCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CoreCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.WorkerService/CoreCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CoreCount(ctx, req.(*CoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_QuitWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).QuitWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.WorkerService/QuitWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).QuitWorker(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoreCount",
			Handler:    _WorkerService_CoreCount_Handler,
		},
		{
			MethodName: "QuitWorker",
			Handler:    _WorkerService_QuitWorker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunServer",
			Handler:       _WorkerService_RunServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunClient",
			Handler:       _WorkerService_RunClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}
