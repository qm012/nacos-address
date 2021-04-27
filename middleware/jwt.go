package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github/qm012/nacos-adress/global"
	"github/qm012/nacos-adress/model"
	"github/qm012/nacos-adress/service"
	"go.uber.org/zap"
	"net/http"
)

const (
	BearerSchema = "Bearer "
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) == 0 && len(BearerSchema) == len(authHeader) {
			err := errors.New("authorization is null")
			ctx.JSON(http.StatusOK, model.NewFailedResult(err))
			ctx.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]

		claims, err := service.NewJwtService().ValidateToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, model.NewFailedCodeResult(http.StatusUnauthorized, err))
			return
		}
		global.Log.Info("verification information",
			zap.Any("Claims[username]", claims.Username),
			zap.Any("Claims[Admin]", claims.Admin),
			zap.Any("Claims[Issuer]", claims.Issuer),
			zap.Any("Claims[IssuedAt]", claims.IssuedAt),
			zap.Any("Claims[ExpiresAt]", claims.ExpiresAt))
	}
}
