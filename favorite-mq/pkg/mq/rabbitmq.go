package rabbitmq

import (
	"douyin/favorite-mq/pkg/parse"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Connect *amqp.Connection
	Channel *amqp.Channel
}

var (
	MQueue *Rabbitmq
	once   *sync.Once
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
