package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// TCP serverni portda tinglash
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("TCP serverni ishga tushirishda xato: %v", err)
	}
	defer listener.Close()

	fmt.Println("TCP server :8080 portda ishlamoqda...")

	for {
		// Yangi ulanishlarni qabul qilish
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Ulanish qabul qilishda xato: %v", err)
			continue 
		}

		// Har bir ulanish uchun al ICC goroutinada ishlov berish
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	// Ma'lumotni o‘qish
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Ma'lumot o‘qishda xato: %v", err)
		return
	}

	// O‘qilgan ma'lumotni qaytarish
	response := fmt.Sprintf("Qabul qilingan xabar: %s", string(buffer[:n]))
	conn.Write([]byte(response))
	fmt.Printf("Klientdan xabar: %s\n", string(buffer[:n]))
}