package main

import (
	"fgstgu/go/src/c/gin/c02/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(Cors())
	v1 := router.Group("api/v1")
	{
		v1.POST("/user", user.CreateUser)
		v1.GET("/user", user.GetUsers)
		v1.GET("/user/:id", user.GetUser)
		v1.PUT("/user/:id", user.UpdateUser)
		v1.DELETE("/user/:id", user.DeleteUser)
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Cors func.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	}
}
