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

// SpotifyServiceClient is the client API for SpotifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotifyServiceClient interface {
	StopSong(ctx context.Context, in *SpotifyStopInfo, opts ...grpc.CallOption) (*SpotifyStopInfo, error)
}

type spotifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotifyServiceClient(cc grpc.ClientConnInterface) SpotifyServiceClient {
	return &spotifyServiceClient{cc}
}

func (c *spotifyServiceClient) StopSong(ctx context.Context, in *SpotifyStopInfo, opts ...grpc.CallOption) (*SpotifyStopInfo, error) {
	out := new(SpotifyStopInfo)
	err := c.cc.Invoke(ctx, "/spotify.SpotifyService/StopSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotifyServiceServer is the server API for SpotifyService service.
// All implementations must embed UnimplementedSpotifyServiceServer
// for forward compatibility
type SpotifyServiceServer interface {
	StopSong(context.Context, *SpotifyStopInfo) (*SpotifyStopInfo, error)
	mustEmbedUnimplementedSpotifyServiceServer()
}

// UnimplementedSpotifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpotifyServiceServer struct {
}

func (UnimplementedSpotifyServiceServer) StopSong(context.Context, *SpotifyStopInfo) (*SpotifyStopInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSong not implemented")
}
func (UnimplementedSpotifyServiceServer) mustEmbedUnimplementedSpotifyServiceServer() {}

// UnsafeSpotifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotifyServiceServer will
// result in compilation errors.
type UnsafeSpotifyServiceServer interface {
	mustEmbedUnimplementedSpotifyServiceServer()
}

func RegisterSpotifyServiceServer(s *grpc.Server, srv SpotifyServiceServer) {
	s.RegisterService(&_SpotifyService_serviceDesc, srv)
}

func _SpotifyService_StopSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotifyStopInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyServiceServer).StopSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.SpotifyService/StopSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyServiceServer).StopSong(ctx, req.(*SpotifyStopInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpotifyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotify.SpotifyService",
	HandlerType: (*SpotifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StopSong",
			Handler:    _SpotifyService_StopSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/spotify_service.proto",
}
