package favorite_dal

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrInvalidUserID = fmt.Errorf("invalid user id")

	ErrInvalidVideoID = fmt.Errorf("invalid video id")

	ErrInvalidVideoUserID = fmt.Errorf("invalid video user id")

	ErrRepeatFavorite = fmt.Errorf("repeat favorite")

	ErrRepeatUnFavorite = fmt.Errorf("repeat unfavorite")
)

func (f *Favorite) CreateFavorite() error {
	var err error
	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if f.VideoUserID <= 0 {
		return ErrInvalidVideoUserID
	}

	if err = create_favorite(f); err != nil {
		return err
	}

	return nil
}

func create_favorite(f *Favorite) error {
	return FavoriteDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).First(f).Error; err == nil {
			return ErrRepeatFavorite
		}

		if err := tx.Create(f).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", f.UserID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", f.VideoUserID).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (f *Favorite) DeleteFavorite() error {
	var err error
	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = delete_favorite(f); err != nil {
		return err
	}

	return nil
}

func delete_favorite(f *Favorite) error {
	return FavoriteDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).First(f).Error; err != nil {
			return ErrRepeatUnFavorite
		}

		if err := tx.Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Unscoped().Delete(f).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", f.UserID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", f.VideoUserID).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
