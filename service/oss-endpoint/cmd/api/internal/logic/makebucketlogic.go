package logic

import (
	"context"
	utils "go-zero-resource/common/util"
	"go-zero-resource/service/oss-endpoint/cmd/api/internal/svc"
	"go-zero-resource/service/oss-endpoint/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MakeBucketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) MakeBucketLogic {
	return MakeBucketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeBucketLogic) MakeBucket(req types.MakeBucketReq) (*types.EmptyReply, error) {
	// todo: add your logic here and delete this line
	//if true {
	//	return nil, errorx.NewDefaultError("error")
	//}
	err := utils.Verify(req, utils.BucketNameVerify)
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
