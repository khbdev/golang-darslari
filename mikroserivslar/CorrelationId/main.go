package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



func CorrelationMiddleware() gin.HandlerFunc {
return  func(ctx *gin.Context) {
	corrID := ctx.GetHeader("X-Correlation-ID")
	if corrID == "" {
		corrID = uuid.New().String()
	}

	ctx.Set("correlation_id", corrID)

	ctx.Writer.Header().Set("X-Correlation-ID", corrID)

	log.Printf(`{"service":"gateway","msg":"Request received","correlation_id":"%s"}`, corrID)
	
	ctx.Next()
	log.Printf(`{"service":"gateway","msg":"Request finished","correlation_id":"%s"}`, corrID)
 }
}





func main(){
r := gin.Default()

r.Use(CorrelationMiddleware())


r.GET("/play", func(ctx *gin.Context) {
	corrID := ctx.GetString("correlation_id")
	log.Printf(`{"service":"payment","msg":"Processing payment","correlation_id":"%s"}`, corrID)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "payment succes", 
		"correlation_id": corrID,
	})
})

r.Run(":8081")

}