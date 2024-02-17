/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-16 23:24:53
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 22:47:27
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\cmd\mq-cmd\favorite-mq\main.go
 */
package main

import (
	"context"
	"douyin/internal/pkg/kitex_client"
	"douyin/internal/pkg/mq"
	"douyin/internal/pkg/parse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parse.Init("../../../deployment/config/config.yaml")
	mq.Init()
	kitex_client.Init()

	// 创建队列
	q, err := mq.MQueue.Channel.QueueDeclare(
		"favorite",
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		panic(err)
	}

	messges, err := mq.MQueue.Channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range messges {
			msg_list := strings.Split(string(d.Body), "-")
			if len(msg_list) != 4 {
				fmt.Println("err msg list")
			}

			kitex_code, _ := strconv.ParseInt(msg_list[0], 10, 16)
			user_id, _ := strconv.ParseInt(msg_list[1], 10, 64)
			to_user_id, _ := strconv.ParseInt(msg_list[2], 10, 64)
			video_id, _ := strconv.ParseInt(msg_list[3], 10, 64)

			// 创建相关用户点赞信息字段
			if err := kitex_client.FavoriteActionRpc(context.Background(), int16(kitex_code), user_id, video_id); err != nil {
				fmt.Println(err.Error())
				continue
			}

			// 用户点赞数和被点赞的用户的获赞数自增
			if err := kitex_client.UserInfoActionRpc(context.Background(), int16(kitex_code), user_id, &to_user_id); err != nil {
				fmt.Println(err.Error())
				continue
			}

			// 被点赞的作品的获赞数自增
			if err := kitex_client.VideoInfoActionRpc(context.Background(), int16(kitex_code), video_id); err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Println(d.Body)
		}
	}()
	fmt.Println("退出请按 CTRL+C\n")
	<-forever
}
