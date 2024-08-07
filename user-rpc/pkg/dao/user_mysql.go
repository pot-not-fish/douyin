/*
 * @Author: LIKE_A_STAR
 * @Date: 2023-11-10 19:40:46
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-29 23:00:43
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\user-rpc\pkg\dao\user_mysql.go
 */
package dao

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

var (
	// AES key
	key = "3kJfnpr5zp6vw1lfS04u5z3nbHLPRQ5m"

	// 用户名被注册
	ErrNameExist = fmt.Errorf("name has been register")

	// 用户名或密码为空
	ErrNameOrPdEmpty = fmt.Errorf("name or password is empty")

	// 用户名或密码存在非法字符
	ErrFormatNameOrPw = fmt.Errorf("error format name or password")

	// 用户的id为空
	ErrUserIdEmpty = fmt.Errorf("user_id is empty")

	// 输入的密码超过16位
	ErrInvalidPassword = fmt.Errorf("password is too long")

	// user数据库的指针为空
	ErrNullUserDb = fmt.Errorf("user's db pointer is null")
)

/**
 * @method
 * @description 用于创建user表的一个字段
 * @param
 * @return error
 */
func (user *User) CreateUser() error {
	if userDb == nil {
		return ErrNullUserDb
	}

	// 必须要传入用户名和密码才能创建用户
	if len(user.Name) == 0 || len(user.Password) == 0 {
		return ErrNameOrPdEmpty
	}

	re := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	if !re.MatchString(user.Password) {
		return ErrFormatNameOrPw
	}

	// AES加密保存密码
	cipherPassword, err := AESEncrypto(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(cipherPassword)

	// 保证用户名唯一
	if err := userDb.Where("name = ?", user.Name).First(user); err == nil {
		return ErrNameExist
	}

	// md5生成用户唯一标识id
	// raw_id := []byte(user.Name + time.Now().String())
	// hax := md5.Sum(raw_id)
	// user.UserId = fmt.Sprintf("%x", hax)

	if err := userDb.Create(user).Error; err != nil {
		return err
	}

	if err := user.CreateUserCache(); err != nil {
		return err
	}

	return nil
}

/**
 * @method
 * @description 用于查询用户的账号和密码
 * @param
 * @return error
 */
func (user *User) RetrieveAccount() error {
	if userDb == nil {
		return ErrNullUserDb
	}

	if len(user.Name) == 0 || len(user.Password) == 0 {
		return ErrNameOrPdEmpty
	}

	// AES加密后验证用户名是否相同
	cipher_password, err := AESEncrypto(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(cipher_password)

	if err := userDb.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

/**
 * @method
 * @description 查找用户的信息
 * @param
 * @return
 */
func (user *User) RetrieveUser() error {
	var err error

	if userDb == nil {
		return ErrNullUserDb
	}

	if user.ID == 0 {
		return ErrUserIdEmpty
	}

	if err = user.RetrieveUserCache(); err == nil {
		return nil
	}

	if err = userDb.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return err
	}

	go UpdateUserCache(int64(user.ID))

	return nil
}

/**
 * @method
 * @description 将数据库的信息同步给缓存
 * @param
 * @return
 */
func UpdateUserCache(user_id int64) error {
	var err error

	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id == 0 {
		return ErrUserIdEmpty
	}

	user := User{Model: gorm.Model{ID: uint(user_id)}}
	if err = userDb.Where("id = ?", user_id).First(&user).Error; err != nil {
		return err
	}

	go user.CreateUserCache()

	return nil
}

/**
 * @function
 * @description 查找多个用户
 * @param
 * @return
 */
func RetreiveUsers(userid_list []int64) ([]User, error) {
	if userDb == nil {
		return nil, ErrNullUserDb
	}

	var users = make([]User, len(userid_list))
	for k, v := range userid_list {
		users[k].ID = uint(v)
		if err := users[k].RetrieveUserCache(); err == nil {
			continue
		}

		if err := userDb.Where("id = ?", v).First(&users[k]).Error; err != nil {
			return nil, err
		}

		go UpdateUserCache(v)
	}

	return users, nil
}

/**
 * @function
 * @description AES加密
 * @param
 * @return
 */
func AESEncrypto(src string) (string, error) {
	// 如果字符小于16位，则填充字符^，注意该字符不能作为输入明文的字符
	// 业务需要，前面输入保证这里字符小于16位，所以这里没有判断字符大于16位的情况
	if len(src) > 16 {
		return "", ErrInvalidPassword
	}

	if len(src) <= 16 {
		src += string(bytes.Repeat([]byte("^"), 16-len(src)))
	}

	c, err := aes.NewCipher([]byte(key))
	ciphertext := make([]byte, len(src))
	if err != nil {
		return "", err
	}
	c.Encrypt(ciphertext, []byte(src))
	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}

/**
 * @function
 * @description AES解密
 * @param
 * @return
 */
func AESDecrypto(src string) (string, error) {
	decode, err := base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	plaintext := make([]byte, len(decode))
	c.Decrypt(plaintext, []byte(decode))
	return strings.Trim(string(plaintext), "^"), nil
}

/**
 * @method
 * @description 用户的作品数量+1
 * @param
 * @return
 */
func IncWorkCount(user_id int64) error {
	var err error

	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id == 0 {
		return ErrUserIdEmpty
	}

	if err = IncWorkCountCache(user_id); err == nil {
		go inc_work_count(user_id)
		return nil
	}

	if err = inc_work_count(user_id); err != nil {
		return err
	}

	go UpdateUserCache(user_id)

	return nil
}

func inc_work_count(user_id int64) error {
	var err error

	if err = userDb.Model(&User{}).Where("id = ?", user_id).Update("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}

/**
 * @function
 * @description 用户的点赞量+1
 * @param
 * @return
 */
func IncFavorite(user_id, favorite_id int64) error {
	var err error

	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id == 0 || favorite_id == 0 {
		return ErrUserIdEmpty
	}

	if err = IncFavoriteCache(user_id, favorite_id); err == nil {
		go inc_favorite(user_id, favorite_id)
		return nil
	}

	if err = inc_favorite(user_id, favorite_id); err != nil {
		return err
	}

	go UpdateUserCache(user_id)
	go UpdateUserCache(favorite_id)

	return nil
}

func inc_favorite(user_id, favorite_id int64) error {
	return userDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", favorite_id).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
 * @function
 * @description 用户点赞量-1
 * @param
 * @return
 */
func DecFavorite(user_id, favorite_id int64) error {
	var err error
	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id == 0 || favorite_id == 0 {
		return ErrUserIdEmpty
	}

	if err = DecFavoriteCache(user_id, favorite_id); err == nil {
		go dec_favorite(user_id, favorite_id)
		return nil
	}

	if err = dec_favorite(user_id, favorite_id); err != nil {
		return err
	}

	go UpdateUserCache(user_id)
	go UpdateUserCache(favorite_id)

	return nil
}

func dec_favorite(user_id, favorite_id int64) error {
	return userDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", user_id).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", favorite_id).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func IncRelation(user_id, follow_id int64) error {
	var err error
	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id <= 0 || follow_id <= 0 {
		return ErrUserIdEmpty
	}

	if err = IncRelationCache(user_id, follow_id); err == nil {
		go inc_relation(user_id, follow_id)
		return nil
	}

	// 如果缓存更新失败，则尝试更新数据库，再同步缓存
	if err = inc_relation(user_id, follow_id); err != nil {
		return err
	}

	go UpdateUserCache(user_id)
	go UpdateUserCache(follow_id)

	return nil
}

func inc_relation(user_id, follow_id int64) error {
	return userDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", follow_id).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

func DecRelation(user_id, follow_id int64) error {
	var err error
	if userDb == nil {
		return ErrNullUserDb
	}

	if user_id <= 0 || follow_id <= 0 {
		return ErrUserIdEmpty
	}

	if err = DecRelationCache(user_id, follow_id); err == nil {
		go dec_relation(user_id, follow_id)
		return nil
	}

	if err = dec_relation(user_id, follow_id); err != nil {
		return err
	}

	go UpdateUserCache(user_id)
	go UpdateUserCache(follow_id)

	return nil
}

func dec_relation(user_id, follow_id int64) error {
	return userDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", follow_id).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
