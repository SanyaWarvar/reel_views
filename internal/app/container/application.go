package container

import "rv/internal/application/auth"

func (c *Container) getApplication() *applications {
	if c.applications == nil {
		c.applications = &applications{c: c}
	}
	return c.applications
}

type applications struct {
	c *Container

	auth *auth.Service
}

func (s *applications) getAuthApplicationService() *auth.Service {
	if s.auth == nil {
		s.auth = auth.NewAuthService(
			s.c.getTransactionManager(),
			s.c.getLogger(),
		)
	}
	return s.auth
}
