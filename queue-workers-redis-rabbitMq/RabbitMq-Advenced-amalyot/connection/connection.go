package connection

import (
	"log"

	"github.com/streadway/amqp"
)


func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("❌ RabbitMQ ga ulana olmadim: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Channel ochilmadi: %v", err)
	}

	return conn, ch
}
