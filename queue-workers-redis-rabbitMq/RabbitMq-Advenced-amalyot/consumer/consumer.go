package consumer

import (
	"log"

	"github.com/streadway/amqp"
)

// ConsumeQueue berilgan queue dan xabarlarni oladi
func ConsumeQueue(ch *amqp.Channel, queueName string, handler func([]byte) error) {
	msgs, err := ch.Consume(
		queueName, // qaysi queue
		"",        // consumer tag
		false,     // auto-ack yo'q, qo'lda ack qilamiz
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Queue consume qilinmadi: %v", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("📩 Message keldi: %s", msg.Body)

			// handler xatolik bilan tugasa → Nack qilib DLX ga yuboramiz
			if err := handler(msg.Body); err != nil {
				log.Printf("⚠️ Handler xatolik: %v", err)
				msg.Nack(false, false) // qayta urinish emas, DLX ga ketadi
			} else {
				msg.Ack(false)
			}
		}
	}()
}
