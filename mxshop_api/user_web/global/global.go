package global

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_web/config"
	"mxshop_api/user_web/proto"
	"net/http"
)

var Sc config.ServerConfig
var NacosConf config.ServerConfig
var UserSrvClient proto.UserClient

func Check(ctx *gin.Context, err error, serverName ...string) {
	if serverName == nil {
		serverName[0] = "[Server]"
	}
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"server": serverName[0], "msg": err.Error()},
	)
	ctx.Abort()
	return
}

func SuccessResp(ctx *gin.Context, data ...map[string]any) {
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"msg": "Success", "data": data},
	)
	return
}

func FailedResp(ctx *gin.Context, info string) {
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"msg": "Failed", "info": info},
	)
	return
}
