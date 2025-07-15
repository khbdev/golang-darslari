package main

import "fmt"


func main(){

ages()
swetchesTest("Chorshanbas")
} 

func swetchesTest(d string) *string{
	days := d

	switch days{
	case "Dushanba":
		fmt.Println("Bugun dushanba")
		case "Seshanba":
		fmt.Println("Bugun Seshanba")
		case "Chorshanba":
		fmt.Println("Bugun Chorshanba")
		default:
			fmt.Println("Xato")
	}
	return &days
}



func ages(){
	var age int = 16

if age == 16 {
	fmt.Println("Sizning yoshingiz 16")
} else if age == 17{
	fmt.Println("Sizning yoshingiz 17")
}else {
	fmt.Println("bunday yosh qiymatdan yoq")
}
}