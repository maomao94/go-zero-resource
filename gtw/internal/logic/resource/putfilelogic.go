package resource

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileLogic {
	return &PutFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutFileLogic) PutFile(req *types.PutFileReq) (resp *types.File, err error) {
	// todo: add your logic here and delete this line

	return
}
