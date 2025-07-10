package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Fayl yaratish
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("❌ Xato:", err)
		return
	}
	defer file.Close()

	// 2. Oddiy yozish
	file.WriteString("👋 Salom, bu oddiy yozish\n")

	// 3. Bufer bilan yozish
	writer := bufio.NewWriter(file)
	writer.WriteString("📝 Bu satr bufer orqali yozildi\n")
	writer.WriteString("✅ Jarayon tugadi\n")

	// 4. Buferni faylga “oqizish”
	writer.Flush()

	date, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Xatolik", err)
	}

  
	fmt.Println(string(date))

}
