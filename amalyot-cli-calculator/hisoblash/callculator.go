package hisoblash

import "fmt"

func Calculator(a, b int, amal string) *int {
	var natija int
	switch amal {
	case "+":
		natija = a + b
	case "-":
		natija = a - b
	case "*":
		natija = a * b
	case "/":
		if b == 0 {
			fmt.Println("Xatolik: 0 ga bo‘lish mumkin emas.")
			return nil
		}
		natija = a / b
	default:
		fmt.Println("Xatolik: noto‘g‘ri amal.")
		return nil
	}
	return &natija
}
