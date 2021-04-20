package main

import (
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/initialize"
	"go.uber.org/zap"
)

func main() {

	initialize.InitServerConfig()
	initialize.InitLogger()

	r := initialize.Routers()

	if err := r.Run(":8849"); err != nil {
		global.Log.Fatal("startup failed", zap.Int16("port", 8849))
	}
}
