package api

import (
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/model"
	"net/http"
)

// add server list
func AddNacosServerList(ctx *gin.Context) {
	var nacosCo model.NacosCo
	if err := ctx.ShouldBind(&nacosCo); err != nil {
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
		return
	}
	if err := storageMgr.Add(nacosCo.ClusterIps); err != nil {
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
		return
	}
	ctx.JSON(http.StatusOK, model.NewSuccessResult())

}

// delete an item in the server list
func DeleteNacosServerList(ctx *gin.Context) {
	var nacosCo model.NacosCo
	if err := ctx.ShouldBind(&nacosCo); err != nil {
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
		return
	}
	if err := storageMgr.Delete(nacosCo.ClusterIps); err != nil {
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
		return
	}
	ctx.JSON(http.StatusOK, model.NewSuccessResult())
}

// delete all server list
func DeleteAllNacosServerList(ctx *gin.Context) {
	if err := storageMgr.DeleteAll(); err != nil {
		ctx.JSON(http.StatusOK, model.NewFailedResult(err))
		return
	}
	ctx.JSON(http.StatusOK, model.NewSuccessResult())
}
