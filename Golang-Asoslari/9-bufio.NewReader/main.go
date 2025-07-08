package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Faylni ochamiz
	file, err := os.Open("date.txt")
	if err != nil {
		fmt.Println("❌ Xato:", err)
		return
	}
	defer file.Close() // Oxirida faylni yopamiz

	// 2. O‘quvchini (reader) yaratamiz
	reader := bufio.NewReader(file)

	fmt.Println("📖 Fayldagi satrlar:")

	// 3. Fayl tugamaguncha har bir satrni o‘qiymiz
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // EOF bo‘lsa chiqib ketamiz
		}
		fmt.Print(line) // satrni chiqaramiz
	}
}
