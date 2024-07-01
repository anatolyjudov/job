package storage

import "github.com/anatolyjudov/job/storage/storage_driver"

func StorageServiceFactory(driver string) *StorageService {

	driverInstance := storage_driver.SqliteDriver{}

	return &StorageService{
		driver: driverInstance,
	}
}
