package logic

import (
	"context"
	"go-zero-resource/common/ossx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveBucketLogic {
	return RemoveBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveBucketLogic) RemoveBucket(req types.RemoveBucketReq) error {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return err
	} else {
		err := template.RemoveBucket(req.TenantId, req.BucketName)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
