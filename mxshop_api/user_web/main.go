package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop_api/user_web/global"
	"mxshop_api/user_web/initialize"
	"mxshop_api/user_web/utils"
	"mxshop_api/user_web/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 注册验证器
	//if v, isSafe := binding.Validator.Engine().(*validator.Validate); isSafe == true {
	//	_ = v.RegisterValidation("mobile", webValidator.ValidateMobile)
	//}

	// 初始化配置
	initialize.InitConfig()
	// 初始化全局日志
	initialize.InitLogger()
	// 初始化用户中台grpc client 负载均衡
	initialize.InitGrpcUserSrvClient()
	// 注册user-srv
	fmt.Println(global.NacosConf)
	c := register.NewClient(global.NacosConf.UserWeb.Host, global.NacosConf.UserWeb.Port)
	global.NacosConf.UserSrv.ID = fmt.Sprintf("%v", uuid.NewV4())
	err := c.Register(global.NacosConf.UserWeb.Name, global.NacosConf.UserWeb.Tag, global.NacosConf.UserWeb.ID)
	if err != nil {
		zap.S().Info(err)
	}
	//初始化路由
	Router := initialize.Routers()
	go func() {
		if initialize.GetEnvParam("IS_DEBUG") {
			zap.S().Info(fmt.Sprintf("服务器启动，端口为 %d", global.NacosConf.UserWeb.Port))
			if err := Router.Run(fmt.Sprintf(":%d", global.NacosConf.UserWeb.Port)); err != nil {
				zap.S().Info(fmt.Sprintf("服务器失败，msg： %v", err.Error()))
			}
		} else {
			FreePort, _ := utils.GetFreePort()
			zap.S().Info(fmt.Sprintf("服务器启动，端口为 %d", FreePort))
			if err := Router.Run(fmt.Sprintf(":%d", FreePort)); err != nil {
				zap.S().Info(fmt.Sprintf("服务器失败，msg： %v", err.Error()))
			}
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = c.DeRegister(global.NacosConf.UserWeb.ID)
	if err != nil {
		panic(err)
	}
}

//func GetEnvParams(env string) any {
//	viper.AutomaticEnv()
//	value := viper.Get(env)
//	if value, isSafe := value.(bool); isSafe {
//		return value
//	}
//	return value
//}
//
//func main() {
//	v := viper.New()
//	v.SetConfigFile("./user_web/debug-config.yaml")
//
//	if err := v.ReadInConfig(); err != nil {
//		fmt.Println(err)
//	}
//
//	Sc := config.ServerConfig{}
//	err := v.Unmarshal(&Sc)
//	if err != nil {
//		fmt.Println(err)
//		panic(err)
//	}
//	fmt.Println(v.Get("Jwt"))
//	fmt.Println(Sc)
//	fmt.Println(GetEnvParams("IS_DEBUG"))
//	v.WatchConfig()
//
//	v.OnConfigChange(func(in fsnotify.Event) {
//		fmt.Println("zcasdasd")
//		_ = v.ReadInConfig()
//
//		sc := config.ServerConfig{}
//		_ = v.Unmarshal(&sc)
//
//		fmt.Println(sc)
//	})
//
//	time.Sleep(time.Second * 300)
//}
