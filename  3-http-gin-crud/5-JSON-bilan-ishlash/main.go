package main

import (
    "github.com/gin-gonic/gin"
)

type User struct {
    ID    int    `json:"id" binding:"required"`
    Name  string `json:"name" binding:"required"`
    Email string `json:"email"`
}

func main() {
    r := gin.Default()

   
    r.POST("/users", func(c *gin.Context) {
        var user User

        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }
   
        c.JSON(201, gin.H{
            "message": "Foydalanuvchi yaratildi",
            "user":    user,
        })
    })

    r.Run() 
}