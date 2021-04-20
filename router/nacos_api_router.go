package router

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/api"
)

// init nacos client and server api router
func InitNacosApiRouter(router *gin.RouterGroup) {
	apiRouter := router.Group("/nacos")
	{
		apiRouter.GET("/serverlist", api.NacosClient)
		apiRouter.GET("/server/serverlist", api.NacosServer)
	}
}
