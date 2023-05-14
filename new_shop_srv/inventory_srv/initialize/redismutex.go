package initialize

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"new_shop_srv/inventory_srv/global"
)

func RedisMutexInit() {
	client := goredislib.NewClient(
		&goredislib.Options{
			Addr: fmt.Sprintf("%v:%v", global.NacosConf.RedisConfig.Host, global.NacosConf.RedisConfig.Port),
		},
	)
	pool := goredis.NewPool(client)
	global.Rs = redsync.New(pool)
}
