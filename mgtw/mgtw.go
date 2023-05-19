package main

import (
	"flag"
	"fmt"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/server"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/mgtw/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mgtw.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	restServer := rest.MustNewServer(c.RestConfig)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMgtwServer(grpcServer, server.NewMgtwServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceGroup.Add(restServer)
	serviceGroup.Add(s)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting server at %s:%d...\n", c.RestConfig.Host, c.RestConfig.Port)
	s.Start()
}
