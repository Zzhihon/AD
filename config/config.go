package config

import (
	"os"
)

var (
	RabbitMQURL string
	MySQLDSN    string
)

func LoadConfig() {
	RabbitMQURL = os.Getenv("RABBITMQ_URL")
	MySQLDSN = os.Getenv("MYSQL_DSN")
}
