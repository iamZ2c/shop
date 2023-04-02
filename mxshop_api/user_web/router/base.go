package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_web/api"
)

func InitBaseRouter(group *gin.RouterGroup) {
	baseRouter := group.Group("base")
	{
		baseRouter.GET("captcha", api.GetCaptcha)
		baseRouter.POST("genSmsCode", api.SendSms)
	}
}
