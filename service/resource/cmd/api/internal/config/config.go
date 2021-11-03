package config

import (
	"go-zero-resource/common/gormx"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	gormx.MysqlConf
	cache.CacheConf
}
