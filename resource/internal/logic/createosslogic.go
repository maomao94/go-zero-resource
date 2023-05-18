package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/ctxdata"
	"github.com/hehanpeng/go-zero-resource/model"
	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOssLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOssLogic {
	return &CreateOssLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOssLogic) CreateOss(in *pb.CreateOssReq) (*pb.Empty, error) {
	userId := ctxdata.GetUserIdFromMetadata(l.ctx)
	_, err := l.svcCtx.TOssModel.Insert(l.ctx, &model.TOss{
		CreateUser: userId,
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
		Status:     1,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
