package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global variables
var (
	rdb *redis.Client
	db  *gorm.DB
	ctx = context.Background()
)

// Title model
type Title struct {
	ID    uint   `gorm:"primaryKey" json:"ID"`
	Title string `json:"Title"`
}

func main() {
	// 1️⃣ Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // agar Docker ichida bo'lsa: "redis:6379"
	})

	// 2️⃣ GORM MySQL connection
	dsn := "root:yangi_parol@tcp(127.0.0.1:3306)/amalyotApi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// 3️⃣ AutoMigrate
	if err := db.AutoMigrate(&Title{}); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// 4️⃣ Gin router
	r := gin.Default()

	// GET /titles - read-through caching
	r.GET("/titles", func(c *gin.Context) {
		val, err := rdb.Get(ctx, "titles").Result()
		if err == nil {
			// Cache hit
			var titles []Title
			if err := json.Unmarshal([]byte(val), &titles); err == nil {
				c.JSON(200, gin.H{"cached": true, "data": titles})
				return
			}
			log.Println("Cache unmarshal failed:", err)
		} else {
			log.Println("Cache miss:", err)
		}

		// Cache miss yoki unmarshalling xato → DB'dan o'qish
		var titles []Title
		if err := db.Find(&titles).Error; err != nil {
			c.JSON(500, gin.H{"error": "DB query failed"})
			return
		}

		// Redis'ga yozish (TTL = 60s)
		data, _ := json.Marshal(titles)
		if err := rdb.Set(ctx, "titles", data, 60*time.Second).Err(); err != nil {
			log.Println("Redis SET failed:", err)
		}

		c.JSON(200, gin.H{"cached": false, "data": titles})
	})

	// POST /titles - yangi title qo'shish + cache invalidate
	r.POST("/titles", func(c *gin.Context) {
		var t Title
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&t).Error; err != nil {
			c.JSON(500, gin.H{"error": "DB insert failed"})
			return
		}

		// Cache tozalash
		if err := rdb.Del(ctx, "titles").Err(); err != nil {
			log.Println("Redis DEL failed:", err)
		}

		c.JSON(200, t)
	})

	// 5️⃣ Run server
	r.Run(":8082")
}
