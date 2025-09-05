package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func connectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("RabbitMQ ga ulanishda xatolik: %s", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Channel ochishda xatolik: %s", err)
    }

    return conn, ch
}

func declareQueue(ch *amqp.Channel, name string) amqp.Queue {
    q, err := ch.QueueDeclare(
        name,
        true,  // durable
        false, // auto-delete
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        log.Fatalf("Queue yaratishda xatolik: %s", err)
    }
    return q
}

func publishMessage(ch *amqp.Channel, q amqp.Queue, body string) {
   for i := 0; i < 10; i++ {
	 err := ch.Publish(
        "",       // default exchange
        q.Name,   // routing key = queue nomi
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    if err != nil {
        log.Fatalf("Xabar yuborishda xatolik: %s", err)
    }
    fmt.Println("Xabar yuborildi:", body)
   }
}


func consumeMessages(ch *amqp.Channel, q amqp.Queue) {
    msgs, err := ch.Consume(
        q.Name,
        "",    // consumer tag
        false,  // auto-ack
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Queue’dan xabar olishda xatolik: %s", err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            fmt.Println("Xabar olindi:", string(d.Body))
        }
    }()

    fmt.Println("Consumer ishga tushdi. Ctrl+C bilan to‘xtatish mumkin.")
    <-forever
}

func main() {
    conn, ch := connectRabbitMQ()
    defer conn.Close()
    defer ch.Close()

    q := declareQueue(ch, "email_queue")

    publishMessage(ch, q, "Salom, Azizbek! Email yuboriladi.")

    consumeMessages(ch, q)
}