package relation_dal

import (
	"douyin/relation-rpc/pkg/parse"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model

	FollowID   int64
	FollowerID int64
}

var (
	relationDb *gorm.DB
)

func Init() {
	var err error

	username := parse.ConfigStructure.Mysql.Username // 使用者名字 如root
	password := parse.ConfigStructure.Mysql.Password
	host := parse.ConfigStructure.Mysql.Host
	port := parse.ConfigStructure.Mysql.Port
	dbname := "relation" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	relationDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	relationDb.AutoMigrate(&Relation{})
}

var (
	redisDB *redis.Client
)

/**
 * @function
 * @description 用于初始化redis的连接
 * @param
 * @return
 */
func InitRedis() {
	var err error
	redisDB = redis.NewClient(&redis.Options{
		Addr:     parse.ConfigStructure.Redis.Address,
		Password: parse.ConfigStructure.Redis.Password,
		DB:       0,
	})

	_, err = redisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
