package v1

import (
	"rv/internal/endpoint/controller/http/api/v1/auth"

	"github.com/gin-gonic/gin"
)

type Dispatcher struct {
	apiPath string

	auth *auth.Controller
}

func NewDispatcher(
	apiPath string,

	auth *auth.Controller,
) *Dispatcher {
	return &Dispatcher{
		apiPath: apiPath,
		auth:    auth,
	}
}

func (d *Dispatcher) Init(router *gin.RouterGroup) {
	api := router.Group("/v1")
	{
		d.auth.Init(api)
		/*authorizedGroup := api.Group("", authorization)
		{

		}*/
	}
}
