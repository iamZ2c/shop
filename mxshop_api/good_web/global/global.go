package global

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/good_web/config"
	"mxshop_api/good_web/proto"
	"net/http"
	"strconv"
)

var Sc config.ServerConfig
var NacosConf config.ServerConfig
var GoodsSrvClient proto.GoodsClient

func Check(ctx *gin.Context, err error, serverName string) {

	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"server": serverName, "msg": err.Error()},
	)
	ctx.Abort()

}

func SuccessResp(ctx *gin.Context, data ...map[string]any) {
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"msg": "Success", "data": data},
	)

}

func FailedResp(ctx *gin.Context, info string) {
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"msg": "Failed", "info": info},
	)
	return
}

func Str2Int(s string) int32 {
	c, err := strconv.Atoi(s)
	if err != nil {
		panic(c)
	}
	return int32(c)
}
