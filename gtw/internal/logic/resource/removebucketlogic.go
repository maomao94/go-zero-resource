package resource

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveBucketLogic {
	return &RemoveBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveBucketLogic) RemoveBucket(req *types.RemoveBucketReq) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.RemoveBucket(l.ctx, &pb.RemoveBucketReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
