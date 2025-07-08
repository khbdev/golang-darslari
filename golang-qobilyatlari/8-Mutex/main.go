package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex // 🔐 Mutex yaratamiz

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()     // ✋ Kirishdan oldin qulfla
		counter++     // 🧮 Havfsiz inkrement
		mu.Unlock()   // ✅ Ish bitdi, boshqalarga ruxsat ber
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter =", counter) // Doim 10000 ✅
}
