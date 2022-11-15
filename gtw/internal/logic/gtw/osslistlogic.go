package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/jinzhu/copier"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssListLogic {
	return &OssListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssListLogic) OssList(req *types.OssListReq) (resp *types.OssListReply, err error) {
	ossListResp, err := l.svcCtx.ResourceRpc.OssList(l.ctx, &pb.OssListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		OrderBy:  "create_time desc",
		TenantId: req.TenantId,
		Category: req.Category,
	})
	if err != nil {
		return nil, err
	}
	var ossList []types.Oss
	if len(ossListResp.Oss) > 0 {
		for _, pbOss := range ossListResp.Oss {
			var oss types.Oss
			_ = copier.Copy(&oss, pbOss)
			ossList = append(ossList, oss)
		}
	}
	return &types.OssListReply{
		List:     ossList,
		Total:    ossListResp.Total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
