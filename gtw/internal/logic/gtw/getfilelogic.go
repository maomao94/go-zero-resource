package gtw

import (
	"context"
	"gtw/resource/pb"
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
	writer http.ResponseWriter
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer http.ResponseWriter) *GetFileLogic {
	return &GetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *GetFileLogic) GetFile(req *types.GetFileReq) (resp *types.EmptyReply, err error) {
	getFileResp, err := l.svcCtx.ResourceRpc.GetFile(l.ctx, &pb.GetFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Filename:   req.Filename,
	})
	if err != nil {
		return nil, err
	}
	l.writer.Header().Set("Content-Disposition", "attachment; filename=\""+getFileResp.Filename+"\"")
	l.writer.Header().Set("Content-Type", getFileResp.ContentType)
	//reader := bytes.NewReader(getFileResp.Stream)
	//l.writer.Header().Set("Content-Length", strconv.FormatInt(reader.Size(), 10))
	l.writer.Header().Set("Content-Length", strconv.Itoa(len(getFileResp.Stream)))
	l.writer.Write(getFileResp.Stream)
	//if _, err := io.Copy(l.writer, reader); err != nil {
	//	return nil, err
	//}
	return
}
