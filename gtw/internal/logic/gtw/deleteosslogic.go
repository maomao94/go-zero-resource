package gtw

import (
	"context"

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

func (l *DeleteOssLogic) DeleteOss(req *types.OssDelete) error {
	// todo: add your logic here and delete this line

	return nil
}
