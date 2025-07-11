package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)


func main(){
	
}

func BenchmarkHelloHandler(b *testing.B) {
    router := gin.New()
    router.GET("/hello", func(c *gin.Context) {
        c.String(200, "hello")
    })

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/hello", nil)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        router.ServeHTTP(w, req)
    }
}