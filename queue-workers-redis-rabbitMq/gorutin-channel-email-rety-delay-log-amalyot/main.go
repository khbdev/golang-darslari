package main

import (
	"fmt"
	"math/rand"
	"time"
)


func sendEmail(jobName string, retries int) string {
	if rand.Intn(2) == 0 { 
		fmt.Println("Email ish", jobName, "muvaffaqiyatsiz bo'ldi")
		if retries > 0 {
			fmt.Println("Email ish", jobName, "1 soniyadan keyin qayta urinadi... (qolgan urinishlar:", retries, ")")
			time.Sleep(1 * time.Second)     
			return sendEmail(jobName, retries-1) 
		} else {
			return fmt.Sprint("Email ish", jobName, "barcha urinishlardan keyin muvaffaqiyatsiz bo'ldi")
		}
	}
	fmt.Println("Email ish", jobName, "muvaffaqiyatli yuborildi (simulyatsiya)")
	return fmt.Sprint("Email ish", jobName, "muvaffaqiyatli yakunlandi")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Email Queue simulyatsiyasi boshlandi")

	emails := []string{"email1", "email2",}
	results := make(chan string, len(emails)) 


	for _, email := range emails {
		emailJob := email
		go func() {
			results <- sendEmail(emailJob, 3)
		}()
	}


	for i := 0; i < len(emails); i++ {
		fmt.Println(<-results)
	}
}
