package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()



func main(){
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})


	err := rdb.LPush(ctx, "jobs", "email #1").Err()
	if err != nil {
	   panic(err)
	}
	fmt.Println("Job navatga qoshildi: email #1")
}