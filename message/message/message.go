// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

package message

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty     = pb.Empty
	KqSendReq = pb.KqSendReq
	PingResp  = pb.PingResp

	Message interface {
		Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error)
		KqSend(ctx context.Context, in *KqSendReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultMessage struct {
		cli zrpc.Client
	}
)

func NewMessage(cli zrpc.Client) Message {
	return &defaultMessage{
		cli: cli,
	}
}

func (m *defaultMessage) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResp, error) {
	client := pb.NewMessageClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultMessage) KqSend(ctx context.Context, in *KqSendReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewMessageClient(m.cli.Conn())
	return client.KqSend(ctx, in, opts...)
}