package svc

import (
	"github.com/hehanpeng/go-zero-resource/sys/internal/config"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type ServiceContext struct {
	Config config.Config
	SsoSvc httpc.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SsoSvc: httpc.NewService("ssoSvc"),
	}
}
