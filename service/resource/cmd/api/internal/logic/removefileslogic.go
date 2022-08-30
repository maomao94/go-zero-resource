package logic

import (
	"context"
	"go-zero-resource/common/ossx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveFilesLogic {
	return RemoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFilesLogic) RemoveFiles(req types.RemoveFilesReq) error {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return err
	} else {
		err := template.RemoveFiles(req.TenantId, req.BucketName, req.FileNames)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
