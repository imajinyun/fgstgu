package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAPI struct.
type UserAPI struct {
	UserService UserService
}

// ProvideUserAPI returns a UserAPI.
func ProvideUserAPI(u UserService) UserAPI {
	return UserAPI{UserService: u}
}

// FindAll method for UserAPI returns all users.
func (u UserAPI) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, gin.H{"users": ToUserTransfers(users)})
}

// FindByID method for UserAPI returns a user by id.
func (u UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(int64(id))
	c.JSON(http.StatusOK, gin.H{"user": ToUserTransfer(user)})
}

// Create method for UserAPI will create user to the DB.
func (u UserAPI) Create(c *gin.Context) {
	var userTransfer UserTransfer
	err := c.BindJSON(&userTransfer)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	createUser := u.UserService.Save(ToUser(userTransfer))
	c.JSON(http.StatusOK, gin.H{"user": ToUserTransfer(createUser)})
}

// Update method for UserAPI will update user by id.
func (u UserAPI) Update(c *gin.Context) {
	var userTransfer UserTransfer
	err := c.BindJSON(&userTransfer)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(int64(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	user.Name = userTransfer.Name
	user.Age = userTransfer.Age
	log.Println(user)
	u.UserService.Save(user)
	c.Status(http.StatusOK)
}

// Delete method for UserAPI will delete user by id.
func (u *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := u.UserService.FindByID(int64(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	u.UserService.Delete(user)
	c.Status(http.StatusOK)
}
