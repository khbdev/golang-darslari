package main

import "github.com/gin-gonic/gin"



func main(){
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Message": "Welcome To Hello world"})
	})

	r.Run(":8084")
}