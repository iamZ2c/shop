package global

import (
	"gorm.io/gorm"
	"new_shop_srv/user_srv/config"
)

var DB *gorm.DB
var NacosConf config.ServerConfig
