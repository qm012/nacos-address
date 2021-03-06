package api

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/model"
	"github/qm012/nacos-adress/service"
	"go.uber.org/zap"
	"net/http"
)

var storageMgr *service.StorageMgr

func Init() {
	storageMgr = service.NewStorageMgr()
	jwtService = service.NewJwtService()
}

// nacos client get serverList
func NacosClient(ctx *gin.Context) {
	ipStr, err := storageMgr.Get()
	if err != nil {
		global.Log.Error("get ipStr error", zap.Error(err))
		ctx.String(http.StatusOK, err.Error())
		return
	}
	ctx.String(http.StatusOK, ipStr)
}

// nacosserver get serverList
func NacosServer(ctx *gin.Context) {
	ipStr, err := storageMgr.Get()
	if err != nil {
		result := model.NewFailedResult(err)
		ctx.JSON(http.StatusOK, result)
		return
	}
	result := model.NewSuccessResultData(ipStr)
	ctx.JSON(http.StatusOK, result)
}
