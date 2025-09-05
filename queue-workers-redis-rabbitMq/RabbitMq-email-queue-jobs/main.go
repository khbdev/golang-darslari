package main

import (
	
	"queue/connection"

	"queue/producter"
)



func main(){
	rmq := connection.NewRabbitMQ()
	defer rmq.Close()
   setupRabbitMQ(rmq)

   producter.PublishEmails(rmq.Ch)

//    go consumer.ConsumEmails(rmq.Ch, "consumer-1")
//    go consumer.ConsumEmails(rmq.Ch, "consumer-2")
//    	log.Println(" [*] Kutyapmiz. CTRL+C bosib chiqish mumkin.")
// 	forever := make(chan bool)
// 	<-forever
	
}


func setupRabbitMQ(rmq *connection.RabbitMQ) {
	err := rmq.Ch.ExchangeDeclare("email_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	_, err = rmq.Ch.QueueDeclare("email_queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = rmq.Ch.QueueBind("email_queue", "email", "email_exchange", false, nil)
	if err != nil {
		panic(err)
	}
}
