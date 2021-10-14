package main

import (
	"flag"
	"fmt"
    "github.com/tal-tech/go-zero/rest/httpx"
    "go-zero-resource/common/errorx"
    "net/http"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

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
