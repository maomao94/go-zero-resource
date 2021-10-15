package svc

import (
	"go-zero-resource/service/resource/cmd/api/gormx"
	"go-zero-resource/service/resource/cmd/api/internal/config"
	"go-zero-resource/service/resource/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var (
	CachedDb *gormx.CachedConn
)

type ServiceContext struct {
	Config           config.Config
	resourceOssModel model.ResourceOssModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	CachedDb = gormx.Gormx(c)
	gormx.MysqlTables(CachedDb.Db)
	return &ServiceContext{
		Config:           c,
		resourceOssModel: model.NewResourceOssModel(conn, c.CacheRedis),
	}
}
