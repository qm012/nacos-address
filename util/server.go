package util

import (
	"github/qm012/nacos-adress/config"
	"github/qm012/nacos-adress/global"
	"go.uber.org/zap"
	"os"
	"strconv"
)

type StorageModel int8

const (
	StorageModelRedis StorageModel = iota
	StorageModelCache
	Cluster    = "cluster"
	Standalone = "standalone"
)

var (
	accountUsername string
	accountPassword string
)

func GetStorageModel() StorageModel {
	if getStandaloneMode() {
		if IsSetRedis() {
			return StorageModelRedis
		}
		return StorageModelCache
	}
	return StorageModelRedis
}

func getStandaloneMode() bool {
	mode := os.Getenv("APP_MODE")
	global.Log.Info("get mod value", zap.String("APP_MODE", mode))
	if len(mode) == 0 {
		if len(global.Server.App.Mode) == 0 {
			return true
		}
		return global.Server.App.Mode == Standalone
	}
	return mode == Standalone
}

func IsSetRedis() bool {
	if len(os.Getenv("REDIS_HOST")) != 0 {
		return true
	}
	return len(global.Server.Redis.Address) != 0
}

func getUsername() string {
	username := os.Getenv("ACCOUNT_USERNAME")
	if len(username) != 0 {
		return username
	}
	if len(global.Server.Account.Username) != 0 {
		return global.Server.Account.Username
	}
	return "nacos"
}
func getPassword() string {
	password := os.Getenv("ACCOUNT_PASSWORD")
	if len(password) != 0 {
		return password
	}
	if len(global.Server.Account.Password) != 0 {
		return global.Server.Account.Password
	}
	return "nacos"
}

func VerifyAccount(username, password string) bool {
	if len(accountUsername) == 0 {
		accountUsername = getUsername()
	}
	if len(accountPassword) == 0 {
		accountPassword = getPassword()
	}
	return accountUsername == username &&
		accountPassword == password
}

func GetRedisConfig() config.Redis {
	rds := config.Redis{}
	addr := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	var db int
	redisDb := os.Getenv("REDIS_DB")
	if len(redisDb) != 0 {
		if atoi, err := strconv.Atoi(redisDb); err == nil {
			db = atoi
		}
	} else {
		db = global.Server.Redis.DB
	}

	rds.Address = If(len(addr) != 0, addr, global.Server.Redis.Address).(string)
	rds.Password = If(len(password) != 0, password, global.Server.Redis.Password).(string)
	rds.DB = db
	return rds
}
