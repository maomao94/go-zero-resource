package message

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendOneMsgToUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendOneMsgToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendOneMsgToUserLogic {
	return &SendOneMsgToUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendOneMsgToUserLogic) SendOneMsgToUser(req *types.SendOneMsgToUserReq) (resp *types.SendOneMsgToUserRes, err error) {
	sendOneMsgToUserResp, err := l.svcCtx.MessageRpc.SendOneMsgToUser(l.ctx, &pb.SendOneMsgToUserReq{
		FromUserId: req.FromUserId,
		ToUerId:    req.ToUserId,
		Msg:        req.Msg,
	})
	if err != nil {
		return nil, err
	}
	return &types.SendOneMsgToUserRes{Seq: sendOneMsgToUserResp.Seq}, nil
}
