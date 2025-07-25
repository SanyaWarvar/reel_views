package user

import (
	"context"
	"rv/internal/domain/dto/request"
	resp "rv/internal/domain/dto/response"
	"rv/internal/domain/dto/user"
	apperrors "rv/internal/errors"
	"rv/pkg/apperror"
	"rv/pkg/applogger"
	"rv/pkg/constants"
	"rv/pkg/response"
	"rv/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userService interface {
	ChangeProfilePicture(ctx context.Context, req request.ChangeProfilePicture, host string) (*resp.ChangePictureResponse, error)
	GetUserById(ctx context.Context, userId uuid.UUID, host string) (*user.User, error)
}

type Controller struct {
	lgr     applogger.Logger
	builder *response.Builder

	userService userService
}

func NewController(logger applogger.Logger, builder *response.Builder, userService userService) *Controller {
	return &Controller{
		lgr:     logger,
		builder: builder,

		userService: userService,
	}
}

func (h *Controller) Init(api, authApi *gin.RouterGroup) {
	user := api.Group("/user")
	userAuth := authApi.Group("/user")
	{
		userAuth.POST("/picture", h.changeProfilePicture)
		user.GET("/profile/:id", h.getUserById)
		userAuth.GET("/profile/me", h.getMe)
	}
}

// @Summary change_profile_picture
// @Description сменить аватарку пользователя
// @Tags user
// @Produce json
// @Param data body request.ChangeProfilePicture true "data"
// @Param X-Request-Id header string true "Request id identity"
// @Param Authorization header string true "auth token"
// @Success 200 {object} response.Response{data=resp.ChangePictureResponse}
// @Failure 400 {object} response.Response{} "possible codes: invalid_token, invalid_authorization_header"
// @Failure 400 {object} response.Response{} "possible codes: bind_body, invalid_X-Request-Id"
// @Failure 422 {object} response.Response{} "possible codes: user_not_found"
// @Router /rl/api/v1/user/picture [post]
func (h *Controller) changeProfilePicture(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.ChangeProfilePicture
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperrors.InvalidAuthorizationHeader)
		return
	}
	req.UserId = userId
	picUrl, err := h.userService.ChangeProfilePicture(ctx, req, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, picUrl))
}

// @Summary get_user_by_id
// @Description получить юзера по айди
// @Tags user
// @Produce json
// @Param id path int true "id"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=user.User}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Failure 422 {object} response.Response{} "possible codes: user_not_found"
// @Router /rl/api/v1/user/profile/{id} [get]
func (h *Controller) getUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.NewBadRequestError(err.Error(), constants.BindBodyError))
		return
	}

	user, err := h.userService.GetUserById(ctx, id, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, user))
}

// @Summary get_me
// @Description получить данные о своем профиле
// @Tags user
// @Produce json
// @Param Authorization header string true "id"
// @Param X-Request-Id header string true "Request id identity"
// @Success 200 {object} response.Response{data=user.User}
// @Failure 400 {object} response.Response{} "possible codes: bind_path, invalid_X-Request-Id"
// @Failure 401 {object} response.Response{} "possible codes: invalid authorization header"
// @Router /rl/api/v1/user/profile/me [get]
func (h *Controller) getMe(c *gin.Context) {
	ctx := c.Request.Context()
	userId, err := util.GetUserId(ctx)
	if err != nil {
		_ = c.Error(apperrors.InvalidAuthorizationHeader)
		return
	}
	user, err := h.userService.GetUserById(ctx, userId, c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatusJSON(h.builder.BuildSuccessResponseBody(ctx, user))
}
