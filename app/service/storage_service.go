package service

import "github.com/anatolyjudov/job/app/model"

type StorageService struct {
	driver model.StorageInterface
}

func (s StorageService) SaveUser(user *model.User) bool {
	return s.driver.SaveUser(user)
}
