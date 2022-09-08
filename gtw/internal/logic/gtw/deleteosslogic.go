package gtw

import (
	"context"
	"gtw/resource/pb"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOssLogic {
	return &DeleteOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOssLogic) DeleteOss(req *types.OssDelete) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.DeleteOss(l.ctx, &pb.DeleteOssReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
