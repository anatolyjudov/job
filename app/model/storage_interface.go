package model

type StorageInterface interface {
	Init() bool
	SaveUser(user User) bool
}
