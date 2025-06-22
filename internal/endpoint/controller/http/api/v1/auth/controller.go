package auth

import (
	"context"
	"rv/internal/domain/dto/request"
	resp "rv/internal/domain/dto/response"
	"rv/pkg/apperror"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userService interface {
	RegisterUser(ctx context.Context, credentials request.RegisterCredentials) (uuid.UUID, error)
}

type authService interface {
	SendConfirmationCode(ctx context.Context, req request.SendCodeRequest) (*resp.SendCodeResponse, error)
	ConfirmCode(ctx context.Context, req request.ConfimationCodeRequest) error
}

type Controller struct {
	lgr     applogger.Logger
	builder *response.Builder

	userService userService
	authService authService
}

func NewController(logger applogger.Logger, builder *response.Builder, userService userService, authService authService) *Controller {
	return &Controller{
		lgr:     logger,
		builder: builder,

		userService: userService,
		authService: authService,
	}
}

func (h *Controller) Init(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", h.register)
		//auth.POST("/login", h.login)
		auth.POST("/code", h.sendCode)
		auth.POST("/confirm", h.confirmCode)
		//auth.POST("/refresh", h.refreshTokens)
		//auth.POST("/forgot", h.forgotPassword)
	}
}

func (h *Controller) register(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.RegisterCredentials
	err := c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}
	userId, err := h.userService.RegisterUser(ctx, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, userId))
}

func (h *Controller) sendCode(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.SendCodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}
	resp, err := h.authService.SendConfirmationCode(ctx, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, resp))
}

func (h *Controller) confirmCode(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.ConfimationCodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}
	err = h.authService.ConfirmCode(ctx, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, nil))
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
