package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}


func main(){
		dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	db.AutoMigrate(&User{})

	var users []User
	db.Where("age = ?", 25).Find(&users)

	fmt.Println("25 yoshli foydalanuvchilar: ")

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}
}