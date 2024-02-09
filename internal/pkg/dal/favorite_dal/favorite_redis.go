/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-26 17:40:54
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-07 16:34:36
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\favorite_dal\favorite_redis.go
 */
package favorite_dal

import (
	"douyin/internal/pkg/dal"
	"fmt"

	"github.com/go-redis/redis"
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

	ok, err := dal.RedisDB.SIsMember("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Result()
	if err != nil {
		return err
	}
	if ok {
		return ErrRepeatFavorite
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

	ok, err := dal.RedisDB.SIsMember("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Result()
	if err != nil {
		return err
	}
	if !ok {
		return ErrRepeatUnFavorite
	}

	if err = dal.RedisDB.SRem("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Err(); err != nil {
		return err
	}

	return nil
}

func (f *Favorite) IsFavoriteCache() (bool, error) {
	var err error
	if dal.RedisDB == nil {
		return false, ErrNullRedisDb
	}

	if f.UserID <= 0 {
		return false, ErrInvalidUserID
	}

	if f.VideoID <= 0 {
		return false, ErrInvalidVideoID
	}

	ok, err := dal.RedisDB.SIsMember("favorite", fmt.Sprintf("%d-%d", f.UserID, f.VideoID)).Result()
	if err != nil {
		return false, err
	}

	return ok, nil
}

func IsFavorite(user_id int64, video_id_list []int64) ([]bool, error) {
	var err error
	if dal.RedisDB == nil {
		return nil, ErrNullRedisDb
	}

	if len(video_id_list) == 0 {
		return nil, ErrEmptyUserID
	}

	is_favorite_list := make([]bool, 0, len(video_id_list))
	_, err = dal.RedisDB.Pipelined(func(p redis.Pipeliner) error {
		var is_favorite bool
		for _, v := range video_id_list {
			ok, err := dal.RedisDB.SIsMember("favorite", fmt.Sprintf("%d-%d", user_id, v)).Result()
			if err != nil {
				return err
			}
			is_favorite = ok
			is_favorite_list = append(is_favorite_list, is_favorite)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return is_favorite_list, nil
}

func RetrieveFavorite(user_id int64) ([]int64, error) {
	var err error
	if FavoriteDb == nil {
		return nil, ErrNullDB
	}

	if user_id < 0 {
		return nil, ErrInvalidUserID
	}

	var favorite_list []Favorite
	if err = FavoriteDb.Order("created_at desc").Where("user_id = ?", user_id).Find(&favorite_list).Error; err != nil {
		return nil, err
	}

	var favorite_video_list []int64
	for _, v := range favorite_list {
		favorite_video_list = append(favorite_video_list, v.VideoID)
	}

	return favorite_video_list, nil
}
