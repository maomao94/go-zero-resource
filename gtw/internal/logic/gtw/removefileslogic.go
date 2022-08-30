package gtw

import (
	"context"
	"gtw/resource/pb"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFilesLogic {
	return &RemoveFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFilesLogic) RemoveFiles(req *types.RemoveFilesReq) (resp *types.EmptyReply, err error) {
	_, err = l.svcCtx.ResourceRpc.RemoveFiles(l.ctx, &pb.RemoveFilesReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filenames,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyReply{}, nil
}
