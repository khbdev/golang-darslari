package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("date.txt")
	if err != nil {
		fmt.Println("❌ Xato:", err)
		return
	}

	fmt.Println("⏳ 2 soniya kutyapmiz...")
	time.Sleep(2 * time.Second)
	fmt.Println("✅ 2 soniya o‘tdi!")

	fmt.Println("📄 Fayl ichidagi matn:")
	fmt.Println(string(data))
}
