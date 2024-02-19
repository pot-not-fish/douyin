/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-16 12:32:02
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 23:37:20
 * @Description:
 * @FilePath: \vscode programd:\vscode\goWorker\src\douyin\internal\pkg\mq\rabbitmq.go
 */
package mq

import (
	"douyin/internal/pkg/parse"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Connect *amqp.Connection
	Channel *amqp.Channel
}

var MQueue *Rabbitmq

/**
 * @function
 * @description 初始化MQueue
 * @param
 * @return
 */
func Init() {
	var err error

	if parse.ConfigStructure == nil {
		panic(fmt.Errorf("config structure nullptr"))
	}

	MQueue = new(Rabbitmq)
	username := parse.ConfigStructure.Rabbitmq.Username
	password := parse.ConfigStructure.Rabbitmq.Password
	address := parse.ConfigStructure.Rabbitmq.Address
	url := fmt.Sprintf("amqp://%s:%s@%s", username, password, address)
	MQueue.Connect, err = amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	MQueue.Channel, err = MQueue.Connect.Channel()
	if err != nil {
		panic(err)
	}
}

/**
 * @method
 * @description 发送端配置
 * @param
 * @return
 */
func PublishSimple(queue string) {
	if MQueue == nil {
		panic(fmt.Errorf("MQueue is nil"))
	}

	_, err := MQueue.Channel.QueueDeclare(
		queue,
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
}

/**
 * @method
 * @description 发送数据
 * @param
 * @return
 */
func Publish(exchange string, queue string, message string) error {
	err := MQueue.Channel.Publish(exchange, queue, false, false, amqp.Publishing{
		Timestamp:    time.Now(),
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(message),
	})
	if err != nil {
		return err
	}

	return nil
}
