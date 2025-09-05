package consumer

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)


func ConsumEmails(ch *amqp.Channel, name string) {
	ch.Qos(2, 0, false)
	

	msgs, err := ch.Consume(
		"email_queue", 
		name, 
		false, // auto-ack yo‘q
		false,
		false,
		false,
		nil,

	)

	if err != nil {
			log.Fatalf("Consumer yaratishda xato: %s", err)
	}
		for msg := range msgs {
		log.Printf("[%s] Xabar oldi: %s", name, msg.Body)

		// Simulyatsiya: user2 ishlamaydi, DLQ ga ketadi
		if string(msg.Body) == "user2@example.com" {
			log.Printf("[%s] Email yuborishda xato! DLQ ga tushdi: %s", name, msg.Body)
			msg.Ack(false) // xabarni olib tashlaymiz
			continue
		}

		// Muvaffaqiyatli ishlov
		time.Sleep(500 * time.Millisecond)
		msg.Ack(false)
		log.Printf("[%s] Email yuborildi: %s", name, msg.Body)
	}
}