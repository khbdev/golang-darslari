package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type User struct{
	gorm.Model
    Name  string
    Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	UserID uint
	Content string
}

func main(){
	// DSN format: username:password@tcp(host:port)/dbname?parseTime=true
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL bilan ulanishda xatolik:", err)
	}

	fmt.Println("MySQL bilan muvaffaqiyatli ulandi!")


	db.AutoMigrate(&User{}, &Post{})

	user := User{
		Name: "Valis",
		Posts: []Post{
			{Content: "Birinchi post"},
        {Content: "Ikkinchi post"},
		},

	}
	db.Create(&user)
	fmt.Println(" muvaffaqiyatli !")

var users []User
db.Preload("Posts").Find(&users)

for _, u := range users {
    fmt.Printf("User: %s\n", u.Name)
    for _, p := range u.Posts {
        fmt.Printf("  Post: %s\n", p.Content)
    }
}
}