package logic

import (
	"context"
	"go-zero-resource/common/ossx"
	"go-zero-resource/common/utils"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) MakeBucketLogic {
	return MakeBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeBucketLogic) MakeBucket(req types.MakeBucketReq) error {
	err := utils.Verify(req, utils.TenantIdVerify)
	if err != nil {
		return err
	}
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return err
	} else {
		bool, err := template.BucketExists(req.TenantId, req.BucketName)
		if err != nil {
			return err
		}
		if !bool {
			err = template.MakeBucket(req.TenantId, req.BucketName)
			if err != nil {
				return err
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
}
