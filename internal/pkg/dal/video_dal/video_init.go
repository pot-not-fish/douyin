/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-25 17:18:02
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-27 17:02:27
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\video_dal\video_init.go
 */
package video_dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	VideoDb *gorm.DB
)

type Video struct {
	gorm.Model

	UserID        int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CommentCount  int64 `gorm:"default:0"`
	FavoriteCount int64 `gorm:"default:0"`
}

/**
 * @function
 * @description 用于初始化video数据库的连接
 * @param
 * @return
 */
func Init() {
	var err error

	username := "root" // 使用者名字 如root
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "video" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	VideoDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	VideoDb.AutoMigrate(&Video{})
}
