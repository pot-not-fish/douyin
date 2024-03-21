/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-03-01 17:13:01
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-03-13 10:13:00
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\relation-mq\main.go
 */
package main

import (
	"context"
	"douyin/relation-mq/pkg/kitex_client"
	"douyin/relation-mq/pkg/mq"
	"douyin/relation-mq/pkg/parse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parse.Init("./local.yaml")
	// parse.Init("./config.yaml")

	mq.Init()

	kitex_client.Init()

	// 创建队列
	q, err := mq.MQueue.Channel.QueueDeclare(
		"follow",
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
			if len(msg_list) != 3 {
				fmt.Println("err msg list")
				continue
			}

			kitex_code, _ := strconv.ParseInt(msg_list[0], 10, 16)
			user_id, _ := strconv.ParseInt(msg_list[1], 10, 16)
			to_user_id, _ := strconv.ParseInt(msg_list[2], 10, 16)

			// 删除用户关注
			if err = kitex_client.RelationActionRpc(context.Background(), int16(kitex_code), user_id, to_user_id); err != nil {
				fmt.Println(err)
				continue
			}

			// 用户信息自减
			if err = kitex_client.UserInfoActionRpc(context.Background(), int16(kitex_code), user_id, &to_user_id); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(d.Body)
		}
	}()
	fmt.Println("退出请按 CTRL+C\n")
	<-forever
}
