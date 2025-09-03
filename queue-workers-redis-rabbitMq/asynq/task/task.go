package task

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeHello = "hello:task"

func NewHelloTask(name string) *asynq.Task {
	payload, _ := json.Marshal(map[string]interface{}{"name": name})
	return asynq.NewTask(TypeHello, payload)
}
