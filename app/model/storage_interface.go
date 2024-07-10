package model

type StorageInterface interface {
	SaveUser(user *User) bool
	GetUsers() []*User
	GetUserById(id int) *User
	DeleteUser(user *User) bool
}
