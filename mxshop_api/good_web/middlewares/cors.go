package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() func(ctx *gin.Context) {
	/*
		解决跨域问题
	*/
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,x-token")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		//设置缓存时间
		ctx.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
