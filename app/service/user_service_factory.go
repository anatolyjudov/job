package service

func UserServiceFactory() *UserService {

	storageService := SqliteStorageServiceFactory()

	return &UserService{
		storage: storageService,
	}
}
