package svc

import (
	"github.com/hehanpeng/go-zero-resource/gtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/resource/resource"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ResourceRpc resource.Resource
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ResourceRpc: resource.NewResource(zrpc.MustNewClient(c.ResourceRpcConf)),
	}
}
