package gtw

import (
	"bytes"
	"context"
	"gtw/resource/pb"
	"io"
	"net/http"
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
	getFileResp, err := l.svcCtx.ResourceRpc.GetFile(l.ctx, &pb.GetFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filename,
	})
	if err != nil {
		return nil, err
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+getFileResp.Filename+"\"")
	w.Header().Set("Content-Type", getFileResp.ContentType)
	reader := bytes.NewReader(getFileResp.Stream)
	w.Header().Set("Content-Length", strconv.FormatInt(reader.Size(), 10))
	if _, err := io.Copy(w, reader); err != nil {
		return nil, err
	}
	return
}
