package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"net"
	"new_shop_srv/user_srv/global"
	"new_shop_srv/user_srv/handler"
	"new_shop_srv/user_srv/initialize"
	"new_shop_srv/user_srv/proto"
	"new_shop_srv/user_srv/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitConfig()
	initialize.SqlClientInit()
	IP := flag.String("IP", global.NacosConf.UserSrv.Host, "this is ip address")
	Port := flag.Int64("Port", int64(global.NacosConf.UserSrv.Port), "this is port")
	flag.Parse()

	go func() {
		server := grpc.NewServer()
		proto.RegisterUserServer(server, &handler.UserServer{})
		lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
		fmt.Println(fmt.Sprintf("Server Start With %s:%d", *IP, *Port))
		err := server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// 注册consul
	c := register.NewClient(global.NacosConf.UserSrv.Host, global.NacosConf.UserSrv.Port)
	global.NacosConf.UserSrv.ID = fmt.Sprintf("%v", uuid.NewV4())
	err := c.Register(global.NacosConf.UserSrv.Name, global.NacosConf.UserSrv.Tag, global.NacosConf.UserSrv.ID)
	if err != nil {
		panic(err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = c.DeRegister(global.NacosConf.UserSrv.ID)
	if err != nil {
		panic(err)
	}
}
