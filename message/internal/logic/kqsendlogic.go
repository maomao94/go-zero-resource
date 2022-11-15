package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/message/internal/svc"
	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KqSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKqSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KqSendLogic {
	return &KqSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KqSendLogic) KqSend(in *pb.KqSendReq) (*pb.Empty, error) {
	err := l.svcCtx.KafkaTestPusher.Push(in.Msg)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
