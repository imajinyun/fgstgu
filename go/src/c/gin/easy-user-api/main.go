package main

import (
	"fgstgu/go/src/c/gin/easy-user-api/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test/test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user.User{})
	db.Create(&user.User{Name: "A1001", Age: 43})
	db.Create(&user.User{Name: "B1002", Age: 33})
	db.Create(&user.User{Name: "C1003", Age: 28})
	return db
}

func main() {
	db := initDatabase()
	user := InitUserAPI(db)

	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func (c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/users", user.FindAll)
		v1.GET("/users/:id", user.FindByID)
		v1.POST("/users", user.Create)
		v1.PUT("/users/:id", user.Update)
		v1.DELETE("/users/:id", user.Delete)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
