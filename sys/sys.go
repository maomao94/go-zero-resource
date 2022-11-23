package main

import (
	"flag"
	"fmt"
	interceptor "github.com/hehanpeng/go-zero-resource/common/Interceptor/rpcserver"
	"github.com/hehanpeng/go-zero-resource/sys/internal/config"
	"github.com/hehanpeng/go-zero-resource/sys/internal/server"
	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSysServer(grpcServer, server.NewSysServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
