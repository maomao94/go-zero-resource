package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOssDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssDetailLogic {
	return &OssDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OssDetailLogic) OssDetail(in *pb.OssDetailReq) (*pb.OssDetailResp, error) {
	oss, err := l.svcCtx.TOssModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var respOss pb.Oss
	_ = copier.Copy(&respOss, oss)
	respOss.CreateTime = oss.CreateTime.Unix()
	return &pb.OssDetailResp{Oss: &respOss}, nil
}
