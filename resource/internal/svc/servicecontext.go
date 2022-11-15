package svc

import (
	"github.com/hehanpeng/go-zero-resource/model"
	"github.com/hehanpeng/go-zero-resource/resource/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
