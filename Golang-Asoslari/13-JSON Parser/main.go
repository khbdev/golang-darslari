package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	start := time.Now()

	file, err := os.Open("large_data.json")
	if err != nil {
		fmt.Println("Xato:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("O‘qishda xato:", err)
		return
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Parse xatosi:", err)
		return
	}

	fmt.Println("Foydalanuvchilar soni:", len(users))
	fmt.Println("1-user:", users[0])
	fmt.Println("Oxirgi user:", users[len(users)-1])

	duration := time.Since(start)
	fmt.Println("Umumiy vaqt:", duration)
}
