package logic

import (
	"context"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gormx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateOssLogic {
	return UpdateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOssLogic) UpdateOss(req types.OssUpdate) error {
	var resourceOss gormx.ResourceOss
	if err := service.ResourceOssApp.UpdateResourceOss(resourceOss); err != nil {
		return err
	} else {
		return nil
	}
	return nil
}
