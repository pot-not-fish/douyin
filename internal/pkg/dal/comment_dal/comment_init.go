package comment_dal

import (
	"douyin/internal/pkg/parse"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Content string
	UserID  int64
	VideoID int64
}

var (
	CommentDb *gorm.DB
)

func Init() {
	var err error

	username := parse.ConfigStructure.Mysql.Username // 使用者名字 如root
	password := parse.ConfigStructure.Mysql.Password
	host := parse.ConfigStructure.Mysql.Host
	port := parse.ConfigStructure.Mysql.Port
	dbname := "comment" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	CommentDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	CommentDb.AutoMigrate(&Comment{})
}
