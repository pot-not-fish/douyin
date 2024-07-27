package dao

import "time"

type Favorite struct {
	ID        int64
	CreatedAt time.Time

	UserID  string
	VideoID string
}
