package svc

import (
	"go-zero-resource/service/resource/cmd/api/gormx"
	"go-zero-resource/service/resource/cmd/api/internal/config"

	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/tools/goctl/model/sql/test/model"
)

type ServiceContext struct {
	Config    config.Config
	Db        gorm.DB
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//conn := sqlx.NewMysql(c.Mysql.DataSource)
	Db := gormx.Gormx(c)
	gormx.MysqlTables(Db)
	return &ServiceContext{
		Config: c,
	}
}
