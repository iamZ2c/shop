package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop_api/user_web/config"
	"mxshop_api/user_web/global"
)

func GetEnvParam(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func initNacosConfig(sc config.ServerConfig) config.ServerConfig {
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
		DataId: "user-srv",
		Group:  "test"})

	if err != nil {
		panic(err)
	}
	s := &config.ServerConfig{}
	err = json.Unmarshal([]byte(content), s)
	global.NacosConf = *s
	return global.NacosConf
}

func InitConfig() config.ServerConfig {
	// viper仅仅读取nacos配置
	v := viper.New()
	CnfPrefix := "./user_web/%s-config.yaml"
	if GetEnvParam("IS_DEBUG") == true {
		v.SetConfigFile(fmt.Sprintf(CnfPrefix, "debug"))
	} else {
		// TODO记得改回来pord
		v.SetConfigFile(fmt.Sprintf(CnfPrefix, "prod"))
	}
	if err := v.ReadInConfig(); err != nil {
		zap.S().Infow(fmt.Sprintf("读取配置失败，msg:%v", err))
		panic(err)
	}
	Sc := config.ServerConfig{}
	err := v.Unmarshal(&Sc)
	NacosConfig := initNacosConfig(Sc)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return NacosConfig
}
