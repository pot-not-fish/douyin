/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-28 00:01:10
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-28 19:53:51
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\comment-rpc\pkg\dao\comment_init.go
 */
package dao

import (
	"douyin/comment-rpc/pkg/parse"
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
	commentDb *gorm.DB
)

func Init() {
	var err error

	username := parse.ConfigStructure.Mysql.Username // 使用者名字 如root
	password := parse.ConfigStructure.Mysql.Password
	host := parse.ConfigStructure.Mysql.Host
	port := parse.ConfigStructure.Mysql.Port
	dbname := "comment" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	commentDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	commentDb.AutoMigrate(&Comment{})
}
