// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gRPC/proto/reactionService.proto

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
	ReactionService_RegisterAction_FullMethodName = "/reaction.ReactionService/RegisterAction"
	ReactionService_LaunchReaction_FullMethodName = "/reaction.ReactionService/LaunchReaction"
)

// ReactionServiceClient is the client API for ReactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReactionServiceClient interface {
	RegisterAction(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	LaunchReaction(ctx context.Context, in *LaunchRequest, opts ...grpc.CallOption) (*LaunchResponse, error)
}

type reactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReactionServiceClient(cc grpc.ClientConnInterface) ReactionServiceClient {
	return &reactionServiceClient{cc}
}

func (c *reactionServiceClient) RegisterAction(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, ReactionService_RegisterAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) LaunchReaction(ctx context.Context, in *LaunchRequest, opts ...grpc.CallOption) (*LaunchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LaunchResponse)
	err := c.cc.Invoke(ctx, ReactionService_LaunchReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReactionServiceServer is the server API for ReactionService service.
// All implementations must embed UnimplementedReactionServiceServer
// for forward compatibility.
type ReactionServiceServer interface {
	RegisterAction(context.Context, *ReactionRequest) (*ReactionResponse, error)
	LaunchReaction(context.Context, *LaunchRequest) (*LaunchResponse, error)
	mustEmbedUnimplementedReactionServiceServer()
}

// UnimplementedReactionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReactionServiceServer struct{}

func (UnimplementedReactionServiceServer) RegisterAction(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAction not implemented")
}
func (UnimplementedReactionServiceServer) LaunchReaction(context.Context, *LaunchRequest) (*LaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchReaction not implemented")
}
func (UnimplementedReactionServiceServer) mustEmbedUnimplementedReactionServiceServer() {}
func (UnimplementedReactionServiceServer) testEmbeddedByValue()                         {}

// UnsafeReactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReactionServiceServer will
// result in compilation errors.
type UnsafeReactionServiceServer interface {
	mustEmbedUnimplementedReactionServiceServer()
}

func RegisterReactionServiceServer(s grpc.ServiceRegistrar, srv ReactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedReactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReactionService_ServiceDesc, srv)
}

func _ReactionService_RegisterAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).RegisterAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactionService_RegisterAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).RegisterAction(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_LaunchReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).LaunchReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactionService_LaunchReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).LaunchReaction(ctx, req.(*LaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReactionService_ServiceDesc is the grpc.ServiceDesc for ReactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reaction.ReactionService",
	HandlerType: (*ReactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAction",
			Handler:    _ReactionService_RegisterAction_Handler,
		},
		{
			MethodName: "LaunchReaction",
			Handler:    _ReactionService_LaunchReaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/reactionService.proto",
}
