package user_dal

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	ErrEmptyCacheUserID = fmt.Errorf("cache: empty user id")

	ErrEmptyCacheName = fmt.Errorf("cache: emtpy name")

	ErrEmptyCacheAvatar = fmt.Errorf("cache: empty avatar")

	ErrEmptyCacheBackground = fmt.Errorf("cache empty background")

	ErrEmptyCacheSignature = fmt.Errorf("cache empty signature")

	ErrEmptyRelationID = fmt.Errorf("cache: empty relation id")

	ErrEmptyFavoriteID = fmt.Errorf("cache: empty favorite id")
)

/**
 * @method
 * @description 创建用户信息的缓存
 * @param
 * @return
 */
func (u *User) CreateUserCache() error {
	if u.ID <= 0 {
		return ErrEmptyCacheUserID
	}

	if u.Name == "" {
		return ErrEmptyCacheName
	}

	if u.Avatar == "" {
		return ErrEmptyCacheAvatar
	}

	if u.Background == "" {
		return ErrEmptyCacheBackground
	}

	if u.Signature == "" {
		return ErrEmptyCacheSignature
	}

	_, err := redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		id := fmt.Sprintf("%d", u.ID)
		if err := pipe.HSet("user_name", id, u.Name).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_follow_count", id, u.FollowCount).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_follower_count", id, u.FollowerCount).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_avatar", id, u.Avatar).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_background", id, u.Background).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_signature", id, u.Signature).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_total_favorited", id, u.TotalFavorited).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_work_count", id, u.WorkCount).Err(); err != nil {
			return err
		}

		if err := pipe.HSet("user_favorite_count", id, u.FavoriteCount).Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

/**
 * @method
 * @description 在缓存中查找用户的信息
 * @param
 * @return
 */
func (u *User) RetrieveUserCache() error {
	var err error

	if u.ID <= 0 {
		return ErrEmptyCacheUserID
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		id := fmt.Sprintf("%d", u.ID)

		u.Name, err = pipe.HGet("user_name", id).Result()
		if err != nil {
			return err
		}

		follow_count, err := pipe.HGet("user_follow_count", id).Result()
		if err != nil {
			return err
		}
		u.FollowCount, err = strconv.ParseInt(follow_count, 10, 64)
		if err != nil {
			return err
		}

		follower_count, err := pipe.HGet("user_follower_count", id).Result()
		if err != nil {
			return err
		}
		u.FollowerCount, err = strconv.ParseInt(follower_count, 10, 64)
		if err != nil {
			return err
		}

		u.Avatar, err = pipe.HGet("user_avatar", id).Result()
		if err != nil {
			return err
		}

		u.Background, err = pipe.HGet("user_background", id).Result()
		if err != nil {
			return err
		}

		u.Signature, err = pipe.HGet("user_signature", id).Result()
		if err != nil {
			return err
		}

		total_favorited, err := pipe.HGet("user_total_favorited", id).Result()
		if err != nil {
			return err
		}
		u.TotalFavorited, err = strconv.ParseInt(total_favorited, 10, 64)
		if err != nil {
			return err
		}

		work_count, err := pipe.HGet("user_work_count", id).Result()
		if err != nil {
			return err
		}
		u.WorkCount, err = strconv.ParseInt(work_count, 10, 64)
		if err != nil {
			return err
		}

		favorite_count, err := pipe.HGet("user_favorite_count", id).Result()
		if err != nil {
			return err
		}
		u.FavoriteCount, err = strconv.ParseInt(favorite_count, 10, 64)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户添加关注，缓存的用户关注数和粉丝数自增
 * @param
 * @return
 */
func IncRelationCache(follower_id, follow_id int64) error {
	var err error

	if follower_id <= 0 || follow_id <= 0 {
		return ErrEmptyRelationID
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		err = pipe.HIncrBy("user_follow_count", fmt.Sprintf("%d", follower_id), 1).Err()
		if err != nil {
			return err
		}

		err = pipe.HIncrBy("user_follower_count", fmt.Sprintf("%d", follow_id), 1).Err()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户添加关注，缓存的用户关注数和粉丝数自减
 * @param
 * @return
 */
func DecRelationCache(follower_id, follow_id int64) error {
	var err error

	if follower_id <= 0 || follow_id <= 0 {
		return ErrEmptyRelationID
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		err = pipe.HIncrBy("user_follow_count", fmt.Sprintf("%d", follower_id), -1).Err()
		if err != nil {
			return err
		}

		err = pipe.HIncrBy("user_follower_count", fmt.Sprintf("%d", follow_id), -1).Err()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户点赞
 * @param
 * @return
 */
func IncFavoriteCache(user_id, favorite_id int64) error {
	var err error

	if user_id <= 0 || favorite_id <= 0 {
		return ErrEmptyFavoriteID
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		err = pipe.HIncrBy("user_total_favorited", fmt.Sprintf("%d", favorite_id), 1).Err()
		if err != nil {
			return err
		}

		err = pipe.HIncrBy("user_favorite_count", fmt.Sprintf("%d", user_id), 1).Err()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户取消点赞
 * @param
 * @return
 */
func DecFavoriteCache(user_id, favorite_id int64) error {
	var err error

	if user_id <= 0 || favorite_id <= 0 {
		return ErrEmptyFavoriteID
	}

	_, err = redisDB.Pipelined(func(pipe redis.Pipeliner) error {
		err = pipe.HIncrBy("user_total_favorited", fmt.Sprintf("%d", favorite_id), -1).Err()
		if err != nil {
			return err
		}

		err = pipe.HIncrBy("user_favorite_count", fmt.Sprintf("%d", user_id), -1).Err()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户发布作品，作品量自增
 * @param
 * @return
 */
func IncWorkCountCache(user_id int64) error {
	var err error

	if user_id <= 0 {
		return ErrEmptyCacheUserID
	}

	err = redisDB.HIncrBy("user_work_count", fmt.Sprintf("%d", user_id), 1).Err()
	if err != nil {
		return err
	}

	return nil
}
