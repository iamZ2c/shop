package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"new_shop_srv/goods_srv/global"
	"new_shop_srv/goods_srv/handler"
	"new_shop_srv/goods_srv/initialize"
	"new_shop_srv/goods_srv/proto"
	"new_shop_srv/goods_srv/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitConfig()
	initialize.SqlClientInit()
	IP := flag.String("IP", global.NacosConf.GoodsSrv.Host, "this is ip address")
	Port := flag.Int64("Port", int64(global.NacosConf.GoodsSrv.Port), "this is port")
	flag.Parse()

	go func() {
		server := grpc.NewServer()
		proto.RegisterGoodsServer(server, &handler.GoodsServer{})
		lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
		fmt.Println(fmt.Sprintf("Server Start With %s:%d", *IP, *Port))
		err := server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	// 注册consul
	c := register.NewClient(global.NacosConf.GoodsSrv.Host, global.NacosConf.GoodsSrv.Port)
	err := c.Register(global.NacosConf.GoodsSrv.Name, global.NacosConf.GoodsSrv.Tag, global.NacosConf.GoodsSrv.ID)
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = c.DeRegister(global.NacosConf.GoodsSrv.ID)
	if err != nil {
		panic(err)
	}
}
