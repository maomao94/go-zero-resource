package resource

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
)

type DownloadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
	w      http.ResponseWriter
}

func NewDownloadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) *DownloadFileLogic {
	return &DownloadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
		w:      w,
	}
}

func (l *DownloadFileLogic) DownloadFile(req *types.DownloadFileReq) error {
	stat, err := os.Stat(req.Path)
	if err != nil {
		return err
	}
	//bytes, err := os.ReadFile(req.Path)
	//if err != nil {
	//	return err
	//}
	l.w.Header().Set("Content-Disposition", "attachment; filename=\""+stat.Name()+"\"")
	//l.w.Header().Set("Content-Type", http.DetectContentType(bytes))
	//l.w.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))
	//n, err := l.w.Write(bytes)
	//if err != nil {
	//	return err
	//}
	//if n < int(stat.Size()) {
	//	return io.ErrClosedPipe
	//}
	http.ServeFile(l.w, l.r, req.Path)
	return nil
}
