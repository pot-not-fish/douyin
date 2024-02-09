package relation_dal

import (
	"douyin/internal/pkg/parse"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	username := parse.ConfigStructure.Mysql.Username // 使用者名字 如root
	password := parse.ConfigStructure.Mysql.Password
	host := parse.ConfigStructure.Mysql.Host
	port := parse.ConfigStructure.Mysql.Port
	dbname := "relation" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	RelationDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	RelationDb.AutoMigrate(&Relation{})
}
