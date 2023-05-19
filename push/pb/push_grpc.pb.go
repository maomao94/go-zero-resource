// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: push.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GtwClient is the client API for Gtw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GtwClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error)
	PushOneMsgToUser(ctx context.Context, in *PushOneMsgToUserReq, opts ...grpc.CallOption) (*PushOneMsgToUserRes, error)
}

type gtwClient struct {
	cc grpc.ClientConnInterface
}

func NewGtwClient(cc grpc.ClientConnInterface) GtwClient {
	return &gtwClient{cc}
}

func (c *gtwClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/push.gtw/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gtwClient) PushOneMsgToUser(ctx context.Context, in *PushOneMsgToUserReq, opts ...grpc.CallOption) (*PushOneMsgToUserRes, error) {
	out := new(PushOneMsgToUserRes)
	err := c.cc.Invoke(ctx, "/push.gtw/pushOneMsgToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GtwServer is the server API for Gtw service.
// All implementations must embed UnimplementedGtwServer
// for forward compatibility
type GtwServer interface {
	Ping(context.Context, *Empty) (*PingResp, error)
	PushOneMsgToUser(context.Context, *PushOneMsgToUserReq) (*PushOneMsgToUserRes, error)
	mustEmbedUnimplementedGtwServer()
}

// UnimplementedGtwServer must be embedded to have forward compatible implementations.
type UnimplementedGtwServer struct {
}

func (UnimplementedGtwServer) Ping(context.Context, *Empty) (*PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGtwServer) PushOneMsgToUser(context.Context, *PushOneMsgToUserReq) (*PushOneMsgToUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushOneMsgToUser not implemented")
}
func (UnimplementedGtwServer) mustEmbedUnimplementedGtwServer() {}

// UnsafeGtwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GtwServer will
// result in compilation errors.
type UnsafeGtwServer interface {
	mustEmbedUnimplementedGtwServer()
}

func RegisterGtwServer(s grpc.ServiceRegistrar, srv GtwServer) {
	s.RegisterService(&Gtw_ServiceDesc, srv)
}

func _Gtw_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.gtw/ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gtw_PushOneMsgToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushOneMsgToUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwServer).PushOneMsgToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.gtw/pushOneMsgToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwServer).PushOneMsgToUser(ctx, req.(*PushOneMsgToUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gtw_ServiceDesc is the grpc.ServiceDesc for Gtw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gtw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "push.gtw",
	HandlerType: (*GtwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _Gtw_Ping_Handler,
		},
		{
			MethodName: "pushOneMsgToUser",
			Handler:    _Gtw_PushOneMsgToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push.proto",
}
