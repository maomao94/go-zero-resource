package logic

import (
	"context"
	"go-zero-resource/common/ossx"
	"io"
	"net/http"
	"path"
	"strconv"

	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetFileLogic {
	return GetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileLogic) GetFile(req types.GetFileReq, w http.ResponseWriter) error {
	template, err := ossx.Template(req.TenantId, req.Code, l.svcCtx.Config.Oss.TenantMode)
	if err != nil {
		return err
	} else {
		object, err := template.GetObject(req.TenantId, req.BucketName, req.FileName)
		if err != nil {
			return err
		} else {
			fileHeader := make([]byte, 512)
			fileStat, err := object.Stat()
			if err != nil {
				return err
			}
			object.Read(fileHeader)
			//获取文件名称带后缀
			//fileNameWithSuffix := path.Base(fileStat.Key)
			//获取文件的后缀(文件类型)
			//fileType := path.Ext(fileNameWithSuffix)
			//获取文件名称(不带后缀)
			//fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
			w.Header().Set("Content-Disposition", "attachment; filename=\""+path.Base(fileStat.Key)+"\"")
			w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
			w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size, 10))
			object.Seek(0, 0)
			if _, err := io.Copy(w, object); err != nil {
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
}
