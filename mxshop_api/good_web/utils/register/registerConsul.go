package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"reflect"
)

type ConsulClient interface {
	Register(name string, tags []string, id string) error
	DeRegister(serverId string) error
}

type Client struct {
	Host string
	Port int
}

func NewClient(host string, port int) ConsulClient {
	return Client{
		Host: host,
		Port: port,
	}
}

func (c Client) Deregister(serverId string) error {
	fmt.Println(serverId)
	return nil
}

func (c Client) Register(name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	Registration := new(api.AgentServiceRegistration)
	Registration.Name = name
	Registration.ID = id
	Registration.Tags = tags
	Registration.Address = c.Host
	Registration.Port = c.Port

	// 心跳检测先不急
	//Check := &api.AgentServiceCheck{
	//	HTTP:                           "http://192.168.10.27:8899/health",
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "300s",
	//}
	//Registration.Check = Check

	err = client.Agent().ServiceRegister(Registration)
	if err != nil {
		return err
	}
	fmt.Println("注册成功")
	return nil
}

func (c Client) DeRegister(id string) error {
	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)

	err = client.Agent().ServiceDeregister(id)
	if err != nil {
		return err
	}
	fmt.Println("退出注册成功")
	return nil
}

func AllServices() {
	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	Services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	fmt.Println(Services)
}

// FilterService 返回map对象，包含过滤出来的services信息
func FilterService(ServiceName string) map[string]*api.AgentService {
	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	filter := fmt.Sprintf(`Service == "%v"`, ServiceName)
	Service, err := client.Agent().ServicesWithFilter(filter)
	fmt.Println(reflect.TypeOf(Service))
	return Service
}

//func main() {
//	S := []string{"shop-web"}
//	Register("0.0.0.0", 50053, "shop-web", S, "shop-web")
//	c := FilterService("shop-web")
//
//	for _, value := range c {
//		if value == nil {
//			fmt.Println("dddd")
//		}
//		fmt.Println(reflect.TypeOf(value.Address))
//	}
//}
