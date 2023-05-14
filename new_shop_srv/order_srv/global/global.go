package global

import (
	"gorm.io/gorm"
	"new_shop_srv/order_srv/config"
)

var DB *gorm.DB
var NacosConf config.ServerConfig
