package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`      // Primary key
	Name      string    `gorm:"size:255"`        // Max uzunlik 255
	Email     string    `gorm:"size:100;unique"` // Unikal email
	CreatedAt time.Time // Yaratilgan vaqt
}

func createUser(db *gorm.DB, name, email string)error {
  user := User{
        Name:  name,
        Email: email,
    }

    result := db.Create(&user)
    return result.Error
}

func allUsers(db *gorm.DB) ([]User, error) {
    var users []User

    result := db.Find(&users)
    return users, result.Error
}

func deleteUser(db *gorm.DB, id uint) error {
    result := db.Delete(&User{}, id)
    return result.Error
}

func main() {
 dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/GOLANGdatabase?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalln("MySQLga ulanishda xatolik:", err)
    }
	log.Println("MySQLga muvaffaqiyatli ulandik!")



	err = createUser(db, "Azizbek", "khbcoder@gmail.comsss")

	if err != nil {
		log.Fatal("Foydalanuvchi qo'shish xatosi:", err)
	}
	fmt.Println("Yangi foydalanuvchi qo'shildi: Azizbek")


	users, err := allUsers(db)
	if err != nil {
		log.Fatal("Foydalanuvchilarni olish xatosi:", err)
	}
	fmt.Print("Hamma Users")

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, CreatedAt: %s\n",
			user.ID, user.Name, user.Email, user.CreatedAt)
	}

	err = deleteUser(db, 1)

		if err != nil {
		log.Fatal("Foydalanuvchi o'chirish xatosi:", err)
	}
	fmt.Println("ID=1 bo'lgan foydalanuvchi o'chirildi")

		users, err = allUsers(db)
	if err != nil {
		log.Fatal("Foydalanuvchilarni olish xatosi:", err)
	}
	fmt.Println("Yangilangan foydalanuvchilar ro'yxati:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, CreatedAt: %s\n",
			user.ID, user.Name, user.Email, user.CreatedAt)
	}
	
}
