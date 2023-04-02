package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_web/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// 健康检查
	Router.GET("health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":   200,
			"status": "health",
		})
	})

	// 版本号
	GroupV1 := Router.Group("v1")
	{
		router.InitBaseRouter(GroupV1)
		router.UserRouter(GroupV1)
	}

	return Router
}
