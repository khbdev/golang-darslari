package main 

import "fmt"


func main(){
// 	a := 10
// 	b := 5

// 	fmt.Println(a + b)  // 13
// fmt.Println(a - b)  // 7
// fmt.Println(a * b)  // 30
// fmt.Println(a / b)  // 3 (int bo‘lsa qoldiqsiz bo‘ladi)
// fmt.Println(a % b)  // 1 (modul: qoldiq)


// fmt.Println(a == b)  // false
// fmt.Println(a != b)  // true
// fmt.Println(a > b)   // true

age := 30

// fmt.Println(age > 18 && age < 30)  // true
// fmt.Println(age < 18 || age > 25)  // true
// fmt.Println(!(age < 18))     

if age < 20 {
	fmt.Println("Voyaga yetmgaga")
} else if age < 30 {
	fmt.Println("yosh")
} else {
	 fmt.Println("katta yosh")
}


grade := "A"

switch grade {
case "A":
	fmt.Println("Zor!")

case "B":
	fmt.Println("Yaxshi")

case "C":
	fmt.Println("Boladi")
	} 


	for i := 1; i <= 555; i++{
		fmt.Println(i)
	}

	arr := []string{"Go", "Rust", "Python"}

for index, value := range arr {
    fmt.Printf("%d: %s\n", index, value)
}



numbers()

for100()

}

func numbers() {
	var num int

	fmt.Print("Son kiriting: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Juft son")
	} else {
		fmt.Print("Toq Son")
	}
}


func for100(){
for i := 0; i < 100; i++ {
	fmt.Println(i)
}



var temp int
fmt.Print("Haroratni kiriting: ")
fmt.Scanln(&temp)


switch {
case temp <= 0:
    fmt.Println("Sovuq")
case temp <= 25:
    fmt.Println("Yoqimli ob-havo")
default:
    fmt.Println("Issiq")
}
}