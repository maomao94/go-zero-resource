package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/handler"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

var configFile = flag.String("f", "etc/gtw.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusBadRequest, e.Data()
		case (interface {
			GRPCStatus() *status.Status
		}):
			code := errorx.CodeFromGrpcError(err)
			return code, errorx.FromError(ctx, err)
		default:
			return http.StatusBadRequest, errorx.FromError(ctx, err)
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
