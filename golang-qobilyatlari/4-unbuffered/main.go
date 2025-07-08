package main

import (
	"fmt"
	"time"
)

func ketish(ch chan string){
	time.Sleep(time.Second * 2)
	fmt.Println("habar yuborilyabdi")
	ch <- "Azizbek menga ishon sen albatdan testdan 100 otasan"
	fmt.Print("Habar yuborildi")	
}

func olish(ch chan string){
	fmt.Println("habar kutulyabdi ")
	msg := <- ch
	fmt.Print("Habar Oldindi", msg)	
}

func main(){
	ch := make(chan string)

	go ketish(ch)
	go olish(ch)

	fmt.Scanln()
	
}