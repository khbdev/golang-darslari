package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)
var msg string = "salom"

func forlet() []string {
	var message []string
	for i := 0; i < 1000; i++ {
	message = append(message, fmt.Sprintf("%d: %s", i+1, msg))
}
return message
}

func main(){
	r := gin.Default()

	r.GET("/hello",  func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messages": forlet(),
		})
	})

	r.Run(":8001")
}