package logic

import (
	"context"
	"go-zero-resource/common/api"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/mode"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OssListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OssListLogic {
	return OssListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssListLogic) OssList(req types.OssListReq) (*types.PageResult, error) {
	// todo: add your logic here and delete this line
	search := mode.ResourceOssSearch{
		PageInfo: api.PageInfo{
			PageSize: req.PageSize,
			Page:     req.Page,
		},
	}
	service.ResourceOssApp.GetResourceOssInfoList(search)
	return &types.PageResult{}, nil
}
