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
	// 2. Report queue ni consume qilamiz
	internal.ConsumeQueue(ch, "report_queue", handleReport)

	log.Println("🚀 Report consumer ishlayapti...")
	select {} // processni ochiq tutamiz
}

// Report handler
func handleReport(body []byte) error {
	log.Printf("📊 Report handler ishladi: %s", string(body))

	// Agar xato qilmoqchi bo‘lsak:
	if string(body) == "fail" {
		return fmt.Errorf("report error")
	}

	return nil
}
