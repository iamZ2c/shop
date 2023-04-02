package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/router"
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
		router.GoodsRouter(GroupV1)
	}

	return Router
}
