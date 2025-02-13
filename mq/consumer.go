package mq

import (
	"AD/storage"
	"github.com/streadway/amqp"
	"log"
)

// 消费者 - 监听队列
func StartConsumer() {
	// 1. 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://hzh:qweasd123@8.138.167.80:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err) // 连接失败，直接退出
	}
	defer conn.Close() // 确保连接关闭

	// 2. 创建 Channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err) // 创建 Channel 失败，直接退出
	}
	defer ch.Close() // 确保 Channel 关闭

	// 确保队列存在
	q, err := ch.QueueDeclare(
		"image_queue", // 队列名
		true,          // durable: 是否持久化
		false,         // autoDelete: 是否自动删除
		false,         // exclusive: 是否为专属队列
		false,         // noWait: 是否等待服务器确认
		nil,           // arguments: 其他参数
	)
	if err != nil {
		log.Fatalf("队列声明失败: %v", err)
	}
	log.Printf("队列 %s 已声明", q.Name)

	// 3. 消费队列
	msgs, err := ch.Consume(
		"image_queue", // 队列名称
		"",            // 消费者标签
		true,          // 自动应答
		false,         // 独占
		false,         // 不等待
		false,         // 不额外参数
		nil,           // 额外参数
	)
	if err != nil {
		log.Fatalf("Failed to consume the queue: %v", err) // 消费队列失败，直接退出
	}

	log.Println("等待 RabbitMQ 消息...")

	// 4. 处理消息
	for msg := range msgs {
		result, err := processImage(msg.Body)
		if err != nil {
			log.Println("AI server error:", err)
			continue
		}

		// 存储结果
		// storage.SavePrediction(*result)
		log.Println(result.ImageID)
		log.Println(result.Probability)
		log.Println("success")
	}
}

// 调用 AI 服务器
func processImage(imageData []byte) (*storage.Prediction, error) {
	//aiResp, err := http.Post("http://ai-server/predict", "application/octet-stream", bytes.NewReader(imageData))
	//if err != nil {
	//	return nil, err
	//}
	//defer aiResp.Body.Close()
	//
	//var result PredictionResult
	//json.NewDecoder(aiResp.Body).Decode(&result)
	//
	//return &result, nil
	return &storage.Prediction{
		ImageID:     "123",
		Probability: "0",
	}, nil
}
