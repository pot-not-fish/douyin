package favorite_dal

import (
	"fmt"
)

var (
	// 用户id无效
	ErrInvalidUserID = fmt.Errorf("invalid user id")

	// 视频id无效
	ErrInvalidVideoID = fmt.Errorf("invalid video id")

	// 视频所有者的用户id无效
	ErrInvalidVideoUserID = fmt.Errorf("invalid video user id")

	// 重复点赞
	ErrRepeatFavorite = fmt.Errorf("repeat favorite")

	// 重复取消点赞
	ErrRepeatUnFavorite = fmt.Errorf("repeat unfavorite")

	// 数据库指针未初始化
	ErrNullDB = fmt.Errorf("nullptr database")
)

func (f *Favorite) CreateFavorite() error {
	var err error
	if FavoriteDb == nil {
		return ErrNullDB
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if f.VideoUserID <= 0 {
		return ErrInvalidVideoUserID
	}

	if err = FavoriteDb.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).First(f).Error; err == nil {
		return ErrRepeatFavorite
	}

	if err = FavoriteDb.Create(f).Error; err != nil {
		return err
	}

	return nil
}

func (f *Favorite) DeleteFavorite() error {
	var err error
	if FavoriteDb == nil {
		return ErrNullDB
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = FavoriteDb.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).First(f).Error; err != nil {
		return ErrRepeatUnFavorite
	}

	if err = FavoriteDb.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Unscoped().Delete(f).Error; err != nil {
		return err
	}

	return nil
}
