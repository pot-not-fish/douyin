/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 00:04:46
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-01 00:04:56
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\pkg\parse\parse.go
 */
package parse

import "github.com/spf13/viper"

type Config struct {
	Redis RedisConfig
	Mysql MysqlConfig
	Etcd  EtcdConfig
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
