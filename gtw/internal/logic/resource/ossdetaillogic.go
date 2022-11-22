package resource

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	pb2 "github.com/hehanpeng/go-zero-resource/sys/pb"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssDetailLogic {
	return &OssDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssDetailLogic) OssDetail(req *types.BaseReq) (resp *types.Oss, err error) {
	ossDetailResp, err := l.svcCtx.ResourceRpc.OssDetail(l.ctx, &pb.OssDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var respOss types.Oss
	_ = copier.Copy(&respOss, ossDetailResp.Oss)
	err = mr.Finish(func() error {
		getUserInfoResp, err := l.svcCtx.SysRpc.GetUserInfo(l.ctx, &pb2.GetUserInfoReq{Id: respOss.CreateUser})
		if err != nil {
			return err
		}
		respOss.CreateNickname = getUserInfoResp.User.Nickname
		return nil
	}, func() error {
		getUserInfoResp, err := l.svcCtx.SysRpc.GetUserInfo(l.ctx, &pb2.GetUserInfoReq{Id: respOss.UpdateUser})
		if err != nil {
			return err
		}
		respOss.UpdateNickname = getUserInfoResp.User.Nickname
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &respOss, nil
}
