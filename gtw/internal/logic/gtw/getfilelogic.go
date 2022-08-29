package gtw

import (
	"context"
	"io"
	"net/http"
	"path"
	"strconv"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileLogic {
	return &GetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileLogic) GetFile(req *types.GetFileReq, w http.ResponseWriter) (resp *types.EmptyReply, err error) {
	w.Header().Set("Content-Disposition", "attachment; filename=\""+path.Base(fileStat.Key)+"\"")
	w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size, 10))
	object.Seek(0, 0)
	if _, err := io.Copy(w, object); err != nil {
		if err != nil {
			return err
		}
	}
	return
}
