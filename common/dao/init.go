package dao

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DatabasePool map[string]*gorm.DB

	CacheDB *redis.Client

	randnum = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Init() {
	var err error

	// mysql连接池初始化
	DatabasePool = make(map[string]*gorm.DB)
	// 分库test-1 test-2 test-3
	for i := 1; i <= 3; i++ {
		username := MySQLConfig.Username
		password := MySQLConfig.Password
		host := MySQLConfig.Host
		port := MySQLConfig.Port
		dbname := fmt.Sprintf("test-%d", i)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
		DatabasePool[dbname], err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}

	// redis连接初始化
	CacheDB = redis.NewClient(&redis.Options{
		Addr:     RedisConfig.Address,
		Password: RedisConfig.Password,
		DB:       0,
	})
	_, err = CacheDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
