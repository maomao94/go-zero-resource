package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/model"

	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOssLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOssLogic {
	return &UpdateOssLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOssLogic) UpdateOss(in *pb.UpdateOssReq) (*pb.Empty, error) {
	err := l.svcCtx.TOssModel.Update(l.ctx, &model.TOss{
		Id:         in.Id,
		TenantId:   in.TenantId,
		Category:   in.Category,
		OssCode:    in.OssCode,
		Endpoint:   in.Endpoint,
		AccessKey:  in.AccessKey,
		SecretKey:  in.SecretKey,
		BucketName: in.BucketName,
		AppId:      in.AppId,
		Region:     in.Region,
		Remark:     in.Remark,
		Status:     in.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
