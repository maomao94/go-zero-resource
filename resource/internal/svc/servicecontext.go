package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gtw/model"
	"gtw/resource/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	TOssModel model.TOssModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		TOssModel: model.NewTOssModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
