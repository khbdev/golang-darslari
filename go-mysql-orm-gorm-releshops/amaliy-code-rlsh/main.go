package main

import (
	"amaliy/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main(){
	// DSN (Data Source Name) format: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"

	// GORM orqali ulanish
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ MySQL bazasiga ulanishda xatolik: " + err.Error())
	}

	fmt.Println("✅ MySQL bazasiga muvaffaqiyatli ulandik!")

	// Hozircha faqat connection test
	_ = db

	err = db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Post{},
		&models.Tags{},
	)
	if err != nil {
		panic("❌ AutoMigrate xatolikka uchradi: " + err.Error())
	}

	user := models.User{
	Name: "Azizbek",

	Profile: &models.Profile{
		Bio: "Backend Golang Developer 👨‍💻",
	},

	Posts: []models.Post{
		{
			Content: "GORM bilan ishlash zo'r ekan!",
			Tags: []models.Tags{
				{Name: "golang"},
				{Name: "gorm"},
			},
		},
		{
			Content: "Bugun Laravelda API yozdim!",
			Tags: []models.Tags{
				{Name: "laravel"},
				{Name: "backend"},
			},
		},
	},
}
if err := db.Create(&user).Error; err != nil {
	fmt.Printf("❌ User yaratishda xatolik: %v", err)
}
	fmt.Println("Malumot bazaga yozildi")

	fmt.Println("✅ AutoMigrate muvaffaqiyatli bajarildi!")

	var users []models.User

if err := db.Preload("Profile").Preload("Posts.Tags").Find(&users).Error; err != nil {
	log.Fatalf("❌ Ma'lumotlarni olishda xatolik: %v", err)
}
for _, user := range users {
	fmt.Println("👤 User:", user.Name)
	fmt.Println("   🧠 Bio:", user.Profile.Bio)

	for _, post := range user.Posts {
		fmt.Println("   ✍️ Post:", post.Content)
		for _, tag := range post.Tags {
			fmt.Println("      🏷️ Tag:", tag.Name)
		}
	}
}

}