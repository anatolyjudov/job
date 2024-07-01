package storage_driver

type SqliteDriver struct{}

func (d SqliteDriver) Init() bool {
	return true
}
