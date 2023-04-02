package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"reflect"
)

func main() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "24ce0a70-b0e3-4e59-a8e1-dfc6ec0431b8",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.10.12",
			Port:   8848,
		},
	}

	c, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := c.GetConfig(vo.ConfigParam{
		DataId: "user-srv",
		Group:  "test"})

	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	s := &ServerConfig{}
	_ = json.Unmarshal([]byte(content), &ServerConfig{})
	fmt.Println(reflect.TypeOf(s))
}
