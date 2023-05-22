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
	pub := l.svcCtx.PubContainer.PubMap
	if len(pub) > 0 {
		l.Logger.Infof("len(pubMap)=%d", len(pub))
		for _, v := range pub {
			_, err := v.PushOneWsMsgToUser(l.ctx, &pb2.PushOneMsgToUserReq{
				Seq:        seq,
				FromUserId: 10001,
				ToUserId:   10001,
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
