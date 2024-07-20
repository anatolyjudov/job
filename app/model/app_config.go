package model

type AppConfig struct {
	repositories []*Repository
	context      *AppContext
}

func (c *AppConfig) Repositories() []*Repository {
	return c.repositories
}

func (c *AppConfig) GetContext() *AppContext {
	return c.context
}
