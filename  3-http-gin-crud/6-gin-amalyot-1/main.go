package main

import "github.com/gin-gonic/gin"


func main(){
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
    "message": "User shu inson: " + id,
		})
	})

	router.Run(":8001")
}