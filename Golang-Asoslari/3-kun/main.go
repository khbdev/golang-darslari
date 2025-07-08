package main

import "fmt"

func main() {
	// Paramern funksiya chaqirish
	var name string

	fmt.Print("Ismingizni Kiritng: ")
	fmt.Scanln(&name)

	ism := ismFunk(name)

	println(ism)
	//

	qoshish, ayrish := hisoblash(10, 5)
	fmt.Println("Qoshish: ", qoshish)
	fmt.Println("ayrish: ", ayrish)


	qongiroq(5)

	for i := 0; i < 7; i++ {
		fmt.Print(fibonacci(i), " ")
	}

	 a := 5
    setTo100(&a)
    fmt.Println("a:", a) // 100
}

// Paremetr funksiya

func ismFunk(ism string) string {
	return "Ismingiz, " + ism
}
// 

// multiple return - bir nechta natija qaytarish

func hisoblash(a int, b int) (int, int) {
	return a + b, a - b
}

// recursion - ozini ozi chaqiradigan funskiya

func qongiroq(xabarlarSoni int){
	if xabarlarSoni == 0 {
		fmt.Println("Dostim Javob berdi!")
		return
	}
	fmt.Println("Qongiroq qilingi........... dostim javob bermadi")
	qongiroq(xabarlarSoni -1)
}

// Fibonacci - bu funksiyasi

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func setTo100(x *int) {
    *x = 100
}