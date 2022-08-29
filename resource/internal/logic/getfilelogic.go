package logic

import (
	"context"
	"gtw/common/ossx"
	"gtw/model"
	"gtw/resource/internal/svc"
	"gtw/resource/pb"
	"net/http"

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
	ossTemplate, err := ossx.Template(in.TenantId, in.Code, l.svcCtx.Config.Oss.TenantMode, func(tenantId, code string) (oss *model.TOss, err error) {
		return l.svcCtx.TOssModel.FindOneByTenantIdOssCode(l.ctx, in.TenantId, in.Code)
	})
	if err != nil {
		return nil, err
	}
	object, err := ossTemplate.GetObject(in.TenantId, in.BucketName, in.Filename)
	if err != nil {
		return nil, err
	}
	fileHeader := make([]byte, 512)
	fileStat, err := object.Stat()
	if err != nil {
		return nil, err
	}
	object.Read(fileHeader)
	contentType := http.DetectContentType(fileHeader)
	filename := fileStat.Key
	//获取文件名称带后缀
	//fileNameWithSuffix := path.Base(fileStat.Key)
	//获取文件的后缀(文件类型)
	//fileType := path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	//fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	return &pb.GetFileResp{
		Filename:    filename,
		ContentType: contentType,
		Stream:      nil,
	}, nil
}
