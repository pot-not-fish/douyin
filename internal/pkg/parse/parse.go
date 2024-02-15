/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-09 16:44:23
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-13 19:11:06
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\parse\parse.go
 */
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
