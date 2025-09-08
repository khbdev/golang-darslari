package main

import (
	"fmt"
	"log"
	"queueu/connection"
	"queueu/consumer"
	"queueu/producer"

	"github.com/streadway/amqp"
)



func main(){
		conn, ch := connection.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	setupRabbitMQ(ch)

	consumer.ConsumeQueue(ch, "signup_queue", handleSignup)
	consumer.ConsumeQueue(ch, "report_queue", handleReport)
		producer.PublishMessage(ch, "main_exchange", "signup.email", "Welcome to TodoList!")
	producer.PublishMessage(ch, "main_exchange", "report.daily", "Daily report here!")
	log.Println("🚀 Consumers ishlayapti...")
	select {}
}



func handleSignup(body []byte) error {
	log.Printf("👤 Signup handler ishladi: %s", string(body))
	// Misol uchun error qaytaramiz → DLX ga tushadi
	if string(body) == "fail" {
		return fmt.Errorf("signup error")
	}
	return nil
}

// Report consumer uchun
func handleReport(body []byte) error {
	log.Printf("📊 Report handler ishladi: %s", string(body))
	return nil
}

func setupRabbitMQ(ch *amqp.Channel) {
	// Main exchange (topic)
	ch.ExchangeDeclare("main_exchange", "topic", true, false, false, false, nil)

	// DLX exchange
	ch.ExchangeDeclare("dlx_exchange", "direct", true, false, false, false, nil)

	// Signup queue
	ch.QueueDeclare("signup_queue", true, false, false, false,
		amqp.Table{"x-dead-letter-exchange": "dlx_exchange"},
	)

	// Report queue
	ch.QueueDeclare("report_queue", true, false, false, false,
		amqp.Table{"x-dead-letter-exchange": "dlx_exchange"},
	)

	// Retry queue (DLX uchun)
	ch.QueueDeclare("retry_queue", true, false, false, false,
		amqp.Table{
			"x-dead-letter-exchange": "main_exchange", // qaytib main ga ketadi
		   "x-message-ttl": int64(10000),
		},
	)

	// Bindlar
	ch.QueueBind("signup_queue", "signup.*", "main_exchange", false, nil)
	ch.QueueBind("report_queue", "report.*", "main_exchange", false, nil)
	ch.QueueBind("retry_queue", "retry", "dlx_exchange", false, nil)

	log.Println("✅ RabbitMQ setup tugadi")
}