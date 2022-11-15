package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOssLogic {
	return &CreateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOssLogic) CreateOss(req *types.OssCreate) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.CreateOss(l.ctx, &pb.CreateOssReq{
		TenantId:   req.TenantId,
		Category:   req.Category,
		OssCode:    req.OssCode,
		Endpoint:   req.Endpoint,
		AccessKey:  req.AccessKey,
		SecretKey:  req.SecretKey,
		BucketName: req.BucketName,
		AppId:      req.AppId,
		Region:     req.Region,
		Remark:     req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
