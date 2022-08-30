package logic

import (
	"context"
	"go-zero-resource/service/resource/cmd/api/service"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) OssDetailLogic {
	return OssDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssDetailLogic) OssDetail(req types.BaseResult) (*types.Oss, error) {
	if err, oss := service.ResourceOssApp.GetResourceOss(req.Id); err != nil {
		return nil, err
	} else {
		return &types.Oss{
			TenantId:   oss.TenantId,
			Category:   oss.Category,
			OssCode:    oss.OssCode,
			Endpoint:   oss.Endpoint,
			AccessKey:  oss.AccessKey,
			SecretKey:  oss.SecretKey,
			BucketName: oss.BucketName,
			AppId:      oss.AppId,
			Region:     oss.Region,
			Remark:     oss.Remark,
			Status:     oss.Status,
		}, nil
	}
}
