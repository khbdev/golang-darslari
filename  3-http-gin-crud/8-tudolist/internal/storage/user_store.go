package storage

import (
	"gin/internal/models"
	"sync"
)

var (
	users  = []models.User{}
	userID = 1
	muUser sync.Mutex
)

func CreateUser(u models.User) models.User {
	muUser.Lock()
	defer muUser.Unlock()

	u.ID = userID
	userID++
	users = append(users, u)
	return u
}

func FindUserByEmail(email string) *models.User {
	for _, u := range users {
		if u.Email == email {
			return &u
		}
	}
	return nil
}

func FindUserByApiKey(apiKey string) *models.User {
	for _, u := range users {
		if u.ApiKey == apiKey {
			return &u
		}
	}
	return nil
}
