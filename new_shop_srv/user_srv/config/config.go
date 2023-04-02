package config

type RedisConfig struct {
	Host   string `mapstructrue:"host"`
	Port   int    `mapstructrue:"port"`
	Expire int    `mapstructrue:"expire"`
}

// ServerConfig 注意yaml文件中的键名称必须与字段名称一样
type ServerConfig struct {
	UserSrv      UserSrvConfig `mapstructrue:"user_srv" json:"user_srv"`
	JWT          JWTConfig     `mapstructrue:"JWT" json:"Jwt"`
	RedisConfig  RedisConfig   `mapstructrue:"redisconfig" json:"redisconfig"`
	ConsulConfig ConsulConfig  `mapstructrue:"consulconfig" json:"consulconfig"`
	UserWeb      UserWebConfig `mapstructrue:"userWeb" json:"userWeb"`
	NacosConfig  NacosConfig   `mapstructrue:"nacosconfig" json:"nacosconfig"`
	MysqlConfig  MysqlConfig   `mapstructrue:"mysqlconfig" json:"mysqlconfig"`
}

type UserWebConfig struct {
	Host string   `mapstructrue:"host" json:"host"`
	Port int      `mapstructrue:"port"`
	Name string   `mapstructrue:"name" json:"name"`
	Tag  []string `json:"tag"`
	ID   string   `json:"id"`
}

type UserSrvConfig struct {
	Host string   `mapstructrue:"host" json:"host"`
	Port int      `mapstructrue:"port" json:"port"`
	Name string   `mapstructrue:"name" json:"name"`
	Tag  []string `json:"tag"`
	ID   string   `json:"id"`
}

type JWTConfig struct {
	SigningKey string `mapstructrue:"SigningKey" json:"signingKey"`
}

type ConsulConfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port int    `mapstructrue:"port" json:"port"`
}

// Nserverconfig Nacos server地址配置
type Nserverconfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port int    `mapstructrue:"port" json:"port"`
}

// Nclientconfig Nacos 客户端配置
type Nclientconfig struct {
	Namespace string `mapstructrue:"namespace" json:"namespace"`
	TimeOutMs int    `mapstructrue:"timeoutms" json:"timeoutms"`
	Logdir    string `mapstructrue:"logdir" json:"logdir"`
	Cachedir  string `mapstructrue:"cachedir" json:"cachedir"`
	Loglevel  string `mapstructrue:"loglevel" json:"loglevel"`
}
type NacosConfig struct {
	Nserverconfig Nserverconfig `mapstructrue:"nserverconfig" json:"nserverconfig"`
	Nclientconfig Nclientconfig `mapstructrue:"nclientconfig" json:"nclientconfig"`
}

type MysqlConfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port uint64 `mapstructrue:"port" json:"port"`
}
