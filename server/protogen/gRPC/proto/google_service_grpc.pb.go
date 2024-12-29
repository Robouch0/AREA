// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
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
	GoogleService_SendEmailMe_FullMethodName     = "/google.GoogleService/SendEmailMe"
	GoogleService_DeleteEmailMe_FullMethodName   = "/google.GoogleService/DeleteEmailMe"
	GoogleService_MoveToTrash_FullMethodName     = "/google.GoogleService/MoveToTrash"
	GoogleService_MoveFromTrash_FullMethodName   = "/google.GoogleService/MoveFromTrash"
	GoogleService_WatchGmailEmail_FullMethodName = "/google.GoogleService/WatchGmailEmail"
	GoogleService_WatchMeTrigger_FullMethodName  = "/google.GoogleService/WatchMeTrigger"
)

// GoogleServiceClient is the client API for GoogleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoogleServiceClient interface {
	// Send an email with the current user email
	SendEmailMe(ctx context.Context, in *EmailRequestMe, opts ...grpc.CallOption) (*EmailRequestMe, error)
	// Delete one of user's email based on the subject of the mail
	DeleteEmailMe(ctx context.Context, in *DeleteEmailRequestMe, opts ...grpc.CallOption) (*DeleteEmailRequestMe, error)
	// Move to trash an email
	MoveToTrash(ctx context.Context, in *TrashEmailRequestMe, opts ...grpc.CallOption) (*TrashEmailRequestMe, error)
	// Move out of trash an email
	MoveFromTrash(ctx context.Context, in *TrashEmailRequestMe, opts ...grpc.CallOption) (*TrashEmailRequestMe, error)
	// Watch email of the user currently logged
	WatchGmailEmail(ctx context.Context, in *EmailTriggerReq, opts ...grpc.CallOption) (*EmailTriggerReq, error)
	// Function that handle the payload sent by google gmail
	WatchMeTrigger(ctx context.Context, in *GmailTriggerReq, opts ...grpc.CallOption) (*GmailTriggerReq, error)
}

type googleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoogleServiceClient(cc grpc.ClientConnInterface) GoogleServiceClient {
	return &googleServiceClient{cc}
}

func (c *googleServiceClient) SendEmailMe(ctx context.Context, in *EmailRequestMe, opts ...grpc.CallOption) (*EmailRequestMe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmailRequestMe)
	err := c.cc.Invoke(ctx, GoogleService_SendEmailMe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleServiceClient) DeleteEmailMe(ctx context.Context, in *DeleteEmailRequestMe, opts ...grpc.CallOption) (*DeleteEmailRequestMe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEmailRequestMe)
	err := c.cc.Invoke(ctx, GoogleService_DeleteEmailMe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleServiceClient) MoveToTrash(ctx context.Context, in *TrashEmailRequestMe, opts ...grpc.CallOption) (*TrashEmailRequestMe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrashEmailRequestMe)
	err := c.cc.Invoke(ctx, GoogleService_MoveToTrash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleServiceClient) MoveFromTrash(ctx context.Context, in *TrashEmailRequestMe, opts ...grpc.CallOption) (*TrashEmailRequestMe, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrashEmailRequestMe)
	err := c.cc.Invoke(ctx, GoogleService_MoveFromTrash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleServiceClient) WatchGmailEmail(ctx context.Context, in *EmailTriggerReq, opts ...grpc.CallOption) (*EmailTriggerReq, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmailTriggerReq)
	err := c.cc.Invoke(ctx, GoogleService_WatchGmailEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleServiceClient) WatchMeTrigger(ctx context.Context, in *GmailTriggerReq, opts ...grpc.CallOption) (*GmailTriggerReq, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GmailTriggerReq)
	err := c.cc.Invoke(ctx, GoogleService_WatchMeTrigger_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleServiceServer is the server API for GoogleService service.
// All implementations must embed UnimplementedGoogleServiceServer
// for forward compatibility.
type GoogleServiceServer interface {
	// Send an email with the current user email
	SendEmailMe(context.Context, *EmailRequestMe) (*EmailRequestMe, error)
	// Delete one of user's email based on the subject of the mail
	DeleteEmailMe(context.Context, *DeleteEmailRequestMe) (*DeleteEmailRequestMe, error)
	// Move to trash an email
	MoveToTrash(context.Context, *TrashEmailRequestMe) (*TrashEmailRequestMe, error)
	// Move out of trash an email
	MoveFromTrash(context.Context, *TrashEmailRequestMe) (*TrashEmailRequestMe, error)
	// Watch email of the user currently logged
	WatchGmailEmail(context.Context, *EmailTriggerReq) (*EmailTriggerReq, error)
	// Function that handle the payload sent by google gmail
	WatchMeTrigger(context.Context, *GmailTriggerReq) (*GmailTriggerReq, error)
	mustEmbedUnimplementedGoogleServiceServer()
}

// UnimplementedGoogleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoogleServiceServer struct{}

func (UnimplementedGoogleServiceServer) SendEmailMe(context.Context, *EmailRequestMe) (*EmailRequestMe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailMe not implemented")
}
func (UnimplementedGoogleServiceServer) DeleteEmailMe(context.Context, *DeleteEmailRequestMe) (*DeleteEmailRequestMe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailMe not implemented")
}
func (UnimplementedGoogleServiceServer) MoveToTrash(context.Context, *TrashEmailRequestMe) (*TrashEmailRequestMe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToTrash not implemented")
}
func (UnimplementedGoogleServiceServer) MoveFromTrash(context.Context, *TrashEmailRequestMe) (*TrashEmailRequestMe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveFromTrash not implemented")
}
func (UnimplementedGoogleServiceServer) WatchGmailEmail(context.Context, *EmailTriggerReq) (*EmailTriggerReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchGmailEmail not implemented")
}
func (UnimplementedGoogleServiceServer) WatchMeTrigger(context.Context, *GmailTriggerReq) (*GmailTriggerReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchMeTrigger not implemented")
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

func _GoogleService_SendEmailMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequestMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).SendEmailMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_SendEmailMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).SendEmailMe(ctx, req.(*EmailRequestMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleService_DeleteEmailMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailRequestMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).DeleteEmailMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_DeleteEmailMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).DeleteEmailMe(ctx, req.(*DeleteEmailRequestMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleService_MoveToTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrashEmailRequestMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).MoveToTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_MoveToTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).MoveToTrash(ctx, req.(*TrashEmailRequestMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleService_MoveFromTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrashEmailRequestMe)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).MoveFromTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_MoveFromTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).MoveFromTrash(ctx, req.(*TrashEmailRequestMe))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleService_WatchGmailEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailTriggerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).WatchGmailEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_WatchGmailEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).WatchGmailEmail(ctx, req.(*EmailTriggerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleService_WatchMeTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailTriggerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleServiceServer).WatchMeTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleService_WatchMeTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleServiceServer).WatchMeTrigger(ctx, req.(*GmailTriggerReq))
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
			MethodName: "SendEmailMe",
			Handler:    _GoogleService_SendEmailMe_Handler,
		},
		{
			MethodName: "DeleteEmailMe",
			Handler:    _GoogleService_DeleteEmailMe_Handler,
		},
		{
			MethodName: "MoveToTrash",
			Handler:    _GoogleService_MoveToTrash_Handler,
		},
		{
			MethodName: "MoveFromTrash",
			Handler:    _GoogleService_MoveFromTrash_Handler,
		},
		{
			MethodName: "WatchGmailEmail",
			Handler:    _GoogleService_WatchGmailEmail_Handler,
		},
		{
			MethodName: "WatchMeTrigger",
			Handler:    _GoogleService_WatchMeTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/google_service.proto",
}
