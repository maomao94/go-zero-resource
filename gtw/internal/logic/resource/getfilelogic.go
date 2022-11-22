package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"io"
	"net/http"
	"strconv"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *GetFileLogic {
	return &GetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
	}
}

func (l *GetFileLogic) GetFile(req *types.GetFileReq) error {
	getFileResp, err := l.svcCtx.ResourceRpc.GetFile(l.ctx, &pb.GetFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filename,
	})
	if err != nil {
		return err
	}
	l.w.Header().Set("Content-Disposition", "attachment; filename=\""+getFileResp.Filename+"\"")
	l.w.Header().Set("Content-Type", getFileResp.ContentType)
	//reader := bytes.NewReader(getFileResp.Stream)
	//l.writer.Header().Set("Content-Length", strconv.FormatInt(reader.Size(), 10))
	l.w.Header().Set("Content-Length", strconv.Itoa(len(getFileResp.Stream)))
	//if _, err := io.Copy(l.writer, reader); err != nil {
	//	return nil, err
	//}
	n, err := l.w.Write(getFileResp.Stream)
	if err != nil {
		return err
	}
	if n < len(getFileResp.Stream) {
		return io.ErrClosedPipe
	}
	return nil
}
