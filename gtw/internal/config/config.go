package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	SysRpcConf      zrpc.RpcClientConf
	ResourceRpcConf zrpc.RpcClientConf
	MessageRpcConf  zrpc.RpcClientConf
	HelloRpcConf    zrpc.RpcClientConf
}
