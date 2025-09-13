package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// 	"github.com/redis/go-redis/v9"
// )



// var ctx = context.Background()

// var upgrader = websocket.Upgrader{
//     CheckOrigin: func(r *http.Request) bool {
//         return true // ⚡️ barcha originlarga ruxsat
//     },
// }

// func main(){
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})



// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		conn, err := upgrader.Upgrade(w, r, nil)

// 		if err != nil {
// 			fmt.Println("Upgrade error:", err)
// 			return
// 		}
// 		defer conn.Close()

// 		sub := rdb.Subscribe(ctx, "todos")
// 		ch := sub.Channel()


// 		for  msg := range ch {
// 				fmt.Println("📩 New Todo from Redis:", msg.Payload)
// 			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
// 			if err != nil {
// 				fmt.Println("Write error:", err)
// 				return
// 			}
// 		}
// 	})


// 	fmt.Println("🚀 Server running on :8085")
// 	http.ListenAndServe(":8085", nil)
// }