package handler

import (
	"mikroservice/model"
	"mikroservice/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetProducts(c *gin.Context) {
	products := storage.GetProducts()
	c.JSON(http.StatusOK, products)
}


func CreateProduct(c *gin.Context) {
	var newProduct model.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage.AddProduct(newProduct)
	c.JSON(http.StatusOK, newProduct)
}
