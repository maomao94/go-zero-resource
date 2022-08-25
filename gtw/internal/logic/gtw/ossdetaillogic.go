package gtw

import (
	"context"
	"gtw/resource/pb"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type OssDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssDetailLogic {
	return &OssDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssDetailLogic) OssDetail(req *types.BaseResult) (resp *types.Oss, err error) {
	ossDetailResp, err := l.svcCtx.ResourceRpc.OssDetail(l.ctx, &pb.OssDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var respOss types.Oss
	_ = copier.Copy(&respOss, ossDetailResp.Oss)
	return &respOss, nil
}
