package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"queue/task"

	"github.com/hibiken/asynq"
)

func HandleHelloTask(ctx context.Context, t *asynq.Task) error {
	var payload struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	fmt.Printf("👋 Salom, %s! Task bajarildi.\n", payload.Name)
	return nil
}

func main() {
	redisOpt := asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}

	srv := asynq.NewServer(redisOpt, asynq.Config{Concurrency: 10})

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeHello, HandleHelloTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Worker serveri xato: %v", err)
	}
}
