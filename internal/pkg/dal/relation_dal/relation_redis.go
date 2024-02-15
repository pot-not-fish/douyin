/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-01-26 17:40:38
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-13 16:36:42
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\dal\relation_dal\relation_redis.go
 */
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

	if err = dal.RedisDB.SAdd("relation", fmt.Sprintf("%d-%d", r.FollowerID, r.FollowID)).Err(); err != nil {
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

	if err = dal.RedisDB.SRem("relation", fmt.Sprintf("%d-%d", r.FollowerID, r.FollowID)).Err(); err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 查询是否关注
 * @param
 * @return
 */
func IsFollow(user_id int64, follow_id int64) (bool, error) {
	if RelationDb == nil {
		return false, ErrNullDB
	}

	if user_id <= 0 || follow_id <= 0 {
		return false, ErrEmptyUserID
	}

	if user_id == 0 || follow_id == user_id {
		return false, nil
	}

	ok, err := dal.RedisDB.SIsMember("relation", fmt.Sprintf("%d-%d", user_id, follow_id)).Result()
	if err != nil {
		return false, err
	}

	return ok, nil
}
