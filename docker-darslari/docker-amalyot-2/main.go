package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Faqat GET method ishlaydi!", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server 8085-portda ishlamoqda...")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal(err)
	}
}
