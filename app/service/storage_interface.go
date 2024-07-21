package service

import "github.com/anatolyjudov/job/app/model"

type StorageInterface interface {
	SaveUser(user *model.User) bool
	GetUsers() []*model.User
	GetUserById(id int) *model.User
	DeleteUser(user *model.User) bool
}
