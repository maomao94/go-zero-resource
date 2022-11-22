package svc

import (
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		SysRpc:      sys.NewSys(zrpc.MustNewClient(c.SysRpcConf)),
		ResourceRpc: resource.NewResource(zrpc.MustNewClient(c.ResourceRpcConf)),
		MessageRpc:  message.NewMessage(zrpc.MustNewClient(c.MessageRpcConf)),
	}
}
