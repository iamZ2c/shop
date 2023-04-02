package api

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"mxshop_api/user_web/forms"
	"mxshop_api/user_web/global"

	"net/http"
	"strings"
	"time"

	//"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	//"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
)

func GenerateSmsCode(width int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		zap.S().Errorw("[SendSms]", "err", err)
		return
	}

	smsCode := GenerateSmsCode(6)

	// 阿里云发送验证码
	//client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "LTAI5tBKqi3bCEzvQ2ayR7e3", "B9M5dPq9aL4PJmh6HAckJ41y81U90d")
	//if err != nil {
	//	panic(err)
	//}
	//request := requests.NewCommonRequest()
	//request.Method = "POST"
	//request.Scheme = "https" // https | http
	//request.Domain = "dysmsapi.aliyuncs.com"
	//request.Version = "2017-05-25"
	//request.ApiName = "SendSms"
	//request.QueryParams["RegionId"] = "cn-beijing"
	//request.QueryParams["PhoneNumbers"] = "19823522875"                //手机号
	//request.QueryParams["SignName"] = "IAM2CC"                         //阿里云验证过的项目名 自己设置
	//request.QueryParams["TemplateCode"] = "SMS_264830373"              //阿里云的短信模板号 自己设置
	//request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	//response, err := client.ProcessCommonRequest(request)
	//fmt.Println(fmt.Sprintf("response is %v", response))
	//fmt.Print(client.DoAction(request, response))
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//将验证码保存起来 - redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.NacosConf.RedisConfig.Host, global.NacosConf.RedisConfig.Port),
	})
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(fmt.Sprintf("pong:%v,", pong))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.NacosConf.RedisConfig.Expire)*time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
