package main

import "encoding/json"
import "fmt"


type User struct{
		Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}




func main() {
 	jsonUmarshal := `{
		"Name": "Azizbek Xasanov",
		"Email": "azizbek@example.com",
		"Age": 25
	}`

	var user User

	err := json.Unmarshal([]byte(jsonUmarshal), &user)
if err != nil {
	fmt.Println("Xato", err)
   return
}

	fmt.Println("User name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Yosh:", user.Age)
}