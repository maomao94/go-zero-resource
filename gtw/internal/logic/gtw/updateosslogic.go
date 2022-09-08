package gtw

import (
	"context"
	"gtw/resource/pb"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOssLogic {
	return &UpdateOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOssLogic) UpdateOss(req *types.OssUpdate) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.UpdateOss(l.ctx, &pb.UpdateOssReq{
		Id:         req.Id,
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
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
