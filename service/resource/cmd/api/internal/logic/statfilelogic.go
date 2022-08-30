package logic

import (
	"context"
	"go-zero-resource/common/ossx"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatFileLogic {
	return StatFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatFileLogic) StatFile(req types.StatFileReq) (*types.OssFile, error) {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return nil, err
	} else {
		file, err := template.StatFile(req.TenantId, req.BucketName, req.FileName)
		if err != nil {
			return nil, err
		} else {
			return &types.OssFile{
				Link:        file.Link,
				Name:        file.Name,
				Length:      file.Length,
				PutTime:     file.PutTime.Format("2006-01-02 15:04:05"),
				ContentType: file.ContentType,
			}, nil
		}
	}
}
