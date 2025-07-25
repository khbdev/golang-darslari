package main

import (
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Salom go")
	})

	http.HandleFunc("/Azizbek", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Salom mening ismim Azizbek")
	})

	fmt.Println("Server ishga tushdi 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		fmt.Println("Hatolik yuz berdi", err)
	}
}