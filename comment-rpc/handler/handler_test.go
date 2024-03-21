package handler

import (
	"context"
	"douyin/comment-rpc/comment_rpc"
	"douyin/comment-rpc/pkg/dao"
	"douyin/comment-rpc/pkg/parse"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentAction(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	commentservice := new(CommentServiceImpl)
	resp, err := commentservice.CommentAction(
		context.Background(),
		&comment_rpc.CommentActionReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}

func TestCommentList(t *testing.T) {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	dao.Init()

	commentservice := new(CommentServiceImpl)
	resp, err := commentservice.CommentList(
		context.Background(),
		&comment_rpc.CommentListReq{
			// payload
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	assert.EqualValues(t, 0, resp.Code)
}
