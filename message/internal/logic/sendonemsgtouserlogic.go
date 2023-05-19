package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/message/internal/svc"
	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendOneMsgToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendOneMsgToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendOneMsgToUserLogic {
	return &SendOneMsgToUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendOneMsgToUserLogic) SendOneMsgToUser(in *pb.SendOneMsgToUserReq) (*pb.SendOneMsgToUserRes, error) {
	// todo: add your logic here and delete this line

	return &pb.SendOneMsgToUserRes{}, nil
}
