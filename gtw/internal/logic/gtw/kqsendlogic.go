package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KqSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKqSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KqSendLogic {
	return &KqSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KqSendLogic) KqSend(req *types.KqSendReq) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.MessageRpc.KqSend(l.ctx, &pb.KqSendReq{Msg: req.Msg})
	if err != nil {
		return nil, err
	}
	return
}
