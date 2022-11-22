package resource

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOssLogic {
	return &UpdateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOssLogic) UpdateOss(req *types.OssUpdate) (resp *types.EmptyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
