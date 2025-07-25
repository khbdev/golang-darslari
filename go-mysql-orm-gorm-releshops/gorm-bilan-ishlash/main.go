package main

import (
	"fmt"
	"log"
	"time"


	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
)


type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	CreatedAt time.Time
}



func main(){
dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/GOLANGdatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM ulanish xatosi:", err)
	}
	fmt.Println("GORM orqali MySQL'ga muvaffaqiyatli ulanildi!")


	db.AutoMigrate(&User{})

	newUser := User{Name: "Vali", Email: "vali@gmail.coms"}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatal("Foydalanuvchi qo'shish xatosi:", result.Error)
	}
	fmt.Println("Yangi foydalanuvchi qo'shildi:", newUser.Name)

	var users []User
	db.Find(&users)
	fmt.Println("Foydalnuvchi Royhati:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, CreatedAt: %s\n",
			user.ID, user.Name, user.Email, user.CreatedAt)
	}
}