package main

import (
	"amaliy/models"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Posts{}, &models.Tag{})

	user := models.User{
		Name: "Ali",
		Profile: models.Profile{
			Bio: "Go developer",
		},
		Posts: []models.Posts{
			{
				Content: "First post",
				Tags: []models.Tag{
					{Name: "tech"},
					{Name: "go"},
				},
			},
		},
	}

	db.Create(&user)

	var Users models.User
	db.Preload("Profile").Preload("Posts.Tags").First(&Users, "name = ?", "Ali")


	fmt.Println("👤 User:")
	fmt.Printf("  Name: %s\n", Users.Name)

	fmt.Println("\n📄 Profile:")
	fmt.Printf("  Bio: %s\n", Users.Profile.Bio)

	fmt.Println("\n📝 Posts:")
	for i, post := range Users.Posts {
		fmt.Printf("  ▶ Post #%d: %s\n", i+1, post.Content)
		fmt.Printf("     📌 Tags: ")
		for _, tag := range post.Tags {
			fmt.Printf("#%s ", tag.Name)
		}
		fmt.Println()
	}


}
