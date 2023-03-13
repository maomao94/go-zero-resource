package logic

import (
	"bufio"
	"context"
	"github.com/golang-module/carbon/v2"
	"github.com/hehanpeng/go-zero-resource/common/ossx"
	"github.com/hehanpeng/go-zero-resource/model"
	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"os"
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
	fileStat, err := object.Stat()
	if err != nil {
		return nil, err
	}
	//获取文件名称带后缀
	//fileNameWithSuffix := path.Base(fileStat.Key)
	//获取文件的后缀(文件类型)
	//fileType := path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	//fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	//stream, err := ioutil.ReadAll(object)
	//if err != nil {
	//	return nil, err
	//}
	typeFile := "tempFile"
	dayStr := carbon.Now().Format("20060102")
	dirPath := l.svcCtx.Config.NfsRootPath + "/" + typeFile + "/" + dayStr
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	path := dirPath + "/" + fileStat.Key
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buffer := bufio.NewReader(object)
	_, err = io.Copy(f, buffer)
	if err != nil {
		return nil, err
	}
	return &pb.GetFileResp{
		Filename:    fileStat.Key,
		ContentType: fileStat.ContentType,
		Path:        path,
	}, nil
}
