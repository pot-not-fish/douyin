package favorite_dal

import (
	"douyin/internal/pkg/parse"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model

	UserID  int64
	VideoID int64
}

var (
	FavoriteDb *gorm.DB
)

func Init() {
	var err error

	username := parse.ConfigStructure.Mysql.Username // 使用者名字 如root
	password := parse.ConfigStructure.Mysql.Password
	host := parse.ConfigStructure.Mysql.Host
	port := parse.ConfigStructure.Mysql.Port
	dbname := "favorite" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	FavoriteDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	FavoriteDb.AutoMigrate(&Favorite{})
}
