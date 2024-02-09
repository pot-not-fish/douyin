package parse

import (
	"github.com/spf13/viper"
)

type Config struct {
	Cos   Cos
	Redis RedisConfig
	Mysql MysqlConfig
}

type RedisConfig struct {
	Address  string
	Password string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Cos struct {
	BucketURL  string
	ServiceURL string
	SecretID   string
	SecretKey  string
}

var ConfigStructure *Config

func Init() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("../deployment/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	ConfigStructure = new(Config)
	if err := viper.Unmarshal(ConfigStructure); err != nil {
		panic(err)
	}
}
