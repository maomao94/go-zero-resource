package resource

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"io"
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
	if err != nil {
		return nil, err
	}
	typeFile := "tempFile"
	if req.MfsType == 2 {
		typeFile = "bizFile"
	}
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
	//b := make([]byte, 1024)
	//for {
	//	n, err := buffer.Read(b)
	//	if err != nil && err != io.EOF {
	//		//有一个特殊问题，当一个文件读读完，遇到文件末尾时，它也会返回一个错误，但是此时我已经读到文件末尾EOF，这个错误应该不算错误，所以应该把读到文件末尾这个错误给去掉。
	//		return nil, err
	//	}
	//	if err == io.EOF {
	//		break
	//	}
	//	_, err = f.Write(b[:n])
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	_, err = io.Copy(f, buffer)
	if err != nil {
		return nil, err
	}
	return &types.FileInfo{
		Name:        fileHeader.Filename,
		Path:        path,
		Size:        fileHeader.Size,
		ContextType: fileHeader.Header.Get("Content-Type"),
		Url:         l.svcCtx.Config.DownloadUrl + path,
	}, nil
}
