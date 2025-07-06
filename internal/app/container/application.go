package container

import (
	"rv/internal/application/auth"
	"rv/internal/application/movies"
	"rv/internal/application/reviews"
	userApp "rv/internal/application/user"
)

func (c *Container) getApplication() *applications {
	if c.applications == nil {
		c.applications = &applications{c: c}
	}
	return c.applications
}

type applications struct {
	c *Container

	user   *userApp.Service
	auth   *auth.Service
	movie  *movies.Service
	review *reviews.Service
}

func (s *applications) getUserApplicationService() *userApp.Service {
	if s.user == nil {
		s.user = userApp.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),

			s.c.getServices().getUserService(),
			s.c.getServices().getFileService(),
		)
	}
	return s.user
}

func (s *applications) getAuthApplicationService() *auth.Service {
	if s.auth == nil {
		s.auth = auth.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),

			s.c.getServices().getUserService(),
			s.c.getServices().getSMTPService(),
			s.c.getServices().getTokenService(),
		)
	}
	return s.auth
}

func (s *applications) getMoviesApplicationService() *movies.Service {
	if s.movie == nil {
		s.movie = movies.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),

			s.c.getServices().getMoviesService(),
		)
	}
	return s.movie
}

func (s *applications) getReviewsApplicationService() *reviews.Service {
	if s.review == nil {
		s.review = reviews.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),

			s.c.getServices().getReviewsService(),
		)
	}
	return s.review
}
