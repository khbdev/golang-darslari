package main

import (
	"fmt"
	"sync"
)

func ish(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Gorutina %d ish boshladi\n", n) 
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go ish(i, &wg)
	}
	wg.Wait()
	fmt.Println("Hammasi shu")
}
