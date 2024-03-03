package relation_dal

import (
	"fmt"
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
	if relationDb == nil {
		return ErrNullDB
	}

	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = r.CreateRelationeCache(); err != nil {
		return err
	}

	go relationDb.Create(r)

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
	if relationDb == nil {
		return ErrNullDB
	}

	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = r.DeleteRelationCache(); err != nil {
		return err
	}

	go relationDb.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).Unscoped().Delete(r)

	return nil
}

/**
 * @function
 * @description 查询关注列表
 * @param
 * @return
 */
func RetrieveFollow(user_id int64) ([]int64, error) {
	var err error
	if relationDb == nil {
		return nil, ErrNullDB
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var relation_list []Relation
	if err = relationDb.Order("created_at desc").Where("follower_id = ?", user_id).Find(&relation_list).Error; err != nil {
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
	if relationDb == nil {
		return nil, ErrNullDB
	}

	if user_id <= 0 {
		return nil, ErrInvalidUserID
	}

	var relation_list []Relation
	if err = relationDb.Order("created_at desc").Where("follow_id = ?", user_id).Find(&relation_list).Error; err != nil {
		return nil, err
	}

	var follower_id []int64
	for _, v := range relation_list {
		follower_id = append(follower_id, v.FollowerID)
	}
	return follower_id, nil
}
