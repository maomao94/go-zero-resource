package logic

import (
	"context"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveBucketLogic {
	return RemoveBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveBucketLogic) RemoveBucket(req types.RemoveBucketReq) error {
	// todo: add your logic here and delete this line

	return nil
}
