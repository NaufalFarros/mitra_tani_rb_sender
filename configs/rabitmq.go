package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMq() (*amqp.Connection, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	amqp, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}

	return amqp, nil
}
