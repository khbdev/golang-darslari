package main

import (


	"github.com/gin-gonic/gin"
)

type Book struct
{
	Name string `json:"name"`
	Description string `json:"description"`
}


func main(){
	r := gin.Default()

	r.POST("/book", func(ctx *gin.Context) {
    var book Book
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{
		"data" : "Yangi Kitob nomi" + book.Name, 
		"description": "Kitob tavsifi" + book.Description,
	})
	})
	r.Run(":8001")
}