package common

import "go-micro.dev/v4/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig, err
}