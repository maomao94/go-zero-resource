package svc

import (
	interceptor "github.com/hehanpeng/go-zero-resource/common/Interceptor"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/hello/pb"
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
	HelloRpc  pb.HelloClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sys.NewSys(zrpc.MustNewClient(
			c.SysRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.UnaryMetadataInterceptor))),
		ResourceRpc: resource.NewResource(zrpc.MustNewClient(
			c.ResourceRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.UnaryMetadataInterceptor))),
		MessageRpc: message.NewMessage(zrpc.MustNewClient(
			c.MessageRpcConf, zrpc.WithUnaryClientInterceptor(interceptor.UnaryMetadataInterceptor))),
			hellorpc
	}
}
