package handler

import (
	"context"
	dao "douyin/user-rpc/pkg/dao"
	"douyin/user-rpc/pkg/parse"
	"douyin/user-rpc/user_rpc"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserList(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	userservice := new(UserServiceImpl)
	resp, err := userservice.UserList(
		context.Background(),
		&user_rpc.UserListReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestUserAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	userservice := new(UserServiceImpl)
	resp, err := userservice.UserAction(
		context.Background(),
		&user_rpc.UserActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestUserInfoAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	userservice := new(UserServiceImpl)
	resp, err := userservice.UserInfoAction(
		context.Background(),
		&user_rpc.UserInfoActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}
