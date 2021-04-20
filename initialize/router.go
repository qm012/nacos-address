package initialize

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/middleware"
	"github/qm012/nacos-adress/router"
)

func Routers() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// public api
	public := r.Group("")
	{
		router.InitNacosApiRouter(public)
	}
	// private api option
	private := r.Group("")
	private.Use(middleware.JwtAuth())
	{
		router.InitNacosHandleRouter(private)
	}
	global.Log.Info("init router success")
	return r
}
