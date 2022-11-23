package svc

import (
	"github.com/hehanpeng/go-zero-resource/common/Interceptor/rpcclient"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/hello/hello"
	"github.com/hehanpeng/go-zero-resource/message/message"
	"github.com/hehanpeng/go-zero-resource/resource/resource"
	"github.com/hehanpeng/go-zero-resource/sys/sys"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	SysRpc      sys.Sys
	ResourceRpc resource.Resource
	MessageRpc  message.Message
	HelloRpc    hello.Hello
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sys.NewSys(zrpc.MustNewClient(
			c.SysRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor))),
		ResourceRpc: resource.NewResource(zrpc.MustNewClient(
			c.ResourceRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor))),
		MessageRpc: message.NewMessage(zrpc.MustNewClient(
			c.MessageRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor))),
		HelloRpc: hello.NewHello(zrpc.MustNewClient(
			c.HelloRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor))),
	}
}
