package mq

import (
	"github.com/streadway/amqp"
	"log"
)

// 生产者 - 发送任务
func PublishTask(imageData []byte) error {
	conn, err := amqp.Dial("amqp://hzh:qweasd123@8.138.167.80:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	ch, _ := conn.Channel()
	defer conn.Close()
	defer ch.Close()

	return ch.Publish(
		"", "image_queue", false, false,
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        imageData,
		})
}
