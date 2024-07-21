package service

import "github.com/anatolyjudov/job/app/model"

type StorageService struct {
	driver StorageInterface
}

func (s StorageService) SaveUser(user *model.User) bool {
	return s.driver.SaveUser(user)
}

func (s StorageService) DeleteUser(user *model.User) bool {
	return s.driver.DeleteUser(user)
}

func (s StorageService) GetUsers() []*model.User {
	return s.driver.GetUsers()
}

func (s StorageService) GetUserById(id int) *model.User {
	return s.driver.GetUserById(id)
}
