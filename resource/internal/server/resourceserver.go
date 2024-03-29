// Code generated by goctl. DO NOT EDIT.
// Source: resource.proto

package server

import (
	"context"

	"github.com/hehanpeng/go-zero-resource/resource/internal/logic"
	"github.com/hehanpeng/go-zero-resource/resource/internal/svc"
	"github.com/hehanpeng/go-zero-resource/resource/pb"
)

type ResourceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedResourceServer
}

func NewResourceServer(svcCtx *svc.ServiceContext) *ResourceServer {
	return &ResourceServer{
		svcCtx: svcCtx,
	}
}

func (s *ResourceServer) Ping(ctx context.Context, in *pb.Empty) (*pb.PingResp, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *ResourceServer) OssDetail(ctx context.Context, in *pb.OssDetailReq) (*pb.OssDetailResp, error) {
	l := logic.NewOssDetailLogic(ctx, s.svcCtx)
	return l.OssDetail(in)
}

func (s *ResourceServer) OssList(ctx context.Context, in *pb.OssListReq) (*pb.OssListResp, error) {
	l := logic.NewOssListLogic(ctx, s.svcCtx)
	return l.OssList(in)
}

func (s *ResourceServer) CreateOss(ctx context.Context, in *pb.CreateOssReq) (*pb.Empty, error) {
	l := logic.NewCreateOssLogic(ctx, s.svcCtx)
	return l.CreateOss(in)
}

func (s *ResourceServer) UpdateOss(ctx context.Context, in *pb.UpdateOssReq) (*pb.Empty, error) {
	l := logic.NewUpdateOssLogic(ctx, s.svcCtx)
	return l.UpdateOss(in)
}

func (s *ResourceServer) DeleteOss(ctx context.Context, in *pb.DeleteOssReq) (*pb.Empty, error) {
	l := logic.NewDeleteOssLogic(ctx, s.svcCtx)
	return l.DeleteOss(in)
}

func (s *ResourceServer) MakeBucket(ctx context.Context, in *pb.MakeBucketReq) (*pb.Empty, error) {
	l := logic.NewMakeBucketLogic(ctx, s.svcCtx)
	return l.MakeBucket(in)
}

func (s *ResourceServer) RemoveBucket(ctx context.Context, in *pb.RemoveBucketReq) (*pb.Empty, error) {
	l := logic.NewRemoveBucketLogic(ctx, s.svcCtx)
	return l.RemoveBucket(in)
}

func (s *ResourceServer) StatFile(ctx context.Context, in *pb.StatFileReq) (*pb.StatFileResp, error) {
	l := logic.NewStatFileLogic(ctx, s.svcCtx)
	return l.StatFile(in)
}

func (s *ResourceServer) PutFile(ctx context.Context, in *pb.PutFileReq) (*pb.PutFileResp, error) {
	l := logic.NewPutFileLogic(ctx, s.svcCtx)
	return l.PutFile(in)
}

func (s *ResourceServer) GetFile(ctx context.Context, in *pb.GetFileReq) (*pb.GetFileResp, error) {
	l := logic.NewGetFileLogic(ctx, s.svcCtx)
	return l.GetFile(in)
}

func (s *ResourceServer) RemoveFile(ctx context.Context, in *pb.RemoveFileReq) (*pb.Empty, error) {
	l := logic.NewRemoveFileLogic(ctx, s.svcCtx)
	return l.RemoveFile(in)
}

func (s *ResourceServer) RemoveFiles(ctx context.Context, in *pb.RemoveFilesReq) (*pb.Empty, error) {
	l := logic.NewRemoveFilesLogic(ctx, s.svcCtx)
	return l.RemoveFiles(in)
}
