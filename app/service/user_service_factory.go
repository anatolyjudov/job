package service

func UserServiceFactory() *UserService {
	storageService := StorageServiceFactory()

	return &UserService{
		storage: storageService,
	}
}
