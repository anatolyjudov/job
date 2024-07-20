package model

type ConfigStorageInterface interface {
	Read() (*AppConfig, error)
}
