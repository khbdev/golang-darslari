package main

import (
	"database/sql"
	"fmt"
	"jwt/auth"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func main(){
	if err := godotenv.Load(); err != nil {
		fmt.Println("Xatolik Ulanib bomadi", err)
	}
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Mysql Xatolik ulanib bomadi", err)
	}
		  if err := db.Ping(); err != nil {
        fmt.Println("DB ulanishda xatolik:", err)
    }
	fmt.Println("mysql mufaqilatli")

	auth.DB = db

	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/login", auth.LoginHandler)

	fmt.Println("server 8002 portda ishga tushdi")
	http.ListenAndServe(":8002", nil)

}