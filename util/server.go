package util

import (
	"github/qm012/nacos-adress/global"
	"go.uber.org/zap"
	"os"
	"strconv"
)

type StorageModel int8

const (
	StorageModelRedis StorageModel = iota
	StorageModelCache
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

	standalone := os.Getenv("standalone")
	if len(standalone) == 0 {
		return global.Server.Mode.Standalone
	}
	alone, err := strconv.ParseBool(standalone)
	if err != nil {
		global.Log.Info("Get env standalone parse", zap.Error(err))
		return true
	}
	return alone
}

func IsSetRedis() bool {
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
