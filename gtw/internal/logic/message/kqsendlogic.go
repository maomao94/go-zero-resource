package message

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
