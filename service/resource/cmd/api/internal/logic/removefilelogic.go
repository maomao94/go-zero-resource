package logic

import (
	"context"
	"go-zero-resource/common/ossx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveFileLogic {
	return RemoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFileLogic) RemoveFile(req types.RemoveFileReq) error {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return err
	} else {
		err := template.RemoveFile(req.TenantId, req.BucketName, req.FileName)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
