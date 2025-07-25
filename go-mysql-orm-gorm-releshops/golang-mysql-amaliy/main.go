package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

 

func main(){
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/golangAmaliy"

	db, err := sql.Open("mysql", dsn)
if  err != nil {
	fmt.Println("Xatolik", err)
	return
}

defer db.Close()


err = db.Ping()
if err != nil {
	fmt.Println("Xatolik", err)
	return
}
fmt.Println("Ulandi")


name := "Azizbek"
email := "khbcoder@gmail.comss"

query := "INSERT INTO users (name, email) VALUES (?, ?)"
result, err := db.Exec(query, name, email)
if err != nil {
		fmt.Println("INSERT xatolik:", err)
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("ID olishda xatolik:", err)
		return
	}

		rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("So‘rovda xatolik:", err)
	}
	defer rows.Close()


		for rows.Next() {
		var id int
		var name, email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println("Scan xatoligi:", err)
			continue
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	fmt.Println("Foydalanuvchi qo‘shildi. ID:", lastInsertID)
}
