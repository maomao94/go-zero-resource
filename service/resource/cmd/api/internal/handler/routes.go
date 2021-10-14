// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-resource/service/resource/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oss/list",
				Handler: ossListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/make-bucket",
				Handler: makeBucketHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oss/endpoint/put-file",
				Handler: putFileHandler(serverCtx),
			},
		},
	)
}
