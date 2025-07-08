package main

import (
	"fmt"
)

func main() {

	Users := map[string]string{
		"user1": "Azizbek Xasanov",
		"user2": "Javohir berdiyev",
		"user3": "Abduqodir khusanov",
	}

	for key, value := range Users {
		fmt.Println(key, ":", value)
	}

	ball := make(map[string]int)

	ball["Azizbek"] = 10
	ball["Soliha"] = 20
	ball["Javohir"] = 70

	fmt.Println(ball["Azizbek"])
	fmt.Println(ball["Soliha"])
	fmt.Println(ball["Javohir"])

}
