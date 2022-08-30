package logic

import (
	"context"
	"go-zero-resource/common/api"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gorm_model"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OssListLogic {
	return OssListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssListLogic) OssList(req types.OssListReq) (*types.PageResult, error) {
	pageInfo := gorm_model.ResourceOssSearch{
		PageInfo: api.PageInfo{
			PageSize: req.PageSize,
			Page:     req.Page,
		},
		ResourceOss: gorm_model.ResourceOss{
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
		},
	}
	if err, list, total := service.ResourceOssApp.GetResourceOssInfoList(pageInfo); err != nil {
		return nil, err
	} else {
		return &types.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, nil
	}
}
