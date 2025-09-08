// internal/consumer.go
package internal

import (
	"log"

	"github.com/streadway/amqp"
)

// ConsumeQueue - berilgan queue nomidan xabarlarni iste'mol qiladi
func ConsumeQueue(ch *amqp.Channel, queueName string, handler func([]byte) error) {
	msgs, err := ch.Consume(
		queueName, // qaysi queue
		"",        // consumer tag
		false,     // auto-ack = false (qo‘lda ack qilamiz)
		false,     // exclusive = false
		false,     // no-local = false
		false,     // no-wait = false
		nil,       // args
	)
	if err != nil {
		log.Fatalf("❌ Queue consume qilinmadi: %v", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("📩 [%s] Message keldi: %s", queueName, msg.Body)

			// handler ishlaydi
			if err := handler(msg.Body); err != nil {
				log.Printf("⚠️ [%s] Handler xatolik: %v", queueName, err)
				msg.Nack(false, false) // DLX ga tushadi
			} else {
				msg.Ack(false)
			}
		}
	}()
}
