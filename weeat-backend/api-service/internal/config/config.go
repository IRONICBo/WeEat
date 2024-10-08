package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	AIServiceDiscovery discov.EtcdConf
}
