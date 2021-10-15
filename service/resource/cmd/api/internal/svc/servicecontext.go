package svc

import (
	"go-zero-resource/service/resource/cmd/api/gormx"
	"go-zero-resource/service/resource/cmd/api/internal/config"
	"go-zero-resource/service/resource/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type ServiceContext struct {
	Config           config.Config
	resourceOssModel model.ResourceOssModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	db := gormx.Gormx(c)
	gormx.MysqlTables(db)
	DB = db
	return &ServiceContext{
		Config:           c,
		resourceOssModel: model.NewResourceOssModel(conn, c.CacheRedis),
	}
}
