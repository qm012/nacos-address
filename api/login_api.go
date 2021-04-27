package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/model"
	"github/qm012/nacos-adress/service"
	"github/qm012/nacos-adress/util"
	"net/http"
)

var (
	jwtService service.JWTService
)

func Login(ctx *gin.Context) {
	var credentials model.Credentials
	if err := ctx.ShouldBind(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "parameter error"})
		return
	}
	isAuthenticated := util.VerifyAccount(credentials.Username, credentials.Password)
	if isAuthenticated {
		token := jwtService.GenerateToken(credentials.Username, true)
		ctx.JSON(http.StatusOK, model.NewSuccessResultData(token))
		return
	}
	ctx.JSON(http.StatusOK, model.NewFailedResult(fmt.Errorf("please enter the correct user name and password")))
}
