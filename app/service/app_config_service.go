package service

import (
	"log"

	"github.com/anatolyjudov/job/app/model"
)

type AppConfigService struct {
	storage      ConfigStorageInterface
	cachedConfig *model.AppConfig
}

func (s *AppConfigService) Get() *model.AppConfig {

	if s.cachedConfig != nil {
		return s.cachedConfig
	}

	config, err := s.storage.Read()

	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	s.cachedConfig = config

	return config
}
