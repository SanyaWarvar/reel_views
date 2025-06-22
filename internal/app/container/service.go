package container

func (c *Container) getServices() *services {
	if c.services == nil {
		c.services = &services{c: c}
	}
	return c.services
}

type services struct {
	c *Container
}

/*
func (s *services) getMatchService() *match.Service {
	if s.match == nil {
		s.match = match.NewService(
			s.c.getLogger(),
			s.c.getRepositories().getMatchRepository(),
		)
	}
	return s.match
}*/
