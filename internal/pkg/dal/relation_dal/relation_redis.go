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

/**
 * @function
 * @description 查询关注列表
 * @param
 * @return
 */
func RetrieveFollow(user_id int64) ([]int64, error) {
	var err error
	if RelationDb == nil {
		return nil, ErrNullDB
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var relation_list []Relation
	if err = RelationDb.Order("created_at desc").Where("follower_id = ?", user_id).Find(&relation_list).Error; err != nil {
		return nil, err
	}

	var follow_id_list []int64
	for _, v := range relation_list {
		follow_id_list = append(follow_id_list, v.FollowID)
	}
	return follow_id_list, nil
}

/**
 * @function
 * @description 查询粉丝列表
 * @param
 * @return
 */
func RetrieveFollower(user_id int64) ([]int64, error) {
	var err error
	if RelationDb == nil {
		return nil, ErrNullDB
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var relation_list []Relation
	if err = RelationDb.Order("created_at desc").Where("follow_id = ?", user_id).Find(&relation_list).Error; err != nil {
		return nil, err
	}

	var follower_id []int64
	for _, v := range relation_list {
		follower_id = append(follower_id, v.FollowerID)
	}
	return follower_id, nil
}
