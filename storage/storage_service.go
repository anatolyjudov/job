package storage

import "github.com/anatolyjudov/job/app/model"

type StorageService struct {
	driver model.StorageInterface
}

func (s StorageService) Init() bool {
	return s.driver.Init()
}
