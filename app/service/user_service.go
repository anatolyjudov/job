package service

import "github.com/anatolyjudov/job/app/model"

type UserService struct {
	storage *StorageService
}

func (us *UserService) Add(name string, avatar string) *model.User {

	user := model.User{}
	user.SetName(name).SetAvatarFromString(avatar)

	if us.storage.SaveUser(&user) {
		return &user
	}

	return &user
}
