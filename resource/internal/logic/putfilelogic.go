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
	file, err := ossTemplate.PutFile(nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var pbFile pb.File
	_ = copier.Copy(&pbFile, file)
	return &pb.PutFileResp{File: &pbFile}, nil
}
