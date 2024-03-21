package handler

import (
	"context"
	"douyin/relation-rpc/follow_rpc"
	"douyin/relation-rpc/pkg/dao"
	"douyin/relation-rpc/pkg/parse"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelationAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	relationservice := new(FollowServiceImpl)
	resp, err := relationservice.RelationAction(
		context.Background(),
		&follow_rpc.RelationActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestIsFollow(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	relationservice := new(FollowServiceImpl)
	resp, err := relationservice.IsFollow(
		context.Background(),
		&follow_rpc.IsFollowReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestRelationList(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	relationservice := new(FollowServiceImpl)
	resp, err := relationservice.RelationList(
		context.Background(),
		&follow_rpc.RelationListReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}
