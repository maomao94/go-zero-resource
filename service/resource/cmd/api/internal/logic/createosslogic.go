package logic

import (
	"context"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gorm_model"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CreateOssLogic) CreateOss(req types.OssCreate) error {
	resourceOss := gorm_model.ResourceOss{
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
	if err := service.ResourceOssApp.CreateResourceOss(resourceOss); err != nil {
		return err
	} else {
		return nil
	}
}
