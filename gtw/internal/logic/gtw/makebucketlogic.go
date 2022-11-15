package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeBucketLogic {
	return &MakeBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeBucketLogic) MakeBucket(req *types.MakeBucketReq) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.MakeBucket(l.ctx, &pb.MakeBucketReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
