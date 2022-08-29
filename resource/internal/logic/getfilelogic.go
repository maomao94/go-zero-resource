package logic

import (
	"context"

	"gtw/resource/internal/svc"
	"gtw/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileLogic {
	return &GetFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileLogic) GetFile(in *pb.GetFileReq) (*pb.GetFileResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFileResp{}, nil
}
