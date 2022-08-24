package gtw

import (
	"context"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatFileLogic {
	return &StatFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatFileLogic) StatFile(req *types.StatFileReq) (resp *types.OssFile, err error) {
	// todo: add your logic here and delete this line

	return
}
