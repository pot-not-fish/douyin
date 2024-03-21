/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-10 19:25:53
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-12 18:44:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\user-rpc\pkg\dao\user_init.go
 */
package dao

import (
	"douyin/user-rpc/pkg/parse"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userDb    *gorm.DB
	onceMySQL *sync.Once
)

type User struct {
	gorm.Model

	Name     string
	Password string

	FollowCount   int64
	FollowerCount int64

	Avatar     string `gorm:"default:https://i2.hdslb.com/bfs/face/9075d1c862aa031471e601aa10a60da678108556.jpg@240w_240h_1c_1s_!web-avatar-search-videos.webp"`
	Background string `gorm:"default:https://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"`
	Signature  string `gorm:"default:这是一段个性签名"`

	WorkCount int64

	TotalFavorited int64 `gorm:"default:0"`
	FavoriteCount  int64 `gorm:"default:0"`
}

/**
 * @function
 * @description 连接数据库user，并且创建所有的数据表
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
		dbname   = "user" // 数据库名字
		dsn      = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	)

	onceMySQL.Do(func() {
		userDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		userDb.AutoMigrate(&User{})
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
