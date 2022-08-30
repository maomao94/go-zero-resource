package gormx

import (
	"go-zero-resource/service/resource/model/gorm_model"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// see doc/sql-cache.md
const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

var (
	exclusiveCalls = syncx.NewSingleFlight()
	stats          = cache.NewStat("gorm_model")
)

type (
	ExecFn         func(db *gorm.DB) (int64, error)
	IndexQueryFn   func(db *gorm.DB, v interface{}) (interface{}, error)
	PrimaryQueryFn func(db *gorm.DB, v, primary interface{}) error
	QueryFn        func(db *gorm.DB, v interface{}) error

	CachedConn struct {
		Cache cache.Cache
		Db    *gorm.DB
	}

	MysqlConf struct {
		DataSource   string
		MaxIdleConns int    // 空闲中的最大连接数
		MaxOpenConns int    // 打开到数据库的最大连接数
		LogMode      string // 是否开启Gorm全局日志
		Logx         bool   // 是否通过log-zero写入日志文件
	}
)

func (cc CachedConn) DelCache(keys ...string) error {
	return cc.Cache.Del(keys...)
}

func (cc CachedConn) GetCache(key string, v interface{}) error {
	return cc.Cache.Get(key, v)
}

func (cc CachedConn) Exec(exec ExecFn, keys ...string) (int64, error) {
	rows, err := exec(cc.Db)
	if err != nil {
		return rows, err
	}

	if err := cc.DelCache(keys...); err != nil {
		return rows, err
	}

	return rows, nil
}

func (cc CachedConn) QueryRow(v interface{}, key string, query QueryFn) error {
	return cc.Cache.Take(v, key, func(v interface{}) error {
		return query(cc.Db, v)
	})
}

func (cc CachedConn) QueryRowIndex(v interface{}, key string, keyer func(primary interface{}) string,
	indexQuery IndexQueryFn, primaryQuery PrimaryQueryFn) error {
	var primaryKey interface{}
	var found bool

	if err := cc.Cache.TakeWithExpire(&primaryKey, key, func(val interface{}, expire time.Duration) (err error) {
		primaryKey, err = indexQuery(cc.Db, v)
		if err != nil {
			return
		}

		found = true
		return cc.Cache.SetWithExpire(keyer(primaryKey), v, expire+cacheSafeGapBetweenIndexAndPrimary)
	}); err != nil {
		return err
	}

	if found {
		return nil
	}

	return cc.Cache.Take(v, keyer(primaryKey), func(v interface{}) error {
		return primaryQuery(cc.Db, v, primaryKey)
	})
}

func (cc CachedConn) SetCache(key string, v interface{}) error {
	return cc.Cache.Set(key, v)
}

//func (cc CachedConn) Transact(fn func(sqlx.Session) error) error {
//	return cc.Db.Transact(fn)
//}

func Gormx(config MysqlConf, cacheConf cache.CacheConf) *CachedConn {
	switch "mysql" {
	case "mysql":
		return GormMysql(config, cacheConf)
	default:
		return GormMysql(config, cacheConf)
	}
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		gorm_model.ResourceOss{},
	)
	if err != nil {
		logx.Errorf("register table failed %s", err)
		os.Exit(0)
	}
	logx.Info("register table success")
}

func GormMysql(config MysqlConf, cacheConf cache.CacheConf) *CachedConn {
	//dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       config.DataSource, // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,              // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,              // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,              // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,             // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(config.LogMode, config.Logx)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		return &CachedConn{
			Cache: cache.New(cacheConf, exclusiveCalls, stats, gorm.ErrRecordNotFound),
			Db:    db,
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
