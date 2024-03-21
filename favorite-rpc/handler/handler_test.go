package handler

import (
	"context"
	"douyin/favorite-rpc/favorite_rpc"
	"douyin/favorite-rpc/pkg/dao"
	"douyin/favorite-rpc/pkg/parse"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavoriteAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	favoriteservice := new(FavoriteServiceImpl)
	resp, err := favoriteservice.FavoriteAction(
		context.Background(),
		&favorite_rpc.FavoriteActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestIsFavorite(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	favoriteservice := new(FavoriteServiceImpl)
	resp, err := favoriteservice.IsFavorite(
		context.Background(),
		&favorite_rpc.IsFavoriteReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestIsFavoriteVideo(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	favoriteservice := new(FavoriteServiceImpl)
	resp, err := favoriteservice.FavoriteVideo(
		context.Background(),
		&favorite_rpc.FavoriteVideoReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}
