package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource   string
		MaxIdleConns int    // 空闲中的最大连接数
		MaxOpenConns int    // 打开到数据库的最大连接数
		LogMode      string // 是否开启Gorm全局日志
		Logx         bool   // 是否通过log-zero写入日志文件
	}
	CacheRedis cache.CacheConf
}
