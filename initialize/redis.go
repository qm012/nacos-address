package initialize

import (
	"github.com/go-redis/redis"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/util"
	"go.uber.org/zap"
)

func InitRedis() {
	if util.GetStorageModel() != util.StorageModelRedis {
		global.Log.Info("redis no setting,use cache model storage")
		return
	}
	redisConfig := util.GetRedisConfig()
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := global.Rdb.Ping().Result()
	if err != nil {
		global.Log.Fatal("redis init failed", zap.Error(err))
		return
	}
	global.Log.Info("redis init success")
}
