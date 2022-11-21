package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/model"
	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOssLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOssLogic {
	return &DeleteOssLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOssLogic) DeleteOss(in *pb.DeleteOssReq) (*pb.Empty, error) {
	err := l.svcCtx.TOssModel.DeleteSoft(l.ctx, &model.TOss{Id: in.Id})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
