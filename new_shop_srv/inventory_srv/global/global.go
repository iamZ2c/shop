package global

import (
	"github.com/go-redsync/redsync/v4"
	"gorm.io/gorm"
	"new_shop_srv/inventory_srv/config"
)

var DB *gorm.DB
var NacosConf config.ServerConfig
var Rs *redsync.Redsync
