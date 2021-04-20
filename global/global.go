package global

import (
	"github.com/go-redis/redis"
	"github/qm012/nacos-adress/config"
	"go.uber.org/zap"
)

var (
	Log    *zap.Logger
	Rdb    *redis.Client
	Server *config.Server
)
