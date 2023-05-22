package logic

import (
	"context"
	"strconv"

	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/mgtw/pb"

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
	cli, err := l.svcCtx.ClientManager.GetUserClient(12138, strconv.FormatInt(in.ToUserId, 10))
	if err != nil {
		return nil, err
	}
	err = cli.SendSeqMsg(strconv.FormatInt(in.Seq, 10), []byte(in.Msg))
	if err != nil {
		return nil, err
	}
	return &pb.PushOneMsgToUserRes{}, nil
}
