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
)

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
