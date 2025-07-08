package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Habar 1"
		close(ch1) // Kanalni yopamiz
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Habar 2"
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1 dan:", msg1)

		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2 dan:", msg2)
		}
	}
}
