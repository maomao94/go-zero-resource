package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/jinzhu/copier"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatFileLogic {
	return &StatFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatFileLogic) StatFile(req *types.StatFileReq) (resp *types.OssFile, err error) {
	statFileResp, err := l.svcCtx.ResourceRpc.StatFile(l.ctx, &pb.StatFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filename,
	})
	if err != nil {
		return nil, err
	}
	var respOssFile types.OssFile
	_ = copier.Copy(&respOssFile, statFileResp.OssFile)
	return &respOssFile, nil
}
