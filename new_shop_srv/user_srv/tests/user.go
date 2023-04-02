//package main
//
//import (
//	"context"
//	"fmt"
//	// 进程内轮询负载均衡
//	_ "github.com/mbobakov/grpc-consul-resolver"
//	"google.golang.org/grpc"
//	"log"
//	"new_shop_srv/user_srv/proto"
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
//func TestGetUserList() {
//	list, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
//		Pn:    1,
//		PSize: 2,
//	})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(list)
//	//for _, user := range list.Data {
//	//	checkResp, _ := userClient.CheckUser(context.Background(), &proto.PasswordCheckInfo{
//	//		Password:          "admin123",
//	//		EncryptedPassword: user.Password,
//	//	})
//	//	fmt.Println(user.NickName)
//	//	fmt.Println(checkResp)
//	//}
//
//}
//func main() {
//	Init()
//	TestGetUserList()
//	TestGetUserList()
//	//conn.Close()
//}
