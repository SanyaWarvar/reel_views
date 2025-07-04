package v1

import (
	"rv/internal/endpoint/controller/http/api/v1/auth"
	"rv/internal/endpoint/controller/http/api/v1/user"

	"github.com/gin-gonic/gin"
)

type Dispatcher struct {
	apiPath string

	auth *auth.Controller
	user *user.Controller
}

func NewDispatcher(
	apiPath string,

	auth *auth.Controller,
	user *user.Controller,
) *Dispatcher {
	return &Dispatcher{
		apiPath: apiPath,
		auth:    auth,
		user:    user,
	}
}

func (d *Dispatcher) Init(router *gin.RouterGroup, authorization gin.HandlerFunc) {
	api := router.Group("/v1")
	{
		d.auth.Init(api)
		authorizedGroup := api.Group("", authorization)
		{
			d.user.Init(api, authorizedGroup)
		}
	}
}
