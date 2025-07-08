package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	user := User{
		Name:  "Azizbek Xasanov",
		Email: "aziz@example.com",
		Age:   25,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("❌ Xato:", err)
		return
	}

	// []byte ni stringga aylantiramiz
	fmt.Println("📦 JSON format:")
	fmt.Println(string(jsonData))
}
