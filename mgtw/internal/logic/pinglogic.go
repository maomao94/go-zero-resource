package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/mgtw/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.Empty) (*pb.PingResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PingResp{}, nil
}
