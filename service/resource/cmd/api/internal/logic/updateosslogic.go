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
	err := utils.Verify(req, utils.IdVerify)
	if err != nil {
		return err
	}
	resourceOss := gorm_model.ResourceOss{
		GVA_MODEL: gorm_model.GVA_MODEL{
			ID: req.Id,
		},
		TenantId:   req.TenantId,
		Category:   req.Category,
		OssCode:    req.OssCode,
		Endpoint:   req.Endpoint,
		AccessKey:  req.AccessKey,
		SecretKey:  req.SecretKey,
		BucketName: req.BucketName,
		AppId:      req.AppId,
		Region:     req.Region,
		Remark:     req.Remark,
		Status:     req.Status,
	}
	if err := service.ResourceOssApp.UpdateResourceOss(resourceOss); err != nil {
		return err
	} else {
		return nil
	}
	return nil
}
