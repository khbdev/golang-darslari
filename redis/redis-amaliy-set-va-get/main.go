package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)


var ctx = context.Background()

func main(){
rgb := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

_, err := rgb.Ping(ctx).Result()
if err != nil {
	fmt.Println("Xatolik", err)
}
fmt.Println("Redis Ulandik")
var name string

fmt.Print("Name kiriting: ")
fmt.Scanln(&name)
err = rgb.Set(ctx, "username", &name, 0).Err()
if err != nil {
	log.Fatalf("SET error: %v", err)
}
vel, err := rgb.Get(ctx, "username").Result()
if err != nil {
	log.Fatalf("Xatolik: %d", err)
}
fmt.Println("mana", vel)


err = rgb.Set(ctx, "khasanov", "Azizbek", 3*time.Second).Err()
if err != nil {
	log.Fatalf("Xatolik %d", err )
}

lastname, err := rgb.Get(ctx, "khasanov").Result()
if err != nil {
	log.Fatalf("Xatolik: %d", err)
}
fmt.Println("lastname", lastname)
}