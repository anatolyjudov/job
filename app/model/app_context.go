package model

type AppContext struct {
	repository Repository
}

func (c *AppContext) CurrentRepository() *Repository {
	return &c.repository
}
