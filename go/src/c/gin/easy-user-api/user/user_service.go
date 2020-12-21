package user

// UserService struct.
type UserService struct {
	UserRepository UserRepository
}

// ProvideUserService returns a UserService.
func ProvideUserService(u UserRepository) UserService {
	return UserService{UserRepository: u}
}

// FindAll returns all users.
func (u *UserService) FindAll() []User {
	return u.UserRepository.FindAll()
}

// FindByID returns a user.
func (u *UserService) FindByID(id int64) User {
	return u.UserRepository.FindByID(id)
}

// Save will persist the user to the DB.
func (u *UserService) Save(user User) User {
	u.UserRepository.Save(user)
	return user
}

// Delete will delete the user from the DB.
func (u *UserService) Delete(user User) {
	u.UserRepository.Delete(user)
}
