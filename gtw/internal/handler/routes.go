// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	gtw "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/gtw"
	message "github.com/hehanpeng/go-zero-resource/gtw/internal/handler/message"
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
		rest.WithPrefix("/gtw/v1/sys"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/message/kq/send",
				Handler: message.KqSendHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gtw/v1/message"),
	)
}
