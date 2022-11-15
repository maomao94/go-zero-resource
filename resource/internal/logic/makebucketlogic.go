package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/ossx"
	"github.com/hehanpeng/go-zero-resource/model"

	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeBucketLogic {
	return &MakeBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeBucketLogic) MakeBucket(in *pb.MakeBucketReq) (*pb.Empty, error) {
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	bool, err := ossTemplate.BucketExists(in.TenantId, in.BucketName)
	if err != nil {
		return nil, err
	}
	if !bool {
		err = ossTemplate.MakeBucket(in.TenantId, in.BucketName)
		if err != nil {
			return nil, err
		}
	}
	return &pb.Empty{}, nil
}
