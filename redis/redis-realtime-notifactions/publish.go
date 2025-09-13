package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Oddiy publish qilish
	err := rdb.Publish(ctx, "todos", `{"id":1,"title":"Learn Go"}`).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Todo publish qilindi!")
}