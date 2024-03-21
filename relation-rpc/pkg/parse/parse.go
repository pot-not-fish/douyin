/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-29 19:50:44
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-29 19:51:54
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-rpc\pkg\parse\parse.go
 */
package parse

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Redis    RedisConfig
	Mysql    MysqlConfig
	Rabbitmq RabbitmqConfig
	Etcd     EtcdConfig
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

type RabbitmqConfig struct {
	Username string
	Password string
	Address  string
}

type EtcdConfig struct {
	Host string
}

var (
	ConfigStructure *Config
	once            *sync.Once
)

func Init(path string) {
	// 根据引用config的文件位置不同，需要传入不同的路径
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	once.Do(func() {
		ConfigStructure = new(Config)
		if err := viper.Unmarshal(ConfigStructure); err != nil {
			panic(err)
		}
	})
}
