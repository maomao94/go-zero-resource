package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/jinzhu/copier"

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
	countBuilder := l.svcCtx.TOssModel.CountBuilder("1")
	if len(in.TenantId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"tenant_id": in.TenantId,
		})
		countBuilder = countBuilder.Where(squirrel.Eq{
			"tenant_id": in.TenantId,
		})
	}
	if in.Category > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"category": in.Category,
		})
		countBuilder = countBuilder.Where(squirrel.Eq{
			"category": in.Category,
		})
	}
	total, err := l.svcCtx.TOssModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}
	list, err := l.svcCtx.TOssModel.FindPageListByPage(l.ctx, whereBuilder, in.Page, in.PageSize, in.OrderBy)
	if err != nil {
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
	return &pb.OssListResp{Oss: respOss, Total: total}, nil
}
