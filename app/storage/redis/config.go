package redis

import (
	configgo "github.com/nextmillenniummedia/config-go"
)

type RedisConfig struct {
	Addr     []string `config:"required"`
	Mode     string   `config:"enum=standalone|cluster,default=standalone"`
	DB       int      `config:"default=0"`
	Username string   `config:""`
	Password string   `config:""`
}

const (
	CONNECTION_MODE_STANDALONE = "standalone"
	CONNECTION_MODE_CLUSTER    = "cluster"
)

func GetConfig() (config RedisConfig, err error) {
	settings := configgo.Setting{
		Title:  "Redis",
		Prefix: "REDIS",
	}
	err = configgo.InitConfig(&config, settings).Process()
	return
}
