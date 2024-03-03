package dao

import (
	"fmt"
)

var (
	// 用户id无效
	ErrInvalidUserID = fmt.Errorf("invalid user id")

	// 视频id无效
	ErrInvalidVideoID = fmt.Errorf("invalid video id")

	// 重复点赞
	ErrRepeatFavorite = fmt.Errorf("repeat favorite")

	// 重复取消点赞
	ErrRepeatUnFavorite = fmt.Errorf("repeat unfavorite")

	// 数据库指针未初始化
	ErrNullDB = fmt.Errorf("nullptr database")

	// 用户id列表为空
	ErrEmptyUserID = fmt.Errorf("empty user id")
)

func (f *Favorite) CreateFavorite() error {
	var err error
	if favoriteDb == nil {
		return ErrNullDB
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = f.CreateFavoriteCache(); err != nil {
		return err
	}

	go favoriteDb.Create(f)

	return nil
}

func (f *Favorite) DeleteFavorite() error {
	var err error
	if favoriteDb == nil {
		return ErrNullDB
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = f.DeleteFavoriteCache(); err != nil {
		return err
	}

	go favoriteDb.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Unscoped().Delete(f)

	return nil
}
