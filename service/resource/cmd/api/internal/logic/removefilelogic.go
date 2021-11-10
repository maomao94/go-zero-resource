package logic

import (
	"context"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveFileLogic {
	return RemoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFileLogic) RemoveFile(req types.RemoveFileReq) error {
	// todo: add your logic here and delete this line

	return nil
}
