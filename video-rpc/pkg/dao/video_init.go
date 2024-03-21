/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-25 17:18:02
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-12 18:51:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\video-rpc\pkg\dao\video_init.go
 */
package dao

import (
	"douyin/video-rpc/pkg/parse"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	videoDb   *gorm.DB
	onceMySQL *sync.Once
)

type Video struct {
	gorm.Model

	UserID        int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CommentCount  int64 `gorm:"default:0"`
	FavoriteCount int64 `gorm:"default:0"`
}

/**
 * @function
 * @description 用于初始化video数据库的连接
 * @param
 * @return
 */
func Init() {
	var (
		err      error
		username = parse.ConfigStructure.Mysql.Username // 使用者名字 如root
		password = parse.ConfigStructure.Mysql.Password
		host     = parse.ConfigStructure.Mysql.Host
		port     = parse.ConfigStructure.Mysql.Port
		dbname   = "video" // 数据库名字
		dsn      = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	)

	onceMySQL.Do(func() {
		videoDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		videoDb.AutoMigrate(&Video{})
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
