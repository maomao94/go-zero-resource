package svc

import (
	"go-zero-resource/service/resource/cmd/api/gormx"
	"go-zero-resource/service/resource/cmd/api/internal/config"

	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/tools/goctl/model/sql/test/model"
)

var (
	DB *gorm.DB
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//conn := sqlx.NewMysql(c.Mysql.DataSource)
	db := gormx.Gormx(c)
	gormx.MysqlTables(db)
	DB = db
	return &ServiceContext{
		Config: c,
	}
}
