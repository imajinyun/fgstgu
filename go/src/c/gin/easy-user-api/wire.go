//+build wireinject

package main

import (
	"fgstgu/go/src/c/gin/easy-user-api/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitUserAPI returns a UserAPI.
func InitUserAPI(db *gorm.DB) user.UserAPI {
	wire.Build(
		user.ProvideUserRepository,
		user.ProvideUserService,
		user.ProvideUserAPI,
	)
	return user.UserAPI{}
}
