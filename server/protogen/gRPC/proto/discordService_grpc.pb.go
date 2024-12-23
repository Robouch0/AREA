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

// DiscordServiceClient is the client API for DiscordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscordServiceClient interface {
	CreateMessage(ctx context.Context, in *CreateMsg, opts ...grpc.CallOption) (*CreateMsg, error)
	EditMessage(ctx context.Context, in *EditMsg, opts ...grpc.CallOption) (*EditMsg, error)
	DeleteMessage(ctx context.Context, in *DeleteMsg, opts ...grpc.CallOption) (*DeleteMsg, error)
	CreateReaction(ctx context.Context, in *CreateReact, opts ...grpc.CallOption) (*CreateReact, error)
	DeleteAllReactions(ctx context.Context, in *DeleteAllReact, opts ...grpc.CallOption) (*DeleteAllReact, error)
	DeleteReactions(ctx context.Context, in *DeleteReact, opts ...grpc.CallOption) (*DeleteReact, error)
}

type discordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscordServiceClient(cc grpc.ClientConnInterface) DiscordServiceClient {
	return &discordServiceClient{cc}
}

func (c *discordServiceClient) CreateMessage(ctx context.Context, in *CreateMsg, opts ...grpc.CallOption) (*CreateMsg, error) {
	out := new(CreateMsg)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordServiceClient) EditMessage(ctx context.Context, in *EditMsg, opts ...grpc.CallOption) (*EditMsg, error) {
	out := new(EditMsg)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/EditMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordServiceClient) DeleteMessage(ctx context.Context, in *DeleteMsg, opts ...grpc.CallOption) (*DeleteMsg, error) {
	out := new(DeleteMsg)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/DeleteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordServiceClient) CreateReaction(ctx context.Context, in *CreateReact, opts ...grpc.CallOption) (*CreateReact, error) {
	out := new(CreateReact)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/CreateReaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordServiceClient) DeleteAllReactions(ctx context.Context, in *DeleteAllReact, opts ...grpc.CallOption) (*DeleteAllReact, error) {
	out := new(DeleteAllReact)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/DeleteAllReactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordServiceClient) DeleteReactions(ctx context.Context, in *DeleteReact, opts ...grpc.CallOption) (*DeleteReact, error) {
	out := new(DeleteReact)
	err := c.cc.Invoke(ctx, "/discord.DiscordService/DeleteReactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscordServiceServer is the server API for DiscordService service.
// All implementations must embed UnimplementedDiscordServiceServer
// for forward compatibility
type DiscordServiceServer interface {
	CreateMessage(context.Context, *CreateMsg) (*CreateMsg, error)
	EditMessage(context.Context, *EditMsg) (*EditMsg, error)
	DeleteMessage(context.Context, *DeleteMsg) (*DeleteMsg, error)
	CreateReaction(context.Context, *CreateReact) (*CreateReact, error)
	DeleteAllReactions(context.Context, *DeleteAllReact) (*DeleteAllReact, error)
	DeleteReactions(context.Context, *DeleteReact) (*DeleteReact, error)
	mustEmbedUnimplementedDiscordServiceServer()
}

// UnimplementedDiscordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscordServiceServer struct {
}

func (UnimplementedDiscordServiceServer) CreateMessage(context.Context, *CreateMsg) (*CreateMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedDiscordServiceServer) EditMessage(context.Context, *EditMsg) (*EditMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedDiscordServiceServer) DeleteMessage(context.Context, *DeleteMsg) (*DeleteMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedDiscordServiceServer) CreateReaction(context.Context, *CreateReact) (*CreateReact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReaction not implemented")
}
func (UnimplementedDiscordServiceServer) DeleteAllReactions(context.Context, *DeleteAllReact) (*DeleteAllReact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllReactions not implemented")
}
func (UnimplementedDiscordServiceServer) DeleteReactions(context.Context, *DeleteReact) (*DeleteReact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReactions not implemented")
}
func (UnimplementedDiscordServiceServer) mustEmbedUnimplementedDiscordServiceServer() {}

// UnsafeDiscordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscordServiceServer will
// result in compilation errors.
type UnsafeDiscordServiceServer interface {
	mustEmbedUnimplementedDiscordServiceServer()
}

func RegisterDiscordServiceServer(s *grpc.Server, srv DiscordServiceServer) {
	s.RegisterService(&_DiscordService_serviceDesc, srv)
}

func _DiscordService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).CreateMessage(ctx, req.(*CreateMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscordService_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/EditMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).EditMessage(ctx, req.(*EditMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscordService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/DeleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).DeleteMessage(ctx, req.(*DeleteMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscordService_CreateReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).CreateReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/CreateReaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).CreateReaction(ctx, req.(*CreateReact))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscordService_DeleteAllReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllReact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).DeleteAllReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/DeleteAllReactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).DeleteAllReactions(ctx, req.(*DeleteAllReact))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscordService_DeleteReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServiceServer).DeleteReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discord.DiscordService/DeleteReactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServiceServer).DeleteReactions(ctx, req.(*DeleteReact))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discord.DiscordService",
	HandlerType: (*DiscordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _DiscordService_CreateMessage_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _DiscordService_EditMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _DiscordService_DeleteMessage_Handler,
		},
		{
			MethodName: "CreateReaction",
			Handler:    _DiscordService_CreateReaction_Handler,
		},
		{
			MethodName: "DeleteAllReactions",
			Handler:    _DiscordService_DeleteAllReactions_Handler,
		},
		{
			MethodName: "DeleteReactions",
			Handler:    _DiscordService_DeleteReactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/discordService.proto",
}
