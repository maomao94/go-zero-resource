// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	gtw "gtw/gtw/internal/handler/gtw"
	"gtw/gtw/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: gtw.PingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/detail",
				Handler: gtw.OssDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/list",
				Handler: gtw.OssListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/create",
				Handler: gtw.CreateOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/update",
				Handler: gtw.UpdateOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/delete",
				Handler: gtw.DeleteOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/makeBucket",
				Handler: gtw.MakeBucketHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeBucket",
				Handler: gtw.RemoveBucketHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/statFile",
				Handler: gtw.StatFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/putFile",
				Handler: gtw.PutFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/getFile",
				Handler: gtw.GetFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeFile",
				Handler: gtw.RemoveFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeFiles",
				Handler: gtw.RemoveFilesHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gtw/v1"),
	)
}