package main

import (
	"flag"
	"fmt"
	"go-zero-resource/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-resource/service/resource/cmd/api/internal/config"
	"go-zero-resource/service/resource/cmd/api/internal/handler"
	"go-zero-resource/service/resource/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/resource-api.yaml", "the config file")

//go:generate goctl api go -api resource.api -dir . -home ../../../../template
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 支持跨域
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
			w.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			next(w, r)
		}
	})

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Fail()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
