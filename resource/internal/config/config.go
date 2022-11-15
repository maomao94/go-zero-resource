package config

import (
	"github.com/hehanpeng/go-zero-resource/common/ossx/osssconfig"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
	Oss   osssconfig.OssConf
}
