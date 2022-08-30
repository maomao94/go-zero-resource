package logic

import (
	"context"
	"gtw/common/ossx"
	"gtw/model"

	"gtw/resource/internal/svc"
	"gtw/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFileLogic {
	return &RemoveFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFileLogic) RemoveFile(in *pb.RemoveFileReq) (*pb.Empty, error) {
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	err = ossTemplate.RemoveFile(in.TenantId, in.BucketName, in.Filename)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
