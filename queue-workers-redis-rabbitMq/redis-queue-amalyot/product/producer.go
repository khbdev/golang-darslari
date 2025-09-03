package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)



var ctx = context.Background()


func main(){
		rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

      

		job := "send_email:khbcoder@gmail.com"

		err := rdb.LPush(ctx, "jobs", job).Err()

			if err != nil {
		panic(err)
	}
	rdb.Expire(ctx, "jobs", 60*time.Second)
		fmt.Println("Job navbatga qo‘shildi:", job)
}