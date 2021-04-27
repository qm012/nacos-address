package main

import (
	"fmt"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/initialize"
	"go.uber.org/zap"
)

func main() {

	initialize.InitServerConfig()
	initialize.InitLogger()
	initialize.InitRedis()
	initialize.InitStorageMgr()

	r := initialize.Routers()

	info := `
	nacos-address console

	http://127.0.0.1:8849/index

	see and operate data by URL
`
	fmt.Print(info)
	if err := r.Run(":8849"); err != nil {
		global.Log.Fatal("startup failed", zap.Int16("port", 8849))
	}
}
