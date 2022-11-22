package logic

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	return &pb.GetUserInfoResp{User: &pb.User{
		Id:       1,
		Mobile:   "15651013267",
		Nickname: "何汉鹏",
		Sex:      1,
		Avatar:   "http://",
	}}, nil
}
