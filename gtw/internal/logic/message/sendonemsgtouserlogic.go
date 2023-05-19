package message

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/hehanpeng/go-zero-resource/message/pb"
	"strconv"

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
	fromUserId, err := strconv.ParseInt(req.FromUserId, 10, 64)
	if err != nil {
		return nil, err
	}
	toUserId, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		return nil, err
	}
	sendOneMsgToUserResp, err := l.svcCtx.MessageRpc.SendOneMsgToUser(l.ctx, &pb.SendOneMsgToUserReq{
		FromUserId: fromUserId,
		ToUerId:    toUserId,
		Msg:        req.Msg,
	})
	if err != nil {
		return nil, err
	}
	return &types.SendOneMsgToUserRes{Seq: strconv.FormatInt(sendOneMsgToUserResp.Seq, 10)}, nil
}
