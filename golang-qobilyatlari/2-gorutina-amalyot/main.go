package main

import (
	"fmt"
	"sync"
	"time"
)


func salom(name string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Salom: ", name)
}

func age(number int, wg * sync.WaitGroup){
defer wg.Done()
fmt.Println("Yosh:", number)
}

func main(){
var wg sync.WaitGroup

var n int
fmt.Print("Nechta ism kiritmoqchisiz? ")
fmt.Scanln(&n)

names := make([]string, n)

for i := 0; i < n; i++ {
	fmt.Printf("%d-ism kiritng:", i+1)
	fmt.Scanln(&names[i])
}
for _, name := range names{
	wg.Add(2)
	go salom(name, &wg)
	go age(16, &wg)
}
wg.Wait()
fmt.Println("Barcha salomlar tugadi")
time.Sleep(time.Millisecond * 900)
}