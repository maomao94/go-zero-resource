package svc

import (
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/config"
	"github.com/hehanpeng/go-zero-resource/mgtw/manager"
)

type ServiceContext struct {
	Config        config.Config
	ClientManager *manager.ClientManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ClientManager: manager.NewClientManager(),
	}
}
