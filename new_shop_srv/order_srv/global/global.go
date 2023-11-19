package global

import (
	"gorm.io/gorm"
	"new_shop_srv/order_srv/config"
	"new_shop_srv/order_srv/proto"
)

var (
	DB        *gorm.DB
	NacosConf config.ServerConfig

	GoodsSrvClient  proto.GoodsClient
	InventoryClient proto.InventoryClient
)
