package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"new_shop_srv/order_srv/global"
	"new_shop_srv/order_srv/initialize"
	"new_shop_srv/order_srv/proto"
	"new_shop_srv/order_srv/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitConfig()
	initialize.SqlClientInit()
	initialize.InitSrvClient()
	IP := flag.String("IP", global.NacosConf.OrderSrv.Host, "this is ip address")
	Port := flag.Int64("Port", int64(global.NacosConf.OrderSrv.Port), "this is port")
	flag.Parse()

	go func() {
		server := grpc.NewServer()
		proto.RegisterOrderServer(server, &proto.UnimplementedOrderServer{})
		lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
		fmt.Println(fmt.Sprintf("Server Start With %s:%d", *IP, *Port))
		err := server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	// 注册consul
	c := register.NewClient(global.NacosConf.OrderSrv.Host, global.NacosConf.OrderSrv.Port)
	err := c.Register(global.NacosConf.OrderSrv.Name, global.NacosConf.OrderSrv.Tag, global.NacosConf.OrderSrv.ID)
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = c.DeRegister(global.NacosConf.OrderSrv.ID)
	if err != nil {
		panic(err)
	}
}
