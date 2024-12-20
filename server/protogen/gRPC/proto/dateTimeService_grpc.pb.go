// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DateTimeServiceClient is the client API for DateTimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DateTimeServiceClient interface {
	LaunchCronJob(ctx context.Context, in *TriggerTimeRequest, opts ...grpc.CallOption) (*TriggerTimeResponse, error)
}

type dateTimeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDateTimeServiceClient(cc grpc.ClientConnInterface) DateTimeServiceClient {
	return &dateTimeServiceClient{cc}
}

func (c *dateTimeServiceClient) LaunchCronJob(ctx context.Context, in *TriggerTimeRequest, opts ...grpc.CallOption) (*TriggerTimeResponse, error) {
	out := new(TriggerTimeResponse)
	err := c.cc.Invoke(ctx, "/dateTime.DateTimeService/LaunchCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DateTimeServiceServer is the server API for DateTimeService service.
// All implementations must embed UnimplementedDateTimeServiceServer
// for forward compatibility
type DateTimeServiceServer interface {
	LaunchCronJob(context.Context, *TriggerTimeRequest) (*TriggerTimeResponse, error)
	mustEmbedUnimplementedDateTimeServiceServer()
}

// UnimplementedDateTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDateTimeServiceServer struct {
}

func (UnimplementedDateTimeServiceServer) LaunchCronJob(context.Context, *TriggerTimeRequest) (*TriggerTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchCronJob not implemented")
}
func (UnimplementedDateTimeServiceServer) mustEmbedUnimplementedDateTimeServiceServer() {}

// UnsafeDateTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DateTimeServiceServer will
// result in compilation errors.
type UnsafeDateTimeServiceServer interface {
	mustEmbedUnimplementedDateTimeServiceServer()
}

func RegisterDateTimeServiceServer(s *grpc.Server, srv DateTimeServiceServer) {
	s.RegisterService(&_DateTimeService_serviceDesc, srv)
}

func _DateTimeService_LaunchCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateTimeServiceServer).LaunchCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dateTime.DateTimeService/LaunchCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateTimeServiceServer).LaunchCronJob(ctx, req.(*TriggerTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DateTimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dateTime.DateTimeService",
	HandlerType: (*DateTimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchCronJob",
			Handler:    _DateTimeService_LaunchCronJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/dateTimeService.proto",
}
