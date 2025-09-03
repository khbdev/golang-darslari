package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	sub := rdb.Subscribe(ctx, "job_events")

	ch := sub.Channel()

	for msg := range ch {
		fmt.Println("📢 Event:", msg.Payload)
	}
}
