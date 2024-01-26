package favorite_dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	UserID        int64
	VideoID       int64
	FavoriteCount int64 `gorm:"default:0"`
}

type User struct {
	gorm.Model

	UserID         int64
	TotalFavorited int64 `gorm:"default:0"`
	FavoriteCount  int64 `gorm:"default:0"`
}

type Favorite struct {
	gorm.Model

	UserID      int64
	VideoID     int64
	VideoUserID int64 // video所对应的user的id
}

var (
	FavoriteDb *gorm.DB
)

func Init() {
	var err error

	username := "root" // 使用者名字 如root
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "favorite" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	FavoriteDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	FavoriteDb.AutoMigrate(&Video{}, &User{}, &Favorite{})
}
