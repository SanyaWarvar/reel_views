package container

import (
	v1 "rv/internal/endpoint/controller/http/api/v1"
	"rv/internal/endpoint/controller/http/api/v1/auth"
	"rv/internal/endpoint/controller/http/api/v1/movies"
	"rv/internal/endpoint/controller/http/api/v1/reviews"
	"rv/internal/endpoint/controller/http/api/v1/user"
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

			user.NewController(
				c.getLogger(),
				c.getResponseBuilder(),
				c.getApplication().getUserApplicationService(),
			),

			movies.NewController(
				c.getLogger(),
				c.getResponseBuilder(),
				c.getApplication().getMoviesApplicationService(),
			),

			reviews.NewController(
				c.getLogger(),
				c.getResponseBuilder(),
				c.getApplication().getReviewsApplicationService(),
			),
		)
	}
	return c.httpDispatcher
}
