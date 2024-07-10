package service

import "github.com/anatolyjudov/job/app/model"

type UserService struct {
	storage *StorageService
}

func (us *UserService) Delete(user *model.User) {
	us.storage.DeleteUser(user)
}

func (us *UserService) Update(user *model.User) {
}

func (us *UserService) List() []*model.User {
	return us.storage.GetUsers()
}

func (us *UserService) Add(name string, avatar string) *model.User {

	user := model.User{}
	user.SetName(name).SetAvatarFromString(avatar)

	if us.storage.SaveUser(&user) {
		return &user
	}

	return &user
}
