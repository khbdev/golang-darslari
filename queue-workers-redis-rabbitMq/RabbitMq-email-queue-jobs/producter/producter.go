package producter

import (
	"log"

	"github.com/streadway/amqp"
)

func PublishEmails(ch *amqp.Channel) {
	emails := []string{
		"user1@example.com",
		"user2@example.com", // buni consumer DLQ sifatida tashlaydi
		"user3@example.com",
		"user4@example.com",
	}


for i := 0; i < 10; i++ {
	for _, email := range emails {
	err := ch.Publish(
		"email_exchange",
		"email", 
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(email),
		},
	)
	if err != nil {
			log.Printf("Xabar yuborishda xato: %s", err)
	} else {
			log.Printf("Yuborildi: %s", email)
	}
}
}
}