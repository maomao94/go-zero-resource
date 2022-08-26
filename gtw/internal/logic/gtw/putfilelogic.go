package gtw

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gtw/resource/pb"
	"io/ioutil"
	"mime/multipart"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileLogic {
	return &PutFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutFileLogic) PutFile(req *types.PutFileReq, fileHeader *multipart.FileHeader) (resp *types.File, err error) {
	if fileHeader == nil {
		return nil, errors.New("fileHeader error")
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stream, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	putFileResp, err := l.svcCtx.ResourceRpc.PutFile(l.ctx, &pb.PutFileReq{
		TenantId:   req.TenantId,
		Code:       req.Code,
		BucketName: req.BucketName,
		Stream:     stream,
		Filename:   fileHeader.Filename,
	})
	if err != nil {
		return nil, err
	}
	var respFile types.File
	_ = copier.Copy(&respFile, putFileResp.File)
	return &respFile, nil
}
