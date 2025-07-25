package container

import (
	"rv/internal/domain/services/file"
	moviesSrv "rv/internal/domain/services/movies"
	"rv/internal/domain/services/reviews"
	smtpSrv "rv/internal/domain/services/smtp"
	tokenSrv "rv/internal/domain/services/token"
	userSrv "rv/internal/domain/services/user"
)

func (c *Container) getServices() *services {
	if c.services == nil {
		c.services = &services{c: c}
	}
	return c.services
}

type services struct {
	c *Container

	user   *userSrv.Service
	smtp   *smtpSrv.Service
	token  *tokenSrv.Service
	file   *file.Service
	movie  *moviesSrv.Service
	review *reviews.Service
}

func (s *services) getUserService() *userSrv.Service {
	if s.user == nil {
		s.user = userSrv.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),
			s.c.getRepositories().getUserRepository(),
		)
	}
	return s.user
}

func (s *services) getSMTPService() *smtpSrv.Service {
	if s.smtp == nil {
		s.smtp = smtpSrv.NewService(
			s.c.getLogger(),
			smtpSrv.NewConfig(
				s.c.getConfig().Email.OwnerEmail,
				s.c.getConfig().Email.OwnerPassword,
				s.c.getConfig().Email.Address,
				s.c.getConfig().Email.CodeLenght,
				s.c.getConfig().Email.CodeExp,
				s.c.getConfig().Email.MinTTL,
			),
			s.c.getCaches().getSmtpCache(),
		)
	}
	return s.smtp
}

func (s *services) getTokenService() *tokenSrv.Service {
	if s.token == nil {
		s.token = tokenSrv.NewService(
			s.c.getConfig().Jwt.AccessTTL,
			s.c.getConfig().Jwt.RefreshTTL,
			s.c.getConfig().Jwt.JwtSecret,
			s.c.getRepositories().getTokenRepository(),
		)

	}
	return s.token
}

func (s *services) getFileService() *file.Service {
	if s.file == nil {
		s.file = file.NewService(
			s.c.getLogger(),
			s.c.getRepositories().getFileRepository(),
		)

	}
	return s.file
}

func (s *services) getMoviesService() *moviesSrv.Service {
	if s.movie == nil {
		s.movie = moviesSrv.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),
			s.c.getRepositories().getMoviesRepository(),
			s.c.getRepositories().getReviewsRepository(),
		)
	}
	return s.movie
}

func (s *services) getReviewsService() *reviews.Service {
	if s.review == nil {
		s.review = reviews.NewService(
			s.c.getTransactionManager(),
			s.c.getLogger(),
			s.c.getRepositories().getReviewsRepository(),
		)
	}
	return s.review
}
