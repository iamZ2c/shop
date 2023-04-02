package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func GrpcErr2GinErr(err error, ctx *gin.Context) {
	if err, isSafe := status.FromError(err); isSafe {
		switch err.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": err.Message(),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求参数错误",
			})
		case codes.Internal:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务器内部错误",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": fmt.Sprintf("其他错误，%s", err),
			})
		}
		return
	}

}
