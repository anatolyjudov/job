package service

import "github.com/anatolyjudov/job/app/model"

type ConfigStorageInterface interface {
	Read() (*model.AppConfig, error)
}
