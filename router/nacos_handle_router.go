package router

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/api"
)

func InitNacosHandleRouter(router *gin.RouterGroup) {
	handleRouter := router.Group("/nacos/serverlist")
	{
		handleRouter.POST("", api.AddNacosServerList)
		handleRouter.DELETE("", api.DeleteNacosServerList)
		handleRouter.DELETE("/all", api.DeleteAllNacosServerList)
	}
}
