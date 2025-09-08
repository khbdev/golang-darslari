package main

import (
	"Mq/internal"
	"fmt"
	"log"
)

func main() {
	// 1. RabbitMQ ga ulanamiz
	conn := internal.Connect()
	defer conn.Close()

	ch := internal.SetupChannel(conn)
	defer ch.Close()

	// 2. Signup queue ni consume qilamiz
	internal.ConsumeQueue(ch, "signup_queue", handleSignup)

	log.Println("🚀 Signup consumer ishlayapti...")
	select {} // processni ochiq tutamiz
}

// Signup handler
func handleSignup(body []byte) error {
	log.Printf("👤 Signup handler ishladi: %s", string(body))

	// Misol uchun error qilish → DLX ga tushadi → 5s dan keyin retry bo‘ladi
	if string(body) == "fail" {
		return fmt.Errorf("signup error")
	}
	return nil
}
