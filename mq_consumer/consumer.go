package mq_consumer

import (
	"AD/service"
	"github.com/streadway/amqp"
	"log"
)

// StartConsumer 启动消费者，监听 RabbitMQ 队列
func StartConsumer(uploadService *service.PredicService) {
	conn, err := amqp.Dial("amqp://hzh:qweasd123@8.138.167.80:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"image_queue", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume the queue: %v", err)
	}

	log.Println("等待 RabbitMQ 消息...")

	for msg := range msgs {
		filePath := string(msg.Body)
		retries := 0

		for retries < 3 {
			err := uploadService.ProcessPrediction(filePath)
			if err == nil {
				break
			}
			retries++
			log.Printf("处理失败，重试次数: %d, 错误: %v", retries, err)
		}

		if retries == 3 {
			log.Printf("任务失败超过三次，文件路径: %s", filePath)
		}
	}
}
