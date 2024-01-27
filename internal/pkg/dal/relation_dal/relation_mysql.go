package relation_dal

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	// 关注者的id无效
	ErrInvalidFollowID = fmt.Errorf("invalid follow id")

	// 粉丝的id无效
	ErrInvalidFollowerID = fmt.Errorf("invalid follower id")

	// 重复关注
	ErrRepeatFollow = fmt.Errorf("repeat follow")

	// 重复取消关注
	ErrRepeatUnFollow = fmt.Errorf("repeat unfollow")

	//数据库指针未初始化
	ErrNullDB = fmt.Errorf("nullptr database")

	// 查询是否关注的用户列表不等
	ErrInEqualList = fmt.Errorf("invalid user list")

	// 用户id列表为空
	ErrEmptyUserID = fmt.Errorf("empty user id")

	// 无效用户id
	ErrInvalidUserID = fmt.Errorf("invalid user id")
)

/**
 * @method
 * @description 用户添加关注
 * @param
 * @return
 */
func (r *Relation) CreateRelation() error {
	var err error
	if RelationDb == nil {
		return ErrNullDB
	}

	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = RelationDb.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err == nil {
		return ErrRepeatFollow
	}

	if err = RelationDb.Create(r).Error; err != nil {
		return err
	}

	return nil
}

/**
 * @method
 * @description 用户取消关注
 * @param
 * @return
 */
func (r *Relation) DeleteRelation() error {
	var err error
	if RelationDb == nil {
		return ErrNullDB
	}

	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = RelationDb.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err != nil {
		return ErrRepeatUnFollow
	}

	if err = RelationDb.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).Unscoped().Delete(r).Error; err != nil {
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
	var err error
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
	err = RelationDb.Transaction(func(tx *gorm.DB) error {
		for k, v := range follower_id_list {
			is_follow := false
			if v == 0 {
				is_follow_list = append(is_follow_list, is_follow)
				continue
			}
			if err = tx.Where("follow_id = ? AND follower_id = ?", follow_id_list[k], v).Error; err == nil {
				is_follow = true
			}
			is_follow_list = append(is_follow_list, is_follow)
		}
		return nil
	})

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
		is_follow := false
		if user_id == 0 {
			relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, is_follow)
			relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowID)
			continue
		}
		if err = RelationDb.Where("follower_id = ? AND follow_id = ?", user_id, v.FollowID).First(&Relation{}).Error; err != nil {
			is_follow = true
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
		is_follow := false
		if user_id == 0 {
			relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, is_follow)
			relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowerID)
			continue
		}
		if err = RelationDb.Where("follower_id = ? AND follow_id = ?", user_id, v.FollowerID).First(&Relation{}).Error; err != nil {
			is_follow = true
		}
		relation_info_list.IsFollowList = append(relation_info_list.IsFollowList, is_follow)
		relation_info_list.RelationList = append(relation_info_list.RelationList, v.FollowerID)
	}
	return &relation_info_list, nil
}
