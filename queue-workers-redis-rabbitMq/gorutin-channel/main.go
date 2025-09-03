package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())


	resultCh := make(chan string)

	go func() {
		const restirys = 3
		const delay = 2 * time.Second

		for i := 0; i < restirys; i++ {
			fmt.Printf("Urinish %d: job bajarilmoqda... \n", i+1)

			success := rand.Intn(2) == 1
			if success {
				resultCh <- "Job Mufaqiyatli bajarildi"
				return
			}

			fmt.Printf("Urinish %d muvaffaqiyatsiz. %v kechikish... \n", i+1, delay)
			time.Sleep(delay)

		
		}
			resultCh <- "Barcha urinishlar Mufaqiyatsiz"
	}()

	resulent := <-resultCh
	fmt.Println(resulent)
}