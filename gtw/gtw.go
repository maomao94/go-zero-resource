package main

import (
	"flag"
	"fmt"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/handler"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
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

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusInternalServerError, e.Data()
		case (interface {
			GRPCStatus() *status.Status
		}):
			errcode
			return errcode, e.GRPCStatus().Message()
		default:
			return http.StatusInternalServerError, e
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
