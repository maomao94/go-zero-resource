package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/push/internal/svc"
	"github.com/hehanpeng/go-zero-resource/push/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushOneWsMsgToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushOneWsMsgToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushOneWsMsgToUserLogic {
	return &PushOneWsMsgToUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushOneWsMsgToUserLogic) PushOneWsMsgToUser(in *pb.PushOneMsgToUserReq) (*pb.PushOneMsgToUserRes, error) {
	// todo: add your logic here and delete this line

	return &pb.PushOneMsgToUserRes{}, nil
}
