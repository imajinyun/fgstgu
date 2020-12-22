package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User struct.
type User struct {
	ID    int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name  string `gorm:"not null" form:"name" json:"name"`
	Email string `gorm:"not null" form:"email" json:"email"`
}

// InitDatabase returns a db instance.
func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test/test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if !db.Migrator().HasTable(&User{}) {
		log.Println("👌 Initialize database...")
		db.Migrator().CreateTable(&User{})
		db.Set("gorm:user", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
	}
	return db
}

// GetUsers returns a user entities.
// -> curl -i http://127.0.0.1:8080/api/v1/user
func GetUsers(c *gin.Context) {
	db := InitDatabase()

	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser returns a user entity.
// -> curl -i http://127.0.0.1:8080/api/v1/user/1
func GetUser(c *gin.Context) {
	db := InitDatabase()

	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)

	if user.ID > 0 {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}
}

// CreateUser will create a user info.
// -> curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Jack\", \"email\": \"jack@gmail.com\" }" http://localhost:8080/api/v1/user
func CreateUser(c *gin.Context) {
	db := InitDatabase()

	var user User
	c.Bind(&user)
	if user.Name != "" && user.Email != "" {
		db.Create(&user)
		c.JSON(http.StatusCreated, gin.H{"success": user})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "fields are empty"})
	}
}

// UpdateUser will update a user info.
// -> curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\": \"Tom\", \"Email\": \"tom@qq.com\" }" http://localhost:8080/api/v1/user/1
func UpdateUser(c *gin.Context) {
	db := InitDatabase()

	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)

	if user.Name != "" && user.Email != "" {
		if user.ID != 0 {
			var newUser User
			c.Bind(&newUser)
			result := User{
				ID:    user.ID,
				Name:  newUser.Name,
				Email: newUser.Email,
			}
			db.Save(&result)
			c.JSON(http.StatusOK, gin.H{"success": result})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		}
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "fields are empty"})
	}
}

// DeleteUser will delete a user.
// -> curl -i -X DELETE http://localhost:8080/api/v1/user/1
func DeleteUser(c *gin.Context) {
	db := InitDatabase()

	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)

	if user.ID > 0 {
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"success": "😂 user #" + id + " deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "🔥 user not found"})
	}
}

// OptionsUser func.
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
