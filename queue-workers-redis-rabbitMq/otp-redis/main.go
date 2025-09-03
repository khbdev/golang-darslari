package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
	})
	return rdb
}



func produceOTP(rdb *redis.Client, userID string) {
	// Random OTP yaratamiz
	otp := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)

	// Redis’da TTL bilan saqlaymiz
	key := "otp:" + userID
	rdb.Set(ctx, key, otp, 5*time.Minute) // 5 daqiqaga amal qiladi

	// Queue’ga tashlaymiz (kim uchun OTP yaratildi)
	rdb.LPush(ctx, "otp_queue", userID)

	fmt.Println("📩 Producer: OTP yaratildi va queue’ga qo‘yildi:", userID, otp)
}

func consumeOTP(rdb *redis.Client) {
	for {
		// Queue’dan bitta userID olish (blocking)
		userID, err := rdb.BRPop(ctx, 0*time.Second, "otp_queue").Result()
		if err != nil {
			fmt.Println("Consumer xato:", err)
			continue
		}

		// BRPop qaytaradi: [queueName, value]
		if len(userID) < 2 {
			continue
		}
		uid := userID[1]

		// Redis’dan OTP olib ko‘ramiz
		key := "otp:" + uid
		otp, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			fmt.Println("❌ OTP topilmadi yoki eskirgan:", uid)
		} else if err != nil {
			fmt.Println("Redis xato:", err)
		} else {
			fmt.Println("✅ Consumer: OTP mavjud:", uid, otp)
		}
	}
}

func main() {
	rdb := connectRedis()

	// Producer → bitta user uchun OTP yaratadi
	go func() {
		for i := 1; i <= 3; i++ {
			userID := fmt.Sprintf("user%d", i)
			produceOTP(rdb, userID)
			time.Sleep(2 * time.Second) // oraliq
		}
	}()

	// Consumer → doim queue’ni kuzatib turadi
	consumeOTP(rdb)
}

