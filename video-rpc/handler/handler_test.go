package handler

import (
	"context"
	"douyin/video-rpc/pkg/dao"
	"douyin/video-rpc/pkg/parse"
	"douyin/video-rpc/video_rpc"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoFeed(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	videoservice := new(VideoServiceImpl)
	resp, err := videoservice.VideoFeed(
		context.Background(),
		&video_rpc.VideoFeedReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestVideoAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	videoservice := new(VideoServiceImpl)
	resp, err := videoservice.VideoAction(
		context.Background(),
		&video_rpc.VideoActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestVideoInfoAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	dao.InitRedis()

	videoservice := new(VideoServiceImpl)
	resp, err := videoservice.VideoInfoAction(
		context.Background(),
		&video_rpc.VideoInfoActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}
