package router

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/api"
)

func InitLoginRouter(router *gin.RouterGroup) {
	r := router.Group("/login")
	{
		r.POST("", api.Login)
	}
}
