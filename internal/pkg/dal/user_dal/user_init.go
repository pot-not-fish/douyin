/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-10 19:25:53
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-01-23 10:48:29
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\user_dal\user_init.go
 */
package user_dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserDb *gorm.DB
)

type User struct {
	gorm.Model

	Name     string
	Password string

	FollowCount   int64 `gorm:"default:0"`
	FollowerCount int64 `gorm:"default:0"`

	Avatar     string `gorm:"default:https://i2.hdslb.com/bfs/face/9075d1c862aa031471e601aa10a60da678108556.jpg@240w_240h_1c_1s_!web-avatar-search-videos.webp"`
	Background string `gorm:"default:https://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"`
	Signature  string `gorm:"default:这是一段个性签名"`

	TotalFavorited int64 `gorm:"default:0"`
	WorkCount      int64 `gorm:"default:0"`
	FavoriteCount  int64 `gorm:"default:0"`
}

type Relation struct {
	gorm.Model

	FollowId   int64
	FollowerId int64
}

/**
 * @function
 * @description 连接数据库user，并且创建所有的数据表
 * @param
 * @return
 */
func Init() {
	var err error

	username := "root" // 使用者名字 如root
	password := "840231514Xk"
	host := "159.75.186.106"
	port := 3306
	dbname := "user" // 数据库名字

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	UserDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	UserDb.AutoMigrate(&User{}, &Relation{})
}
