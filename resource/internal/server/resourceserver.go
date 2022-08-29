// Code generated by goctl. DO NOT EDIT!
// Source: resource.proto

package server

import (
	"context"

	"gtw/resource/internal/logic"
	"gtw/resource/internal/svc"
	"gtw/resource/pb"
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

func (s *ResourceServer) MakeBucket(ctx context.Context, in *pb.MakeBucketReq) (*pb.Empty, error) {
	l := logic.NewMakeBucketLogic(ctx, s.svcCtx)
	return l.MakeBucket(in)
}

func (s *ResourceServer) PutFile(ctx context.Context, in *pb.PutFileReq) (*pb.PutFileResp, error) {
	l := logic.NewPutFileLogic(ctx, s.svcCtx)
	return l.PutFile(in)
}

func (s *ResourceServer) GetFile(ctx context.Context, in *pb.GetFileReq) (*pb.GetFileResp, error) {
	l := logic.NewGetFileLogic(ctx, s.svcCtx)
	return l.GetFile(in)
}
