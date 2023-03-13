package resource

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
	"github.com/jinzhu/copier"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

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
	typeFile := "tempFile"
	dayStr := carbon.Now().Format("20060102")
	dirPath := l.svcCtx.Config.NfsRootPath + "/" + typeFile + "/" + dayStr
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	u, _ := uuid.NewUUID()
	path := dirPath + "/" + strings.Replace(fmt.Sprintf("%s", u), "-", "", -1) + path.Ext(fileHeader.Filename)
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buffer := bufio.NewReader(file)
	_, err = io.Copy(f, buffer)
	if err != nil {
		return nil, err
	}
	putFileResp, err := l.svcCtx.ResourceRpc.PutFile(l.ctx, &pb.PutFileReq{
		TenantId:    req.TenantId,
		Code:        req.Code,
		BucketName:  req.BucketName,
		Path:        path,
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
