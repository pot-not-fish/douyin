package relation_dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserID        int64
	FollowCount   int64 `gorm:"default:0"`
	FollowerCount int64 `gorm:"default:0"`
}

type Relation struct {
	gorm.Model

	FollowID   int64
	FollowerID int64
}

var (
	RelationDb *gorm.DB
)

func Init() {
	var err error

	username := "root" // 使用者名字 如root
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "relation" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	RelationDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	RelationDb.AutoMigrate(&Relation{}, &User{})
}
