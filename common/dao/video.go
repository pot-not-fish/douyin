package dao

import (
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type Video struct {
	ID        int64
	CreatedAt time.Time

	VideoID     string
	Title       string
	Description string

	FavoriteCount int
	CommentCount  int
}

type VideoDao struct{}

func (v VideoDao) CreateVideo(mysql *gorm.DB, cache *redis.Client) error {

	return nil
}
