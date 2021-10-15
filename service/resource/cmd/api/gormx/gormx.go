package gormx

import (
	"database/sql"
	"go-zero-resource/service/resource/cmd/api/internal/config"
	"go-zero-resource/service/resource/model/gormx"
	"os"

	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/tal-tech/go-zero/core/stores/cache"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	// ExecFn defines the sql exec method.
	ExecFn func(conn sqlx.SqlConn) (sql.Result, error)
	// IndexQueryFn defines the query method that based on unique indexes.
	IndexQueryFn func(conn sqlx.SqlConn, v interface{}) (interface{}, error)
	// PrimaryQueryFn defines the query method that based on primary keys.
	PrimaryQueryFn func(conn sqlx.SqlConn, v, primary interface{}) error
	// QueryFn defines the query method.
	QueryFn func(conn sqlx.SqlConn, v interface{}) error

	CachedDb struct {
		Db    *gorm.DB
		cache *cache.Cache
	}
)

func Gormx(config config.Config) *CachedDb {
	switch "mysql" {
	case "mysql":
		return GormMysql(config)
	default:
		return GormMysql(config)
	}
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		gormx.ResourceOss{},
	)
	if err != nil {
		logx.Errorf("register table failed %s", err)
		os.Exit(0)
	}
	logx.Info("register table success")
}

func GormMysql(config config.Config) *CachedDb {
	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       config.Mysql.DataSource, // DSN data source name
		DefaultStringSize:         191,                     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(config.Mysql.LogMode, config.Mysql.Logx)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.Mysql.MaxOpenConns)
		return &CachedDb{
			Db:    db,
			cache: nil,
		}
	}
}

func gormConfig(logMode string, logx bool) *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	setLogx(logx)
	switch logMode {
	case "silent", "Silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(logger.Info)
	default:
		config.Logger = Default.LogMode(logger.Info)
	}
	return config
}
