package logic

import (
	"context"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gormx"

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
	var resourceOss gormx.ResourceOss
	if err := service.ResourceOssApp.CreateResourceOss(resourceOss); err != nil {
		return err
	} else {
		return nil
	}
}
