package logic

import (
	"context"
	"fmt"
	"go-zero-resource/common/api"
	"go-zero-resource/common/errorx"
	"go-zero-resource/service/resource/cmd/api/service"
	"go-zero-resource/service/resource/model/gormx"

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
	pageInfo := gormx.ResourceOssSearch{
		PageInfo: api.PageInfo{
			PageSize: req.PageSize,
			Page:     req.Page,
		},
	}
	if err, list, total := service.ResourceOssApp.GetResourceOssInfoList(pageInfo); err != nil {
		return nil, errorx.NewDefaultError(fmt.Sprintf("获取失败: %v", err))
	} else {
		return &types.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, nil
	}
}
