package gtw

import (
	"context"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

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

func (l *OssDetailLogic) OssDetail(req *types.BaseResult) (resp *types.Oss, err error) {
	// todo: add your logic here and delete this line

	return
}
