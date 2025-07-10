package main

import "fmt"

func hello(channel chan string){
  channel <- "Salom azizbek"
}

func main(){
	channel := make(chan string)
  
	go hello(channel)

	msg := <- channel
	fmt.Println(msg)

}