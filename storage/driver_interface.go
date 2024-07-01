package storage

type StorageDriver interface {
	Init() bool
}
