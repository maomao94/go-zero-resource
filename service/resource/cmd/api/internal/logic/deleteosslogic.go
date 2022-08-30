package logic

import (
	"context"
	"go-zero-resource/common/utils"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gorm_model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteOssLogic {
	return DeleteOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOssLogic) DeleteOss(req types.OssDelete) error {
	err := utils.Verify(req, utils.IdVerify)
	if err != nil {
		return err
	}
	resourceOss := gorm_model.ResourceOss{
		GVA_MODEL: gorm_model.GVA_MODEL{
			ID: req.Id,
		},
	}
	if err := service.ResourceOssApp.DeleteResourceOss(resourceOss); err != nil {
		return err
	} else {
		return nil
	}
}
