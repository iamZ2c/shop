package main

import (
	"context"
	"google.golang.org/grpc"
	"new_shop_srv/inventory_srv/proto"
)

//
//import (
//	"context"
//	"fmt"
//	// 进程内轮询负载均衡
//	_ "github.com/mbobakov/grpc-consul-resolver"
//	"google.golang.org/grpc"
//	"log"
//	"new_shop_srv/inventory_srv/proto"
//)
//
//var conn *grpc.ClientConn
//var userClient proto.UserClient
//var err error
//
//func Init() {
//	//conn, err = grpc.Dial("0.0.0.0:50053", grpc.WithInsecure())
//	//if err != nil {
//	//	panic(err)
//	//}
//	conn, err := grpc.Dial(
//		"consul://192.168.10.12:8500/usr-srv?tag=iam2cc",
//		grpc.WithInsecure(),
//		// 进程内轮询负载均衡
//		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	//defer conn.Close()
//	userClient = proto.NewUserClient(conn)
//}
//
func TestGetUserList() {
	conn, _ := grpc.Dial("192.168.10.12:2346", grpc.WithInsecure())
	client := proto.NewInventoryClient(conn)
	for i := 1; i <= 31; i++ {
		client.SetInv(context.Background(), &proto.InvInfo{
			GoodsId: int32(i),
			Num:     20,
		})
	}
}
func main() {
	TestGetUserList()
	//conn.Close()
}
