package mq

import (
	"douyin/relation-mq/pkg/parse"
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Connect *amqp.Connection
	Channel *amqp.Channel
}

var (
	MQueue *Rabbitmq

	once *sync.Once
)

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

	once.Do(func() {
		var (
			username = parse.ConfigStructure.Rabbitmq.Username
			password = parse.ConfigStructure.Rabbitmq.Password
			address  = parse.ConfigStructure.Rabbitmq.Address
			url      = fmt.Sprintf("amqp://%s:%s@%s", username, password, address)
		)
		MQueue = new(Rabbitmq)
		MQueue.Connect, err = amqp.Dial(url)
		if err != nil {
			panic(err)
		}

		MQueue.Channel, err = MQueue.Connect.Channel()
		if err != nil {
			panic(err)
		}
	})
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
