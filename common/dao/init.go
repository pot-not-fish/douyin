package dao

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	randnum = rand.New(rand.NewSource(time.Now().UnixNano()))

	uuid1 = uuid.New()
)

var (
	DatabasePool map[string]*MySQLConn
)

type MySQLConn struct {
	DB *gorm.DB
}

func MySQLConnInit(dbname string) (*MySQLConn, error) {
	var err error
	mysqlConn := new(MySQLConn)
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	mysqlConn.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return mysqlConn, nil
}

var (
	CacheDB *redis.Client
)

func CacheInit() error {
	var err error
	CacheDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = CacheDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
