package initialize

import (
	"github/qm012/nacos-adress/api"
	"github/qm012/nacos-adress/global"
)

func InitStorageMgr() {
	global.Log.Info("init api mgr/service start")
	api.Init()
	global.Log.Info("init api mgr/service end")
}
