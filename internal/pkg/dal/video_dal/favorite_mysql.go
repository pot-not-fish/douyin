package video_dal

import "gorm.io/gorm"

/**
 * @method
 * @description 创建用户喜欢视频的字段，视频信息中点赞量增加1
 * @param
 * @return
 */
func (f *Favorite) CreateFavorite() error {
	if VideoDb == nil {
		return ErrNullVideoDb
	}

	if f.UserId <= 0 {
		return ErrInvalidUserId
	}

	if f.VideoId == 0 {
		return ErrEmptyVideoId
	}

	if err := create_favorite(f); err != nil {
		return err
	}

	return nil
}

func create_favorite(f *Favorite) error {
	return VideoDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).First(f).Error; err == nil {
			return ErrInvalidFavorite
		}

		if err := tx.Create(f).Error; err != nil {
			return err
		}

		if err := tx.Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (f *Favorite) DeleteFavorite() error {
	if VideoDb == nil {
		return ErrNullVideoDb
	}

	if f.UserId <= 0 {
		return ErrInvalidUserId
	}

	if f.VideoId == 0 {
		return ErrEmptyVideoId
	}

	if err := delete_favorite(f); err != nil {
		return err
	}

	return nil
}

func delete_favorite(f *Favorite) error {
	return VideoDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).First(f).Error; err != nil {
			return ErrInvalidFavorite
		}

		if err := tx.Delete(f).Error; err != nil {
			return err
		}

		if err := tx.Model(&Video{}).Where("id = ?", f.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func RetrieveFavorite(user_id, video_id int64) (bool, error) {
	var err error

	if VideoDb == nil {
		return false, ErrNullVideoDb
	}

	if user_id <= 0 {
		return false, ErrInvalidUserId
	}

	if video_id == 0 {
		return false, ErrEmptyVideoId
	}

	if err = VideoDb.Where("user_id = ? AND video_id = ?", user_id, video_id).First(&Favorite{}).Error; err != nil {
		return false, nil
	}

	return true, nil
}

func RetrieveFavoriteVideo(user_id int64) (*[]Video, error) {
	if VideoDb == nil {
		return nil, ErrNullVideoDb
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserId
	}

	var favorites []Favorite
	if err := VideoDb.Order("created_at desc").Where("user_id = ?", user_id).Find(&favorites).Error; err != nil {
		return nil, err
	}

	var videos = make([]Video, len(favorites))

	for k, v := range favorites {
		if err := VideoDb.Where("id = ?", v.VideoId).First(&videos[k]).Error; err != nil {
			return nil, err
		}
	}

	return &videos, nil
}
