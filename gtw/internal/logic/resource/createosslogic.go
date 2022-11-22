package resource

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOssLogic {
	return &CreateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOssLogic) CreateOss(req *types.OssCreate) (resp *types.EmptyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
