/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-09 16:44:23
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-28 20:08:31
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\comment-rpc\pkg\parse\parse.go
 */
package parse

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MysqlConfig

	Etcd EtcdConfig
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
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
