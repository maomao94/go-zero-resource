package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	pb2 "github.com/hehanpeng/go-zero-resource/mgtw/pb"
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
	rpcList := l.svcCtx.PubContainer.MGtwRpcList
	if len(rpcList) > 0 {
		for _, v := range rpcList {
			_, err := v.PushOneWsMsgToUser(l.ctx, &pb2.PushOneMsgToUserReq{
				Seq: seq,
			})
			if err != nil {
				return nil, err
			}
		}
		return &pb.SendOneMsgToUserResp{Seq: seq}, nil
	} else {
		return nil, errorx.NewEnumErrorf(errorx.Code_Err, "发送失败")
	}
}
