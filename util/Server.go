package util

import "github/qm012/nacos-adress/global"

func GetStandaloneMode() bool {
	return global.Server.Mode.Standalone
}

func IsSetRedis() bool {
	return len(global.Server.Redis.Address) == 0
}
