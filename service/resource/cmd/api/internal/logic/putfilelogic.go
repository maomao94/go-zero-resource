package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
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

func (l *PutFileLogic) PutFile() error {
	// todo: add your logic here and delete this line

	return nil
}
