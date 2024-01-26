package user_dal

import "gorm.io/gorm"

/**
 * @method
 * @description 用户添加关注，需要填写双方的id
 * @param
 * @return
 */
func (r *Relation) CreateRelation() error {
	var err error

	if UserDb == nil {
		return ErrNullUserDb
	}

	if r.FollowID == 0 {
		return ErrEmptyFollowId
	}

	if r.FollowerID == 0 {
		return ErrEmptyFollowerId
	}

	if r.FollowID == r.FollowerID {
		return ErrInvalidRelation
	}

	if err = IncRelationCache(r.FollowerID, r.FollowID); err == nil {
		go create_relation(r)
		return nil
	}

	if err = create_relation(r); err != nil {
		return err
	}
	go func() {
		var follower = User{Model: gorm.Model{ID: uint(r.FollowerID)}}
		var follow = User{Model: gorm.Model{ID: uint(r.FollowID)}}
		follower.UpdateUserCache()
		follow.UpdateUserCache()
	}()

	return nil
}

/**
 * @function
 * @description 数据库事物操作，创建关注的mysql字段
 * @param
 * @return
 */
func create_relation(r *Relation) error {
	return UserDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err == nil {
			return ErrRepeatRelation
		}

		if err := tx.Create(r).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", r.FollowID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", r.FollowerID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
 * @method
 * @description 用户取消关注，需要双方id
 * @param
 * @return
 */
func (r *Relation) DeleteRelation() error {
	var err error

	if UserDb == nil {
		return ErrNullUserDb
	}

	if r.FollowID == 0 {
		return ErrEmptyFollowId
	}

	if r.FollowerID == 0 {
		return ErrEmptyFollowerId
	}

	if err = DecRelationCache(r.FollowerID, r.FollowID); err == nil {
		go delete_relation(r)
		return nil
	}

	if err = delete_relation(r); err != nil {
		return err
	}
	go func() {
		var follower = User{Model: gorm.Model{ID: uint(r.FollowerID)}}
		var follow = User{Model: gorm.Model{ID: uint(r.FollowID)}}
		follower.UpdateUserCache()
		follow.UpdateUserCache()
	}()

	return nil
}

/**
 * @function
 * @description 数据库事物操作，删除关注字段
 * @param
 * @return
 */
func delete_relation(r *Relation) error {
	return UserDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("follow_id = ? AND follower_id = ?", r.FollowID, r.FollowerID).First(r).Error; err != nil {
			return ErrInvalidRelation
		}

		if err := tx.Delete(r).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", r.FollowID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", r.FollowerID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
 * @function
 * @description 查找用户的关注列表
 * @param
 * @return
 */
func RetrieveFollows(user_id int64) ([]Relation, error) {
	if UserDb == nil {
		return nil, ErrNullUserDb
	}

	var follow_list []Relation

	if user_id == 0 {
		return nil, ErrUserIdEmpty
	}

	if err := UserDb.Order("created_at desc").Where("follower_id = ?", user_id).Find(&follow_list).Error; err != nil {
		return nil, err
	}

	return follow_list, nil
}

/**
 * @function
 * @description 查找用户的粉丝列表
 * @param
 * @return
 */
func RetrieveFollowers(user_id int64) ([]Relation, error) {
	if UserDb == nil {
		return nil, ErrNullUserDb
	}

	var follower_list []Relation

	if user_id == 0 {
		return nil, ErrUserIdEmpty
	}

	if err := UserDb.Order("created_at desc").Where("follow_id = ?", user_id).Find(&follower_list).Error; err != nil {
		return nil, err
	}

	return follower_list, nil
}

/**
 * @function
 * @description 查找是否存在关注的情况
 * @param
 * @return
 */
func IsFollow(follower_id, follow_id int64) bool {
	if UserDb == nil {
		return false
	}

	if follow_id == 0 || follower_id == 0 {
		return false
	}

	if err := UserDb.Where("follow_id = ? AND follower_id = ?", follow_id, follower_id).First(&Relation{}).Error; err != nil {
		return false
	}

	return true
}
