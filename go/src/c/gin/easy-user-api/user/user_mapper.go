package user

// ToUser converts UserTransfer to User.
func ToUser(ut UserTransfer) User {
	return User{Name: ut.Name, Age: ut.Age}
}

// ToUserTransfer converts User to User
func ToUserTransfer(user User) UserTransfer {
	return UserTransfer{ID: user.ID, Name: user.Name, Age: user.Age}
}

// ToUserTransfers converts User list to UserTransfer list.
func ToUserTransfers(users []User) []UserTransfer {
	transfers := make([]UserTransfer, len(users))
	for i, item := range users {
		transfers[i] = ToUserTransfer(item)
	}
	return transfers
}
