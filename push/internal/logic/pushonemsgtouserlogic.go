package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/push/internal/svc"
	"github.com/hehanpeng/go-zero-resource/push/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushOneMsgToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushOneMsgToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushOneMsgToUserLogic {
	return &PushOneMsgToUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushOneMsgToUserLogic) PushOneMsgToUser(in *pb.PushOneMsgToUserReq) (*pb.PushOneMsgToUserRes, error) {
	// todo: add your logic here and delete this line

	return &pb.PushOneMsgToUserRes{}, nil
}
