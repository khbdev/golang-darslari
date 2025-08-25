package main

import (
	"fmt"
	"math/rand"
	"time"
)


func doJob(name string, retries int) string {
	if rand.Intn(2) == 0 {
		fmt.Println("Ish", name, "muvaffaqiyatsiz bo'ldi")
		if retries > 0 {
			fmt.Println("Ish", name, "1 soniyadan keyin qayta urinadi... (qolgan urinishlar:", retries, ")")
			time.Sleep(1 * time.Second)    
			return doJob(name, retries-1)   
		} else {
			return fmt.Sprint("Ish", name, "barcha urinishlardan keyin ham muvaffaqiyatsiz bo'ldi")
		}
	}
	fmt.Println("Ish", name, "muvaffaqiyatli bajarildi")
	return fmt.Sprint("Ish", name, "muvaffaqiyatli yakunlandi")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Background job simulyatsiyasi boshlandi")

	jobs := []string{"ish1", "ish2", "ish3", "ish4"}
	results := make(chan string, len(jobs))


	for _, job := range jobs {
		jobName := job
		go func() {
			results <- doJob(jobName, 3) 
		}()
	}


	for i := 0; i < len(jobs); i++ {
		fmt.Println(<-results)
	}
}
