package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_web/api"
)

func UserRouter(Router *gin.RouterGroup) {
	group := Router.Group("user")

	{
		//group.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		group.GET("list", api.GetUserList)
		group.POST("login_by_password", api.LoginByPassWord)
		group.POST("register", api.Register)
	}

}
