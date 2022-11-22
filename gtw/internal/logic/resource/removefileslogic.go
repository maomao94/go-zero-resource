package resource

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
