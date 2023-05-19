package logic

import (
	"context"
	"time"

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

func (l *SendOneMsgToUserLogic) SendOneMsgToUser(in *pb.SendOneMsgToUserReq) (*pb.SendOneMsgToUserResp, error) {
	seq := time.Now().UnixNano()
	return &pb.SendOneMsgToUserResp{Seq: seq}, nil
}
