/*
 * @Author: LIKE_A_STAR
 * @Date: 2024-02-16 12:32:02
 * @LastEditors: LIKE_A_STAR
 * @LastEditTime: 2024-02-17 22:32:40
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

/**
 * @method
 * @description 订阅端配置
 * @param
 * @return
 */
func (r *Rabbitmq) SubscribeSet(exchange string) {
	// 创建交换机
	err := r.Channel.ExchangeDeclare(
		exchange, // 交换机的名字
		"fanout", // 交换机类型
		true,     // 是否持久化
		false,    // 是否自动删除
		false,    // 是否内置交换机
		false,    // 是否等待服务器确认
		nil,      // 其他配置
	)
	if err != nil {
		panic(err)
	}

	// 创建队列
	q, err := r.Channel.QueueDeclare(
		"",    // 队列名称
		false, // 是否持久化
		false, // 是否自动删除
		true,  // 排他（只对创建队列的连接可见）
		false, // 是否等待服务器确认
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = r.Channel.QueueBind(
		q.Name,   // 队列名称
		"",       // 绑定key
		exchange, // 交换机名称
		true,     // 是否等待服务器确认
		nil,
	)
	if err != nil {
		panic(err)
	}
}

/**
 * @method
 * @description 订阅消息队列
 * @param
 * @return
 */
func (r *Rabbitmq) Subscribe(queue string) error {
	messges, err := r.Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for d := range messges {
		fmt.Println(d.Body)
	}

	return nil
}
