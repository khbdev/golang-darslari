package main

import (
	"encoding/json"
	"fmt"

	"net/http"

)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func UserFunc(w http.ResponseWriter, r *http.Request){
   var user User

   err := json.NewDecoder(r.Body).Decode(&user)
   if err != nil {
	http.Error(w, "Json natogri", http.StatusBadRequest)
	return
   }

   fmt.Println("Name", user.Name)
   fmt.Println("Age", user.Age)
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
	http.HandleFunc("/users", UserFunc)
	http.HandleFunc("/", oneroute)
        http.HandleFunc("/salom", handler)
    http.ListenAndServe(":8082", nil)
}
