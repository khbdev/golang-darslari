package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)



var ctxs = context.Background()


func main(){
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	for {
		result, err  := rdb.BRPop(ctxs, 0, "jobs").Result()
		if err != nil {
			panic(err)
		}

		fmt.Println("Ichi olindi", result)
	}
}