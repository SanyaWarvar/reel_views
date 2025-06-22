package container

import (
	v1 "rv/internal/endpoint/controller/http/api/v1"
	"rv/internal/endpoint/controller/http/api/v1/auth"
)

func (c *Container) getHTTPDispatcher() *v1.Dispatcher {
	if c.httpDispatcher == nil {
		c.httpDispatcher = v1.NewDispatcher(
			c.getConfig().Internal.Path,

			auth.NewController(
				c.getLogger(),
				c.getResponseBuilder(),
				c.getApplication().getUserApplicationService(),
				c.getApplication().getAuthApplicationService(),
			),
		)
	}
	return c.httpDispatcher
}
