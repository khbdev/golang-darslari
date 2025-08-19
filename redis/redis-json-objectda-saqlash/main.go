package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Xatolik: %v", err)
	}

	fmt.Println("Redisga ulanish muvaffaqiyatli")

	user := User{
		Name:  "Azizbek",
		Age:   17,
		Email: "khbcoder@gmail.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("JSON xatolik: %v", err)
	}


	err = rdb.Set(ctx, "user:1", jsonData, 12*time.Second).Err()
	if err != nil {
		log.Fatalf("Set xatolik: %v", err)
	}

	fmt.Println("User ma'lumotlar Redisga saqlandi. 12 sekund kutamiz...")

	
	time.Sleep(5 * time.Second)
	val, err := rdb.Get(ctx, "user:1").Result()
	if err != nil {
		log.Fatalf("Get xatolik: %v", err)
	}
	fmt.Println("5s keyin:", string(val))

	time.Sleep(8 * time.Second) 
	val, err = rdb.Get(ctx, "user:1").Result()
	if err == redis.Nil {
		fmt.Println("13s keyin: ma'lumot topilmadi (o'chirilgan)")
	} else if err != nil {
		log.Fatalf("Get xatolik: %v", err)
	} else {
		fmt.Println("13s keyin:", string(val))
	}
}
