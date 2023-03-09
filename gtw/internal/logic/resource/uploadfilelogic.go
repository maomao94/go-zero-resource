package resource

import (
	"context"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/zeromicro/go-zero/core/stringx"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq) (resp *types.FileInfo, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	file, fileHeader, err := l.r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		fileHeader.Filename, fileHeader.Size, fileHeader.Header)
	typeFile := "tempFile"
	if req.MfsType == 2 {
		typeFile = "bizFile"
	}
	u, _ := uuid.NewUUID()
	dayStr := carbon.Now().Format("20060102")
	midPath := dayStr + "/" + strings.Replace(fmt.Sprintf("%s", u), "-", "", -1) + path.Ext(fileHeader.Filename)
	path := "/" + typeFile + midPath
	err = os.WriteFile(path, []byte{}, 0o600)
	if err != nil {
		return nil, err
	}
	return &types.FileInfo{
		Name:        fileHeader.Filename,
		Path:        path,
		Size:        fileHeader.Size,
		ContextType: fileHeader.Header.Get("Content-Type"),
	}, nil
}
