package internal

import (
	"log"

	"github.com/streadway/amqp"
)





func Connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://azizbek:12345@localhost:5672/myvhost")
	if err != nil {
		log.Fatalf("❌ RabbitMQ ulanish xatosi: %v", err)
	}
	log.Println("✅ RabbitMQ ga ulandi")
	return conn
}


func SetupChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Kanal ochishda xato: %v", err)
	}
	return ch
}

func DeclareQueues(ch *amqp.Channel) {
	// Main exchange
	err := ch.ExchangeDeclare("email_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Exchange yaratishda xato: %v", err)
	}

	// DLX exchange
	err = ch.ExchangeDeclare("dlx_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ DLX exchange yaratishda xato: %v", err)
	}

	// Signup queue
	_, err = ch.QueueDeclare("signup_queue", true, false, false, false, amqp.Table{
		"x-dead-letter-exchange": "dlx_exchange",
		"x-dead-letter-routing-key": "signup_dlq",
	})
	if err != nil {
		log.Fatalf("❌ signup_queue yaratishda xato: %v", err)
	}

	// Report queue
	_, err = ch.QueueDeclare("report_queue", true, false, false, false, amqp.Table{
		"x-dead-letter-exchange": "dlx_exchange",
		"x-dead-letter-routing-key": "report_dlq",
	})
	if err != nil {
		log.Fatalf("❌ report_queue yaratishda xato: %v", err)
	}

	// Signup DLQ (5s TTL bilan)
	_, err = ch.QueueDeclare("signup_dlq", true, false, false, false, amqp.Table{
		"x-message-ttl": int32(5000),
		"x-dead-letter-exchange": "email_exchange",
		"x-dead-letter-routing-key": "signup",
	})
	if err != nil {
		log.Fatalf("❌ signup_dlq yaratishda xato: %v", err)
	}

	// Report DLQ (5s TTL bilan)
	_, err = ch.QueueDeclare("report_dlq", true, false, false, false, amqp.Table{
		"x-message-ttl": int32(5000),
		"x-dead-letter-exchange": "email_exchange",
		"x-dead-letter-routing-key": "report",
	})
	if err != nil {
		log.Fatalf("❌ report_dlq yaratishda xato: %v", err)
	}

	// Bindings
	ch.QueueBind("signup_queue", "signup", "email_exchange", false, nil)
	ch.QueueBind("report_queue", "report", "email_exchange", false, nil)
	ch.QueueBind("signup_dlq", "signup_dlq", "dlx_exchange", false, nil)
	ch.QueueBind("report_dlq", "report_dlq", "dlx_exchange", false, nil)

	log.Println("✅ Queue va Exchange lar yaratildi")
}