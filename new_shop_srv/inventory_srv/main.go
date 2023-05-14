package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"new_shop_srv/inventory_srv/global"
	"new_shop_srv/inventory_srv/handler"
	"new_shop_srv/inventory_srv/initialize"
	"new_shop_srv/inventory_srv/proto"
	"new_shop_srv/inventory_srv/utils/register"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitConfig()
	initialize.SqlClientInit()
	initialize.RedisMutexInit()
	server := grpc.NewServer()
	conf := global.NacosConf.InventorySrv
	IP := flag.String("IP", conf.Host, "this is ip address")
	Port := flag.Int64("Port", int64(conf.Port), "this is port")
	flag.Parse()
	go func() {
		proto.RegisterInventoryServer(server, &handler.InventoryServer{})
		lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
		fmt.Println(fmt.Sprintf("Server Start With %s:%d", *IP, *Port))
		err := server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	c := register.NewClient(conf.Host, conf.Port)
	c.Register(conf.Name, global.NacosConf.InventorySrv.Tag, conf.ID)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err := c.DeRegister(conf.ID)
	if err != nil {
		panic(err)
	}
}
