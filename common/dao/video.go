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

	// 根据粉丝数，决定缓存的时间
	VideoID string `gorm:"uniqueIndex"`
	UserID  int64
	Title   string

	// 第一次生成根据粉丝数决定缓存时间，后续根据点赞和视频量决定缓存时间
	FavoriteCount int
	CommentCount  int
}

type VideoDao struct{}

var (
	VideoCur int32 = 1
)

// 需要UserID Title
func (v VideoDao) CreateVideo(video *Video) error {
	var (
		err      error
		database = DatabasePool["test"]
	)

	atomic.AddInt32(&VideoCur, 1)
	defer atomic.CompareAndSwapInt32(&VideoCur, 3, 0)
	tablename := fmt.Sprintf("video-%v", atomic.LoadInt32(&VideoCur))
	video.VideoID = fmt.Sprintf("%v-%v", uuid1.String(), tablename)
	if err = database.Table(tablename).Create(video).Error; err != nil {
		return err
	}
	go func() {
		user, err := DefaultDao.UserDao.FirstByID(video.UserID)
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
func (v VideoDao) VideoFeed() ([]*Video, error) {

}

func (v VideoDao) FlushCache(video *Video, timeout time.Duration) error {
	var (
		err   error
		cache = CacheDB
	)
	_, err = cache.Pipelined(func(p redis.Pipeliner) error {
		if err = p.HSet(video.VideoID, "title", video.Title).Err(); err != nil {
			return err
		}
		if err = p.HSet(video.VideoID, "user_id", video.UserID).Err(); err != nil {
			return err
		}
		if err = p.HSet(video.VideoID, "comment_count", video.CommentCount).Err(); err != nil {
			return err
		}
		if err = p.HSet(video.VideoID, "favorite_count", video.FavoriteCount).Err(); err != nil {
			return err
		}
		if err = p.Expire(video.VideoID, timeout).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
