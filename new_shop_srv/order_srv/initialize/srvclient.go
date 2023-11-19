package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"new_shop_srv/order_srv/global"
	"new_shop_srv/order_srv/proto"
)

func InitSrvClient() {
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%v/inventory_srv?tag=inventory-srv", global.NacosConf.ConsulConfig.Host, global.NacosConf.ConsulConfig.Port),
		grpc.WithInsecure(),
		// 进程内轮询负载均衡
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw(
			"[GetInventoryClient] 连接失败",
			"info", err.Error(),
		)
	}
	//conn, err = grpc.Dial("192.168.10.12:2346", grpc.WithInsecure())
	global.InventoryClient = proto.NewInventoryClient(conn)

	goodsconn, err := grpc.Dial(fmt.Sprintf("consul://%s:%v/goods_srv?tag=goods-srv", global.NacosConf.ConsulConfig.Host, global.NacosConf.ConsulConfig.Port),
		grpc.WithInsecure(),
		// 进程内轮询负载均衡
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw(
			"[GetGoodsClient] 连接失败",
			"info", err.Error(),
		)
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsconn)
}
