package logic

import (
	"context"
	"go-zero-resource/common/ossx"
	"io"
	"net/http"
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
	template, err := ossx.Template(req.TenantId, req.Code)
	if err != nil {
		return err
	} else {
		object, err := template.GetObject(req.TenantId, req.BucketName, req.FileName)
		if err != nil {
			return err
		} else {
			fileStat, _ := object.Stat()
			w.Header().Set("Content-Disposition", "attachment; filename="+fileStat.Key)
			w.Header().Set("Content-type", fileStat.ContentType)
			w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size, 10))
			io.Copy(w, object)
			return nil
		}
	}
}
