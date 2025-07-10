package main

import (
	"fmt"

)

func sender(ch chan string) {
	ch <- " 1-xabar"
	ch <- "2-xabar"
	ch <- "3-xabar"

	close(ch)
}

func receiver(ch chan string) {
	for msg := range ch {
		fmt.Println(" Receiver: ", msg)
	}
}

func main() {
	ch := make(chan string, 3) // 3 ta qiymat sig‘adigan buffered channel

	go sender(ch)
receiver(ch)
}
