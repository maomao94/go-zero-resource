package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/ossx"
	"github.com/hehanpeng/go-zero-resource/model"

	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFilesLogic {
	return &RemoveFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFilesLogic) RemoveFiles(in *pb.RemoveFilesReq) (*pb.Empty, error) {
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	err = ossTemplate.RemoveFiles(in.TenantId, in.BucketName, in.Filename)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
