// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gRPC/proto/hugging_face_service.proto

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
	HuggingFaceService_LaunchTextGeneration_FullMethodName       = "/huggingface.HuggingFaceService/LaunchTextGeneration"
	HuggingFaceService_CreateRepoUpdateWebHook_FullMethodName    = "/huggingface.HuggingFaceService/CreateRepoUpdateWebHook"
	HuggingFaceService_CreateNewPRWebHook_FullMethodName         = "/huggingface.HuggingFaceService/CreateNewPRWebHook"
	HuggingFaceService_CreateNewDiscussionWebHook_FullMethodName = "/huggingface.HuggingFaceService/CreateNewDiscussionWebHook"
	HuggingFaceService_TriggerWebHook_FullMethodName             = "/huggingface.HuggingFaceService/TriggerWebHook"
)

// HuggingFaceServiceClient is the client API for HuggingFaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HuggingFaceServiceClient interface {
	LaunchTextGeneration(ctx context.Context, in *TextGenerationReq, opts ...grpc.CallOption) (*TextGenerationRes, error)
	CreateRepoUpdateWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error)
	CreateNewPRWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error)
	CreateNewDiscussionWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error)
	TriggerWebHook(ctx context.Context, in *WebHookTriggerReq, opts ...grpc.CallOption) (*WebHookTriggerReq, error)
}

type huggingFaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHuggingFaceServiceClient(cc grpc.ClientConnInterface) HuggingFaceServiceClient {
	return &huggingFaceServiceClient{cc}
}

func (c *huggingFaceServiceClient) LaunchTextGeneration(ctx context.Context, in *TextGenerationReq, opts ...grpc.CallOption) (*TextGenerationRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextGenerationRes)
	err := c.cc.Invoke(ctx, HuggingFaceService_LaunchTextGeneration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huggingFaceServiceClient) CreateRepoUpdateWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HFWebHookInfo)
	err := c.cc.Invoke(ctx, HuggingFaceService_CreateRepoUpdateWebHook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huggingFaceServiceClient) CreateNewPRWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HFWebHookInfo)
	err := c.cc.Invoke(ctx, HuggingFaceService_CreateNewPRWebHook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huggingFaceServiceClient) CreateNewDiscussionWebHook(ctx context.Context, in *HFWebHookInfo, opts ...grpc.CallOption) (*HFWebHookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HFWebHookInfo)
	err := c.cc.Invoke(ctx, HuggingFaceService_CreateNewDiscussionWebHook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huggingFaceServiceClient) TriggerWebHook(ctx context.Context, in *WebHookTriggerReq, opts ...grpc.CallOption) (*WebHookTriggerReq, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebHookTriggerReq)
	err := c.cc.Invoke(ctx, HuggingFaceService_TriggerWebHook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuggingFaceServiceServer is the server API for HuggingFaceService service.
// All implementations must embed UnimplementedHuggingFaceServiceServer
// for forward compatibility.
type HuggingFaceServiceServer interface {
	LaunchTextGeneration(context.Context, *TextGenerationReq) (*TextGenerationRes, error)
	CreateRepoUpdateWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error)
	CreateNewPRWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error)
	CreateNewDiscussionWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error)
	TriggerWebHook(context.Context, *WebHookTriggerReq) (*WebHookTriggerReq, error)
	mustEmbedUnimplementedHuggingFaceServiceServer()
}

// UnimplementedHuggingFaceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHuggingFaceServiceServer struct{}

func (UnimplementedHuggingFaceServiceServer) LaunchTextGeneration(context.Context, *TextGenerationReq) (*TextGenerationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchTextGeneration not implemented")
}
func (UnimplementedHuggingFaceServiceServer) CreateRepoUpdateWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepoUpdateWebHook not implemented")
}
func (UnimplementedHuggingFaceServiceServer) CreateNewPRWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewPRWebHook not implemented")
}
func (UnimplementedHuggingFaceServiceServer) CreateNewDiscussionWebHook(context.Context, *HFWebHookInfo) (*HFWebHookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewDiscussionWebHook not implemented")
}
func (UnimplementedHuggingFaceServiceServer) TriggerWebHook(context.Context, *WebHookTriggerReq) (*WebHookTriggerReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerWebHook not implemented")
}
func (UnimplementedHuggingFaceServiceServer) mustEmbedUnimplementedHuggingFaceServiceServer() {}
func (UnimplementedHuggingFaceServiceServer) testEmbeddedByValue()                            {}

// UnsafeHuggingFaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HuggingFaceServiceServer will
// result in compilation errors.
type UnsafeHuggingFaceServiceServer interface {
	mustEmbedUnimplementedHuggingFaceServiceServer()
}

func RegisterHuggingFaceServiceServer(s grpc.ServiceRegistrar, srv HuggingFaceServiceServer) {
	// If the following call pancis, it indicates UnimplementedHuggingFaceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HuggingFaceService_ServiceDesc, srv)
}

func _HuggingFaceService_LaunchTextGeneration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextGenerationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuggingFaceServiceServer).LaunchTextGeneration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuggingFaceService_LaunchTextGeneration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).LaunchTextGeneration(ctx, req.(*TextGenerationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuggingFaceService_CreateRepoUpdateWebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HFWebHookInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuggingFaceServiceServer).CreateRepoUpdateWebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuggingFaceService_CreateRepoUpdateWebHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).CreateRepoUpdateWebHook(ctx, req.(*HFWebHookInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuggingFaceService_CreateNewPRWebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HFWebHookInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuggingFaceServiceServer).CreateNewPRWebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuggingFaceService_CreateNewPRWebHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).CreateNewPRWebHook(ctx, req.(*HFWebHookInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuggingFaceService_CreateNewDiscussionWebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HFWebHookInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuggingFaceServiceServer).CreateNewDiscussionWebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuggingFaceService_CreateNewDiscussionWebHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).CreateNewDiscussionWebHook(ctx, req.(*HFWebHookInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuggingFaceService_TriggerWebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebHookTriggerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuggingFaceServiceServer).TriggerWebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HuggingFaceService_TriggerWebHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).TriggerWebHook(ctx, req.(*WebHookTriggerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// HuggingFaceService_ServiceDesc is the grpc.ServiceDesc for HuggingFaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HuggingFaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "huggingface.HuggingFaceService",
	HandlerType: (*HuggingFaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchTextGeneration",
			Handler:    _HuggingFaceService_LaunchTextGeneration_Handler,
		},
		{
			MethodName: "CreateRepoUpdateWebHook",
			Handler:    _HuggingFaceService_CreateRepoUpdateWebHook_Handler,
		},
		{
			MethodName: "CreateNewPRWebHook",
			Handler:    _HuggingFaceService_CreateNewPRWebHook_Handler,
		},
		{
			MethodName: "CreateNewDiscussionWebHook",
			Handler:    _HuggingFaceService_CreateNewDiscussionWebHook_Handler,
		},
		{
			MethodName: "TriggerWebHook",
			Handler:    _HuggingFaceService_TriggerWebHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/hugging_face_service.proto",
}
