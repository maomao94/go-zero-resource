package gtw

import (
	"context"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFilesLogic {
	return &RemoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFilesLogic) RemoveFiles(req *types.RemoveFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
