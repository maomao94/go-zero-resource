package resource

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"net/http"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

type PutFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PutFileLogic {
	return &PutFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PutFileLogic) PutFile(req *types.PutFileReq) (resp *types.File, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	file, fileHeader, err := l.r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		fileHeader.Filename, fileHeader.Size, fileHeader.Header)
	stream, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	putFileResp, err := l.svcCtx.ResourceRpc.PutFile(l.ctx, &pb.PutFileReq{
		TenantId:    req.TenantId,
		Code:        req.Code,
		BucketName:  req.BucketName,
		Stream:      stream,
		Filename:    fileHeader.Filename,
		ContentType: fileHeader.Header.Get("content-type"),
	})
	if err != nil {
		return nil, err
	}
	var respFile types.File
	_ = copier.Copy(&respFile, putFileResp.File)
	return &respFile, nil
}
