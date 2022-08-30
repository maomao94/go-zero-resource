package gtw

import (
	"context"
	"gtw/resource/pb"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFileLogic {
	return &RemoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFileLogic) RemoveFile(req *types.RemoveFileReq) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.RemoveFile(l.ctx, &pb.RemoveFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filename,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
