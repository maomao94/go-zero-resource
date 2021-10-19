package logic

import (
	"context"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PutFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutFileLogic {
	return PutFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutFileLogic) PutFile(req types.PutFileReq) (*types.File, error) {
	// todo: add your logic here and delete this line

	return &types.File{}, nil
}
