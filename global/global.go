package global

import (
	"github.com/go-redis/redis"
	"github/qm012/nacos-adress/config"
	"github/qm012/nacos-adress/service"
	"go.uber.org/zap"
)

var (
	Log        *zap.Logger
	Rdb        *redis.Client
	Server     *config.Server
	StorageMgr *service.StorageMgr
)
