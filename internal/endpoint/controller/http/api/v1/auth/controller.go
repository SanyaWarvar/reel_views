package auth

import (
	"context"
	"rv/internal/domain/auth/dto"
	"rv/pkg/apperror"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service interface {
	RegisterUser(ctx context.Context, credentials dto.AuthCredentials) (uuid.UUID, error)
}

type Controller struct {
	lgr     applogger.Logger
	srv     service
	builder *response.Builder
}

func NewController(logger applogger.Logger, srv service, builder *response.Builder) *Controller {
	return &Controller{
		lgr:     logger,
		srv:     srv,
		builder: builder,
	}
}

func (h *Controller) Init(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("register", h.register)
		//auth.POST("login", h.login)
		//auth.POST("send", h.sendCode)
		//auth.POST("confirm", h.confirmEmail)
		//auth.POST("refresh", h.refreshTokens)
		//auth.POST("forgot", h.forgotPassword)
	}
}

func (h *Controller) register(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.AuthCredentials
	err := c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}
	userId, err := h.srv.RegisterUser(ctx, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, userId))
}

/*
// @Summary Get match demo
// @Description Get match demo
// @Tags Matches
// @Produce json
// @Param Authorization header string true "Access token" default(Bearer <token>)
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=dto.DemoUrl}
// @Failure 401 {object} response.Response{}
// @Failure 403 {object} response.Response{}
// @Failure 410 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /match/:id/demo [get]
func (h *Controller) getDemoUrl(c *gin.Context) {
	id, err := util.UUIDFromString(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err))
		return
	}

	ctx := c.Request.Context()

	users, err := h.srv.GetMatchDemoUrl(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, users))
}
*/
