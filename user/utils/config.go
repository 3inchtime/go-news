package utils

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

type UserConfig struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

const (
	Host = "192.168.1.103"
	Port = 8500
	Prefix = "/go-news/config"
)

var Config UserConfig

func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		return nil, err
	}
	return mysqlConfig, nil
}

func GetConsulConfig(host, prefix string, port int) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(Host+":"+strconv.FormatInt(Port, 10)),
		consul.WithPrefix(Prefix),
		consul.StripPrefix(true),
	)

	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}

	err = conf.Load(consulSource)
	return conf, err
}

func init() {
	consulConfig, err := GetConsulConfig(Host, "/go-news/config", 8500)
	if err != nil {
		logger.Fatal(err)
	}

	// Mysql配置信息
	mysqlConfig, err := GetMysqlFromConsul(consulConfig, "mysql")
	if err != nil {
		logger.Fatal(err)
	}

	Config.Mysql = *mysqlConfig
}
