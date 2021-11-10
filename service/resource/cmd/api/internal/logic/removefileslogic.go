package logic

import (
	"context"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveFilesLogic {
	return RemoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFilesLogic) RemoveFiles(req types.RemoveFilesReq) error {
	// todo: add your logic here and delete this line

	return nil
}
