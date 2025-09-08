package main

import (
	"Mq/internal"
	"log"
)



func main(){
	conn := internal.Connect()
	defer conn.Close()

	ch := internal.SetupChannel(conn)
	defer ch.Close()


	internal.DeclareQueues(ch)

	internal.PublishMessage(ch, "email_exchange", "signup", "hooyyyyyyyyyyyyyyyyyyyyyyyyy soskaaaaaaaaa")
	internal.PublishMessage(ch, "email_exchange", "report", "Daily Report: 10 users registered")

	log.Println("Barcha xabarlar yuborildi")


}