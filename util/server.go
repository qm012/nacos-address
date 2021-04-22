package util

import "github/qm012/nacos-adress/global"

type StorageModel int8

const (
	StorageModelRedis StorageModel = iota
	StorageModelCache
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
	return global.Server.Mode.Standalone
}

func IsSetRedis() bool {
	return len(global.Server.Redis.Address) != 0
}
