package svc

import (
	"go-zero-resource/common/gormx"
	"go-zero-resource/service/resource/cmd/api/internal/config"
	"go-zero-resource/service/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	CachedConn *gormx.CachedConn
)

type ServiceContext struct {
	Config           config.Config
	resourceOssModel model.ResourceOssModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	CachedConn = gormx.Gormx(c.Mysql, c.CacheRedis)
	gormx.MysqlTables(CachedConn.Db)
	return &ServiceContext{
		Config:           c,
		resourceOssModel: model.NewResourceOssModel(conn, c.CacheRedis),
	}
}
