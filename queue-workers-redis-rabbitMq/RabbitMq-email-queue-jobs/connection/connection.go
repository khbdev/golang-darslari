package connection

import (
	"log"

	"github.com/streadway/amqp"
)


type RabbitMQ struct {
	Conn *amqp.Connection
	Ch *amqp.Channel
}


func NewRabbitMQ() *RabbitMQ{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMq ulanish Xatolik")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Channel ochishda xatolik: %s", err)
	}
		return &RabbitMQ{
		Conn: conn,
		Ch:   ch,
	}
}

func (r *RabbitMQ) Close() {
	r.Ch.Close()
	r.Conn.Close()
}