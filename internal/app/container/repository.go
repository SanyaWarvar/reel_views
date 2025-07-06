package container

import (
	"rv/internal/infrastructure/repository/file"
	"rv/internal/infrastructure/repository/movies"
	"rv/internal/infrastructure/repository/reviews"
	tokensRepo "rv/internal/infrastructure/repository/tokens"
	userRepo "rv/internal/infrastructure/repository/user"
)

func (c *Container) getRepositories() *repositories {
	if c.repositories == nil {
		c.repositories = &repositories{c: c}
	}
	return c.repositories
}

type repositories struct {
	c *Container

	user   *userRepo.Repository
	token  *tokensRepo.Repository
	file   *file.Repository
	movie  *movies.Repository
	review *reviews.Repository
}

func (r *repositories) getUserRepository() *userRepo.Repository {
	if r.user == nil {
		r.user = userRepo.NewRepository(r.c.getDBPool())
	}
	return r.user
}

func (r *repositories) getTokenRepository() *tokensRepo.Repository {
	if r.token == nil {
		r.token = tokensRepo.NewRepository(r.c.getDBPool())
	}
	return r.token
}

func (r *repositories) getFileRepository() *file.Repository {
	if r.file == nil {
		r.file = file.NewRepository(r.c.getDBPool())
	}
	return r.file
}

func (r *repositories) getMoviesRepository() *movies.Repository {
	if r.movie == nil {
		r.movie = movies.NewRepository(r.c.getDBPool())
	}
	return r.movie
}

func (r *repositories) getReviewsRepository() *reviews.Repository {
	if r.review == nil {
		r.review = reviews.NewRepository(r.c.getDBPool())
	}
	return r.review
}
