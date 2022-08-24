package gtw

import (
	"context"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

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

func (l *MakeBucketLogic) MakeBucket(req *types.MakeBucketReq) error {
	// todo: add your logic here and delete this line

	return nil
}
