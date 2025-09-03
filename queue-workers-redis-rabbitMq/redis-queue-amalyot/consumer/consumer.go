package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	for {
		// Navbatdan job olish
		result, err := rdb.BRPop(ctx, 0*time.Second, "jobs").Result()
		if err != nil {
			panic(err)
		}
		job := result[1]
		fmt.Println("Consumer oldi:", job)

		// Fake bajarish (50% muvaffaqiyatsiz bo‘lsin)
		if time.Now().Unix()%2 == 0 {
			fmt.Println("❌ Xatolik:", job)
			// Retry hisobini oshiramiz
			retryKey := "retry_count:" + job
			count, _ := rdb.Incr(ctx, retryKey).Result()
			rdb.Expire(ctx, retryKey, 1*time.Hour) // Retry hisobini saqlash muddati

			if count < 3 {
				// Yana jobs queue’ga tashlash
				rdb.LPush(ctx, "jobs", job)
				fmt.Println("↩️  Retryga qaytardi:", job)
			} else {
				// Dead queue
				rdb.LPush(ctx, "dead_jobs", job)
				fmt.Println("💀 Dead queue:", job)
			}

			// Pub/Sub xabar yuboramiz
			rdb.Publish(ctx, "job_events", fmt.Sprintf("fail:%s retry:%d", job, count))
		} else {
			fmt.Println("✅ Muvaffaqiyatli bajarildi:", job)
			// Pub/Sub xabar yuboramiz
			rdb.Publish(ctx, "job_events", fmt.Sprintf("success:%s", job))
		}
	}
}
