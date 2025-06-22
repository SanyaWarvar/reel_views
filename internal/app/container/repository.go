package container

func (c *Container) getRepositories() *repositories {
	if c.repositories == nil {
		c.repositories = &repositories{c: c}
	}
	return c.repositories
}

type repositories struct {
	c *Container
}

/*
func (r *repositories) getCsConfigRepository() *cs_config.Repository {
	if r.csConfig == nil {
		r.csConfig = cs_config.NewRepository(r.c.getDBPool())
	}
	return r.csConfig
}
*/
