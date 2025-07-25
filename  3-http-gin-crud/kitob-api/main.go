package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Book struct — kitob modeli
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// In-memory storage — kitoblar ro‘yxati
var books = []Book{
	{ID: 1, Title: "Go Programming", Author: "John Doe"},
	{ID: 2, Title: "REST APIs", Author: "Jane Smith"},
}

// Logger Middleware — Har bir so‘rovni konsolga yozadi
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// So‘rov vaqtini olish
		start := time.Now()
		// So‘rov yo‘li va metodi
		path := c.Request.URL.Path
		method := c.Request.Method

		// Keyingi handler’ga o‘tish
		c.Next()

		// So‘rov tugaganidan keyin log yozish
		latency := time.Since(start)
		fmt.Printf("[%s] %s %s %v\n", time.Now().Format(time.RFC3339), method, path, latency)
	}
}

// Custom API Key Middleware — X-Api-Key sarlavhasini tekshiradi
func apiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("salom")
		if apiKey != "secret-key" { // "secret-key" o‘rniga haqiqiy kalit qo‘yiladi
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort() // So‘rovni to‘xtatish
			return
		}
		c.Next() // Keyingi handler’ga o‘tish
	}
}

// GET /books — Barcha kitoblar
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// GET /books/:id — Bitta kitob
func getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// POST /books — Yangi kitob qo‘shish
func createBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Auto-increment ID
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// PUT /books/:id — Kitobni yangilash
func updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// DELETE /books/:id — Kitobni o‘chirish

func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}


func main() {
	// Gin router yarat指标


	router := gin.Default()

	// Global Middleware qo‘shish
	router.Use(loggerMiddleware()) // Har bir so‘rov uchun log yozadi

	// Faqat muayyan route’lar uchun API Key middleware
	protected := router.Group("/books")
	protected.Use(apiKeyMiddleware()) // API Key tekshiruvi

	// Routes
	protected.GET("/", getBooks)
	protected.GET("/:id", getBookByID)
	protected.POST("/", createBook)
	protected.PUT("/:id", updateBook)
	protected.DELETE("/:id", deleteBook)

	// Serverni 8080 portda ishga tushirish
	router.Run(":8001")
}