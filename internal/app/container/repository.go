package container

import userRepo "rv/internal/infrastructure/repository/user"

func (c *Container) getRepositories() *repositories {
	if c.repositories == nil {
		c.repositories = &repositories{c: c}
	}
	return c.repositories
}

type repositories struct {
	c *Container

	user *userRepo.Repository
}

func (r *repositories) getUserRepository() *userRepo.Repository {
	if r.user == nil {
		r.user = userRepo.NewRepository(r.c.getDBPool())
	}
	return r.user
}
