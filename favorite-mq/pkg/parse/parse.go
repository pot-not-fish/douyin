/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 23:34:56
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 23:39:03
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\favorite-mq\pkg\parse\parse.go
 */
package parse

import "github.com/spf13/viper"

type Config struct {
	Rabbitmq RabbitmqConfig
	Etcd     EtcdConfig
}

type RabbitmqConfig struct {
	Username string
	Password string
	Address  string
}

type EtcdConfig struct {
	Host string
}

var ConfigStructure *Config

func Init(path string) {
	// 根据引用config的文件位置不同，需要传入不同的路径
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	ConfigStructure = new(Config)
	if err := viper.Unmarshal(ConfigStructure); err != nil {
		panic(err)
	}
}
