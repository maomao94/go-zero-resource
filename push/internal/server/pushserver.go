// Code generated by goctl. DO NOT EDIT.
// Source: push.proto

package server

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/push/internal/logic"
	"github.com/hehanpeng/go-zero-resource/push/internal/svc"
	"github.com/hehanpeng/go-zero-resource/push/pb"
)

type PushServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPushServer
}

func NewPushServer(svcCtx *svc.ServiceContext) *PushServer {
	return &PushServer{
		svcCtx: svcCtx,
	}
}

func (s *PushServer) Ping(ctx context.Context, in *pb.Empty) (*pb.PingResp, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *PushServer) PushOneWsMsgToUser(ctx context.Context, in *pb.PushOneMsgToUserReq) (*pb.PushOneMsgToUserRes, error) {
	l := logic.NewPushOneWsMsgToUserLogic(ctx, s.svcCtx)
	return l.PushOneWsMsgToUser(in)
}
