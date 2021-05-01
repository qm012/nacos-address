package router

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/api"
)

func InitHtmlRouter(router *gin.RouterGroup) {
	html := router.Group("")
	{
		html.GET("/index", api.Index)
	}
}
