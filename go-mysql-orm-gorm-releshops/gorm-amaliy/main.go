package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type User struct {
	ID uint
	Name string
	Email string
}

func main(){
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("MySQL ulanishda xatolik:", err)
		return
	}

	fmt.Println("Ulandi onasini emsin")
	db.AutoMigrate()

	users := []User{
		{Name: "Azizbek", Email: "khbcodergmail.comssssssssss"},
	}
	if err := db.Create(&users).Error; err != nil{
		fmt.Println("❌ Ma'lumot yozishda xatolik:", err)
		return
	}

	fmt.Println("Qoshildi")


	var userss []User
	if err := db.Find(&userss).Error; err != nil {
		log.Fatal("Xatolik", err)
	}

	for _, u := range userss {
		fmt.Printf("ID: %d, Name: %s", u.ID, u.Name,)
	}
}