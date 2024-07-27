package dao

import (
	"douyin/common/cryptox"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type User struct {
	ID int64

	Username string `gorm:"uniqueIndex"`
	Password string

	// 如果fans超过一定数量，则做永久缓存，否则，过期缓存
	FollowCount int
	FansCount   int
	WorkCount   int
}

type UserDao struct{}

// 需要Username Password
func (u UserDao) Create(database *gorm.DB, cache *redis.Client, user *User) error {
	var err error
	user.Password, err = cryptox.EncryptoPassword(user.Password)
	if err != nil {
		return err
	}
	if err = database.Create(user).Error; err != nil {
		return err
	}
	go func() {
		if err = u.FlushCache(cache, user, time.Second*time.Duration(randnum.Intn(1800)+1800)); err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (u UserDao) ValidatePassword(database *gorm.DB, username, password string) error {
	var err error
	user := new(User)
	if err = database.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	epassword, err := cryptox.EncryptoPassword(password)
	if err != nil {
		return err
	}
	if epassword != user.Password {
		return fmt.Errorf("fail to validate password")
	}
	return nil
}

// 需要ID
func (u UserDao) FirstByID(database *gorm.DB, cache *redis.Client, user_id string) (*User, error) {
	var (
		err  error
		user = new(User)
	)
	if user, err = u.FirstByIDCache(cache, user_id); err == nil {
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
		if err = u.FlushCache(cache, user, timeout); err != nil {
			log.Println(err.Error())
		}
	}()
	return user, nil
}

// 需要ID
func (u UserDao) IncWorkCount(database *gorm.DB, cache *redis.Client, user_id string) error {
	var err error
	if err = cache.Get(fmt.Sprintf("work_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(database, cache, user_id)
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
func (u UserDao) IncFavoriteCount(database *gorm.DB, cache *redis.Client, user_id string) error {
	var err error
	if err = cache.Get(fmt.Sprintf("favorite_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(database, cache, user_id)
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

// 需要ID
func (u UserDao) IncTotalFavorited(database *gorm.DB, cache *redis.Client, user_id string) error {
	var err error
	if err = cache.Get(fmt.Sprintf("total_favorited_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(database, cache, user_id)
		if err != nil {
			return err
		}
	}
	if err = cache.Incr(fmt.Sprintf("total_favorited_%v", user_id)).Err(); err != nil {
		return err
	}
	// TODO：为了保证数据库缓存一致性，这里需要引入重试机制或订阅bin log
	go func() {
		if err = database.Model(&User{}).Where("id = ?", user_id).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (u UserDao) IncFollowCount(database *gorm.DB, cache *redis.Client, user_id string) error {
	var err error
	if err = cache.Get(fmt.Sprintf("follow_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(database, cache, user_id)
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

func (u UserDao) IncFansCount(database *gorm.DB, cache *redis.Client, user_id string) error {
	var err error
	if err = cache.Get(fmt.Sprintf("fans_count_%v", user_id)).Err(); err != nil {
		_, err := u.FirstByID(database, cache, user_id)
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

func (u UserDao) FirstByIDCache(cache *redis.Client, user_id string) (*User, error) {
	var (
		err  error
		user = new(User)
	)
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		work_count, err := p.Get(fmt.Sprintf("work_count_%v", user_id)).Result()
		if err != nil {
			return err
		}
		user.WorkCount, err = strconv.Atoi(work_count)
		if err != nil {
			return err
		}
		follow_count, err := p.Get(fmt.Sprintf("follow_count_%v", user_id)).Result()
		if err != nil {
			return err
		}
		user.FollowCount, err = strconv.Atoi(follow_count)
		if err != nil {
			return err
		}
		fans_count, err := p.Get(fmt.Sprintf("fans_count_%v", user_id)).Result()
		if err != nil {
			return err
		}
		user.FansCount, err = strconv.Atoi(fans_count)
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

func (u UserDao) FlushCache(cache *redis.Client, user *User, timeout time.Duration) error {
	var err error
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		if err = p.Set(fmt.Sprintf("work_count_%v", user.ID), user.WorkCount, timeout).Err(); err != nil {
			return err
		}
		if err = p.Set(fmt.Sprintf("follow_count_%v", user.ID), user.FollowCount, timeout).Err(); err != nil {
			return err
		}
		if err = p.Set(fmt.Sprintf("fans_count_%v", user.ID), user.FansCount, timeout).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
