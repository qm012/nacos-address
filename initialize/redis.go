package initialize

import (
	"github.com/go-redis/redis"
	"github/qm012/nacos-adress/global"
)

func InitRedis() error {
	redisConfig := global.Server.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := global.Rdb.Ping().Result()
	return err
}
