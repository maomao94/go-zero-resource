package logic

import (
	"context"
	"errors"

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
	if in.Id == 0 {
		return nil, errors.New("id error")
	}
	if in.Id != 1 {
		return &pb.GetUserInfoResp{User: &pb.User{
			Id:       2,
			Mobile:   "15651013267",
			Nickname: "何汉鹏2",
			Sex:      2,
			Avatar:   "https://",
		}}, nil
	} else {
		return &pb.GetUserInfoResp{User: &pb.User{
			Id:       1,
			Mobile:   "15651013267",
			Nickname: "何汉鹏",
			Sex:      1,
			Avatar:   "http://",
		}}, nil
	}
}
