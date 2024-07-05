package model

type StorageInterface interface {
	SaveUser(user User) bool
}
