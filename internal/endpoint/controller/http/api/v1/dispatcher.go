package v1

import (
	"rv/internal/endpoint/controller/http/api/v1/auth"
	"rv/internal/endpoint/controller/http/api/v1/movies"
	"rv/internal/endpoint/controller/http/api/v1/reviews"
	"rv/internal/endpoint/controller/http/api/v1/user"

	"github.com/gin-gonic/gin"
)

type Dispatcher struct {
	apiPath string

	auth   *auth.Controller
	user   *user.Controller
	movie  *movies.Controller
	review *reviews.Controller
}

func NewDispatcher(
	apiPath string,

	auth *auth.Controller,
	user *user.Controller,
	movie *movies.Controller,
	review *reviews.Controller,
) *Dispatcher {
	return &Dispatcher{
		apiPath: apiPath,
		auth:    auth,
		user:    user,
		movie:   movie,
		review:  review,
	}
}

func (d *Dispatcher) Init(router *gin.RouterGroup, authorization gin.HandlerFunc) {
	api := router.Group("/v1")
	{
		d.auth.Init(api)
		authorizedGroup := api.Group("", authorization)
		{
			d.user.Init(api, authorizedGroup)
			d.movie.Init(api, authorizedGroup)
			d.review.Init(api, authorizedGroup)
		}
	}
}
