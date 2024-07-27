package dao

import "time"

type Relation struct {
	ID        int64
	CreatedAt time.Time

	FollowID string
	FansID   string
}
