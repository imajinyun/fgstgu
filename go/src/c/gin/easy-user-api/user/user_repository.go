package user

import (
	"gorm.io/gorm"
)

// UserRepository struct.
type UserRepository struct {
	DB *gorm.DB
}

// ProvideUserRepository returns a UserRepository.
func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// FindAll returns all users.
func (u *UserRepository) FindAll() []User {
	var users []User
	u.DB.Find(&users)
	return users
}

// FindByID returns a user by ID.
func (u *UserRepository) FindByID(id int64) User {
	var user User
	u.DB.First(&user, id)
	return user
}

// Save will persist the user to the DB.
func (u *UserRepository) Save(user User) User {
	u.DB.Save(&user)
	return user
}

// Delete will delete the user from the DB.
func (u *UserRepository) Delete(user User) {
	u.DB.Delete(&user)
}
