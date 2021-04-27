package initialize

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/middleware"
	"github/qm012/nacos-adress/model"
	"github/qm012/nacos-adress/router"
	"github/qm012/nacos-adress/validate"
	"go.uber.org/zap"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.New()

	if vd, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := vd.RegisterValidation("verify_ip", validate.VerifyIp)
		if err != nil {
			global.Log.Fatal("register validate failed", zap.Error(err))
		}
	}
	r.Static("/static", "./site/static")
	r.LoadHTMLGlob("./site/templates/*")
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.NoRoute(func(ctx *gin.Context) {
		err := errors.New("api no correct")
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
	})

	// public api
	public := r.Group("")
	{
		router.InitNacosApiRouter(public)
		router.InitLoginRouter(public)
		router.InitHtmlRouter(public)
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
