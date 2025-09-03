package main

import (
	"fmt"
	"log"
	"net/http"
	"queue/task"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "name query param kiritilmagan", http.StatusBadRequest)
			return
		}

		task := task.NewHelloTask(name)
		info, err := client.Enqueue(task)
		if err != nil {
			http.Error(w, "Task enqueue xato: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "✅ Task yuborildi! ID=%s, Queue=%s\n", info.ID, info.Queue)
	})

	fmt.Println("🚀 HTTP server 8084-portda ishlayapti...")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
