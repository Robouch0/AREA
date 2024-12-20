// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
<<<<<<< Updated upstream
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gRPC/proto/hugging_face_service.proto
=======
>>>>>>> Stashed changes

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

// HuggingFaceServiceClient is the client API for HuggingFaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HuggingFaceServiceClient interface {
	LaunchTextGeneration(ctx context.Context, in *TextGenerationReq, opts ...grpc.CallOption) (*TextGenerationRes, error)
}

type huggingFaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHuggingFaceServiceClient(cc grpc.ClientConnInterface) HuggingFaceServiceClient {
	return &huggingFaceServiceClient{cc}
}

func (c *huggingFaceServiceClient) LaunchTextGeneration(ctx context.Context, in *TextGenerationReq, opts ...grpc.CallOption) (*TextGenerationRes, error) {
	out := new(TextGenerationRes)
	err := c.cc.Invoke(ctx, "/huggingface.HuggingFaceService/LaunchTextGeneration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuggingFaceServiceServer is the server API for HuggingFaceService service.
// All implementations must embed UnimplementedHuggingFaceServiceServer
// for forward compatibility
type HuggingFaceServiceServer interface {
	LaunchTextGeneration(context.Context, *TextGenerationReq) (*TextGenerationRes, error)
	mustEmbedUnimplementedHuggingFaceServiceServer()
}

// UnimplementedHuggingFaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHuggingFaceServiceServer struct {
}

func (UnimplementedHuggingFaceServiceServer) LaunchTextGeneration(context.Context, *TextGenerationReq) (*TextGenerationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchTextGeneration not implemented")
}
func (UnimplementedHuggingFaceServiceServer) mustEmbedUnimplementedHuggingFaceServiceServer() {}

// UnsafeHuggingFaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HuggingFaceServiceServer will
// result in compilation errors.
type UnsafeHuggingFaceServiceServer interface {
	mustEmbedUnimplementedHuggingFaceServiceServer()
}

func RegisterHuggingFaceServiceServer(s *grpc.Server, srv HuggingFaceServiceServer) {
	s.RegisterService(&_HuggingFaceService_serviceDesc, srv)
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
		FullMethod: "/huggingface.HuggingFaceService/LaunchTextGeneration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuggingFaceServiceServer).LaunchTextGeneration(ctx, req.(*TextGenerationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HuggingFaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "huggingface.HuggingFaceService",
	HandlerType: (*HuggingFaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchTextGeneration",
			Handler:    _HuggingFaceService_LaunchTextGeneration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/hugging_face_service.proto",
}
