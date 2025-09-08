package producer

import (
	"log"

	"github.com/streadway/amqp"
)





func PublishMessage(ch *amqp.Channel, exchange, routingKey, body string){
	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

if err != nil {
    log.Printf("❌ Message publish bo'lmadi: %v", err)
} else {
    log.Printf("✅ Message yuborildi: %s", body)  // ✅ faqat log yozadi
}
}