package service

import (
	"log"

	"github.com/anatolyjudov/job/storage/storage_driver"
)

func StorageServiceFactory(configService *AppConfigService) *StorageService {

	config := configService.Get()

	currentRepository := config.GetContext().CurrentRepository()

	if currentRepository.Type != "sqlite3" {
		log.Fatal("Unknown repository type")
	}

	driverInstance := storage_driver.SqliteDriver{Creds: currentRepository.Creds}

	return &StorageService{
		driver: driverInstance,
	}
}
