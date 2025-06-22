package container

import (
	smtpSrv "rv/internal/domain/smtp"
	userSrv "rv/internal/domain/user"
)

func (c *Container) getServices() *services {
	if c.services == nil {
		c.services = &services{c: c}
	}
	return c.services
}

type services struct {
	c *Container

	user *userSrv.Service
	smtp *smtpSrv.Service
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
