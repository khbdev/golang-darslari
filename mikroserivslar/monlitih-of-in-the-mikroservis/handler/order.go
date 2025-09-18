package handler

import (
	"mikroservice/model"
	"mikroservice/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetOrders(c *gin.Context){
	Orders := storage.GetOrders()
	c.JSON(http.StatusOK, Orders)
}


func CreateOrder(c *gin.Context){
	var createOrders model.Order
	if err := c.ShouldBindJSON(&createOrders); err != nil{
		c.JSON(http.StatusOK, createOrders)
		return 
	}
	storage.AddOrder(createOrders)
		c.JSON(http.StatusOK, createOrders)
	
}