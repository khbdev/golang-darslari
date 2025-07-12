package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // API guruhlash
    api := r.Group("/api/")
    {
        // GET: Query parametr bilan
        api.GET("/users", func(c *gin.Context) {
            name := c.Query("name")
            c.JSON(200, gin.H{
                "message": "Foydalanuvchilar ro'yxati",
                "name":    name,
            })
        })

        // POST: JSON body bilan
        api.POST("/users", func(c *gin.Context) {
            var user struct {
                Name string `json:"name"`
            }
            if err := c.BindJSON(&user); err != nil {
                c.JSON(400, gin.H{"error": err.Error()})
                return
            }
            c.JSON(201, gin.H{
                "message": "Yangi foydalanuvchi: " + user.Name,
            })
        })

        // PUT: Parametr bilan
        api.PUT("/users/:id", func(c *gin.Context) {
            id := c.Param("id")
            c.JSON(200, gin.H{
                "message": "Foydalanuvchi " + id + " yangilandi",
            })
        })

        // DELETE: Parametr bilan
        api.DELETE("/users/:id", func(c *gin.Context) {
            id := c.Param("id")
            c.JSON(200, gin.H{
                "message": "Foydalanuvchi " + id + " o'chirildi",
            })
        })
    }

    r.Run(":8001") // localhost:8080
}