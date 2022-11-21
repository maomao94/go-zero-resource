package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.Empty) (*pb.PingResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PingResp{}, nil
}
