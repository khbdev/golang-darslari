package main 

import "fmt"

// var hello string = "Salom "
// var number int = 23
// var float float64 = 12.3
// var Bool bool = 2 == 4

// const ozgarmas = "bu umuman ozgarmaydi"

func main(){
	var ism string
	var yosh int
	// hellow := "hello 2"

	// fmt.Println(hello)
	// fmt.Println(hellow)
	// fmt.Println(number)
	// fmt.Println(float)
    // fmt.Println(Bool)
	// fmt.Println(ozgarmas)

	fmt.Print("Ismingizni Kiriting: ")
	fmt.Scanln(&ism)


	fmt.Print("Yoshingizni Kiriting: ")
	fmt.Scanln(&yosh)


	fmt.Printf("Salom %s", ism)
	fmt.Printf("sizning yoshingiz %d, 10yildan song shunday boladi ", yosh+10)



}