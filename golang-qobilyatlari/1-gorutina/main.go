package main

import (
	"fmt"
	"sync"
)

func hello(name string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Salom: ", name)
}
  

func main(){
	
var name string

fmt.Print("Salom ismingizni kiritng: ")
fmt.Scanln(&name)

var wg sync.WaitGroup
wg.Add(1)
go hello(name, &wg)


wg.Wait()

}