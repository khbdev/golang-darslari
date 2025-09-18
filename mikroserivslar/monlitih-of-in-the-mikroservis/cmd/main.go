package main

import (
	"mikroservice/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine{
r := gin.Default()

r.GET("/users", handler.GetUsers)
r.POST("/users", handler.CreateUser)

r.GET("/orders", handler.GetOrders)
r.POST("/orders", handler.CreateOrder)
r.GET("/products", handler.GetProducts)
r.POST("/products", handler.CreateProduct)
return  r
}


func main(){

r := SetupRouters()

r.Run(":8083")
}