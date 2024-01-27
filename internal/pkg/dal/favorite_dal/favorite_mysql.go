package favorite_dal

import (
	"fmt"

	"gorm.io/gorm"
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

	// 查询是否点赞的用户列表和视频列表不等
	ErrInEqualList = fmt.Errorf("invalid user and video list")

	// 用户id列表为空
	ErrEmptyUserID = fmt.Errorf("empty user id")
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

func IsFavorite(user_id_list []int64, video_id_list []int64) ([]bool, error) {
	var err error
	if FavoriteDb == nil {
		return nil, ErrNullDB
	}

	if len(user_id_list) == 0 || len(video_id_list) == 0 {
		return nil, ErrEmptyUserID
	}

	if len(user_id_list) != len(video_id_list) {
		return nil, ErrInEqualList
	}

	is_favorite_list := make([]bool, 0, len(video_id_list))
	err = FavoriteDb.Transaction(func(tx *gorm.DB) error {
		for k, v := range video_id_list {
			var is_favorite bool
			if err = tx.Where("user_id = ? AND video_id = ?", user_id_list[k], v).First(&Favorite{}).Error; err != nil {
				is_favorite = false
			} else {
				is_favorite = true
			}
			is_favorite_list = append(is_favorite_list, is_favorite)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return is_favorite_list, nil
}

type FavoriteVideo struct {
	IsFavoriteList []bool
	FavoriteList   []int64
}

func RetrieveFavorite(user_id, owner_id int64) (*FavoriteVideo, error) {
	var err error
	if FavoriteDb == nil {
		return nil, ErrNullDB
	}

	if user_id < 0 {
		return nil, ErrInvalidUserID
	}

	if owner_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var favorite_videos FavoriteVideo
	var favorite_list []Favorite
	if err = FavoriteDb.Order("created_at desc").Where("user_id = ?", owner_id).Find(&favorite_list).Error; err != nil {
		return nil, err
	}

	for _, v := range favorite_list {
		is_favorite := false
		if user_id == 0 {
			favorite_videos.IsFavoriteList = append(favorite_videos.IsFavoriteList, is_favorite)
			favorite_videos.FavoriteList = append(favorite_videos.FavoriteList, v.VideoID)
			continue
		}
		if err = FavoriteDb.Where("user_id = ? AND video_id = ?", user_id, v.VideoID).First(&Favorite{}).Error; err == nil {
			is_favorite = true
		}
		favorite_videos.IsFavoriteList = append(favorite_videos.IsFavoriteList, is_favorite)
		favorite_videos.FavoriteList = append(favorite_videos.FavoriteList, v.VideoID)
	}

	return &favorite_videos, nil
}
