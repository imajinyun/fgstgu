package user

import "gorm.io/gorm"

// User struct.
type User struct {
	gorm.Model
	ID   int64
	Name string
	Age  uint8
}
