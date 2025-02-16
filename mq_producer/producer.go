package mq_producer

import (
	"github.com/streadway/amqp"
	"log"
)

func PublishTask(filePath string) error {
	conn, err := amqp.Dial("amqp://hzh:qweasd123@8.138.167.80:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return err
	}
	defer ch.Close()

	return ch.Publish(
		"", "image_queue", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(filePath),
		})
}
