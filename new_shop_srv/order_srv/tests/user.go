package main

import (
	"context"
	"encoding/json"
	"fmt"
	// 进程内轮询负载均衡
	//_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"new_shop_srv/order_srv/proto"
)

var conn *grpc.ClientConn
var goodClient proto.GoodsClient
var err error

func Init() {
	conn, err = grpc.Dial("192.168.10.12:2345", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	goodClient = proto.NewGoodsClient(conn)

}

func TestGetUserList() {

	resp, err := goodClient.GoodsList(context.Background(), &proto.GoodsFilterReq{
		PriceMin:    0,
		PriceMax:    0,
		IsHot:       false,
		IsNew:       false,
		IsTab:       false,
		TopCategory: 1,
		Pages:       0,
		PagePerNums: 0,
		KeyWords:    "",
		Brand:       0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
func main() {
	//Init()
	//TestGetUserList()
	type GormList []string
	g := GormList{"a", "b"}
	a := []string{"a", "b"}
	c, _ := json.Marshal(a)
	fmt.Println(c)
	err = json.Unmarshal(c, &g)
	fmt.Println(err)
	fmt.Println(g)
}
