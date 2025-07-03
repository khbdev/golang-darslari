package main

import (
	"cli-callculator/hisoblash" // go mod init qilgan bo‘lishing kerak
	"fmt"
	"strconv"
)

func main() {
	var s1, s2, amal string

	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scanln(&s1)

	fmt.Print("Amalni kiriting (+, -, *, /): ")
	fmt.Scanln(&amal)

	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scanln(&s2)

	son1, err1 := strconv.Atoi(s1)
	son2, err2 := strconv.Atoi(s2)

	if err1 != nil || err2 != nil {
		fmt.Println("Xatolik: faqat son kiriting")
		return
	}

	natijaPtr := hisoblash.Calculator(son1, son2, amal)

	if natijaPtr != nil {
		fmt.Printf("Natija: %d %s %d = %d\n", son1, amal, son2, *natijaPtr)
	}
}
