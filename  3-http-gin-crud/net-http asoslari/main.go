package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)
func postHandler(w http.ResponseWriter, r * http.Request){
	bodyBytes, _ := io.ReadAll(r.Body)
	fmt.Println("Body: ", string(bodyBytes))
}

func oneroute(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-type", "application/json")
	date := map[string]string{"xabar": "salom"}
	json.NewEncoder(w).Encode(date)
}

func handler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

date := map[string]string{"xabar": "Salom json"}
 json.NewEncoder(w).Encode(date)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/", oneroute)
        http.HandleFunc("/salom", handler)
    http.ListenAndServe(":8082", nil)
}
