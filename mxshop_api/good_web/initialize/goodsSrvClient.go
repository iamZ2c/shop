package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop_api/good_web/global"
	"mxshop_api/good_web/proto"
)

func InitGrpcGoodsSrvClient() {

	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%v/goods-srv?tag=iam2cc", global.NacosConf.ConsulConfig.Host, global.NacosConf.ConsulConfig.Port),
		grpc.WithInsecure(),
		// 进程内轮询负载均衡
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw(
			"[good-client] 连接失败",
			"info", err.Error(),
		)
	}

	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}
