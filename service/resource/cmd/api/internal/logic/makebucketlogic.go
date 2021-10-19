package logic

import (
	"context"
	"go-zero-resource/service/resource/cmd/api/ossx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	template, err := ossx.Template(req.TenantId, req.Code)
	if err != nil {
		return err
	} else {
		bool, err := template.BucketExists(req.BucketName)
		if err != nil {
			return err
		}
		if !bool {
			err = template.MakeBucket(req.BucketName)
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
