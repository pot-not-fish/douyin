package dal

import "github.com/go-redis/redis"

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
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	_, err = RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
