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
	// todo: add your logic here and delete this line

	return &pb.Empty{}, nil
}
