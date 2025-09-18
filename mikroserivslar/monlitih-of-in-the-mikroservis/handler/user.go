package handler

import (
	"mikroservice/model"
	"mikroservice/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context){
users := storage.GetUsers()
c.JSON(http.StatusOK, users)
}


func CreateUser(c *gin.Context){
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil{
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	storage.AddUser(newUser)
	c.JSON(http.StatusOK, newUser)
}