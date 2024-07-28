package dao

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type User struct {
	ID int64

	// 如果fans超过一定数量，则做永久缓存，否则，过期缓存
	FollowCount    int64
	FansCount      int64
	WorkCount      int64
	FavoriteCount  int64
	TotalFavorited int64
}

type UserDao struct{}

func (u UserDao) Create() (*User, error) {
	var (
		err      error
		database = DatabasePool["test"]
		user     = new(User)
	)
	if err = database.Create(user).Error; err != nil {
		return nil, err
	}
	go func() {
		if err = u.FlushCache(user, time.Second*time.Duration(randnum.Intn(1800)+1800)); err != nil {
			log.Println(err.Error())
		}
	}()
	return user, nil
}

// 需要ID
func (u UserDao) FirstByID(user_id int64) (*User, error) {
	var (
		err      error
		user     = new(User)
		database = DatabasePool["test"]
	)
	if user, err = u.FirstByIDCache(user_id); err == nil {
		return user, nil
	}
	// 缓存层没有找到，在数据库里面找，然后同步到缓存
	if err = database.Where("id = ?", user_id).First(user).Error; err != nil {
		return nil, err
	}
	go func() {
		// 如果粉丝数超过一定数量，则直接持久缓存
		timeout := time.Second * time.Duration(randnum.Intn(1800)+1800)
		if user.FansCount < 500 {
			timeout = 0
		}
		if err = u.FlushCache(user, timeout); err != nil {
			log.Println(err.Error())
		}
	}()
	return user, nil
}

// 需要ID
func (u UserDao) IncWorkCount(user_id int64) error {
	var (
		err      error
		database = DatabasePool["test"]
		cache    = CacheDB
	)
	if err = cache.Get(fmt.Sprintf("work_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(user_id)
		if err != nil {
			return err
		}
	}
	if err = cache.Incr(fmt.Sprintf("work_count_%v", user_id)).Err(); err != nil {
		return err
	}
	// TODO：为了保证数据库缓存一致性，这里需要引入重试机制或订阅bin log
	go func() {
		if err = database.Model(&User{}).Where("id = ?", user_id).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

// 需要ID
func (u UserDao) IncFavoriteCount(user_id int64) error {
	var (
		err      error
		database = DatabasePool["test"]
		cache    = CacheDB
	)
	if err = cache.Get(fmt.Sprintf("favorite_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(user_id)
		if err != nil {
			return err
		}
	}
	if err = cache.Incr(fmt.Sprintf("favorite_count_%v", user_id)).Err(); err != nil {
		return err
	}
	// TODO：为了保证数据库缓存一致性，这里需要引入重试机制或订阅bin log
	go func() {
		if err = database.Model(&User{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (u UserDao) IncFollowCount(user_id int64) error {
	var (
		err      error
		database = DatabasePool["test"]
		cache    = CacheDB
	)
	if err = cache.Get(fmt.Sprintf("follow_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(user_id)
		if err != nil {
			return err
		}
	}
	if err = cache.Incr(fmt.Sprintf("follow_count_%v", user_id)).Err(); err != nil {
		return err
	}
	// TODO：为了保证数据库缓存一致性，这里需要引入重试机制或订阅bin log
	go func() {
		if err = database.Model(&User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (u UserDao) IncFansCount(user_id int64) error {
	var (
		err      error
		database = DatabasePool["test"]
		cache    = CacheDB
	)
	if err = cache.Get(fmt.Sprintf("fans_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(user_id)
		if err != nil {
			return err
		}
	}
	if err = cache.Incr(fmt.Sprintf("fans_count_%v", user_id)).Err(); err != nil {
		return err
	}
	// TODO：为了保证数据库缓存一致性，这里需要引入重试机制或订阅bin log
	go func() {
		if err = database.Model(&User{}).Where("id = ?", user_id).Update("fans_count", gorm.Expr("fans_count + ?", 1)).Error; err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (u UserDao) FirstByIDCache(user_id int64) (*User, error) {
	var (
		err   error
		user  = new(User)
		cache = CacheDB
	)
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		key := fmt.Sprintf("user_%v", user.ID)
		work_count, err := p.HGet(key, "work_count").Result()
		if err != nil {
			return err
		}
		user.WorkCount, err = strconv.ParseInt(work_count, 10, 64)
		if err != nil {
			return err
		}
		follow_count, err := p.HGet(key, "follow_count").Result()
		if err != nil {
			return err
		}
		user.FollowCount, err = strconv.ParseInt(follow_count, 10, 64)
		if err != nil {
			return err
		}
		fans_count, err := p.HGet(key, "fans_count").Result()
		if err != nil {
			return err
		}
		user.FansCount, err = strconv.ParseInt(fans_count, 10, 64)
		if err != nil {
			return err
		}
		favorite_count, err := p.HGet(key, "favorite_count").Result()
		if err != nil {
			return err
		}
		user.FavoriteCount, err = strconv.ParseInt(favorite_count, 10, 64)
		if err != nil {
			return err
		}
		total_favorited, err := p.HGet(key, "total_favorited").Result()
		if err != nil {
			return err
		}
		user.TotalFavorited, err = strconv.ParseInt(total_favorited, 10, 64)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserDao) FlushCache(user *User, timeout time.Duration) error {
	var (
		err   error
		cache = CacheDB
	)
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		key := fmt.Sprintf("user_%v", user.ID)
		if err = p.HSet(key, "work_count", user.WorkCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(key, "follow_count", user.FollowCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(key, "fans_count", user.FansCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(key, "favorite_count", user.FavoriteCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(key, "total_favorited", user.TotalFavorited).Err(); err != nil {
			return err
		}
		if err = p.Expire(key, timeout).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
