package config

import (
	"go-zero-resource/common/gormx"
	"go-zero-resource/common/ossx/osssconfig"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql      gormx.MysqlConf
	CacheRedis cache.CacheConf
	Oss        osssconfig.OssConf
}
