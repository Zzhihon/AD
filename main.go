package main

import (
	"AD/api"
	"AD/mq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	// 读取配置
	//config.LoadConfig()

	// 初始化数据库
	//storage.InitDB("dbname")

	// 启动 RabbitMQ 消费者
	go mq.StartConsumer()

	// 启动 HTTP 服务器
	router.HandleFunc("/upload/", api.UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/ws/", api.WebSocketHandler).Methods(http.MethodPost)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
