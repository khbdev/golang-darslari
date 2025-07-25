package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN — data source name (user:password@tcp(host:port)/dbname)
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/GOLANGdatabase"

	// DB bilan ulanish
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ulanishda xato:", err)
	}
	defer db.Close()

	// Ping bilan tekshirish
	if err := db.Ping(); err != nil {
		log.Fatal("Ping xatosi:", err)
	}

	fmt.Println("✅ MySQL'ga muvaffaqiyatli ulanildi!")

	// Foydalanuvchilarni olish
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("Foydalanuvchilarni olishda xato:", err)
	}
	defer rows.Close()

	fmt.Println("\n📄 Foydalanuvchilar ro'yxati:")
	for rows.Next() {
		var id int
		var name, email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal("User ma'lumotini o'qishda xato:", err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Postlarni olish
	posts, err := db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		log.Fatal("Postlarni olishda xato:", err)
	}
	defer posts.Close()

	fmt.Println("\n📝 Foydalanuvchi postlari:")
	for posts.Next() {
		var id int
		var title, content string

		if err := posts.Scan(&id, &title, &content); err != nil {
			log.Fatal("Post ma'lumotini o'qishda xato:", err)
		}
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", id, title, content)
	}
}
