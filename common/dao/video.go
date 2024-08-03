package dao

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
)

type Video struct {
	ID        int64
	CreatedAt time.Time

	UserID int64
	Title  string

	// 第一次生成根据粉丝数决定缓存时间，后续根据点赞和视频量决定缓存时间
	FavoriteCount int
	CommentCount  int

	// 影响因子
	Important int
}

const createVideoTable = `create table if not exists ` + "`%s`" + `
(
    id bigint      auto_increment primary key,
    created_at     timestamp      default CURRENT_TIMESTAMP not null,
    user_id        int                                      not null,
    title          varchar(255)                             not null,
    favorite_count int                                      not null,
    comment_count  int                                      not null,
    port           int                                      not null,
    important      int                                      not null,
)`

type VideoDao struct{}

type Record struct {
	ID         int64
	VideoCur   int
	VideoCount int64
}

var (
	VideoCount int64 = 0
)

func (v VideoDao) SepVideo() (string, error) {
	atomic.AddInt64(&VideoCount, 1)
	// 方便测试，使用一个小的分表指标
	tablename := fmt.Sprintf("video_%d", atomic.LoadInt64(&VideoCount)/10)
	if VideoCount%10 == 0 {
		database := DatabasePool["test"].DB
		if err := database.Exec(fmt.Sprintf(createVideoTable, tablename)); err != nil {
			return "", nil
		}
	}
	return tablename, nil
}

// 需要UserID Title
func (v VideoDao) CreateVideo(video *Video) error {
	var (
		err      error
		database = DatabasePool["test"].DB
	)

	errcallback := func() {
		atomic.AddInt64(&VideoCount, -1)
	}

	tablename, err := v.SepVideo()
	if err != nil {
		errcallback()
		return err
	}

	if err = database.Table(tablename).Create(video).Error; err != nil {
		errcallback()
		return err
	}
	go func() {
		user, err := DefaultUser.FirstByID(video.UserID)
		if err != nil {
			log.Println(err.Error())
		}
		// 第一次创建视频，根据粉丝数，决定缓存的时间，如果粉丝数达到一定程度，缓存时间翻倍
		timeout := time.Second * time.Duration(randnum.Intn(1800)+1800)
		if user.FansCount > 500 {
			timeout *= 2
		}
		if err = v.FlushCache(video, timeout); err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

// 翻页查询的问题
// 翻页的时候，如果此时有视频插入，则会有问题
func (v VideoDao) VideoFeed(limit int, offset int64) ([]*Video, error) {
	var (
		err      error
		database = DatabasePool["test"].DB
	)
	if limit <= 0 {
		limit = 10
	}
	if offset <= 0 {
		offset = atomic.LoadInt64(&VideoCount)
	}
	offset /= 10
	var videos []*Video
	if err = database.Table(fmt.Sprintf("video_%v", offset)).Where("id < ?", offset).Limit(limit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (v VideoDao) FlushCache(video *Video, timeout time.Duration) error {
	var (
		err   error
		cache = CacheDB
	)
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		if err = p.HSet(fmt.Sprintf("video_%v", video.ID), "title", video.Title).Err(); err != nil {
			return err
		}
		if err = p.HSet(fmt.Sprintf("video_%v", video.ID), "user_id", video.UserID).Err(); err != nil {
			return err
		}
		if err = p.HSet(fmt.Sprintf("video_%v", video.ID), "comment_count", video.CommentCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(fmt.Sprintf("video_%v", video.ID), "favorite_count", video.FavoriteCount).Err(); err != nil {
			return err
		}
		if err = p.Expire(fmt.Sprintf("video_%v", video.ID), timeout).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
