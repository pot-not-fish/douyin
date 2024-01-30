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

type RelationInfo struct {
	RelationList []int64
	IsFollowList []bool
}

/**
 * @function
 * @description 查询关注列表
 * @param
 * @return
 */
func RetrieveFollow(user_id, owner_id int64) (*RelationInfo, error) {
	var err error
	if RelationDb == nil {
		return nil, ErrNullDB
	}

	if user_id < 0 || owner_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var relation_info_list RelationInfo
	var relation_list []Relation
	if err = RelationDb.Order("created_at desc").Where("follower_id = ?", owner_id).Find(&relation_list).Error; err != nil {
		return nil, err
	}

	for _, v := range relation_list {
		is_follow, err := IsFollow(user_id, v.FollowID)
		if err != nil {
			return nil, err
		}
		relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, is_follow)
		relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowID)
	}
	return &relation_info_list, nil
}

/**
 * @function
 * @description 查询粉丝列表
 * @param
 * @return
 */
func RetrieveFollower(user_id, owner_id int64) (*RelationInfo, error) {
	var err error
	if RelationDb == nil {
		return nil, ErrNullDB
	}

	var relation_info_list RelationInfo
	var relation_list []Relation
	if err = RelationDb.Order("created_at desc").Where("follow_id = ?", owner_id).Find(&relation_list).Error; err != nil {
		return nil, err
	}

	for _, v := range relation_list {
		is_follow, err := IsFollow(user_id, v.FollowerID)
		if err != nil {
			return nil, err
		}
		relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, is_follow)
		relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowerID)
	}
	return &relation_info_list, nil
}
