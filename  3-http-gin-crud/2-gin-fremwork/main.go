package main

import "github.com/gin-gonic/gin"

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func main(){
	router := gin.Default()

	router.GET("/about", func(ctx *gin.Context) {
		respone := map[string]string{
			"message": "Salom, bu json javob",
		}
		ctx.JSON(200, respone)
	})
	router.GET("/param/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": "Foydalnuvchi id: " + id,
		})
	})
	router.GET("/query", func(ctx *gin.Context) {
	    query := ctx.Query("name");
		if  query == "" {
			query = "Mehmod"
		}
		ctx.JSON(200, gin.H{
			"xat": "Salom, " + query + "!", 
		})
	})
	router.POST("/user", func(c *gin.Context) {
        var user User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": "Noto'g'ri JSON format"})
            return
        }
        c.JSON(200, gin.H{
            "message": "Foydalanuvchi qabul qilindi",
            "name":    user.Name,
            "age":     user.Age,
        })
    })


	router.Run(":8002")

}