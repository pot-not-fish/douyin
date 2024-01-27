package favorite_dal

import (
	"douyin/internal/pkg/dal"
	"fmt"
)

var (
	ErrNullRedisDb = fmt.Errorf("nullptr redis db")
)

func (f *Favorite) CreateFavoriteCache() error {
	var err error
	if dal.RedisDB == nil {
		return ErrNullRedisDb
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = dal.RedisDB.SAdd("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Err(); err != nil {
		return err
	}

	return nil
}

func (f *Favorite) DeleteFavoriteCache() error {
	var err error
	if dal.RedisDB == nil {
		return ErrNullRedisDb
	}

	if f.UserID <= 0 {
		return ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return ErrInvalidVideoID
	}

	if err = dal.RedisDB.SRem("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Err(); err != nil {
		return err
	}

	return nil
}
