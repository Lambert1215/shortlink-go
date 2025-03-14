package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf       zrpc.RpcClientConf
	GroupRpcConf      zrpc.RpcClientConf
	CoreRpcConf       zrpc.RpcClientConf
	ShortLinkModelUrl string
}
