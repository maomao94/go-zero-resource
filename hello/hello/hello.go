// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package hello

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/hello/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty    = pb.Empty
	PingResp = pb.PingResp

	Hello interface {
		Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error)
	}

	defaultHello struct {
		cli zrpc.Client
	}
)

func NewHello(cli zrpc.Client) Hello {
	return &defaultHello{
		cli: cli,
	}
}

func (m *defaultHello) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error) {
	client := pb.NewHelloClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
