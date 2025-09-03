package main

import (
	"fmt"
	"math/rand"
	"time"
)

type EmailJob struct {
	ID      int
	To      string
	Subject string
	Body    string
	Retries int
}

func worker(id int, jobs chan EmailJob, results chan<- string, ohirgi chan<-string) {
	for job := range jobs {
		fmt.Printf("worker %d: email ketdi #%d to %s\n", id, job.ID, job.To)
		success := sendEmail(job)
		if success {
			results <- fmt.Sprintf("Worker %d: Email #%d email togri keldi ✅", id, job.ID)
			ohirgi <- "Hamma Email togri keldi"
		} else {
			if job.Retries > 0 {
				fmt.Printf("Worker %d: Email #%d mufaqiyatsiz , takrorlanadi 2skundan keyin...\n", id, job.ID)
				time.Sleep(2 * time.Second)
				job.Retries--
				jobs <- job // retry
			} else {
				results <- fmt.Sprintf("Worker %d: Email #%d umuman mufaqiyatsiz boldi ❌", id, job.ID)
               ohirgi <- "Hamma Email fail boldi"
			}
		}
	}
}

// Bu faqat simulate qiladi, haqiqiy email yubormaydi
func sendEmail(job EmailJob) bool {
	// Tasodifiy muvaffaqiyat/fail
	return rand.Intn(1) == 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan EmailJob, 1)
	results := make(chan string, 1)
	ohirgi := make(chan string)

	// 3 ta worker ishga tushadi
	for w := 1; w <= 1; w++ {
		go worker(w, jobs, results, ohirgi)
	}

	// 5 ta email job yaratamiz
	for i := 1; i <= 1; i++ {
		jobs <- EmailJob{
			ID:      i,
			To:      fmt.Sprintf("khbcoder@gmail.com"),
			Subject: "Salom!",
			Body:    "karochi bu email.",
			Retries: 3, 
		}
	}

	for i := 1; i <= 1; i++ {
		fmt.Println(<-results)
	}

	for i := 0; i < 1; i++ {
		fmt.Println(<-ohirgi)
	}
}
