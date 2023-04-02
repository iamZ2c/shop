package global

import (
	"gorm.io/gorm"
	"new_shop_srv/inventory_srv/config"
)

var DB *gorm.DB
var NacosConf config.ServerConfig
