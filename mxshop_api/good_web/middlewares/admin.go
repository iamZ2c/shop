package middlewares

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/models"
	"net/http"
)

func IsAdminAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		if claims, isSafe := claims.(*models.CustomClaims); isSafe {
			Role := claims.AuthorityId
			if Role != 2 {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"msg": "permission deny",
				})
				ctx.Abort()
				return
			}
			ctx.Next()
		}
	}
}
