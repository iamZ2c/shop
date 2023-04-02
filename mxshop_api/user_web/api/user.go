package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop_api/user_web/forms"
	"mxshop_api/user_web/global"
	"mxshop_api/user_web/global/resp"
	"mxshop_api/user_web/middlewares"
	"mxshop_api/user_web/models"
	"mxshop_api/user_web/proto"
	"net/http"
	"strconv"
	"time"
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

func GetUserList(ctx *gin.Context) {
	//// 从注册中心获取 user服务的信息 代码留存demo，负载均衡已经在init 方法里面实现
	//UsrSrvAddr := ""
	//UsrSrvPort := 0
	//FilterService := utils.FilterService("usr-srv")
	//for _, v := range FilterService {
	//	UsrSrvAddr = v.Address
	//	UsrSrvPort = v.Port
	//}
	//
	//zap.S().Debug("获取用户列表页")
	//conn, err := grpc.Dial(fmt.Sprintf("%s:%s", UsrSrvAddr, strconv.Itoa(UsrSrvPort)), grpc.WithInsecure())
	//
	//if err != nil {
	//	zap.S().Errorw(
	//		"[GetUserList] 连接失败",
	//		"info", err.Error(),
	//	)
	//}

	Pn, err := strconv.Atoi(ctx.DefaultQuery("pn", "0"))
	if err != nil {
		zap.S().Errorw(
			"[GetUserList] 分页参数转化失败",
			"info", err.Error(),
		)
	}
	PSize, err := strconv.Atoi(ctx.DefaultQuery("psize", "10"))
	if err != nil {
		zap.S().Errorw(
			"[GetUserList] 分页参数转化失败",
			"info", err.Error(),
		)
	}

	//Client := proto.NewUserClient(conn)
	Client := global.UserSrvClient

	PageInfo := proto.PageInfo{
		Pn:    uint32(Pn),
		PSize: uint32(PSize),
	}
	userList, err := Client.GetUserList(context.Background(), &PageInfo)
	if err != nil {
		GrpcErr2GinErr(err, ctx)
		zap.S().Debug(fmt.Sprintf("请求RPC失败 %s", err.Error()))
		return
	}

	res := make([]interface{}, 0)
	for _, user := range userList.Data {
		userInfo := resp.UserResp{}
		temp := time.Unix(int64(user.BrithDay), 0).Format("2016-02-06")
		userInfo.ID = user.Id
		userInfo.NickName = user.NickName
		userInfo.Birthday = temp
		userInfo.Gender = user.Gender
		userInfo.Mobile = user.Mobile
		res = append(res, userInfo)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"resp": res,
	})
}

func LoginByPassWord(ctx *gin.Context) {
	// 表单验证
	loginByPassWordForm := forms.LoginByPassWordForm{}
	// 会自动判断是form请求还是json
	if err := ctx.ShouldBindJSON(&loginByPassWordForm); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if !store.Verify(loginByPassWordForm.CaptchaId, loginByPassWordForm.VerifyCode, true) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "verify code error",
		})
		return
	}

	ip := "0.0.0.0"
	port := "50053"
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw(
			"[LoginByPassWord] 连接失败",
			"info", err.Error(),
		)
	}
	userClient := proto.NewUserClient(conn)
	userResp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: loginByPassWordForm.Mobile})
	if err != nil {
		zap.S().Errorw(
			"[LoginByPassWord] GetUserByMobile Rpc error, return nil",
			"info", err.Error(),
		)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "account or password error",
		})
		return
	}

	checkRes, err := userClient.CheckUser(context.Background(), &proto.PasswordCheckInfo{
		Password:          loginByPassWordForm.PassWord,
		EncryptedPassword: userResp.Password,
	})
	if err != nil {
		zap.S().Errorw(
			"[PasswordCheckInfo] PasswordCheckInfo Rpc error",
			"info", err.Error(),
		)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "connect error",
		})
		return
	}
	if checkRes.Success {
		// 生成token
		j := middlewares.NewJWT()

		Cc := models.CustomClaims{
			ID:          uint(userResp.Id),
			NickName:    userResp.NickName,
			AuthorityId: uint(userResp.Role),
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(), // 生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*30,
				Issuer:    "iam2cc",
			},
		}
		token, _ := j.CreateToken(Cc)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":       "login success",
			"nick_name": userResp.NickName,
			"id":        userResp.Id, "token": token,
			"expired_at": time.Now().Unix() + 60*60*24*30*1000})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "password or account err"})
		return
	}

}

func Register(ctx *gin.Context) {
	Rf := forms.RegisterForm{}
	err := ctx.ShouldBind(&Rf)
	if err != nil {
		global.Check(ctx, err, "[User-Web-Register]")
		return
	}

	// 查询redis验证码
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

	// 获取redis验证码
	smsCode := rdb.Get(context.Background(), Rf.Mobile)
	res, _ := smsCode.Result()
	if res != Rf.SmsCode {
		fmt.Println(fmt.Sprintf(",Smscode:%v", res))
		global.FailedResp(ctx, "sms_code error")
		return
	}

	ip := "0.0.0.0"
	port := "50053"
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())
	userClient := proto.NewUserClient(conn)
	userInfoResp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: Rf.NickName,
		Password: Rf.PassWord,
		Mobile:   Rf.Mobile,
	})
	if err != nil {
		global.Check(ctx, err, "[User-Rpc-Register]")
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg":  "Success",
			"data": userInfoResp,
		},
	)
	return
}
