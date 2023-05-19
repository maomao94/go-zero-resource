// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	gtw "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/gtw"
	message "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/message"
	resource "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/resource"
	sysuser "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/sys/user"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"

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
		},
		rest.WithPrefix("/gtw/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: sysuser.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: sysuser.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gtw/sys/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/kq/send",
				Handler: message.KqSendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sendOneMsgToUser",
				Handler: message.SendOneMsgToUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gtw/message/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oss/detail",
				Handler: resource.OssDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/list",
				Handler: resource.OssListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/create",
				Handler: resource.CreateOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/update",
				Handler: resource.UpdateOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/delete",
				Handler: resource.DeleteOssHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/makeBucket",
				Handler: resource.MakeBucketHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeBucket",
				Handler: resource.RemoveBucketHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/statFile",
				Handler: resource.StatFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/putFile",
				Handler: resource.PutFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/getFile",
				Handler: resource.GetFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeFile",
				Handler: resource.RemoveFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/removeFiles",
				Handler: resource.RemoveFilesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mfs/uploadFile",
				Handler: resource.UploadFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/mfs/downloadFile",
				Handler: resource.DownloadFileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/gtw/resource/v1"),
	)
}
