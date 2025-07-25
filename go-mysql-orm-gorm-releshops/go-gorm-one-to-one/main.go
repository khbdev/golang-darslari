package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string
	Email   string   `gorm:"unique"`
	Profile Profile  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Profile struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"uniqueIndex"` // 1:1 bog'lash uchun unique bo'lishi shart
	Bio    string
	Age    int
}

func main(){
dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Bazaga ulanishda xatolik:", err)
	}
	// 📛 Eski jadvallarni tozalash
	err = db.Migrator().DropTable(&Profile{}, &User{})
	if err != nil {
		log.Fatal("DropTable xatolik:", err)
	}

		err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("Migratsiyada xatolik:", err)
	}
	user := User{
		Name:  "Azizbek",
		Email: "azizbek@example.com",
		Profile: Profile{
			Bio: "Go backend developer",
			Age: 23,
		},
	}


	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Yaratishda xatolik:", err)
	}

		var fetchedUser []User
	if err := db.Preload("Profile").First(&fetchedUser, user.ID).Error; err != nil {
		log.Fatal("Foydalanuvchini olishda xatolik:", err)
	}
	for _, u := range fetchedUser {
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	fmt.Printf("  → Bio: %s, Age: %d\n", u.Profile.Bio, u.Profile.Age)
	fmt.Println("----------------------------")
}

	
}