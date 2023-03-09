package resource

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"io"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.DownloadFileReq) error {
	f, err := fileutil.OpenDir(req.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	l.w.Header().Set("Content-Disposition", "attachment; filename=\""+f.Name()+"\"")
	l.w.Header().Set("Content-Type", "application/octet-stream")
	l.w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	//if _, err := io.Copy(l.writer, reader); err != nil {
	//	return nil, err
	//}
	n, err := l.w.Write(bytes)
	if err != nil {
		return err
	}
	if n < len(strconv.Itoa(len(bytes))) {
		return io.ErrClosedPipe
	}
	return nil
}
