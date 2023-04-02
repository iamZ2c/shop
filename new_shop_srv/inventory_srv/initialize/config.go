package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"new_shop_srv/inventory_srv/config"
	"new_shop_srv/inventory_srv/global"
)

func GetEnvParam(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	v := viper.New()
	if GetEnvParam("IS_DEBUG") == true {
		v.SetConfigFile("./inventory_srv/debug-config.yaml")
	} else {
		v.SetConfigFile("./prod-config.yaml")
	}
	if err := v.ReadInConfig(); err != nil {
		zap.S().Infow(fmt.Sprintf("读取配置失败，msg:%v", err))
		panic(err)
	}
	Sc := config.ServerConfig{}
	err := v.Unmarshal(&Sc)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	InitNacosConfig(Sc)
}

func InitNacosConfig(sc config.ServerConfig) config.ServerConfig {
	// 分布式全局配置，包括consul负载/注册中心
	clientConfig := constant.ClientConfig{
		NamespaceId:         sc.NacosConfig.Nclientconfig.Namespace,
		TimeoutMs:           uint64(sc.NacosConfig.Nclientconfig.TimeOutMs),
		NotLoadCacheAtStart: true,
		LogDir:              sc.NacosConfig.Nclientconfig.Logdir,
		CacheDir:            sc.NacosConfig.Nclientconfig.Cachedir,
		LogLevel:            sc.NacosConfig.Nclientconfig.Loglevel,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: sc.NacosConfig.Nserverconfig.Host,
			Port:   uint64(sc.NacosConfig.Nserverconfig.Port),
		},
	}

	NacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := NacosClient.GetConfig(vo.ConfigParam{
		DataId: "inventory-srv",
		Group:  "test"})

	if err != nil {
		panic(err)
	}
	s := &config.ServerConfig{}
	err = json.Unmarshal([]byte(content), s)
	global.NacosConf = *s
	return global.NacosConf
}
