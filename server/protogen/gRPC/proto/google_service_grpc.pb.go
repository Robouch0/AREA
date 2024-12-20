// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: gRPC/proto/google_service.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GoogleService_TestCall_FullMethodName = "/google.GoogleService/TestCall"
)

// GoogleServiceClient is the client API for GoogleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoogleServiceClient interface {
	TestCall(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestRequest, error)
}

type googleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoogleServiceClient(cc grpc.ClientConnInterface) GoogleServiceClient {
	return &googleServiceClient{cc}
}

func (c *googleServiceClient) TestCall(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestRequest)
	err := c.cc.Invoke(ctx, GoogleService_TestCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleServiceServer is the server API for GoogleService service.
// All implementations must embed UnimplementedGoogleServiceServer
// for forward compatibility.
type GoogleServiceServer interface {
	TestCall(context.Context, *TestRequest) (*TestRequest, error)
	mustEmbedUnimplementedGoogleServiceServer()
}

// UnimplementedGoogleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoogleServiceServer struct{}

func (UnimplementedGoogleServiceServer) TestCall(context.Context, *TestRequest) (*TestRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCall not implemented")
}
func (UnimplementedGoogleServiceServer) mustEmbedUnimplementedGoogleServiceServer() {}
func (UnimplementedGoogleServiceServer) testEmbeddedByValue()                       {}

// UnsafeGoogleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoogleServiceServer will
// result in compilation errors.
type UnsafeGoogleServiceServer interface {
	mustEmbedUnimplementedGoogleServiceServer()
}

func RegisterGoogleServiceServer(s grpc.ServiceRegistrar, srv GoogleServiceServer) {
	// If the following call pancis, it indicates UnimplementedGoogleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GoogleService_ServiceDesc, srv)
}

func _GoogleService_TestCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).TestCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_TestCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).TestCall(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoogleService_ServiceDesc is the grpc.ServiceDesc for GoogleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoogleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.GoogleService",
	HandlerType: (*GoogleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestCall",
			Handler:    _GoogleService_TestCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/google_service.proto",
}
