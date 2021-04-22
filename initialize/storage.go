package initialize

import (
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/service"
)

func InitStorageMgr() {
	global.Log.Info("init storage mgr start")
	global.StorageMgr = service.NewStorageMgr()
	global.Log.Info("init storage mgr end")
}
