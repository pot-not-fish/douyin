/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 00:01:10
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-12 17:23:12
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\favorite-rpc\pkg\dao\favorite_init.go
 */
package dao

import (
	"douyin/favorite-rpc/pkg/parse"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model

	UserID  int64
	VideoID int64
}

var (
	favoriteDb *gorm.DB
	onceMySQL  *sync.Once
)

func Init() {
	var (
		err      error
		username = parse.ConfigStructure.Mysql.Username // 使用者名字 如root
		password = parse.ConfigStructure.Mysql.Password
		host     = parse.ConfigStructure.Mysql.Host
		port     = parse.ConfigStructure.Mysql.Port
		dbname   = "favorite" // 数据库名字
		dsn      = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	)

	onceMySQL.Do(func() {
		favoriteDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		favoriteDb.AutoMigrate(&Favorite{})
	})
}

var (
	redisDB   *redis.Client
	onceRedis *sync.Once
)

/**
 * @function
 * @description 用于初始化redis的连接
 * @param
 * @return
 */
func InitRedis() {
	var err error

	onceRedis.Do(func() {
		redisDB = redis.NewClient(&redis.Options{
			Addr:     parse.ConfigStructure.Redis.Address,
			Password: parse.ConfigStructure.Redis.Password,
			DB:       0,
		})
	})

	_, err = redisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
