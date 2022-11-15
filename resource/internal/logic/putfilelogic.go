package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/ossx"
	"github.com/hehanpeng/go-zero-resource/model"
	"github.com/jinzhu/copier"

	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileLogic {
	return &PutFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutFileLogic) PutFile(in *pb.PutFileReq) (*pb.PutFileResp, error) {
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	file, err := ossTemplate.PutFileStream(in.TenantId, in.BucketName, in.Filename, in.ContentType, &in.Stream)
	if err != nil {
		return nil, err
	}
	var pbFile pb.File
	_ = copier.Copy(&pbFile, file)
	return &pb.PutFileResp{File: &pbFile}, nil
}
