package handler

import (
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/handler/wss"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/wss",
				Handler: wss.WssHandler(serverCtx),
			},
		},
	)
}
