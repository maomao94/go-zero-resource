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
	l.Logger.Infof("收到消息：fromUserId:%d^toUserId:%d^seq:%d^msg:%s", in.FromUserId, in.ToUserId, in.Seq, in.Msg)
	return &pb.PushOneMsgToUserRes{}, nil
}
