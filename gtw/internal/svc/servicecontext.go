package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gtw/gtw/internal/config"
	"gtw/resource/resource"
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
