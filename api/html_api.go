package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		return
	}
	ctx.HTML(http.StatusOK, "index.html", nil)
}
