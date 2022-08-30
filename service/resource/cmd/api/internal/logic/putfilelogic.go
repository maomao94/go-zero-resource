package logic

import (
	"context"
	"go-zero-resource/common/ossx"
	"mime/multipart"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutFileLogic {
	return PutFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutFileLogic) PutFile(req types.PutFileReq, file *multipart.FileHeader) (*types.File, error) {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return nil, err
	} else {
		f, err := template.PutFile(req.TenantId, req.BucketName, file)
		if err != nil {
			return nil, err
		} else {
			return &types.File{
				Link:         f.Link,
				Domain:       f.Domain,
				Name:         f.Name,
				OriginalName: f.OriginalName,
				AttachId:     f.AttachId,
			}, nil
		}
	}
}
