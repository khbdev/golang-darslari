package storage

import (
	"gin/internal/models"
	"sync"
)

var (
	todos  = []models.Todo{}
	todoID = 1
	muTodo sync.Mutex
)

func CreateTodo(t models.Todo) models.Todo {
	muTodo.Lock()
	defer muTodo.Unlock()

	t.ID = todoID
	todoID++
	todos = append(todos, t)
	return t
}

func GetTodosByUser(userID int) []models.Todo {
	userTodos := []models.Todo{}
	for _, t := range todos {
		if t.UserID == userID {
			userTodos = append(userTodos, t)
		}
	}
	return userTodos
}

func GetTodoByID(id int, userID int) *models.Todo {
	for _, t := range todos {
		if t.ID == id && t.UserID == userID {
			return &t
		}
	}
	return nil
}

func UpdateTodo(id int, updated models.Todo, userID int) bool {
	muTodo.Lock()
	defer muTodo.Unlock()
	for i, t := range todos {
		if t.ID == id && t.UserID == userID {
			updated.ID = id
			updated.UserID = userID
			todos[i] = updated
			return true
		}
	}
	return false
}

func DeleteTodo(id int, userID int) bool {
	muTodo.Lock()
	defer muTodo.Unlock()
	for i, t := range todos {
		if t.ID == id && t.UserID == userID {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
