// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

package server

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/sys/internal/logic"
	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"
)

type ResourceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedResourceServer
}

func NewResourceServer(svcCtx *svc.ServiceContext) *ResourceServer {
	return &ResourceServer{
		svcCtx: svcCtx,
	}
}

func (s *ResourceServer) Ping(ctx context.Context, in *pb.Empty) (*pb.PingResp, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
