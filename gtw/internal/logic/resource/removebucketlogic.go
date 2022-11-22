package resource

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
