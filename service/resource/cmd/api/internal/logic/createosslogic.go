package logic

import (
	"context"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateOssLogic {
	return CreateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOssLogic) CreateOss(req types.Oss) error {
	// todo: add your logic here and delete this line

	return nil
}
