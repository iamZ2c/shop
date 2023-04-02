package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/api/goods"
)

func GoodsRouter(Router *gin.RouterGroup) {
	group := Router.Group("goods")
	{
		//group.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		group.GET("list", goods.List)
		group.POST("new", goods.New)
		group.GET("/:id", goods.Detail)
		group.PATCH("update/:id",goods.Update)
		group.PATCH("status/:id",goods.UpdateGoodsStatus)
		group.DELETE("/:id",goods.DeleteGoods)
	}

}
