package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"gtw/common/ossx"
	"gtw/model"

	"gtw/resource/internal/svc"
	"gtw/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatFileLogic {
	return &StatFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StatFileLogic) StatFile(in *pb.StatFileReq) (*pb.StatFileResp, error) {
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	ossFile, err := ossTemplate.StatFile(in.TenantId, in.BucketName, in.Filename)
	if err != nil {
		return nil, err
	}
	var respOssFile pb.OssFile
	_ = copier.Copy(&respOssFile, ossFile)
	return &pb.StatFileResp{OssFile: &respOssFile}, nil
}
