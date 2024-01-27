package relation_dal

import (
	"douyin/internal/pkg/dal"
	"fmt"
)

var (
	ErrNullRedisDb = fmt.Errorf("nullptr redis db")
)

func (r *Relation) CreateRelationeCache() error {
	var err error
	if dal.RedisDB == nil {
		return ErrNullRedisDb
	}

	if r.FollowID <= 0 || r.FollowerID <= 0 {
		return ErrInvalidUserID
	}

	if err = dal.RedisDB.SAdd("favorite", fmt.Sprintf("%d-%d", r.FollowerID, r.FollowID)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Relation) DeleteRelationCache() error {
	var err error
	if dal.RedisDB == nil {
		return ErrNullRedisDb
	}

	if r.FollowID <= 0 || r.FollowerID <= 0 {
		return ErrInvalidUserID
	}

	if err = dal.RedisDB.SRem("favorite", fmt.Sprintf("%d-%d", r.FollowerID, r.FollowID)).Err(); err != nil {
		return err
	}

	return nil
}
