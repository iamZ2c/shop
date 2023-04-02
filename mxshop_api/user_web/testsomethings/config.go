package main

type RedisConfig struct {
	Host   string `mapstructrue:"host"`
	Port   int    `mapstructrue:"port"`
	Expire int    `mapstructrue:"expire"`
}

// ServerConfig 注意yaml文件中的键名称必须与字段名称一样
type ServerConfig struct {
	UserSrv      UserSrvConfig `mapstructrue:"usersrv" json:"usersrv"`
	JWT          JWTConfig     `mapstructrue:"JWT" json:"Jwt"`
	RedisConfig  RedisConfig   `mapstructrue:"redisconfig" json:"redisconfig"`
	ConsulConfig ConsulConfig  `mapstructrue:"consulconfig" json:"consulconfig"`
	UserWeb      UserWebConfig `mapstructrue:"userWeb" json:"userWeb"`
}

type UserWebConfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port int    `mapstructrue:"port"`
}

type UserSrvConfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port int    `mapstructrue:"port" json:"port"`
	Name string `mapstructrue:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructrue:"SigningKey" json:"signingKey"`
}

type ConsulConfig struct {
	Host string `mapstructrue:"host" json:"host"`
	Port int    `mapstructrue:"port" json:"port"`
}
