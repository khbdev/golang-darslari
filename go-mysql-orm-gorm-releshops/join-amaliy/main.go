package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

type Profile struct {
	ID     uint
	UserID uint
	Bio    string
}

type UserWithProfile struct {
	Name  string
	Email string
	Bio   string
}

func main() {
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB ulanib bo‘lmadi: ", err)
	}

	// Migratsiya
	db.AutoMigrate(&User{}, &Profile{})

	// Faqat bir marta create qilish uchun comment qilsang bo'ladi
	db.Create(&User{Name: "Azizbek", Email: "aziz@example.com"})
	db.Create(&Profile{UserID: 1, Bio: "Mazgi Go developer 😄"})

	// INNER JOIN orqali birlashtirish
	var result []UserWithProfile

	db.Table("users").
		Select("users.name, users.email, profiles.bio").
		Joins("INNER JOIN profiles ON profiles.user_id = users.id").
		Scan(&result)

	for _, user := range result {
		fmt.Printf("Name: %s, Email: %s, Bio: %s\n", user.Name, user.Email, user.Bio)
	}
}
