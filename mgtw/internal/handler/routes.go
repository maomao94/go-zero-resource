package handler

import (
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/handler/ws"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/wsx",
				Handler: ws.WsHandler(serverCtx),
			},
		},
	)
}
