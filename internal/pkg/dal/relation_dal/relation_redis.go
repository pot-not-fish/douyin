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
func IsFollow(follower_id_list, follow_id_list []int64) ([]bool, error) {
	if RelationDb == nil {
		return nil, ErrNullDB
	}

	if len(follow_id_list) == 0 || len(follower_id_list) == 0 {
		return nil, ErrEmptyUserID
	}

	if len(follow_id_list) != len(follower_id_list) {
		return nil, ErrInEqualList
	}

	is_follow_list := make([]bool, 0, len(follow_id_list))
	for k, v := range follower_id_list {
		if v == 0 || v == follow_id_list[k] {
			is_follow_list = append(is_follow_list, false)
			continue
		}
		ok, err := dal.RedisDB.SIsMember("relation", fmt.Sprintf("%d-%d", v, follow_id_list[k])).Result()
		if err != nil {
			return nil, err
		}
		is_follow_list = append(is_follow_list, ok)
	}

	return is_follow_list, nil
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
		if user_id == 0 || user_id == v.FollowID {
			relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, false)
			relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowID)
			continue
		}
		ok, err := dal.RedisDB.SIsMember("relation", fmt.Sprintf("%d-%d", user_id, v.FollowID)).Result()
		if err != nil {
			return nil, err
		}
		relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, ok)
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
		if user_id == 0 || user_id == v.FollowerID {
			relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, false)
			relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowerID)
			continue
		}
		ok, err := dal.RedisDB.SIsMember("relation", fmt.Sprintf("%d-%d", user_id, v.FollowerID)).Result()
		if err != nil {
			return nil, err
		}
		relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, ok)
		relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowerID)
	}
	return &relation_info_list, nil
}
