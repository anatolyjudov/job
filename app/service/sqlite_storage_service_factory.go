package service

import "github.com/anatolyjudov/job/storage/storage_driver"

func SqliteStorageServiceFactory() *StorageService {

	driverInstance := storage_driver.SqliteDriver{}

	return &StorageService{
		driver: driverInstance,
	}
}
