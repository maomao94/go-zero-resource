package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"gtw/model"

	"gtw/resource/internal/svc"
	"gtw/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOssListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssListLogic {
	return &OssListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OssListLogic) OssList(in *pb.OssListReq) (*pb.OssListResp, error) {
	whereBuilder := l.svcCtx.TOssModel.RowBuilder()
	if len(in.TenantId) > 0 {
		whereBuilder.Where(squirrel.Eq{
			"tenant_id": in.TenantId,
		})
	}
	if in.Category > 0 {
		whereBuilder.Where(squirrel.Eq{
			"category": in.Category,
		})
	}
	list, err := l.svcCtx.TOssModel.FindPageListByPage(l.ctx, whereBuilder, in.Page, in.PageSize, in.OrderBy)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	var respOss []*pb.Oss
	if len(list) > 0 {
		for _, oss := range list {
			var pbOss pb.Oss
			_ = copier.Copy(&pbOss, oss)
			pbOss.CreateTime = oss.CreateTime.Unix()
			respOss = append(respOss, &pbOss)
		}
	}
	return &pb.OssListResp{Oss: respOss}, nil
}
