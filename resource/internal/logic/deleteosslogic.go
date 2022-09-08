package logic

import (
	"context"
	"gtw/resource/internal/svc"
	"gtw/resource/pb"

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
	err := l.svcCtx.TOssModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
