package main

import "github.com/gin-gonic/gin"


func main(){
r := gin.Default() 
 r.GET("/", func(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"xabar": "Salom hammaga"})
 })

 r.Run(":8082")
}