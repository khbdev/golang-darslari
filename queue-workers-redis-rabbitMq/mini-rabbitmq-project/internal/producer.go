package internal

import (
	"log"

	"github.com/streadway/amqp"
)

func PublishMessage(ch *amqp.Channel, exchange, routingKey, body string) {
	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Printf("❌ Xabar yuborilmadi: %v", err)
	} else {
		log.Printf("✅ Xabar yuborildi → [%s]: %s", routingKey, body)
	}
}
