package dal

import (
	"douyin/internal/pkg/parse"

	"github.com/go-redis/redis"
)

var (
	RedisDB *redis.Client
)

/**
 * @function
 * @description 用于初始化redis的连接
 * @param
 * @return
 */
func InitRedis() {
	var err error
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     parse.ConfigStructure.Redis.Address,
		Password: parse.ConfigStructure.Redis.Password,
		DB:       0,
	})

	_, err = RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
