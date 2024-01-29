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
	if RelationDb == nil {
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

	go RelationDb.Create(r)

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

	if err = r.DeleteRelationCache(); err != nil {
		return err
	}

	go RelationDb.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).Unscoped().Delete(r)

	return nil
}
