package relation_dal

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrInvalidFollowID = fmt.Errorf("invalid follow id")

	ErrInvalidFollowerID = fmt.Errorf("invalid follower id")

	ErrRepeatFollow = fmt.Errorf("repeat follow")

	ErrRepeatUnFollow = fmt.Errorf("repeat unfollow")
)

func (r *Relation) CreateRelation() error {
	var err error

	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = create_relation(r); err != nil {
		return err
	}

	return nil
}

func create_relation(r *Relation) error {
	return RelationDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err == nil {
			return ErrRepeatFollow
		}

		if err := tx.Create(r).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", r.FollowerID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", r.FollowID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *Relation) DeleteRelation() error {
	var err error
	if r.FollowID <= 0 {
		return ErrInvalidFollowID
	}

	if r.FollowerID <= 0 {
		return ErrInvalidFollowerID
	}

	if err = delete_relation(r); err != nil {
		return err
	}

	return nil
}

func delete_relation(r *Relation) error {
	return RelationDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err != nil {
			return ErrRepeatUnFollow
		}

		if err := tx.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).Unscoped().Delete(r).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", r.FollowerID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("user_id = ?", r.FollowID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
